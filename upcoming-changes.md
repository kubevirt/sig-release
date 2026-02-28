---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Use defined deployment number of replicas as base to fire low count alerts  | [#16806](https://github.com/kubevirt/kubevirt/pull/16806) | [machadovilaca](https://github.com/machadovilaca) |
| Graduate ExpandDisk Feature Gate  | [#16604](https://github.com/kubevirt/kubevirt/pull/16604) | [Dsanatar](https://github.com/Dsanatar) |
| Added support for attestation on the Intel TDX Confidential Computing Platform  | [#15958](https://github.com/kubevirt/kubevirt/pull/15958) | [Aseeef](https://github.com/Aseeef) |
| Handle migration during backup according to migration priority  | [#16877](https://github.com/kubevirt/kubevirt/pull/16877) | [ShellyKa13](https://github.com/ShellyKa13) |
| BugFix: VMs requiring enlightenment are now able to be live migrated after a decentralized live migration  | [#16871](https://github.com/kubevirt/kubevirt/pull/16871) | [awels](https://github.com/awels) |
| Subtract non-schedulable nodes from kubevirt_allocatable_nodes  | [#16865](https://github.com/kubevirt/kubevirt/pull/16865) | [machadovilaca](https://github.com/machadovilaca) |
| fix VMExport failure with long PVC names  | [#16846](https://github.com/kubevirt/kubevirt/pull/16846) | [Aneesh-Hegde](https://github.com/Aneesh-Hegde) |
| fix: Prevent stale VMI backup status update when reusing backup names  | [#16399](https://github.com/kubevirt/kubevirt/pull/16399) | [Acedus](https://github.com/Acedus) |
| Allow disabling Velero hooks in virt-launcher via Annotation  | [#16786](https://github.com/kubevirt/kubevirt/pull/16786) | [alromeros](https://github.com/alromeros) |
| feat: virtctl gained new virt-template / VirtualMachineTemplate related commands (process, create and convert)  | [#16687](https://github.com/kubevirt/kubevirt/pull/16687) | [0xFelix](https://github.com/0xFelix) |
| VEP 165: Containerpath Volumes  | [#16662](https://github.com/kubevirt/kubevirt/pull/16662) | [mhenriks](https://github.com/mhenriks) |
| Remove network-attachment-definition get permissions from virt-controller ClusterRole conditioned by a feature gate.  | [#16821](https://github.com/kubevirt/kubevirt/pull/16821) | [nirdothan](https://github.com/nirdothan) |
| Bug-fix: Correctly detect CDI and Prometheus crds, preventing to misinterpret with different objects.  | [#16643](https://github.com/kubevirt/kubevirt/pull/16643) | [kwonkwonn](https://github.com/kwonkwonn) |
| Fix: live-migration with CBT no longer fails on virtual disk size evaluation errors.  | [#16528](https://github.com/kubevirt/kubevirt/pull/16528) | [Acedus](https://github.com/Acedus) |
| Handle CBT backup abort requests and failures  | [#16426](https://github.com/kubevirt/kubevirt/pull/16426) | [Acedus](https://github.com/Acedus) |
| Add initial CentOS Stream 10 build support with KUBEVIRT_CENTOS_STREAM_VERSION environment variable, these builds will be untested until v1.9.0 and beyond  | [#16582](https://github.com/kubevirt/kubevirt/pull/16582) | [lyarwood](https://github.com/lyarwood) |
| BugFix: storage migration fails with Google Cloud NetApp Volumes  | [#16833](https://github.com/kubevirt/kubevirt/pull/16833) | [akalenyu](https://github.com/akalenyu) |
| Support seamless migration with core passt binding (beta).  | [#16820](https://github.com/kubevirt/kubevirt/pull/16820) | [nirdothan](https://github.com/nirdothan) |
| Support for the deployment of virt-template through virt-operator was added (VEP76)  | [#16655](https://github.com/kubevirt/kubevirt/pull/16655) | [0xFelix](https://github.com/0xFelix) |
| Expose guest panic as a Kubernetes event  | [#16666](https://github.com/kubevirt/kubevirt/pull/16666) | [iholder101](https://github.com/iholder101) |
| Bug fix: VIRT_*_IMAGE environment variable overrides on the virt-operator deployment are now correctly propagated to component deployments (virt-controller, virt-handler, etc.). Previously, changing these env vars had no effect due to the image values being excluded from the install strategy deployment ID hash.  | [#16791](https://github.com/kubevirt/kubevirt/pull/16791) | [lyarwood](https://github.com/lyarwood) |
| `PrefixTargetName` is now allowed as a `VolumeNamePolicy` for `VirtualMachineClone`  | [#16802](https://github.com/kubevirt/kubevirt/pull/16802) | [lyarwood](https://github.com/lyarwood) |
| fix: domain job completion events would not be processed if the domain was paused due to an I/O error.  | [#16778](https://github.com/kubevirt/kubevirt/pull/16778) | [Acedus](https://github.com/Acedus) |
| A VMI.spec.domain.rebootPolicy field can be used to control the method the domain uses to handle reboots originating from inside the VM. Either the hypervisor processes the reboot silently behind the scenes (default) or the user can opt-in to a more visible behavior, where the hypervisor terminates the domain and lets kubevirt to handle the restart according to the runStrategy rules.  | [#16579](https://github.com/kubevirt/kubevirt/pull/16579) | [MarSik](https://github.com/MarSik) |
| Fix LowReadyVirtOperatorsCount use running instead of up and changing kubevirt_virt_operator_ready to use sum and * instead of count and +  | [#16466](https://github.com/kubevirt/kubevirt/pull/16466) | [Ronilerr](https://github.com/Ronilerr) |
| An admin can disable the NAD query logic and use network-resources-injector instead to have less API calls  | [#16734](https://github.com/kubevirt/kubevirt/pull/16734) | [orelmisan](https://github.com/orelmisan) |
| Replaced QuiesceFailed with QuiesceTimeout indication and added 60s Velero pre-backup hook timeout to better handle Windows VSS limitations.  | [#16653](https://github.com/kubevirt/kubevirt/pull/16653) | [noamasu](https://github.com/noamasu) |
| Existing VMs that retain the legacy ordinal naming scheme for secondary interfaces are automatically upgraded without a reboot.  | [#16642](https://github.com/kubevirt/kubevirt/pull/16642) | [orelmisan](https://github.com/orelmisan) |
| Incremental backups supported after VM restart by redefining checkpoints metadata in libvirt  | [#16448](https://github.com/kubevirt/kubevirt/pull/16448) | [ShellyKa13](https://github.com/ShellyKa13) |
| BugFix: vmsnapshot: report volumes being deleted  | [#16621](https://github.com/kubevirt/kubevirt/pull/16621) | [akalenyu](https://github.com/akalenyu) |
| Fix grammar mistakes  | [#16645](https://github.com/kubevirt/kubevirt/pull/16645) | [Ronilerr](https://github.com/Ronilerr) |
| Feature gates can now become explicitly disabled using `kv.spec.configuration.developerConfiguration.disabledFeatureGates`.  | [#16370](https://github.com/kubevirt/kubevirt/pull/16370) | [iholder101](https://github.com/iholder101) |
| Let libvirt lookup the actual disk size if block device to ensure compatibility with encrypted disks.  | [#16366](https://github.com/kubevirt/kubevirt/pull/16366) | [elliot-gustafsson](https://github.com/elliot-gustafsson) |
| Bugfix: Label memorydump-created PVCs to support CDI WebhookPvcRendering  | [#16229](https://github.com/kubevirt/kubevirt/pull/16229) | [noamasu](https://github.com/noamasu) |
| BugFix: Decentralized live migration between volumes with different volumeModes now successfully completes  | [#16637](https://github.com/kubevirt/kubevirt/pull/16637) | [awels](https://github.com/awels) |
| Updated common-instancetypes bundles to v1.6.0  | [#16705](https://github.com/kubevirt/kubevirt/pull/16705) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| Decentralized Live Migration now has a separate condition in VMI and VMIM to indicate any issues  | [#16512](https://github.com/kubevirt/kubevirt/pull/16512) | [awels](https://github.com/awels) |
| Add new `PrefixTargetName` VolumeRestorePolicy for VirtualMachineRestore that creates restored volume names using the format `{targetVMName}-{volumeName}`. This provides predictable, readable names while avoiding collisions when restoring snapshots to different target VMs.  | [#16489](https://github.com/kubevirt/kubevirt/pull/16489) | [lyarwood](https://github.com/lyarwood) |
| Add missing "Direct" and "Extended" options to Hyperv TLBFlush  | [#16404](https://github.com/kubevirt/kubevirt/pull/16404) | [iholder101](https://github.com/iholder101) |
| virt-operator now configures client rate limiting (default: 200 QPS / 400 burst) to improve reconciliation performance when processing large numbers of objects. Rate limits can be customized via --client-qps and --client-burst flags or VIRT_OPERATOR_CLIENT_QPS and VIRT_OPERATOR_CLIENT_BURST environment variables.  | [#16491](https://github.com/kubevirt/kubevirt/pull/16491) | [lyarwood](https://github.com/lyarwood) |
| Fix block volume hotplug breaking autoattachVSOCK  | [#16600](https://github.com/kubevirt/kubevirt/pull/16600) | [woojoong88](https://github.com/woojoong88) |
| Network downward API network-info includes mac addresses  | [#15898](https://github.com/kubevirt/kubevirt/pull/15898) | [bgartzi](https://github.com/bgartzi) |
| The MigrationPriorityQueue feature gate has been promoted from Alpha to Beta.  | [#16558](https://github.com/kubevirt/kubevirt/pull/16558) | [fossedihelm](https://github.com/fossedihelm) |
| Preserve VM Specific fields during update  | [#16585](https://github.com/kubevirt/kubevirt/pull/16585) | [Sreeja1725](https://github.com/Sreeja1725) |
| Introduce `HypervisorConfigurations` field in the `KubevirtConfiguration` CRD.  | [#16326](https://github.com/kubevirt/kubevirt/pull/16326) | [harshitgupta1337](https://github.com/harshitgupta1337) |
| Fixed missing object context in client-go log output after changing verbosity.  | [#16527](https://github.com/kubevirt/kubevirt/pull/16527) | [lukashes](https://github.com/lukashes) |
| Apply CBT to a hotplug volume  | [#16510](https://github.com/kubevirt/kubevirt/pull/16510) | [ShellyKa13](https://github.com/ShellyKa13) |
| Add target-side premigration hook system  | [#16212](https://github.com/kubevirt/kubevirt/pull/16212) | [Barakmor1](https://github.com/Barakmor1) |
| Refactor doc-generator  | [#16511](https://github.com/kubevirt/kubevirt/pull/16511) | [Ronilerr](https://github.com/Ronilerr) |
| Fix ResourceVersion conflicts in VM reconciliation when instancetype controller modifies Status. The instancetype controller now properly propagates ResourceVersion from PatchStatus responses, preventing conflicts in subsequent UpdateStatus calls.  | [#16498](https://github.com/kubevirt/kubevirt/pull/16498) | [lyarwood](https://github.com/lyarwood) |
| The `DisableMDEVConfiguration` feature gate is now deprecated ahead of removal in a future release in favour of a new `kubevirt.spec.configuration.mediatedDevicesConfiguration.enabled` configurable  | [#16220](https://github.com/kubevirt/kubevirt/pull/16220) | [lyarwood](https://github.com/lyarwood) |
| VirtualMachineClone API now includes `VolumeNamePolicy` field to control volume cloning behavior.<br>Currently supports `RandomizeNames` policy which creates new volumes with randomized names (default behavior).<br>This provides an abstraction layer for future flexibility in cloning implementations.  | [#16488](https://github.com/kubevirt/kubevirt/pull/16488) | [lyarwood](https://github.com/lyarwood) |
| Add tolerations for unschedulable taints to hot-plug pods  | [#14661](https://github.com/kubevirt/kubevirt/pull/14661) | [oujonny](https://github.com/oujonny) |
| Label memory-dump PVCs to support CDI WebhookPvcRendering  | [#15113](https://github.com/kubevirt/kubevirt/pull/15113) | [alromeros](https://github.com/alromeros) |
| BugFix: migration metrics missing  | [#16463](https://github.com/kubevirt/kubevirt/pull/16463) | [akalenyu](https://github.com/akalenyu) |
| Scale up KWOK performance test and add virt-controller queue metrics  | [#16024](https://github.com/kubevirt/kubevirt/pull/16024) | [Sreeja1725](https://github.com/Sreeja1725) |
| Macvtap core binding has been removed.  | [#16453](https://github.com/kubevirt/kubevirt/pull/16453) | [nirdothan](https://github.com/nirdothan) |
| The discontinued core SLIRP binding has been completely removed.  | [#16456](https://github.com/kubevirt/kubevirt/pull/16456) | [orelmisan](https://github.com/orelmisan) |
| Prevent false restart-required conditions when the VM and corresponding VMI already share the same firmware UUID.  | [#16329](https://github.com/kubevirt/kubevirt/pull/16329) | [dasionov](https://github.com/dasionov) |
| fix: DataVolumeTemplates with a sourceRef of a DataSource that points to another DataSource now correctly resolves the backing source.  | [#16429](https://github.com/kubevirt/kubevirt/pull/16429) | [Acedus](https://github.com/Acedus) |
| kubevirt_vmi_migration_data_total_bytes is deprecated in favor of kubevirt_vmi_migration_data_bytes_total, in order to comply with the metrics naming conventions.  | [#15975](https://github.com/kubevirt/kubevirt/pull/15975) | [sradco](https://github.com/sradco) |
|   | [#15278](https://github.com/kubevirt/kubevirt/pull/15278) | [sradco](https://github.com/sradco) |
| New VirtLauncherPodsStuckFailed alert  | [#16342](https://github.com/kubevirt/kubevirt/pull/16342) | [sradco](https://github.com/sradco) |
| The KubeVirtVMGuestMemoryPressure<br>and KubeVirtVMGuestMemoryAvailableLow alerts were added to alert when a VM memory is in pressure.  | [#15237](https://github.com/kubevirt/kubevirt/pull/15237) | [sradco](https://github.com/sradco) |
| Fix bug in GuestFilesystemAlmostOutOfSpace, that fired for non relevant file system types.  | [#16351](https://github.com/kubevirt/kubevirt/pull/16351) | [sradco](https://github.com/sradco) |
| Limits the number of guest only interfaces reported on the VMI status to 10. This does not affect the interfaces specified on the spec.  | [#16391](https://github.com/kubevirt/kubevirt/pull/16391) | [frenzyfriday](https://github.com/frenzyfriday) |
| Maintenance: fix release branches potentially failing over identical remote images existing on nodes  | [#16336](https://github.com/kubevirt/kubevirt/pull/16336) | [akalenyu](https://github.com/akalenyu) |
| deprecate --persist flag from virtctl add/remove volume  | [#16280](https://github.com/kubevirt/kubevirt/pull/16280) | [Dsanatar](https://github.com/Dsanatar) |
| Add support for incremental VM backups  | [#16285](https://github.com/kubevirt/kubevirt/pull/16285) | [ShellyKa13](https://github.com/ShellyKa13) |
| Add Ephemeral Hotplug Volume Metric and Alert  | [#15815](https://github.com/kubevirt/kubevirt/pull/15815) | [Dsanatar](https://github.com/Dsanatar) |
| Maintenance: windows lane: W/A wrong nfs image SEEK_DATA impl  | [#16354](https://github.com/kubevirt/kubevirt/pull/16354) | [akalenyu](https://github.com/akalenyu) |
| * Fixed a bug in socket devices that resulted in clusters making use of the Persistent Reservations feature not properly updating their current health.<br>* Fixed a bug in mediated devices that was causing health checks for those devices to not update.  | [#15992](https://github.com/kubevirt/kubevirt/pull/15992) | [Aseeef](https://github.com/Aseeef) |
| Improve boolean flag formatting to parse it correctly.  | [#16355](https://github.com/kubevirt/kubevirt/pull/16355) | [Sreeja1725](https://github.com/Sreeja1725) |
| BugFix: Don't modify VMI CBT status when feature gate is disabled  | [#16343](https://github.com/kubevirt/kubevirt/pull/16343) | [ShellyKa13](https://github.com/ShellyKa13) |
| fix: ensure VMI CBT state remains disabled when the VM has no CBT matcher.  | [#16333](https://github.com/kubevirt/kubevirt/pull/16333) | [Acedus](https://github.com/Acedus) |
| Update dependecy golang.org/x/crypto to v0.45.0  | [#16174](https://github.com/kubevirt/kubevirt/pull/16174) | [dominikholler](https://github.com/dominikholler) |
| Omit LLA from the status report when using masquerade binding.  | [#16242](https://github.com/kubevirt/kubevirt/pull/16242) | [orelmisan](https://github.com/orelmisan) |
| VMBackup: introduce new VM backup API  | [#16081](https://github.com/kubevirt/kubevirt/pull/16081) | [ShellyKa13](https://github.com/ShellyKa13) |
| Update dependecy github.com/opencontainers/selinux to v1.13.0  | [#16173](https://github.com/kubevirt/kubevirt/pull/16173) | [dominikholler](https://github.com/dominikholler) |
| bugfix: prevent cross-vendor migrations  | [#16060](https://github.com/kubevirt/kubevirt/pull/16060) | [dasionov](https://github.com/dasionov) |
| Add event logging for pause and unpause VM operations to align with other VM lifecycle events such as reset  | [#15821](https://github.com/kubevirt/kubevirt/pull/15821) | [SamAlber](https://github.com/SamAlber) |
| VirtualMachinePool now correctly appends index to CloudInit secret references when appendIndexToSecretRefs: true is set, enabling unique cloud-init configurations for each VM in the pool.  | [#15868](https://github.com/kubevirt/kubevirt/pull/15868) | [frank-gen](https://github.com/frank-gen) |
| The `EnableVirtioFsConfigVolumes` feature has graduated to GA and no longer requires the associated feature gate to be enabled.  | [#15913](https://github.com/kubevirt/kubevirt/pull/15913) | [germag](https://github.com/germag) |
| Test Fix: make Alpine ISO mount checks architecture-agnostic  | [#15863](https://github.com/kubevirt/kubevirt/pull/15863) | [HarshithaMS005](https://github.com/HarshithaMS005) |
| Document allowed values for `spec.runStrategy`.  | [#16122](https://github.com/kubevirt/kubevirt/pull/16122) | [dasionov](https://github.com/dasionov) |
| Don't use attachment pods marked for deletion for hotplug volume status updates.  | [#16159](https://github.com/kubevirt/kubevirt/pull/16159) | [Dsanatar](https://github.com/Dsanatar) |
| Allow VMExport with PVCs from Completed Pods  | [#15442](https://github.com/kubevirt/kubevirt/pull/15442) | [Dsanatar](https://github.com/Dsanatar) |
| Migration is using dedicated certificate for mTLS.  | [#15949](https://github.com/kubevirt/kubevirt/pull/15949) | [xpivarc](https://github.com/xpivarc) |
| fix: KSM is enabled in case of node pressure within 3 minutes  | [#16049](https://github.com/kubevirt/kubevirt/pull/16049) | [fossedihelm](https://github.com/fossedihelm) |
| Introduce new API - UtilityVolumes - direct virt-launcher attachment mechanism  | [#15922](https://github.com/kubevirt/kubevirt/pull/15922) | [ShellyKa13](https://github.com/ShellyKa13) |
| kubevirt.io/cpumanager label is advertised for nodes capable of running dedicated VMs.  | [#14892](https://github.com/kubevirt/kubevirt/pull/14892) | [xpivarc](https://github.com/xpivarc) |
| Allow migration when host model changes after libvirt upgrade.  | [#15694](https://github.com/kubevirt/kubevirt/pull/15694) | [Barakmor1](https://github.com/Barakmor1) |
| Add RestartRequired when detaching CD-ROMs from a running VM  | [#15969](https://github.com/kubevirt/kubevirt/pull/15969) | [Dsanatar](https://github.com/Dsanatar) |
| Add GuestFilesystemAlmostOutOfSpace alerts  | [#15714](https://github.com/kubevirt/kubevirt/pull/15714) | [machadovilaca](https://github.com/machadovilaca) |
| Introduce a new subresource `/evacuate/cancel` and `virtctl evacuate-cancel` command to allow users to cancel the evacuation process for a VirtualMachineInstance (VMI). This clears the `evacuationNodeName` field in the VMI's status, stopping the automatic creation of migration resources and fully aborting the eviction cycle.  | [#15957](https://github.com/kubevirt/kubevirt/pull/15957) | [xpivarc](https://github.com/xpivarc) |
| The `MultiArchitecture` feature gate has been deprecated and is no longer used to determine if VirtualMachines with a differing architecture to the control plane should be rejected by the admission webhooks  | [#16023](https://github.com/kubevirt/kubevirt/pull/16023) | [lyarwood](https://github.com/lyarwood) |
| Reject stop requests for paused VMIs.  A paused VMI must be unpaused before it can be stopped.  | [#15405](https://github.com/kubevirt/kubevirt/pull/15405) | [dasionov](https://github.com/dasionov) |
| A decentralized live migration failure is properly propagates between source and target  | [#15716](https://github.com/kubevirt/kubevirt/pull/15716) | [awels](https://github.com/awels) |
| NodeRestriction: Source of node update is now verified  | [#15374](https://github.com/kubevirt/kubevirt/pull/15374) | [xpivarc](https://github.com/xpivarc) |
| Bug fix: KubeVirt.spec.imagetag installation is working again  | [#16050](https://github.com/kubevirt/kubevirt/pull/16050) | [xpivarc](https://github.com/xpivarc) |
| Recording rule kubevirt_vmi_vcpu_count name changes to vmi:kubevirt_vmi_vcpu:count  | [#15968](https://github.com/kubevirt/kubevirt/pull/15968) | [sradco](https://github.com/sradco) |
| Introduce pool.kubevirt.io/v1beta1  | [#15166](https://github.com/kubevirt/kubevirt/pull/15166) | [Sreeja1725](https://github.com/Sreeja1725) |
| VMSnapshot: add SourceIndications status field to list snapshot indications with descriptions for clearer meaning.  | [#15409](https://github.com/kubevirt/kubevirt/pull/15409) | [noamasu](https://github.com/noamasu) |
| Promote IBM Secure Execution Feature to Beta stage.  | [#15934](https://github.com/kubevirt/kubevirt/pull/15934) | [jschintag](https://github.com/jschintag) |
| BugFix: The migration limit was not accurately being used with decentralized live migrations  | [#15767](https://github.com/kubevirt/kubevirt/pull/15767) | [awels](https://github.com/awels) |
| The KubevirtSeccompProfile feature is now in Beta  | [#15970](https://github.com/kubevirt/kubevirt/pull/15970) | [jean-edouard](https://github.com/jean-edouard) |
| promote ImageVolume FG to Beta  | [#15960](https://github.com/kubevirt/kubevirt/pull/15960) | [Barakmor1](https://github.com/Barakmor1) |
| VMPool: Add support for auto-healing startegy  | [#15638](https://github.com/kubevirt/kubevirt/pull/15638) | [Sreeja1725](https://github.com/Sreeja1725) |
| VMpool: Add Scale-in strategy support with Proactive, opportunistic modes and statePreservation  | [#15604](https://github.com/kubevirt/kubevirt/pull/15604) | [Sreeja1725](https://github.com/Sreeja1725) |
| support v0.32.5 code generator  | [#15529](https://github.com/kubevirt/kubevirt/pull/15529) | [Yu-Jack](https://github.com/Yu-Jack) |


_This page is updated daily._
