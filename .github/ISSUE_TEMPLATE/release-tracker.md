---
name: Release Tracker
about: Used to track the work which needs to be done every release cycle
title: ':clipboard:[1.XX] Release Tracker'
labels: sig/release, lifecycle/frozen, kind/tracker 
---

<!--
This template is used for tracking tasks that need to be done in a release cycle.
The issue should be kept open for the entirety of the release cycle until all tasks are completed.
-->

* Release Cycle: `KubeVirt 1.XX`
* [Release Timeline](https://github.com/kubevirt/sig-release/blob/main/releases/v1.XX/schedule.md) <!-- DON'T FORGET TO UPDATE THE LINK -->

## Release tasks

### Milestone: Release Cycle Start

- [ ] Capture feedback from previous release cycle retro
- [ ] Open PR proposal for release schedule, announce, request feedback

  ---

### Milestone: Alpha 0 Tag (`1.XX.0-alpha.0`)

- [ ] Finalize release schedule
- [ ] Announce schedule
- [ ] Remind community about VEP Freeze
- [ ] Start provider for latest k8s beta/stable, add on-demand pre-submit
- [ ] **Tag XX.0-alpha.0**

  ---

### Milestone: VEP Freeze
- [ ] Review all VEPs targeting this release:
  - [ ] Ensure each has an approved design
  - [ ] Update status (implementable / deferred / rejected)
  - [ ] Confirm SIG sign-off on each
- [ ] **VEP Freeze**

  ---

### Milestone: Beta 0 Tag (`1.XX.0-beta.0`)

- [ ] Prepare unconference for next development cycle
- [ ] Mandatory pre-submit lanes
- [ ] Non-standard lanes bumped by stakeholders:
  - [ ] sig-compute
  - [ ] sig-storage
  - [ ] sig-network
  - [ ] other sigs
- [ ] **Tag 1.XX.0-beta.0**

  ---

### Milestone: Code Freeze

- [ ] Begin observation of issues, CI signal, test flakes, critical PRs
- [ ] Notify SIGs about upcoming Code Freeze deadline
- [ ] Bring exceptions to kubevirt-dev
- [ ] Update common-instancetypes bundle in KubeVirt repo
- [ ] Update virt-template bundle in KubeVirt repo
- [ ] Review quarantined tests report
- [ ] **Code Freeze: Branch release-1.XX**
- [ ] Create release-1.XX prow config
- [ ] **Tag v1.XX.0-rc.0 on release-1.XX**

  ---

### Milestone: Release Day

- [ ] Generate api testdata for backward compatibility testing
- [ ] Gather scale & performance data
- [ ] Unconference
- [ ] Code stabilization sign-offs:
  - [ ] sig-compute
  - [ ] sig-network
  - [ ] sig-storage
  - [ ] other sigs
- [ ] Categorise release notes, create User Guide PR with /hold
- [ ] Draft release highlights blog
- [ ] Collect SIG sign-offs on blog
- [ ] Publish scale & performance data
- [ ] Update release notes for last-minute bugfixes
- [ ] **Tag v1.XX.0 on release-1.XX**
- [ ] Ensure docs and blog PRs are approved
- [ ] Announce on kubevirt-dev
- [ ] Publish/unhold release notes and blog
- [ ] Promote through social media

  ---

### Post Release

- [ ] Update Enhancements Tracking board link in ROADMAP
- [ ] Update label-approved-veps script with new TARGET_PROJECT_URL
- [ ] Pin release branch to CDI version after CDI GA
- [ ] After CDI GA: Pin release-1.XX branch to associated CDI version via `KUBEVIRT_CUSTOM_CDI_VERSION` in `automation/test.sh`.

## Further comments

`NONE`