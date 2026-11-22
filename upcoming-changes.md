---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Cloud-init hook sidecars must use CloudInitData since 2019; the legacy CloudInitNoCloudSource payload and fallback code are dropped from the codebase.<br>Cloud-init metadata removed by a hook is no longer silently rebuilt by KubeVirt  | [#19616](https://github.com/kubevirt/kubevirt/pull/19616) | [dankenigsberg](https://github.com/dankenigsberg) |
| virt-operator now preserves additional annotations on the virt-exportproxy Route across reconciliation.  | [#19224](https://github.com/kubevirt/kubevirt/pull/19224) | [davmarro](https://github.com/davmarro) |
| virtctl port-forward now rejects a port mapping with more than two ports, such as `8080:9090:9999`, instead of silently ignoring the extra fields.  | [#19541](https://github.com/kubevirt/kubevirt/pull/19541) | [jay23606](https://github.com/jay23606) |
| Fix a panic in virt-synchronization-controller when a decentralized migration publishes sourceState before targetState.  | [#19823](https://github.com/kubevirt/kubevirt/pull/19823) | [omkarP-bit](https://github.com/omkarP-bit) |
| Live migration no longer switches to post-copy or pauses the guest on completion timeout(when configured) during the initial bulk-copy; the migration is aborted instead. Post-copy/pause assist are triggered only in the convergence phase afterwards.  | [#19390](https://github.com/kubevirt/kubevirt/pull/19390) | [michalskrivanek](https://github.com/michalskrivanek) |
| Bugfix: Prevent VM snapshot export PVCs from being selected as VM backend storage  | [#19741](https://github.com/kubevirt/kubevirt/pull/19741) | [mohitjoshi-hey](https://github.com/mohitjoshi-hey) |
| virt-handler no longer changes the ownership of HostDisk and filesystem PVC directories.  | [#19923](https://github.com/kubevirt/kubevirt/pull/19923) | [Acedus](https://github.com/Acedus) |
| fix, hotplug: retain disk serial capitalization  | [#19669](https://github.com/kubevirt/kubevirt/pull/19669) | [dsanatar](https://github.com/dsanatar) |
| New CR option to allow fallback from migration network to pod network on failure  | [#16940](https://github.com/kubevirt/kubevirt/pull/16940) | [jean-edouard](https://github.com/jean-edouard) |
| Setting blockMultiQueue to false now explicitly configures one request queue for virtio-blk disks and virtio-scsi controllers. API defaults and behavior for omitted settings are unchanged. The fix takes effect when a new VMI is started.  | [#19691](https://github.com/kubevirt/kubevirt/pull/19691) | [davmarro](https://github.com/davmarro) |
| SCC: Revised to least privilege principle  | [#19101](https://github.com/kubevirt/kubevirt/pull/19101) | [xpivarc](https://github.com/xpivarc) |
| virt-handler no longer panics when it gets a lifecycle request for a VMI that is not on its node  | [#19821](https://github.com/kubevirt/kubevirt/pull/19821) | [krishhna24](https://github.com/krishhna24) |
| Node hooks now fall back to the Plugin-level failureStrategy when they do not set their own, like launcher hooks. Plugins with spec.failureStrategy: Ignore no longer block VM start, stop or migration when one of their node hooks without its own failureStrategy fails.  | [#19724](https://github.com/kubevirt/kubevirt/pull/19724) | [iholder101](https://github.com/iholder101) |
| Bump libvirtxml to v1.12007.0 and fix the smmuv3 IOMMU ssidsize attribute, so Grace IOVirtualization VMIs get their SSID size and keep their IOMMU driver attributes when guest definition plugin hooks run.  | [#19723](https://github.com/kubevirt/kubevirt/pull/19723) | [iholder101](https://github.com/iholder101) |
| Disabling decentralized live migration also removes virt-synchronization controller PBD  | [#19787](https://github.com/kubevirt/kubevirt/pull/19787) | [awels](https://github.com/awels) |
| action required: virt-exportproxy is now HPA-scaled (not tied to KubeVirt CR replicas). On upgrade it no longer prefers control-plane placement via spec.infra; node selectors/tolerations prefer spec.workloads when set, otherwise spec.infra. spec.infra.replicas is no longer applied to the export-proxy Deployment—use CustomizeComponents to pin replicas (and opt out of the HPA), or set workloads/infra scheduling as needed.  | [#18459](https://github.com/kubevirt/kubevirt/pull/18459) | [awels](https://github.com/awels) |
| Update the fixed libguestfs appliances bundled in libguestfs-tools to include NTFS-3G and FUSE support. NTFS content can be read using native `virt-*` tools such as `virt-ls` and `virt-cat`; mounting NTFS with `guestfish` (including through `virtctl guestfs`) remains unsupported.  | [#19577](https://github.com/kubevirt/kubevirt/pull/19577) | [etirta](https://github.com/etirta) |
| Fix VMIs with SR-IOV, GPU or PCI host devices failing to start when the Plugins feature gate is enabled and any Plugin exists in the cluster.  | [#19720](https://github.com/kubevirt/kubevirt/pull/19720) | [iholder101](https://github.com/iholder101) |
| manifest-templator now rejects --infra-replicas values above 255 instead of silently wrapping them around.  | [#19602](https://github.com/kubevirt/kubevirt/pull/19602) | [cvas22](https://github.com/cvas22) |
| Fixed a virt-handler panic when a device plugin is restarted, which on nodes with host devices happened on every kubelet restart.  | [#18622](https://github.com/kubevirt/kubevirt/pull/18622) | [alichtl](https://github.com/alichtl) |
| Live migrations are no longer scheduled from a node running an updated virt-handler onto a node that has not yet been updated, preventing migrations from getting stuck during a rolling upgrade.  | [#19134](https://github.com/kubevirt/kubevirt/pull/19134) | [stu-gott](https://github.com/stu-gott) |
| Expired pull mode backups are marked failed instead of succeeded.  | [#19478](https://github.com/kubevirt/kubevirt/pull/19478) | [awels](https://github.com/awels) |
| Beta: IncrementalBackup, and introduce backup.kubevirt.io/v1beta1  | [#19637](https://github.com/kubevirt/kubevirt/pull/19637) | [Acedus](https://github.com/Acedus) |
| CBT overlays are now keyed by volume and PVC name. Pointing a volume at a different PVC triggers a Full backup instead of an incorrect incremental.  | [#19404](https://github.com/kubevirt/kubevirt/pull/19404) | [Acedus](https://github.com/Acedus) |
| Fix for cluster-profiler dump command, to now dump profiling for control-plane pods when the selector field is empty, as earlier was returning no results when run without a label selector  | [#19225](https://github.com/kubevirt/kubevirt/pull/19225) | [ameya-keskar](https://github.com/ameya-keskar) |
| Fixed a bug where a Unix socket listener may be left open if an invalid socket path is specified  | [#19532](https://github.com/kubevirt/kubevirt/pull/19532) | [caxu-rh](https://github.com/caxu-rh) |
| Fixed hotplug attachment pods entering an infinite create/delete loop when<br>mutating admission webhooks inject additional volumes into hp-volume pods.  | [#19442](https://github.com/kubevirt/kubevirt/pull/19442) | [nijinashok](https://github.com/nijinashok) |
| Prevent virt-operator from repeatedly reconciling idle deployments in response to unrelated leader-election Leases and routine Lease renewals.  | [#19417](https://github.com/kubevirt/kubevirt/pull/19417) | [lyarwood](https://github.com/lyarwood) |


_This page is updated daily._
