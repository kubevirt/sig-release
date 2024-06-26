---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
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
| Support Memory Hotplug with Hugepages  | [#12109](https://github.com/kubevirt/kubevirt/pull/12109) | [acardace](https://github.com/acardace) |
| By enabling NodeRestriction feature gate, Kubevirt now authorize virt-handler's requests to modify VMs.  | [#12009](https://github.com/kubevirt/kubevirt/pull/12009) | [xpivarc](https://github.com/xpivarc) |
| The `CommonInstancetypesDeployment` feature and gate are retrospectively moved to Beta from the 1.2.0 release.  | [#11681](https://github.com/kubevirt/kubevirt/pull/11681) | [lyarwood](https://github.com/lyarwood) |
| Add descheduler compatibility  | [#12025](https://github.com/kubevirt/kubevirt/pull/12025) | [fossedihelm](https://github.com/fossedihelm) |
| Bump k8s deps to 0.30.0  | [#12097](https://github.com/kubevirt/kubevirt/pull/12097) | [fossedihelm](https://github.com/fossedihelm) |
| Less privileged virt-operator ClusterRole  | [#12089](https://github.com/kubevirt/kubevirt/pull/12089) | [jean-edouard](https://github.com/jean-edouard) |
| BugFix: Graceful deletion skipped for any delete call to the VM (not VMI) resource  | [#12064](https://github.com/kubevirt/kubevirt/pull/12064) | [akalenyu](https://github.com/akalenyu) |
| Add support for building and running kubevirt on s390x.  | [#10490](https://github.com/kubevirt/kubevirt/pull/10490) | [jschintag](https://github.com/jschintag) |
| Network hotplug feature is declared as Beta.  | [#12079](https://github.com/kubevirt/kubevirt/pull/12079) | [EdDev](https://github.com/EdDev) |
| `LiveUpdates`  of VMs using instance types are now supported with the same caveats as when making changes to a vanilla VM.  | [#11455](https://github.com/kubevirt/kubevirt/pull/11455) | [lyarwood](https://github.com/lyarwood) |
| Create kubevirt_vmi_launcher_memory_overhead_bytes metric  | [#12000](https://github.com/kubevirt/kubevirt/pull/12000) | [machadovilaca](https://github.com/machadovilaca) |
| The 'passt' core network binding is discontinued and removed.  | [#11915](https://github.com/kubevirt/kubevirt/pull/11915) | [ormergi](https://github.com/ormergi) |
| fix starting VM with Manual RunStrategy  | [#12016](https://github.com/kubevirt/kubevirt/pull/12016) | [acardace](https://github.com/acardace) |
| Implement volume migration and introduce the migration updateVolumesStrategy field  | [#11533](https://github.com/kubevirt/kubevirt/pull/11533) | [alicefr](https://github.com/alicefr) |
| Add kubevirt_vmi_last_connection_timestamp_seconds metric  | [#11934](https://github.com/kubevirt/kubevirt/pull/11934) | [assafad](https://github.com/assafad) |
| Introduce export.kibevirt.io/v1beta1  | [#11956](https://github.com/kubevirt/kubevirt/pull/11956) | [mhenriks](https://github.com/mhenriks) |
| BugFix: Fix restore panic in case of volumesnapshot missing  | [#11996](https://github.com/kubevirt/kubevirt/pull/11996) | [ShellyKa13](https://github.com/ShellyKa13) |
| snapshot: Ignore unfreeze error if VMSnapshot deleting  | [#11957](https://github.com/kubevirt/kubevirt/pull/11957) | [mhenriks](https://github.com/mhenriks) |
| Create kubevirt_vmi_info metric  | [#11906](https://github.com/kubevirt/kubevirt/pull/11906) | [machadovilaca](https://github.com/machadovilaca) |
| Infra components control-plane nodes NoSchedule toleration  | [#11969](https://github.com/kubevirt/kubevirt/pull/11969) | [iholder101](https://github.com/iholder101) |
| Introduce snapshot.kibevirt.io/v1beta1  | [#11955](https://github.com/kubevirt/kubevirt/pull/11955) | [mhenriks](https://github.com/mhenriks) |
| Restart of a VM is required when the CPU socket count is reduced  | [#11883](https://github.com/kubevirt/kubevirt/pull/11883) | [orelmisan](https://github.com/orelmisan) |
| add Intel Gaudi to adopters.  | [#11835](https://github.com/kubevirt/kubevirt/pull/11835) | [talcoh2x](https://github.com/talcoh2x) |
| Refactor device plugins to use a base plugin and define a common interface  | [#11344](https://github.com/kubevirt/kubevirt/pull/11344) | [aerosouund](https://github.com/aerosouund) |
| Bug fix: Correctly reflect RestartRequired condition  | [#11973](https://github.com/kubevirt/kubevirt/pull/11973) | [fossedihelm](https://github.com/fossedihelm) |
| Fix RerunOnFailure RunStrategy  | [#11963](https://github.com/kubevirt/kubevirt/pull/11963) | [acardace](https://github.com/acardace) |
| `VirtualMachines` referencing an instance type are now allowed when the `LiveUpdate` feature is enabled and will trigger the `RestartRequired` condition if the reference within the `VirtualMachine` is changed.  | [#11962](https://github.com/kubevirt/kubevirt/pull/11962) | [lyarwood](https://github.com/lyarwood) |
| Update virtctl to use v1beta1 endpoint for both regular and async image upload  | [#11942](https://github.com/kubevirt/kubevirt/pull/11942) | [ido106](https://github.com/ido106) |
| Updated common-instancetypes bundles to v1.0.0  | [#11648](https://github.com/kubevirt/kubevirt/pull/11648) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| Require scheduling infra components onto control-plane nodes  | [#11659](https://github.com/kubevirt/kubevirt/pull/11659) | [iholder101](https://github.com/iholder101) |
| `ControllerRevisions` containing instance types and preferences are now upgraded to their latest available version when the `VirtualMachine` owning them is resync'd by `virt-controller`.  | [#10545](https://github.com/kubevirt/kubevirt/pull/10545) | [lyarwood](https://github.com/lyarwood) |
| The 'macvtap' core network binding is discontinued and removed.  | [#11901](https://github.com/kubevirt/kubevirt/pull/11901) | [EdDev](https://github.com/EdDev) |
| Bugfix: Fix VM manifest rendering in export controller  | [#11922](https://github.com/kubevirt/kubevirt/pull/11922) | [alromeros](https://github.com/alromeros) |
| sidecar-shim: allow stderr log from binary hooks  | [#11908](https://github.com/kubevirt/kubevirt/pull/11908) | [victortoso](https://github.com/victortoso) |
| `spreadOptions` have been introduced to preferences in order to allow for finer grain control of the `preferSpread` `preferredCPUTopology`. This includes the ability to now spread vCPUs across guest visible sockets, cores and threads.  | [#11729](https://github.com/kubevirt/kubevirt/pull/11729) | [lyarwood](https://github.com/lyarwood) |
| Allow to hotplug vcpus for VMs with CPU requests and/or limits set  | [#11655](https://github.com/kubevirt/kubevirt/pull/11655) | [acardace](https://github.com/acardace) |
| The SLIRP core binding is deprecated and removed.  | [#11701](https://github.com/kubevirt/kubevirt/pull/11701) | [EdDev](https://github.com/EdDev) |
| SMBios sidecar can be built out-of-tree  | [#11846](https://github.com/kubevirt/kubevirt/pull/11846) | [victortoso](https://github.com/victortoso) |
| The network-info annotation is now used for mapping between SR-IOV network and the underlying device PCI address  | [#11788](https://github.com/kubevirt/kubevirt/pull/11788) | [ormergi](https://github.com/ormergi) |
| Add the updateVolumeStrategy field  | [#11700](https://github.com/kubevirt/kubevirt/pull/11700) | [alicefr](https://github.com/alicefr) |
| This version of KubeVirt includes upgraded virtualization technology based on libvirt 10.0.0 and QEMU 8.2.0.<br>Each new release of libvirt and QEMU contains numerous improvements and bug fixes.  | [#11256](https://github.com/kubevirt/kubevirt/pull/11256) | [andreabolognani](https://github.com/andreabolognani) |
| Build KubeVirt with go v1.22.2  | [#11482](https://github.com/kubevirt/kubevirt/pull/11482) | [brianmcarey](https://github.com/brianmcarey) |
| Add kubevirt.io/testWorkloadUpdateMigrationAbortion annotation and a mechanism to abort workload updates  | [#11641](https://github.com/kubevirt/kubevirt/pull/11641) | [alicefr](https://github.com/alicefr) |
| Fix the live updates for volumes and disks  | [#11770](https://github.com/kubevirt/kubevirt/pull/11770) | [alicefr](https://github.com/alicefr) |
| Re-adding Cloudflare to our ADOPTERS list  | [#11790](https://github.com/kubevirt/kubevirt/pull/11790) | [aburdenthehand](https://github.com/aburdenthehand) |
| Fix: SEV methods in client-go now satisfy the proxy server configuration, if provided  | [#11718](https://github.com/kubevirt/kubevirt/pull/11718) | [fossedihelm](https://github.com/fossedihelm) |
| Updated go version of the client-go to 1.21  | [#11685](https://github.com/kubevirt/kubevirt/pull/11685) | [fossedihelm](https://github.com/fossedihelm) |
| Extend network binding plugin to support device-info DownwardAPI.  | [#11618](https://github.com/kubevirt/kubevirt/pull/11618) | [AlonaKaplan](https://github.com/AlonaKaplan) |
| Collect VMI OS info from the Guest agent as `kubevirt_vmi_phase_count` metric labels  | [#11283](https://github.com/kubevirt/kubevirt/pull/11283) | [assafad](https://github.com/assafad) |
| Rename rest client metrics to include kubevirt prefix  | [#11676](https://github.com/kubevirt/kubevirt/pull/11676) | [machadovilaca](https://github.com/machadovilaca) |
| New memory statistics added named kubevirt_memory_delta_from_requested_bytes  | [#11557](https://github.com/kubevirt/kubevirt/pull/11557) | [avlitman](https://github.com/avlitman) |
| Improve the handling of ordinal pod interface name for upgrade  | [#11678](https://github.com/kubevirt/kubevirt/pull/11678) | [Vicente-Cheng](https://github.com/Vicente-Cheng) |
| Build the `passt`custom CNI binary statically, for the `passt` network binding plugin.<br>Resolves  https://github.com/kubevirt/kubevirt/issues/11586  | [#11653](https://github.com/kubevirt/kubevirt/pull/11653) | [EdDev](https://github.com/EdDev) |
| Fix kubevirt_vm_created_total being broken down by virt-api pod  | [#11294](https://github.com/kubevirt/kubevirt/pull/11294) | [machadovilaca](https://github.com/machadovilaca) |
| Add e2e tests for metrics  | [#11307](https://github.com/kubevirt/kubevirt/pull/11307) | [machadovilaca](https://github.com/machadovilaca) |
| virtual machines instance will no longer be stuck in an irrecoverable state after an interrupted postcopy migration. Instead, these will fail and could be restarted again.  | [#11479](https://github.com/kubevirt/kubevirt/pull/11479) | [vladikr](https://github.com/vladikr) |
| emission of k8s logs when using programmatic focus with `FIt`  | [#11416](https://github.com/kubevirt/kubevirt/pull/11416) | [dhiller](https://github.com/dhiller) |
| Make 'image' field in hook sidecar annotation optional.  | [#11272](https://github.com/kubevirt/kubevirt/pull/11272) | [dharmit](https://github.com/dharmit) |
| Support HyperV Passthrough: automatically use all available HyperV features  | [#11500](https://github.com/kubevirt/kubevirt/pull/11500) | [iholder101](https://github.com/iholder101) |
| Reduce the downwardMetrics server maximum number of request per second to 1.  | [#11484](https://github.com/kubevirt/kubevirt/pull/11484) | [jcanocan](https://github.com/jcanocan) |
| Allow to hotplug memory for VMs with memory limits set  | [#11498](https://github.com/kubevirt/kubevirt/pull/11498) | [acardace](https://github.com/acardace) |
| Build KubeVirt with Go version 1.21.8  | [#11470](https://github.com/kubevirt/kubevirt/pull/11470) | [brianmcarey](https://github.com/brianmcarey) |
| Improve handling of export resources in virtctl vmexport  | [#11312](https://github.com/kubevirt/kubevirt/pull/11312) | [alromeros](https://github.com/alromeros) |
| Bugfix: Allow vmexport download redirections by printing logs into stderr  | [#11367](https://github.com/kubevirt/kubevirt/pull/11367) | [alromeros](https://github.com/alromeros) |
| Bugfix: Improve handling of IOThreads with incompatible buses  | [#11219](https://github.com/kubevirt/kubevirt/pull/11219) | [alromeros](https://github.com/alromeros) |
| virtctl: It is possible to import volumes from GCS when creating a VM now  | [#11149](https://github.com/kubevirt/kubevirt/pull/11149) | [0xFelix](https://github.com/0xFelix) |
| KubeVirtComponentExceedsRequestedCPU and KubeVirtComponentExceedsRequestedMemory alerts are deprecated; they do not indicate a genuine issue.  | [#11404](https://github.com/kubevirt/kubevirt/pull/11404) | [avlitman](https://github.com/avlitman) |
| add cloudraft to adopters.  | [#11331](https://github.com/kubevirt/kubevirt/pull/11331) | [anjuls](https://github.com/anjuls) |
| add perf-scale benchmarks for release v1.2  | [#11387](https://github.com/kubevirt/kubevirt/pull/11387) | [alaypatel07](https://github.com/alaypatel07) |
| Expose volumesnapshot error in vmsnapshot object  | [#11095](https://github.com/kubevirt/kubevirt/pull/11095) | [ShellyKa13](https://github.com/ShellyKa13) |
| Bug-fix: Fix nil panic if VM update fails  | [#11372](https://github.com/kubevirt/kubevirt/pull/11372) | [xpivarc](https://github.com/xpivarc) |
| BugFix: Ensure DataVolumes created by virt-controller (DataVolumeTemplates) are recreated and owned by the VM in the case of DR and backup/restore.  | [#11267](https://github.com/kubevirt/kubevirt/pull/11267) | [mhenriks](https://github.com/mhenriks) |
| BugFix: Fixed incorrect APIVersion of APIResourceList  | [#10900](https://github.com/kubevirt/kubevirt/pull/10900) | [KarstenB](https://github.com/KarstenB) |
| fix(ksm): set the `kubevirt.io/ksm-enabled` node label to true if the ksm is managed by KubeVirt, instead of reflect the actual ksm value.  | [#11306](https://github.com/kubevirt/kubevirt/pull/11306) | [fossedihelm](https://github.com/fossedihelm) |
| More information in the migration state of VMI / migration objects  | [#11330](https://github.com/kubevirt/kubevirt/pull/11330) | [jean-edouard](https://github.com/jean-edouard) |
| Fix perfscale buckets error  | [#11264](https://github.com/kubevirt/kubevirt/pull/11264) | [machadovilaca](https://github.com/machadovilaca) |
| Extend OWNERS for sig-buildsystem  | [#11183](https://github.com/kubevirt/kubevirt/pull/11183) | [dhiller](https://github.com/dhiller) |
| fix(vmclone): delete vmclone resource when the target vm is deleted  | [#11058](https://github.com/kubevirt/kubevirt/pull/11058) | [fossedihelm](https://github.com/fossedihelm) |
| Bug fix: VM controller doesn't corrupt its cache anymore  | [#11265](https://github.com/kubevirt/kubevirt/pull/11265) | [xpivarc](https://github.com/xpivarc) |
| Fix migration breaking in case the VM has an rng device after hotplugging a block volume on cgroupsv2  | [#11205](https://github.com/kubevirt/kubevirt/pull/11205) | [akalenyu](https://github.com/akalenyu) |
| Bugfix: Improve error reporting when fsfreeze fails  | [#11051](https://github.com/kubevirt/kubevirt/pull/11051) | [alromeros](https://github.com/alromeros) |
| Move some verification from the VMI create validation webhook to the CRD<br>The webhook will no longer return the following errors:<br>* "spec.accessCredentials list exceeds the 256 element limit in length"<br>* "spec.domain.devices.disks list exceeds the 256 element limit in length"<br>* "spec.domain.devices.interfaces list exceeds the 256 element limit in length"<br>* "spec.networks list exceeds the 256 element limit in length"<br>* "spec.volumes list exceeds the 256 element limit in length"  | [#11156](https://github.com/kubevirt/kubevirt/pull/11156) | [nunnatsa](https://github.com/nunnatsa) |
| node-labeller: Remove obsolete functionalities  | [#11146](https://github.com/kubevirt/kubevirt/pull/11146) | [RamLavi](https://github.com/RamLavi) |


_This page is updated daily._
