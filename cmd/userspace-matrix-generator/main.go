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
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	releasePattern  = regexp.MustCompile(`^v([0-9]+)\.([0-9]+)\.([0-9]+)$`)
	defaultPattern  = regexp.MustCompile(`(?m)^build\s+--define=centos_stream_version=([0-9]+)\s*(?:#.*)?$`)
	treePattern     = regexp.MustCompile(`(?ms)^rpmtree\(\s*(.*?)^\)`)
	namePattern     = regexp.MustCompile(`\bname\s*=\s*"([^"]+)"`)
	launcherPattern = regexp.MustCompile(`^launcherbase_(x86_64|aarch64|s390x)(?:_cs([0-9]+))?$`)
	rpmsPattern     = regexp.MustCompile(`(?s)\brpms\s*=\s*\[([^\]]*)\]`)
	rpmPattern      = regexp.MustCompile(`"@([^"/]+)//rpm"`)
	packagePattern  = regexp.MustCompile(`^(.+)-([0-9]+)__([^-]+)-(.+)\.([^.]+)$`)
	basePattern     = regexp.MustCompile(`^([0-9]+)\.`)
	commentPattern  = regexp.MustCompile(`(?m)#.*$`)
	decodeRPM       = strings.NewReplacer("__plus__", "+", "__tilde__", "~", "__caret__", "^")
)

type release struct {
	tag     string
	version [3]int
}

type component struct {
	label    string
	packages []string
}

var components = []component{
	{"libvirt", []string{"libvirt-daemon-driver-qemu"}},
	{"QEMU", []string{"qemu-kvm-core"}},
	{"edk2", []string{"edk2-ovmf", "edk2-aarch64"}},
	{"SeaBIOS", []string{"seabios"}},
	{"swtpm", []string{"swtpm"}},
	{"passt", []string{"passt"}},
	{"virtiofsd", []string{"virtiofsd"}},
}

type rpm struct {
	name, epoch, version, release, arch string
}

func (p rpm) evr() string {
	return p.epoch + ":" + p.version + "-" + p.release
}

type row struct {
	arch, distro, kernel string
	versions             []string
}

func main() {
	repo := flag.String("kubevirt-path", "../kubevirt", "path to a KubeVirt git checkout containing the release tags")
	manifest := flag.String("releases-file", "releases/userspace-releases.json", "JSON array of exact stable release tags to track")
	output := flag.String("output", "releases/userspace-support-matrix.md", "generated Markdown path")
	check := flag.Bool("check", false, "verify that the output matches the tagged sources without writing files")
	flag.Parse()
	if flag.NArg() != 0 {
		log.Fatal("unexpected positional arguments")
	}
	if err := run(*repo, *manifest, *output, *check); err != nil {
		log.Fatal(err)
	}
}

func run(repo, manifest, output string, check bool) error {
	data, err := os.ReadFile(manifest)
	if err != nil {
		return err
	}
	releases, err := parseReleases(data)
	if err != nil {
		return fmt.Errorf("%s: %w", manifest, err)
	}
	generated, err := generate(repo, releases)
	if err != nil {
		return err
	}
	if check {
		current, err := os.ReadFile(output)
		if err != nil {
			return err
		}
		if !bytes.Equal(current, generated) {
			return fmt.Errorf("%s is out of date; rerun userspace-matrix-generator without --check", output)
		}
		return nil
	}
	// Generate all releases successfully before replacing the published matrix.
	return os.WriteFile(output, generated, 0644)
}

func parseReleases(data []byte) ([]release, error) {
	var tags []string
	if err := json.Unmarshal(data, &tags); err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		return nil, fmt.Errorf("at least one release is required")
	}
	var releases []release
	seen := map[[2]int]bool{}
	for _, tag := range tags {
		match := releasePattern.FindStringSubmatch(tag)
		if match == nil {
			return nil, fmt.Errorf("%q is not an exact stable release tag (expected vMAJOR.MINOR.PATCH)", tag)
		}
		r := release{tag: tag}
		for i := range r.version {
			n, err := strconv.Atoi(match[i+1])
			if err != nil {
				return nil, fmt.Errorf("invalid version %q: %w", tag, err)
			}
			r.version[i] = n
		}
		minor := [2]int{r.version[0], r.version[1]}
		if seen[minor] {
			return nil, fmt.Errorf("multiple snapshots for release series %d.%d", minor[0], minor[1])
		}
		seen[minor] = true
		releases = append(releases, r)
	}
	slices.SortFunc(releases, func(a, b release) int { return slices.Compare(b.version[:], a.version[:]) })
	return releases, nil
}

