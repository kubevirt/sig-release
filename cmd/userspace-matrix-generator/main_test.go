/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright the KubeVirt Authors.
 */

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// An unrelated devel tree and an alternative CS10 build must never supply the
// versions reported for a default CS9 release. ARM firmware is a noarch RPM.
const buildFixture = `rpmtree(
    name = "libvirt-devel_x86_64_cs9",
    rpms = ["@libvirt-daemon-driver-qemu-0__99.0-1.el9.x86_64//rpm"],
)

rpmtree(
    name = "launcherbase_x86_64_cs10",
    rpms = [
        "@centos-stream-release-0__10.0-1.el10.x86_64//rpm",
        "@libvirt-daemon-driver-qemu-0__12.0-1.el10.x86_64//rpm",
        "@qemu-kvm-core-18__10.0-1.el10.x86_64//rpm",
    ],
)

rpmtree(
    name = "launcherbase_aarch64_cs9",
    rpms = [
        "@centos-stream-release-0__9.0-38.el9.aarch64//rpm",
        "@libvirt-daemon-driver-qemu-0__11.10.0-19.el9.aarch64//rpm",
        "@qemu-kvm-core-17__10.1.0-20.el9.aarch64//rpm",
        "@edk2-aarch64-0__20241117-8.el9.noarch//rpm",
    ],
)

rpmtree(
    name = "launcherbase_x86_64_cs9",
    rpms = [
        "@centos-stream-release-0__9.0-38.el9.x86_64//rpm",
        "@libvirt-daemon-driver-qemu-0__11.10.0-19.el9.x86_64//rpm",
        "@qemu-kvm-core-17__10.1.0-20.el9.x86_64//rpm",
        "@edk2-ovmf-0__20241117-8.el9.noarch//rpm",
        "@seabios-0__1.16.3-4.el9.x86_64//rpm",
        "@swtpm-0__0.8.0-2.el9.x86_64//rpm",
        "@passt-0__0__caret__20260611.ga9c61ff-1.el9.x86_64//rpm",
        "@virtiofsd-0__1.13.0__tilde__rc1__plus__git-1.el9.x86_64//rpm",
        # "@libvirt-daemon-driver-qemu-0__99.0-1.el9.x86_64//rpm",
    ],
)
`

const bazelrcFixture = "build --define=centos_stream_version=9\nbuild:cs10 --define=centos_stream_version=10\n"

func TestReleaseSnapshots(t *testing.T) {
	releases, err := parseReleases([]byte(`["v1.9.1", "v1.10.0", "v1.7.4"]`))
	if err != nil {
		t.Fatal(err)
	}
	var tags []string
	for _, r := range releases {
		tags = append(tags, r.tag)
	}
	if !reflect.DeepEqual(tags, []string{"v1.10.0", "v1.9.1", "v1.7.4"}) {
		t.Fatalf("releases sorted incorrectly: %v", tags)
	}
	for _, invalid := range []string{
		`[]`, `null`, `["v1.9"]`, `["release-1.9"]`, `["v1.9.0-rc.0"]`,
		`["v1.9.0", "v1.9.1"]`, `["v1.9.1", "v1.9.1"]`, `{"tag":"v1.9.1"}`,
	} {
		t.Run(invalid, func(t *testing.T) {
			if _, err := parseReleases([]byte(invalid)); err == nil {
				t.Fatal("accepted invalid release list")
			}
		})
	}
}

func TestDefaultBuildAndArchitectureVersions(t *testing.T) {
	rows, err := parseRows(buildFixture, bazelrcFixture)
	if err != nil {
		t.Fatal(err)
	}
	if len(rows) != 2 || rows[0].arch != "x86_64" || rows[1].arch != "aarch64" {
		t.Fatalf("unexpected architecture rows: %+v", rows)
	}
	if rows[0].distro != "CentOS Stream 9" || rows[0].kernel != "CentOS Stream 9 / EL 9 distribution kernel" {
		t.Fatalf("incorrect base/kernel guidance: %+v", rows[0])
	}
	want := []string{
		"`0:11.10.0-19.el9`", "`17:10.1.0-20.el9`", "`0:20241117-8.el9`",
		"`0:1.16.3-4.el9`", "`0:0.8.0-2.el9`", "`0:0^20260611.ga9c61ff-1.el9`",
		"`0:1.13.0~rc1+git-1.el9`",
	}
	if !reflect.DeepEqual(rows[0].versions, want) {
		t.Fatalf("wrong RPM versions: got %v, want %v", rows[0].versions, want)
	}
	if rows[1].versions[2] != "`0:20241117-8.el9`" || rows[1].versions[3] != "N/A" {
		t.Fatalf("incorrect ARM firmware versions: %+v", rows[1])
	}

	cs10, err := parseRows(buildFixture, "build --define=centos_stream_version=10\n")
	if err != nil || len(cs10) != 1 || cs10[0].distro != "CentOS Stream 10" || cs10[0].versions[1] != "`18:10.0-1.el10`" {
		t.Fatalf("CS10 default selection: %+v, %v", cs10, err)
	}

	legacy := strings.ReplaceAll(buildFixture, "_cs9", "")
	rows, err = parseRows(legacy, "build --color=yes\n")
	if err != nil || len(rows) != 2 || rows[0].versions[0] != want[0] {
		t.Fatalf("legacy unsuffixed targets: %+v, %v", rows, err)
	}
}

