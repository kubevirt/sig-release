---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| VM supports kubevirt.io/immediate-data-volume-creation: "false" which delays creating DataVolumeTemplates until VM is started  | [#12194](https://github.com/kubevirt/kubevirt/pull/12194) | [mhenriks](https://github.com/mhenriks) |
| Adding newMacAddresses validatewebhook for  VMCloneAPI  | [#11802](https://github.com/kubevirt/kubevirt/pull/11802) | [matthewei](https://github.com/matthewei) |
| Adding support for the `igb` network interface model  | [#11754](https://github.com/kubevirt/kubevirt/pull/11754) | [nickolaev](https://github.com/nickolaev) |
| * Reduced the severity of log messages when a `VolumeSnapshotClass` is not found. When snapshots are not enabled for a volume, the reason will still be displayed in the `status.volumeSnapshotStatuses` field of a `VirtualMachine` resource.  | [#12254](https://github.com/kubevirt/kubevirt/pull/12254) | [jkinred](https://github.com/jkinred) |
| virt-api: unencode authorization extra headers  | [#12460](https://github.com/kubevirt/kubevirt/pull/12460) | [mhenriks](https://github.com/mhenriks) |
| Fix: eviction requests to completed virt-launcher pods cannot trigger a live migration  | [#12451](https://github.com/kubevirt/kubevirt/pull/12451) | [fossedihelm](https://github.com/fossedihelm) |
| The `expand-spec` subresource API now applies defaults to the returned `VirtualMachine` to ensure the `VirtualMachineInstanceSpec` within is closer to the eventual version used when starting the original `VirtualMachine`.  | [#11881](https://github.com/kubevirt/kubevirt/pull/11881) | [lyarwood](https://github.com/lyarwood) |
| This version of KubeVirt includes upgraded virtualization technology based on libvirt 10.5.0 and QEMU 9.0.0.<br>Each new release of libvirt and QEMU contains numerous improvements and bug fixes.  | [#12452](https://github.com/kubevirt/kubevirt/pull/12452) | [andreabolognani](https://github.com/andreabolognani) |
| fix some comments  | [#12425](https://github.com/kubevirt/kubevirt/pull/12425) | [fudancoder](https://github.com/fudancoder) |
| Use optional interface at passt binding sidecar  | [#12354](https://github.com/kubevirt/kubevirt/pull/12354) | [qinqon](https://github.com/qinqon) |
| Drop `ForceRestart` and `ForceStop` methods from client-go  | [#12268](https://github.com/kubevirt/kubevirt/pull/12268) | [fossedihelm](https://github.com/fossedihelm) |
| Network binding plugins: Enable the ability to specify compute memory overhead  | [#12235](https://github.com/kubevirt/kubevirt/pull/12235) | [orelmisan](https://github.com/orelmisan) |
| Fix wrong KubeVirtVMIExcessiveMigrations alert calculation in an upgrade scenario.  | [#12209](https://github.com/kubevirt/kubevirt/pull/12209) | [orenc1](https://github.com/orenc1) |
| Fix: persistent tpm can be used with vmis containing dots in their name  | [#12261](https://github.com/kubevirt/kubevirt/pull/12261) | [fossedihelm](https://github.com/fossedihelm) |
| Add perf-scale benchmarks for release v1.3  | [#12247](https://github.com/kubevirt/kubevirt/pull/12247) | [Sreeja1725](https://github.com/Sreeja1725) |
| BugFix: Grant namespace admin RBAC to passthrough a client USB to a VMI  | [#12181](https://github.com/kubevirt/kubevirt/pull/12181) | [akalenyu](https://github.com/akalenyu) |
| Fix missing performance metrics for VMI resources  | [#12096](https://github.com/kubevirt/kubevirt/pull/12096) | [machadovilaca](https://github.com/machadovilaca) |
| Add unit tests to check for API backward compatibility  | [#11856](https://github.com/kubevirt/kubevirt/pull/11856) | [Sreeja1725](https://github.com/Sreeja1725) |
| Add CPU/Memory utilization of components metrics to kubevirt benchmarks  | [#12116](https://github.com/kubevirt/kubevirt/pull/12116) | [Sreeja1725](https://github.com/Sreeja1725) |
| Virt export route has an edge termination of redirect  | [#12195](https://github.com/kubevirt/kubevirt/pull/12195) | [awels](https://github.com/awels) |
| enable only for VMs with memory >= 1Gi  | [#12212](https://github.com/kubevirt/kubevirt/pull/12212) | [acardace](https://github.com/acardace) |
| Only a single vgpu display option with ramfb will be configured per VMI.  | [#12053](https://github.com/kubevirt/kubevirt/pull/12053) | [vladikr](https://github.com/vladikr) |
| fix RerunOnFailure stuck in Provisioning  | [#12193](https://github.com/kubevirt/kubevirt/pull/12193) | [acardace](https://github.com/acardace) |
| Updated common-instancetypes bundles to v1.0.1  | [#12186](https://github.com/kubevirt/kubevirt/pull/12186) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| VMs with a single socket and NetworkInterfaceMultiqueue enabled require a restart to hotplug additional CPU sockets.  | [#12180](https://github.com/kubevirt/kubevirt/pull/12180) | [0xFelix](https://github.com/0xFelix) |
| All `preferredCPUTopology` constants prefixed with `Prefer` have been deprecated and will be removed in a future version of the `instancetype.kubevirt.io` API.  | [#11927](https://github.com/kubevirt/kubevirt/pull/11927) | [lyarwood](https://github.com/lyarwood) |
| `PreferredDiskDedicatedIoThread` is now only applied to `virtio` disk devices  | [#12169](https://github.com/kubevirt/kubevirt/pull/12169) | [lyarwood](https://github.com/lyarwood) |
| chore: bump virtio-win image version to 0.1.248  | [#12125](https://github.com/kubevirt/kubevirt/pull/12125) | [ksimon1](https://github.com/ksimon1) |
| Memory Hotplug fixes and stabilization  | [#12128](https://github.com/kubevirt/kubevirt/pull/12128) | [acardace](https://github.com/acardace) |
| Bugfix: Implement retry mechanism in export server and vmexport  | [#11911](https://github.com/kubevirt/kubevirt/pull/11911) | [alromeros](https://github.com/alromeros) |
| Introduce validatingAdmissionPolicy to restrict node patches on virt-handler  | [#11982](https://github.com/kubevirt/kubevirt/pull/11982) | [RamLavi](https://github.com/RamLavi) |
| Fix VMPools when `LiveUpdate` as `vmRolloutStrategy` is used.  | [#12119](https://github.com/kubevirt/kubevirt/pull/12119) | [acardace](https://github.com/acardace) |


_This page is updated daily._
