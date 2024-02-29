---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
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
| virtctl: Specifying size when creating a VM and using --volume-import to clone a PVC or a VolumeSnapshot is optional now  | [#11144](https://github.com/kubevirt/kubevirt/pull/11144) | [0xFelix](https://github.com/0xFelix) |
| New cluster-wide `vmRolloutStrategy` setting to define whether changes to VMs should either be always staged or live-updated when possible.  | [#11054](https://github.com/kubevirt/kubevirt/pull/11054) | [jean-edouard](https://github.com/jean-edouard) |
| Introduce a new API to mark a binding plugin as migratable.<br>With this new API a VM using passt binding plugin allows migration.  | [#11064](https://github.com/kubevirt/kubevirt/pull/11064) | [AlonaKaplan](https://github.com/AlonaKaplan) |
| Update runc dependency to v1.1.12  | [#11122](https://github.com/kubevirt/kubevirt/pull/11122) | [brianmcarey](https://github.com/brianmcarey) |
| Refactor monitoring metrics  | [#10982](https://github.com/kubevirt/kubevirt/pull/10982) | [machadovilaca](https://github.com/machadovilaca) |
| Bug fix: Packet drops during the initial phase of VM live migration https://issues.redhat.com/browse/CNV-28040  | [#11069](https://github.com/kubevirt/kubevirt/pull/11069) | [ormergi](https://github.com/ormergi) |
| Reduced VM rescheduling time on node failure  | [#10961](https://github.com/kubevirt/kubevirt/pull/10961) | [jcanocan](https://github.com/jcanocan) |
| fix(vmclone): Generate VM patches from vmsnapshotcontent, instead of current VM  | [#11065](https://github.com/kubevirt/kubevirt/pull/11065) | [fossedihelm](https://github.com/fossedihelm) |
| [Bugfix] Clone VM with WaitForFirstConsumer binding mode PVC now works.  | [#10888](https://github.com/kubevirt/kubevirt/pull/10888) | [fossedihelm](https://github.com/fossedihelm) |
| Update container base image to use current stable debian 12 base  | [#11068](https://github.com/kubevirt/kubevirt/pull/11068) | [brianmcarey](https://github.com/brianmcarey) |
| Fix potential crash when trying to list USB devices on host without any  | [#11047](https://github.com/kubevirt/kubevirt/pull/11047) | [jschintag](https://github.com/jschintag) |
| Expose fs disk information via GuestOsInfo  | [#10970](https://github.com/kubevirt/kubevirt/pull/10970) | [alromeros](https://github.com/alromeros) |
| restrict default cluster role to authenticated only users  | [#11050](https://github.com/kubevirt/kubevirt/pull/11050) | [fossedihelm](https://github.com/fossedihelm) |
| Allow unprivileged users read-only access to VirtualMachineCluster{Instancetypes,Preferences} by default.  | [#11025](https://github.com/kubevirt/kubevirt/pull/11025) | [0xFelix](https://github.com/0xFelix) |
| Refactor monitoring collectors  | [#10853](https://github.com/kubevirt/kubevirt/pull/10853) | [machadovilaca](https://github.com/machadovilaca) |
| Allow `kubevirt.io:default` clusterRole to get,list kubevirts  | [#11001](https://github.com/kubevirt/kubevirt/pull/11001) | [fossedihelm](https://github.com/fossedihelm) |
| Aggregate DVs conditions on VMI (and so VM)  | [#10905](https://github.com/kubevirt/kubevirt/pull/10905) | [tiraboschi](https://github.com/tiraboschi) |
| Bugfix: Reject volume exports when no output is specified  | [#10963](https://github.com/kubevirt/kubevirt/pull/10963) | [alromeros](https://github.com/alromeros) |
| Update monitoring file structure  | [#10962](https://github.com/kubevirt/kubevirt/pull/10962) | [machadovilaca](https://github.com/machadovilaca) |
| Report IP of interfaces using network binding plugin.  | [#10981](https://github.com/kubevirt/kubevirt/pull/10981) | [AlonaKaplan](https://github.com/AlonaKaplan) |
| Updated common-instancetypes bundles to v0.4.0  | [#10922](https://github.com/kubevirt/kubevirt/pull/10922) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| KubeVirt is now built with go 1.21.5  | [#10914](https://github.com/kubevirt/kubevirt/pull/10914) | [brianmcarey](https://github.com/brianmcarey) |
| Change vm.status.PrintableStatus default value to "Stopped"  | [#10846](https://github.com/kubevirt/kubevirt/pull/10846) | [RamLavi](https://github.com/RamLavi) |
| # Create a manifest for a clone with template label filters:<br>virtctl create clone --source-name sourceVM --template-label-filter '*' --template-label-filter '!some/key'<br># Create a manifest for a clone with template annotation filters:<br>virtctl create clone --source-name sourceVM --template-annotation-filter '*' --template-annotation-filter '!some/key'  | [#10787](https://github.com/kubevirt/kubevirt/pull/10787) | [matthewei](https://github.com/matthewei) |
| VMClone: Emit an event in case restore creation fails  | [#10918](https://github.com/kubevirt/kubevirt/pull/10918) | [orelmisan](https://github.com/orelmisan) |
| Fix the value of VMI `Status.GuestOSInfo.Version`  | [#10916](https://github.com/kubevirt/kubevirt/pull/10916) | [orelmisan](https://github.com/orelmisan) |
| Deprecate macvtap  | [#10924](https://github.com/kubevirt/kubevirt/pull/10924) | [AlonaKaplan](https://github.com/AlonaKaplan) |
| vmi status's guestOsInfo adds `Machine`  | [#10898](https://github.com/kubevirt/kubevirt/pull/10898) | [matthewei](https://github.com/matthewei) |
| Raise an error in case passt feature gate or API are used.  | [#10866](https://github.com/kubevirt/kubevirt/pull/10866) | [AlonaKaplan](https://github.com/AlonaKaplan) |
| Built with golang 1.20.12  | [#10879](https://github.com/kubevirt/kubevirt/pull/10879) | [brianmcarey](https://github.com/brianmcarey) |
| IsolateEmulatorThread: Add cluster-wide parity completion setting  | [#10872](https://github.com/kubevirt/kubevirt/pull/10872) | [RamLavi](https://github.com/RamLavi) |
| Refactor monitoring alerts  | [#10700](https://github.com/kubevirt/kubevirt/pull/10700) | [machadovilaca](https://github.com/machadovilaca) |
| Change second emulator thread assign strategy to best-effort.  | [#10839](https://github.com/kubevirt/kubevirt/pull/10839) | [RamLavi](https://github.com/RamLavi) |
| Remove year from generated code copyright  | [#10863](https://github.com/kubevirt/kubevirt/pull/10863) | [dhiller](https://github.com/dhiller) |
| Fix KubeVirt for CRIO 1.28 by using checksums to verify containerdisks when migrating VMIs  | [#10747](https://github.com/kubevirt/kubevirt/pull/10747) | [acardace](https://github.com/acardace) |
| BugFix: Double cloning with filter fails  | [#10860](https://github.com/kubevirt/kubevirt/pull/10860) | [akalenyu](https://github.com/akalenyu) |
| Attachment pod creation is now rate limited  | [#10567](https://github.com/kubevirt/kubevirt/pull/10567) | [awels](https://github.com/awels) |
| Reject VirtualMachineClone creation when target name is equal to source name  | [#10845](https://github.com/kubevirt/kubevirt/pull/10845) | [orelmisan](https://github.com/orelmisan) |
| Requests/Limits can now be configured when using CPU/Memory hotplug  | [#10840](https://github.com/kubevirt/kubevirt/pull/10840) | [acardace](https://github.com/acardace) |
| Add total VMs created metric  | [#10418](https://github.com/kubevirt/kubevirt/pull/10418) | [machadovilaca](https://github.com/machadovilaca) |
| Support macvtap as a binding plugin  | [#10800](https://github.com/kubevirt/kubevirt/pull/10800) | [AlonaKaplan](https://github.com/AlonaKaplan) |
| Fixes device permission when using USB host passthrough  | [#10753](https://github.com/kubevirt/kubevirt/pull/10753) | [victortoso](https://github.com/victortoso) |
| Windows offline activation with ACPI SLIC table  | [#10774](https://github.com/kubevirt/kubevirt/pull/10774) | [victortoso](https://github.com/victortoso) |
| Support multiple CPUs in Housekeeping cgroup  | [#10783](https://github.com/kubevirt/kubevirt/pull/10783) | [RamLavi](https://github.com/RamLavi) |
| Source virt-launcher: Log migration info by default  | [#10809](https://github.com/kubevirt/kubevirt/pull/10809) | [orelmisan](https://github.com/orelmisan) |
| Add v1alpha3 for hooks<br>Fix migration when using sidecars  | [#10046](https://github.com/kubevirt/kubevirt/pull/10046) | [victortoso](https://github.com/victortoso) |
| Refactor monitoring  recording-rules  | [#10651](https://github.com/kubevirt/kubevirt/pull/10651) | [machadovilaca](https://github.com/machadovilaca) |
| Extend kubvirt CR by adding domain attachment option to the network binding plugin API.  | [#10732](https://github.com/kubevirt/kubevirt/pull/10732) | [AlonaKaplan](https://github.com/AlonaKaplan) |
| Added “adm” subcommand under “virtctl”, and “log-verbosity" subcommand under “adm”. The log-verbosity command is:<br>- To show the log verbosity of one or more components.<br>- To set the log verbosity of one or more components.<br>- To reset the log verbosity of all components (reset to the default verbosity (2)).  | [#10244](https://github.com/kubevirt/kubevirt/pull/10244) | [hshitomi](https://github.com/hshitomi) |
| 1. Support "Clone API" to filter VirtualMachine.spec.template.annotation and VirtualMachine.spec.template.label  | [#10658](https://github.com/kubevirt/kubevirt/pull/10658) | [matthewei](https://github.com/matthewei) |
| Fixes SMT Alignment Error in virt-launcher pod by optimizing isolateEmulatorThread feature (BZ#2228103).  | [#10593](https://github.com/kubevirt/kubevirt/pull/10593) | [RamLavi](https://github.com/RamLavi) |
| Restored hotplug attachment pod request/limit to original value  | [#10720](https://github.com/kubevirt/kubevirt/pull/10720) | [awels](https://github.com/awels) |
| Exposing Filesystem Persistent Volumes (PVs)  to the VM using unprivilege virtiofsd.  | [#10657](https://github.com/kubevirt/kubevirt/pull/10657) | [germag](https://github.com/germag) |
| Functional tests for sidecar hook with ConfigMap  | [#10637](https://github.com/kubevirt/kubevirt/pull/10637) | [dharmit](https://github.com/dharmit) |
| Add PVC option to the hook sidecars for supplying additional debugging tools  | [#10598](https://github.com/kubevirt/kubevirt/pull/10598) | [alicefr](https://github.com/alicefr) |
| - documents steps to build the KubeVirt builder container  | [#10526](https://github.com/kubevirt/kubevirt/pull/10526) | [cfilleke](https://github.com/cfilleke) |
| virt-launcher: fix qemu non root log path  | [#10699](https://github.com/kubevirt/kubevirt/pull/10699) | [qinqon](https://github.com/qinqon) |
| BugFix: cgroupsv2 device allowlist is bound to virt-handler internal state/block disk device overwritten on hotplug  | [#10689](https://github.com/kubevirt/kubevirt/pull/10689) | [akalenyu](https://github.com/akalenyu) |
| Remove MigrateVmiDiskTransferRateMetric  | [#10693](https://github.com/kubevirt/kubevirt/pull/10693) | [machadovilaca](https://github.com/machadovilaca) |
| Remove leftover NonRoot feature gate  | [#10615](https://github.com/kubevirt/kubevirt/pull/10615) | [orelmisan](https://github.com/orelmisan) |
| Allow LUN disks to be hotplugged  | [#10529](https://github.com/kubevirt/kubevirt/pull/10529) | [alromeros](https://github.com/alromeros) |
| Remove leftover NonRootExperimental feature gate  | [#10582](https://github.com/kubevirt/kubevirt/pull/10582) | [orelmisan](https://github.com/orelmisan) |
| Disable HTTP/2 to mitigate CVE-2023-44487  | [#10596](https://github.com/kubevirt/kubevirt/pull/10596) | [mhenriks](https://github.com/mhenriks) |
| Fix LowKVMNodesCount not firing  | [#10570](https://github.com/kubevirt/kubevirt/pull/10570) | [machadovilaca](https://github.com/machadovilaca) |
| vmi memory footprint increase by 35M when guest serial console logging is turned on (default on).  | [#10571](https://github.com/kubevirt/kubevirt/pull/10571) | [tiraboschi](https://github.com/tiraboschi) |
| Introduce network binding plugin for Passt networking, interfacing with Kubevirt new network binding plugin API.  | [#10425](https://github.com/kubevirt/kubevirt/pull/10425) | [ormergi](https://github.com/ormergi) |
| Ability to run scripts through hook sidecar  | [#10479](https://github.com/kubevirt/kubevirt/pull/10479) | [dharmit](https://github.com/dharmit) |


_This page is updated daily._
