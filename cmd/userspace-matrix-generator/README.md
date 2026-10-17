# userspace-matrix-generator

Generate the [userspace and host-kernel matrix](../../releases/userspace-support-matrix.md)
from `rpm/BUILD.bazel` and `.bazelrc` at exact KubeVirt release tags. The generator
reads the `launcherbase` RPM trees used to build `virt-launcher`, selecting the
default CentOS Stream configuration and retaining architecture differences and
RPM epochs. It does not execute the build files, build images, or require a
container runtime. Its git reads work offline once the tags are available.

## Release updates

After each stable release, including patch releases:

1. Fetch the published release tags in a KubeVirt checkout:

   ```bash
   git -C ../kubevirt fetch origin --tags
   ```

2. Update `releases/userspace-releases.json` with the latest exact stable tag for
   each supported minor release. Add a series at GA and remove it when support
   ends. Pre-release tags, duplicate series, branches, and floating references
   are rejected. The list is maintained by SIG Release; the generator does not
   infer support status from the number or age of tags.

3. From the sig-release repository root, generate and commit the manifest and
   matrix together:

   ```bash
   go run ./cmd/userspace-matrix-generator --kubevirt-path ../kubevirt
   ```

   The output includes links to the exact source commits for review. The working
   tree and checked-out branch of KubeVirt do not affect the result. A missing tag
   or unsupported dependency layout fails before the output is replaced.

4. Verify the result before publishing:

   ```bash
   go run ./cmd/userspace-matrix-generator --kubevirt-path ../kubevirt --check
   go test ./cmd/userspace-matrix-generator
   ```

The same generation command can be called by release automation after fetching
the tags and updating the tracked release list. `--check` is read-only and exits
unsuccessfully if the published matrix differs from its source data. Use
`--releases-file` and `--output` to override the default manifest and output paths.

## Scope

The table describes the default upstream release build, not every RPM in the
repository or an alternative CentOS Stream build. Required base, libvirt, and QEMU
packages must exist; absent optional components are shown as N/A. Kernel guidance
identifies the matching distribution family, rather than inventing a kernel RPM
version from a container that does not ship a host kernel. This inventory does
not establish a new distribution certification or change release lifecycles.

Other repositories, including the user-guide, can link directly to
`https://github.com/kubevirt/sig-release/blob/main/releases/userspace-support-matrix.md`.
