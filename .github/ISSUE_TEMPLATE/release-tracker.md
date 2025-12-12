---
name: Release Tracker
about: Used to track the work which needs to be done every release cycle
title: ':clipboard:[1.XX] Release Tracker'
labels: sig/release, lifecycle/frozen 
---

<!--
This template is used for tracking tasks that need to be done in a release cycle.
The issue should be kept open for the entirety of the release cycle until all tasks are completed.
-->

* Release Cycle: `KubeVirt 1.XX`
* [Release Timeline](https://github.com/kubevirt/sig-release/blob/main/releases/v1.XX/schedule.md) <!-- DON'T FORGET TO UPDATE THE LINK -->

## Release tasks

### 1. Before the start of the Release Cycle

- [ ] Captured feedback from the previous release cycle retro and planned to incorporate it into the release cycle.
- [ ] Started planning the release schedule by opening a PR proposal, announcing it, and requesting feedback via email
  to [kubevirt-dev](https://groups.google.com/forum/#!forum/kubevirt-dev).

### 2. First weeks of the release cycle up to Virtualization Enhancements Proposal Freeze

- [ ] Release schedule finalized.
  - [ ] Announce schedule.
- [ ] KubeVirt Alpha 0:
  - [ ] Remind the community about Virtualization Enhancements Freeze.
  - [ ] Tag 1.XX.0-alpha.0.
  - [ ] Start the provider for the latest k8s beta/stable version, adding on-demand pre-submit.

### 3. Enhancements freeze up to Release Code Freeze

- [ ] Tag 1.XX.0-beta.0:
  - [ ] Prepare unconference for next version development cycle.
  - [ ] Mandatory pre-submit lanes.
  - [ ] Non-standard lanes (i.e. migrations, multus, ipv6) have been bumped by stakeholders.
    - [ ] sig-compute
    - [ ] sig-storage
    - [ ] sig-network
    - [ ] other sigs
- [ ] Begin casual observation of issues, CI signal, test flakes, and critical PRs.
- [ ] Notify SIGs about upcoming Code Freeze Deadline.
- [ ] Bring exceptions to [kubevirt-dev](https://groups.google.com/forum/#!forum/kubevirt-dev).

### 4. Code Freeze

- [ ] Code Freeze:
  - [ ] Update common-instancetypes bundle in KubeVirt repo.
  - [ ] Review https://storage.googleapis.com/kubevirt-prow/reports/quarantined-tests/kubevirt/kubevirt/index.html
  - [ ] Branch release-1.XX.
  - [ ] Create release-1.XX prow config.
  - [ ] Tag v1.XX.0-rc.0 on release-1.XX.

### 5. After Code Freeze

- [ ] After Code Freeze:
  - [ ] Generate api testdata for the new release and push them to main for backward compatibility testing.
  - [ ] Gather scale & performance data.
  - [ ] Unconference.
  - [ ] Categorise release notes and create User Guide PR with /hold.
  - [ ] Draft release highlights blog.
- [ ] Test and production code stabilization:
  - [ ] sig-compute
  - [ ] sig-network
  - [ ] sig-storage
  - [ ] other sigs

### 6. Around Release Day

- [ ] Collect SIG sign-offs on blog.
- [ ] Publish scale & performance data.

### 7. Release Day

- [ ] Update release notes for any last-minute bugfixes.
- [ ] Tag v1.XX.0 on release-1.XX
- [ ] Ensure docs and blog PRs are approved.
- [ ] Send the email to [kubevirt-dev](https://groups.google.com/forum/#!forum/kubevirt-dev) for announcing.
- [ ] Publish/unhold release notes and blog.
- [ ] Promote through social media.

### 8. Post Release

- [ ] Update KubeVirt Enhancements Tracking board link in [ROADMAP](https://github.com/kubevirt/sig-release/blob/main/ROADMAP.md).
- [ ] Update [label-approved-veps script](https://github.com/kubevirt/kubevirt/blob/main/.github/workflows/label_approved_veps.yaml) with the new `TARGET_PROJECT_URL`.

## Further comments

`NONE`