func gitFile(repo, commit, path string) (string, error) {
	out, err := exec.Command("git", "-C", repo, "show", commit+":"+path).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("read %s at %s: %w: %s", path, commit, err, strings.TrimSpace(string(out)))
	}
	return string(out), nil
}

func generate(repo string, releases []release) ([]byte, error) {
	var b strings.Builder
	b.WriteString(matrixIntro)
	b.WriteString("| KubeVirt release | Architecture | Base userspace | Recommended host kernel | libvirt | QEMU | edk2 | SeaBIOS | swtpm | passt | virtiofsd |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |\n")
	var sources strings.Builder
	for _, r := range releases {
		// Resolve only tags, never a similarly named branch or the working tree.
		out, err := exec.Command("git", "-C", repo, "rev-parse", "--verify", "refs/tags/"+r.tag+"^{commit}").CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("resolve release tag %s (fetch it into %s first): %w: %s", r.tag, repo, err, strings.TrimSpace(string(out)))
		}
		commit := strings.TrimSpace(string(out))
		build, err := gitFile(repo, commit, "rpm/BUILD.bazel")
		if err != nil {
			return nil, err
		}
		bazelrc, err := gitFile(repo, commit, ".bazelrc")
		if err != nil {
			return nil, err
		}
		rows, err := parseRows(build, bazelrc)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", r.tag, err)
		}
		for _, row := range rows {
			cells := []string{"[" + r.tag + "](https://github.com/kubevirt/kubevirt/releases/tag/" + r.tag + ")", row.arch, row.distro, row.kernel}
			cells = append(cells, row.versions...)
			fmt.Fprintf(&b, "| %s |\n", strings.Join(cells, " | "))
		}
		fmt.Fprintf(&sources, "- %s: [RPM dependencies](https://github.com/kubevirt/kubevirt/blob/%s/rpm/BUILD.bazel), [default build configuration](https://github.com/kubevirt/kubevirt/blob/%s/.bazelrc) (commit `%s`).\n", r.tag, commit, commit, commit)
	}
	b.WriteString("\n## Sources\n\n")
	b.WriteString(sources.String())
	return []byte(b.String()), nil
}

