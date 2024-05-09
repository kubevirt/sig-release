---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
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