func TestInvalidDependencyLayouts(t *testing.T) {
	for _, tc := range []struct{ name, build, rc, message string }{
		{"unknown default", buildFixture, "build --define=centos_stream_version=11\n", "no launcherbase"},
		{"missing default", buildFixture, "", "no launcherbase"},
		{"duplicate default", buildFixture, bazelrcFixture + bazelrcFixture, "multiple default"},
		{"duplicate tree", buildFixture + buildFixture, bazelrcFixture, "duplicate launcherbase"},
		{"missing base", strings.ReplaceAll(buildFixture, "centos-stream-release", "other-release"), bazelrcFixture, "no centos-stream-release"},
		{"wrong base", strings.ReplaceAll(buildFixture, "9.0-38", "10.0-38"), bazelrcFixture, "unexpected base"},
		{"missing QEMU", strings.ReplaceAll(buildFixture, "qemu-kvm-core", "other-qemu"), bazelrcFixture, "missing QEMU"},
		{"wrong architecture", strings.ReplaceAll(buildFixture, "20.el9.aarch64", "20.el9.x86_64"), bazelrcFixture, "architecture x86_64"},
		{"malformed reference", strings.ReplaceAll(buildFixture, "qemu-kvm-core-17__", "qemu-kvm-core-17:"), bazelrcFixture, "unrecognized RPM"},
		{"computed RPM list", strings.ReplaceAll(buildFixture, "rpms = [", "rpms = select(["), bazelrcFixture, "no literal RPM list"},
		{"computed RPM entry", strings.ReplaceAll(buildFixture, "rpms = [", "rpms = [optional_rpms,"), bazelrcFixture, "nonliteral RPM dependencies"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parseRows(tc.build, tc.rc)
			if err == nil || !strings.Contains(err.Error(), tc.message) {
				t.Fatalf("got %v, want error containing %q", err, tc.message)
			}
		})
	}
}

func TestGenerateFromTagAndCheckWithoutWriting(t *testing.T) {
	repo := t.TempDir()
	git := func(args ...string) string {
		t.Helper()
		cmd := exec.Command("git", append([]string{"-C", repo}, args...)...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("git %v: %v: %s", args, err, out)
		}
		return strings.TrimSpace(string(out))
	}
	write := func(path, content string) {
		t.Helper()
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
	}
	git("init")
	if err := os.Mkdir(filepath.Join(repo, "rpm"), 0755); err != nil {
		t.Fatal(err)
	}
	write(filepath.Join(repo, "rpm", "BUILD.bazel"), buildFixture)
	write(filepath.Join(repo, ".bazelrc"), bazelrcFixture)
	git("add", ".")
	git("-c", "user.name=Test", "-c", "user.email=test@example.com", "-c", "commit.gpgsign=false", "commit", "-m", "release dependencies")
	git("-c", "tag.gpgsign=false", "tag", "v1.9.1")
	commit := git("rev-parse", "HEAD")
	// Dirty checkout data must not change the release inventory.
	write(filepath.Join(repo, "rpm", "BUILD.bazel"), "invalid working tree")
	dir := t.TempDir()
	manifest, output := filepath.Join(dir, "releases.json"), filepath.Join(dir, "matrix.md")
	write(manifest, `["v1.9.1"]`)
	if err := run(repo, manifest, output, false); err != nil {
		t.Fatal(err)
	}
	generated, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(generated), "`17:10.1.0-20.el9`") || !strings.Contains(string(generated), "/blob/"+commit+"/rpm/BUILD.bazel") {
		t.Fatalf("matrix lacks tagged versions or commit provenance: %s", generated)
	}
	if err := run(repo, manifest, output, true); err != nil {
		t.Fatal(err)
	}
	write(output, "stale matrix")
	if err := run(repo, manifest, output, true); err == nil || !strings.Contains(err.Error(), "out of date") {
		t.Fatalf("check did not detect stale matrix: %v", err)
	}
	// A bad source must not replace previously published output.
	write(manifest, `["v1.9.2"]`)
	if err := run(repo, manifest, output, false); err == nil || !strings.Contains(err.Error(), "resolve release tag") {
		t.Fatalf("missing tag was accepted: %v", err)
	}
	git("branch", "v1.9.2")
	if err := run(repo, manifest, output, false); err == nil {
		t.Fatal("a branch was accepted in place of a release tag")
	}
	current, err := os.ReadFile(output)
	if err != nil || string(current) != "stale matrix" {
		t.Fatalf("check or failed generation modified output: %q, %v", current, err)
	}
}
