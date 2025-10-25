---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Add v1.6.0 perf and scale benchmarks data  | [#15878](https://github.com/kubevirt/kubevirt/pull/15878) | [Sreeja1725](https://github.com/Sreeja1725) |
| Updated common-instancetypes bundles to v1.5.1  | [#15936](https://github.com/kubevirt/kubevirt/pull/15936) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| Fix possible nil pointer caused by migration during kv upgrade  | [#15008](https://github.com/kubevirt/kubevirt/pull/15008) | [fossedihelm](https://github.com/fossedihelm) |
| Experimental support for AMD SEV-SNP that is behind the `WorkloadEncruptionSEV` feature gate.  | [#14845](https://github.com/kubevirt/kubevirt/pull/14845) | [alancaldelas](https://github.com/alancaldelas) |
| Specify correct label selection when creating a service via virtctl expose. The expose command on virtctl v1.7 and above will not work with older KubeVirt versions.  | [#15783](https://github.com/kubevirt/kubevirt/pull/15783) | [orelmisan](https://github.com/orelmisan) |
| Beta: PanicDevices  | [#15830](https://github.com/kubevirt/kubevirt/pull/15830) | [varunrsekar](https://github.com/varunrsekar) |
| It is now possible to configure discard_granularity for VM disks.  | [#15698](https://github.com/kubevirt/kubevirt/pull/15698) | [Acedus](https://github.com/Acedus) |
| Bug fix: Thousands of migrations should not cause failures of active migrations  | [#15867](https://github.com/kubevirt/kubevirt/pull/15867) | [xpivarc](https://github.com/xpivarc) |
| NONR  | [#15552](https://github.com/kubevirt/kubevirt/pull/15552) | [nirdothan](https://github.com/nirdothan) |
| The `DefaultVirtHandler{QPS,Burst}` values are increased to ensure no bottleneck forms within `virt-handler`  | [#15712](https://github.com/kubevirt/kubevirt/pull/15712) | [lyarwood](https://github.com/lyarwood) |
| Fix RestartRequired handling for hotplug volumes  | [#15788](https://github.com/kubevirt/kubevirt/pull/15788) | [mhenriks](https://github.com/mhenriks) |
| Add VirtualMachineInstanceEvictionRequested condition for eviction tracking  | [#15539](https://github.com/kubevirt/kubevirt/pull/15539) | [tiraboschi](https://github.com/tiraboschi) |
| The list of annotations and labels synced from VM.spec.template.metadata to VMI and then to virt-launcher pods can be extended  | [#14902](https://github.com/kubevirt/kubevirt/pull/14902) | [tiraboschi](https://github.com/tiraboschi) |
| Build KubeVirt with go v1.24.7  | [#15784](https://github.com/kubevirt/kubevirt/pull/15784) | [brianmcarey](https://github.com/brianmcarey) |
| fix: prioritize expired cert removal over 50-cert limit in MergeCABundle  | [#15706](https://github.com/kubevirt/kubevirt/pull/15706) | [ksimon1](https://github.com/ksimon1) |
| Support for the `ioThreads` VMI configurable is added to the `instancetype.kubevirt.io/v1beta1` API allowing `supplementalPoolThreadCount` to now be provided by an instance type.  | [#15798](https://github.com/kubevirt/kubevirt/pull/15798) | [lyarwood](https://github.com/lyarwood) |
| Object Graph: Include NADs and ServiceAccounts  | [#15615](https://github.com/kubevirt/kubevirt/pull/15615) | [alromeros](https://github.com/alromeros) |
| Preferences can now express preferred and required architecture values for use within VirtualMachines  | [#15398](https://github.com/kubevirt/kubevirt/pull/15398) | [lyarwood](https://github.com/lyarwood) |
| Bug fix, virt-launcher is properly reaped  | [#15676](https://github.com/kubevirt/kubevirt/pull/15676) | [xpivarc](https://github.com/xpivarc) |
| Replicas of `virt-api` are now scaled depending on the number of nodes within the environment with the `kubevirt.io/schedulable=true` label.  | [#15690](https://github.com/kubevirt/kubevirt/pull/15690) | [lyarwood](https://github.com/lyarwood) |
| BugFix: Restoring naked PVCs from a VMSnapshot are now properly owned by the VM if the restore policy is set to VM  | [#15692](https://github.com/kubevirt/kubevirt/pull/15692) | [awels](https://github.com/awels) |
| Only a single `Signaled Graceful Shutdown` event is now sent to avoid spamming the event recorder during long graceful shutdown attempts  | [#15759](https://github.com/kubevirt/kubevirt/pull/15759) | [lyarwood](https://github.com/lyarwood) |
| The deprecated `instancetype.kubevirt.io/v1alpha{1,2}` API and CRDs have been removed  | [#15400](https://github.com/kubevirt/kubevirt/pull/15400) | [lyarwood](https://github.com/lyarwood) |
| Memory overcommit is now recalculated on migration.<br>Important: deployments that set a memoryOvercommit value below 10 need to bump to 10+ before upgrading.  | [#15681](https://github.com/kubevirt/kubevirt/pull/15681) | [jean-edouard](https://github.com/jean-edouard) |
| build: update to bazel v6.5.0 and rules_oci  | [#13111](https://github.com/kubevirt/kubevirt/pull/13111) | [brianmcarey](https://github.com/brianmcarey) |
| Add VMpool finalizer to ensure proper cleanup  | [#15406](https://github.com/kubevirt/kubevirt/pull/15406) | [Sreeja1725](https://github.com/Sreeja1725) |
| Normalise iface status to ensure test stability of hotplug and hotunplug tests  | [#15669](https://github.com/kubevirt/kubevirt/pull/15669) | [HarshithaMS005](https://github.com/HarshithaMS005) |
| ChangedBlockTracking: enable add/remove of qcow2 overlay if vm matches label selector  | [#14772](https://github.com/kubevirt/kubevirt/pull/14772) | [ShellyKa13](https://github.com/ShellyKa13) |
| Support Istio versions 1.25 and above.  | [#15661](https://github.com/kubevirt/kubevirt/pull/15661) | [nirdothan](https://github.com/nirdothan) |
| bump prometheus operator to 0.80.1  | [#15531](https://github.com/kubevirt/kubevirt/pull/15531) | [Yu-Jack](https://github.com/Yu-Jack) |
| BugFix: Able to cancel in flight decentralized live migrations properly  | [#15605](https://github.com/kubevirt/kubevirt/pull/15605) | [awels](https://github.com/awels) |
| Does Screenshot without the usage of VNC  | [#15238](https://github.com/kubevirt/kubevirt/pull/15238) | [victortoso](https://github.com/victortoso) |
| Update metric kubevirt_vm_container_free_memory_bytes_based_on_rss and kubevirt_vm_container_free_memory_bytes_based_on_working_set_bytes names to kubevirt_vm_container_memory_request_margin_based_on_rss_bytes and kubevirt_vm_container_memory_request_margin_based_on_working_set_bytes so they will be clearer  | [#15504](https://github.com/kubevirt/kubevirt/pull/15504) | [sradco](https://github.com/sradco) |
| Enhance VMPool unit tests to make use of fake client  | [#15503](https://github.com/kubevirt/kubevirt/pull/15503) | [Sreeja1725](https://github.com/Sreeja1725) |
| The `DefaultVirtWebhookClient{QPS,Burst}` values are aligned with `DefaultVirtWebhookClient{QPS,Burst}` to help avoid saturating the webhook client with requests it is unable to serve during mass eviction events  | [#15422](https://github.com/kubevirt/kubevirt/pull/15422) | [lyarwood](https://github.com/lyarwood) |
| Add WithUploadSource builder to libdv  | [#15651](https://github.com/kubevirt/kubevirt/pull/15651) | [dcarrier](https://github.com/dcarrier) |
| BugFix: Windows VM with vTPM that was previously Storage Migrated cannot live migrate  | [#15642](https://github.com/kubevirt/kubevirt/pull/15642) | [akalenyu](https://github.com/akalenyu) |
| Add kubevirt_vm_labels metric which shows vm labels converted to Prometheus labels, and can be configured using config map with ignore and allow lists.  | [#15181](https://github.com/kubevirt/kubevirt/pull/15181) | [avlitman](https://github.com/avlitman) |
| Allow decentralized live migration on L3 networks  | [#15630](https://github.com/kubevirt/kubevirt/pull/15630) | [awels](https://github.com/awels) |
| Fixed priority escalation bug in migration controller  | [#15513](https://github.com/kubevirt/kubevirt/pull/15513) | [jean-edouard](https://github.com/jean-edouard) |
| BugFix: Fix volume migration for VMs with long name  | [#15603](https://github.com/kubevirt/kubevirt/pull/15603) | [akalenyu](https://github.com/akalenyu) |
| Added VolumeOwnershipPolicy to decide how volumes are owned once they are restored.  | [#15344](https://github.com/kubevirt/kubevirt/pull/15344) | [SkalaNetworks](https://github.com/SkalaNetworks) |
| remove ppc64le architecture configuration support  | [#14976](https://github.com/kubevirt/kubevirt/pull/14976) | [dasionov](https://github.com/dasionov) |
| Bugfix: Exclude lost+found from export server  | [#15509](https://github.com/kubevirt/kubevirt/pull/15509) | [alromeros](https://github.com/alromeros) |
| Fix: grpc client in handler rest requests are properly closed  | [#15557](https://github.com/kubevirt/kubevirt/pull/15557) | [fossedihelm](https://github.com/fossedihelm) |
| New VM alerts - VirtualMachineStuckInUnhealthyState, VirtualMachineStuckOnNode  | [#15227](https://github.com/kubevirt/kubevirt/pull/15227) | [sradco](https://github.com/sradco) |
| virtctl: The --local-ssh flag and native ssh and scp clients are removed from virtctl. From now on the local ssh and scp clients on a host are always wrapped by virtctl ssh and scp.  | [#15478](https://github.com/kubevirt/kubevirt/pull/15478) | [0xFelix](https://github.com/0xFelix) |
| Fix incorrect metric name kubevirt_vmi_migration_disk_transfer_rate_bytes to kubevirt_vmi_migration_memory_transfer_rate_bytes  | [#13500](https://github.com/kubevirt/kubevirt/pull/13500) | [brandboat](https://github.com/brandboat) |
| Added virt-launcher to kubevirt_memory_delta_from_requested_bytes metric and cnv_abnormal metrics.  | [#15464](https://github.com/kubevirt/kubevirt/pull/15464) | [avlitman](https://github.com/avlitman) |
| Add `preserve session` option to VNC endpoint  | [#15267](https://github.com/kubevirt/kubevirt/pull/15267) | [victortoso](https://github.com/victortoso) |
| ensure default Firmware.Serial value on newly created vms  | [#15357](https://github.com/kubevirt/kubevirt/pull/15357) | [dasionov](https://github.com/dasionov) |
| BugFix: Unable to delete source VM on failed decentralized live migration  | [#15470](https://github.com/kubevirt/kubevirt/pull/15470) | [awels](https://github.com/awels) |
| Derive eviction-in-progress annotation from VMI eviction status  | [#15423](https://github.com/kubevirt/kubevirt/pull/15423) | [tiraboschi](https://github.com/tiraboschi) |
| virtctl (portfoward|ssh|scp): Drop support for legacy dot syntax. In case the old dot syntax was used virtctl could ask for verification of the host key again. In some cases the known_hosts file might need to be updated manually.  | [#15475](https://github.com/kubevirt/kubevirt/pull/15475) | [0xFelix](https://github.com/0xFelix) |
| bugfix: ensure grace period metadata cache is synced in virt-launcher  | [#15170](https://github.com/kubevirt/kubevirt/pull/15170) | [dasionov](https://github.com/dasionov) |
| bugfix: prevent VMSnapshotContent repeated update with the same error message  | [#15397](https://github.com/kubevirt/kubevirt/pull/15397) | [ShellyKa13](https://github.com/ShellyKa13) |
| Add Command line flag to disable Node Labeller service  | [#15167](https://github.com/kubevirt/kubevirt/pull/15167) | [Sreeja1725](https://github.com/Sreeja1725) |
| Aligning descheduler opt-out annotation name  | [#15365](https://github.com/kubevirt/kubevirt/pull/15365) | [tiraboschi](https://github.com/tiraboschi) |
| This PR adds the following alerts: GuestPeakVCPUQueueHighWarning, GuestPeakVCPUQueueHighCritical  | [#14983](https://github.com/kubevirt/kubevirt/pull/14983) | [sradco](https://github.com/sradco) |
| The `foregroundDeleteVirtualMachine` has been deprecated and replaced with the domain-qualified `kubevirt.io/foregroundDeleteVirtualMachine`.  | [#15096](https://github.com/kubevirt/kubevirt/pull/15096) | [lyarwood](https://github.com/lyarwood) |
| bugfix: Enable vmsnapshot for paused VMs  | [#15001](https://github.com/kubevirt/kubevirt/pull/15001) | [noamasu](https://github.com/noamasu) |
| bugfix: volume hotplug pod is no longer evicted when associated VM can live migrate.  | [#15093](https://github.com/kubevirt/kubevirt/pull/15093) | [Acedus](https://github.com/Acedus) |
| Add GuestAgentInfo info metrics  | [#14879](https://github.com/kubevirt/kubevirt/pull/14879) | [machadovilaca](https://github.com/machadovilaca) |
| bugfix: snapshot and restore now works correctly for VMs after a storage volume migration  | [#15305](https://github.com/kubevirt/kubevirt/pull/15305) | [Acedus](https://github.com/Acedus) |
| Common Names are now enforce for aggregated API  | [#15314](https://github.com/kubevirt/kubevirt/pull/15314) | [xpivarc](https://github.com/xpivarc) |
| Bumped the bundled common-instancetypes to v1.4.0 which add new preferences.  | [#15253](https://github.com/kubevirt/kubevirt/pull/15253) | [0xFelix](https://github.com/0xFelix) |
| BugFix: export fails when VMExport has dots in secret  | [#15182](https://github.com/kubevirt/kubevirt/pull/15182) | [akalenyu](https://github.com/akalenyu) |
| Support for all `*_SHASUM` environment variables has been removed from the `virt-operator` component. Users should instead use the remaining `*_IMAGE` environment variables to request a specific image version using a tag, digest or both.  | [#15061](https://github.com/kubevirt/kubevirt/pull/15061) | [lyarwood](https://github.com/lyarwood) |
| virt-operator won't schedule on worker nodes  | [#15157](https://github.com/kubevirt/kubevirt/pull/15157) | [jean-edouard](https://github.com/jean-edouard) |
| Drop an arbitrary limitation on VM's domain.firmaware.serial. Any string is passed verbatim to smbios. Illegal may be tweaked or ignored based on qemu/smbios version.  | [#15118](https://github.com/kubevirt/kubevirt/pull/15118) | [dankenigsberg](https://github.com/dankenigsberg) |
| Update dependecy golang.org/x/oauth2 to v0.27.0  | [#15098](https://github.com/kubevirt/kubevirt/pull/15098) | [dominikholler](https://github.com/dominikholler) |
| Fix postcopy multifd compatibility during upgrade  | [#15016](https://github.com/kubevirt/kubevirt/pull/15016) | [fossedihelm](https://github.com/fossedihelm) |
| Update dependecy golang.org/x/net to v0.38.0  | [#15100](https://github.com/kubevirt/kubevirt/pull/15100) | [dominikholler](https://github.com/dominikholler) |
| BugFix: export fails when VMExport has dots in name  | [#15099](https://github.com/kubevirt/kubevirt/pull/15099) | [akalenyu](https://github.com/akalenyu) |
| allows virtual machine instances with an instance type to specify memory fields that do not conflict with the instance type  | [#14685](https://github.com/kubevirt/kubevirt/pull/14685) | [seanbanko](https://github.com/seanbanko) |
| Cleanup: libvmi: add consistently named cpu/mem setters  | [#14888](https://github.com/kubevirt/kubevirt/pull/14888) | [akalenyu](https://github.com/akalenyu) |
| Bugfix: Label upload PVCs to support CDI WebhookPvcRendering  | [#15067](https://github.com/kubevirt/kubevirt/pull/15067) | [alromeros](https://github.com/alromeros) |
| HostDisk: KubeVirt no longer performs chown/chmod to compensate for storage that doesn't support fsGroup  | [#15037](https://github.com/kubevirt/kubevirt/pull/15037) | [jean-edouard](https://github.com/jean-edouard) |
| Added support for architecture-specific configuration of `s390x` (IBM Z) in KubeVirt cluster config.  | [#15017](https://github.com/kubevirt/kubevirt/pull/15017) | [nekkunti](https://github.com/nekkunti) |
| The synchronization controller migration network IP address is advertised by the KubeVirt CR  | [#15022](https://github.com/kubevirt/kubevirt/pull/15022) | [awels](https://github.com/awels) |
| Decentralized migration resource now shows the synchronization address  | [#15021](https://github.com/kubevirt/kubevirt/pull/15021) | [awels](https://github.com/awels) |
| Add support for DRA devices such as GPUs and HostDevices.  | [#14365](https://github.com/kubevirt/kubevirt/pull/14365) | [alaypatel07](https://github.com/alaypatel07) |
| Decentralized live migration is available to allow migration across namespaces and clusters  | [#14882](https://github.com/kubevirt/kubevirt/pull/14882) | [awels](https://github.com/awels) |
| Beta: NodeRestriction  | [#14964](https://github.com/kubevirt/kubevirt/pull/14964) | [xpivarc](https://github.com/xpivarc) |
| Possible to trust additional CAs for verifying kubevirt infra structure components  | [#14986](https://github.com/kubevirt/kubevirt/pull/14986) | [awels](https://github.com/awels) |
| Support seamless TCP migration with passt (alpha)  | [#14875](https://github.com/kubevirt/kubevirt/pull/14875) | [nirdothan](https://github.com/nirdothan) |


_This page is updated daily._