func parseRows(build, bazelrc string) ([]row, error) {
	defaultCS := ""
	defaults := defaultPattern.FindAllStringSubmatch(bazelrc, -1)
	for _, match := range defaults {
		defaultCS = match[1]
	}
	if len(defaults) > 1 {
		return nil, fmt.Errorf("multiple default CentOS Stream definitions in .bazelrc")
	}
	// These generated rpmtree rules have literal name and rpms attributes.
	// Reject unsupported layouts instead of taking versions from unrelated trees.
	build = commentPattern.ReplaceAllString(build, "")
	byArch := map[string]row{}
	for _, block := range treePattern.FindAllStringSubmatch(build, -1) {
		name := namePattern.FindStringSubmatch(block[1])
		if name == nil {
			continue
		}
		target := launcherPattern.FindStringSubmatch(name[1])
		if target == nil || target[2] != defaultCS {
			continue
		}
		arch := target[1]
		if _, exists := byArch[arch]; exists {
			return nil, fmt.Errorf("duplicate launcherbase target for %s", arch)
		}
		list := rpmsPattern.FindStringSubmatch(block[1])
		if list == nil {
			return nil, fmt.Errorf("%s has no literal RPM list", name[1])
		}
		if remaining := strings.Trim(rpmPattern.ReplaceAllString(list[1], ""), " \t\r\n,"); remaining != "" {
			return nil, fmt.Errorf("%s has nonliteral RPM dependencies", name[1])
		}
		packages := map[string]rpm{}
		for _, ref := range rpmPattern.FindAllStringSubmatch(list[1], -1) {
			match := packagePattern.FindStringSubmatch(ref[1])
			if match == nil {
				return nil, fmt.Errorf("unrecognized RPM reference %q", ref[1])
			}
			p := rpm{match[1], match[2], decodeRPM.Replace(match[3]), decodeRPM.Replace(match[4]), match[5]}
			if p.arch != arch && p.arch != "noarch" {
				return nil, fmt.Errorf("%s contains RPM %s for architecture %s", name[1], p.name, p.arch)
			}
			if _, exists := packages[p.name]; exists {
				return nil, fmt.Errorf("%s contains multiple RPMs for %s", name[1], p.name)
			}
			packages[p.name] = p
		}
		base, ok := packages["centos-stream-release"]
		if !ok {
			return nil, fmt.Errorf("%s has no centos-stream-release RPM", name[1])
		}
		major := basePattern.FindStringSubmatch(base.version)
		if major == nil || (defaultCS != "" && major[1] != defaultCS) {
			return nil, fmt.Errorf("%s has an unexpected base version %q", name[1], base.version)
		}
		r := row{arch: arch, distro: "CentOS Stream " + major[1], kernel: "CentOS Stream " + major[1] + " / EL " + major[1] + " distribution kernel"}
		for _, c := range components {
			value := "N/A"
			for _, pkg := range c.packages {
				if p, exists := packages[pkg]; exists {
					if value != "N/A" {
						return nil, fmt.Errorf("%s has ambiguous %s packages", name[1], c.label)
					}
					value = "`" + p.evr() + "`"
				}
			}
			if value == "N/A" && (c.label == "libvirt" || c.label == "QEMU") {
				return nil, fmt.Errorf("%s is missing %s", name[1], c.label)
			}
			r.versions = append(r.versions, value)
		}
		byArch[arch] = r
	}
	if len(byArch) == 0 {
		return nil, fmt.Errorf("no launcherbase RPM trees match the default CentOS Stream configuration %q", defaultCS)
	}
	var rows []row
	for _, arch := range []string{"x86_64", "aarch64", "s390x"} {
		if r, exists := byArch[arch]; exists {
			rows = append(rows, r)
		}
	}
	return rows, nil
}

const matrixIntro = `<!-- Generated by cmd/userspace-matrix-generator. Do not edit by hand. -->
# KubeVirt release userspace and host-kernel support matrix

This table records the userspace bundled in the default upstream ` + "`virt-launcher`" + `
build for each tracked release. Each row is an exact patch-release snapshot;
component versions can change between patch releases. The tracked tags are listed
in [userspace-releases.json](userspace-releases.json). See the
[Kubernetes support matrix](k8s-support-matrix.md) for release lifecycle information;
inclusion here does not extend a release's support period.

Use a maintained host kernel from the corresponding CentOS Stream / Enterprise
Linux (EL) distribution family. This is alignment guidance, not a claim that every
kernel or distribution has been tested. The host provides the kernel; it is not
bundled in the container. Kernel versions alone do not capture distribution
backports or guarantee KVM/vhost feature compatibility. Keep source and destination
hosts aligned for live migration. See the
[installation prerequisites](https://kubevirt.io/user-guide/cluster_admin/installation/#host-kernel-and-userland-compatibility)
for host requirements and compatibility guidance.

Versions below use RPM **epoch:version-release**, including the epoch even when
zero. Architecture identifies the launcher build, not the architecture of every
RPM it contains; firmware packages can be ` + "`noarch`" + `. The package names are listed
below. **N/A** means the component is absent from that architecture's
launcher RPM tree.

The component packages are ` + "`libvirt-daemon-driver-qemu`, `qemu-kvm-core`,\n`edk2-ovmf` (x86_64) / `edk2-aarch64` (aarch64), `seabios`, `swtpm`,\n`passt`, and `virtiofsd`" + `. The base userspace is identified by the
` + "`centos-stream-release`" + ` RPM. Alternative build configurations and downstream
images can use different dependencies.

To refresh this generated document, follow the
[generator instructions](../cmd/userspace-matrix-generator/README.md).

`
