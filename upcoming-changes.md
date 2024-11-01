---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
| Enable virt-exportproxy and virt-exportserver image for s390x  | [#12844](https://github.com/kubevirt/kubevirt/pull/12844) | [jschintag](https://github.com/jschintag) |
| VMs admitter: remove validation of vm clone volume from the webhook  | [#12628](https://github.com/kubevirt/kubevirt/pull/12628) | [ShellyKa13](https://github.com/ShellyKa13) |
| Added labels, annotations to VM Export resources and configurable pod readiness timeout  | [#13006](https://github.com/kubevirt/kubevirt/pull/13006) | [chomatdam](https://github.com/chomatdam) |
| GA the `VMLiveUpdateFeatures` feature-gate.  | [#13091](https://github.com/kubevirt/kubevirt/pull/13091) | [acardace](https://github.com/acardace) |
| Add kubevirt_vm_resource_limits metric  | [#13071](https://github.com/kubevirt/kubevirt/pull/13071) | [machadovilaca](https://github.com/machadovilaca) |
| Allow live updating VMs' tolerations  | [#13090](https://github.com/kubevirt/kubevirt/pull/13090) | [acardace](https://github.com/acardace) |
| backend-storage now supports RWO FS  | [#12629](https://github.com/kubevirt/kubevirt/pull/12629) | [jean-edouard](https://github.com/jean-edouard) |
| A new `spec.configuration.instancetype.referencePolicy` configurable has been added to the `KubeVirt` CR with support for `reference` (default), `expand` and `expandAll` policies provided.  | [#13086](https://github.com/kubevirt/kubevirt/pull/13086) | [lyarwood](https://github.com/lyarwood) |
| Fix cache corruption  | [#13064](https://github.com/kubevirt/kubevirt/pull/13064) | [xpivarc](https://github.com/xpivarc) |
| BochsDisplayForEFIGuests is GAed, use  "kubevirt.io/vga-display-efi-x86" annotation on Kubevirt CR before upgrading in case you need retain compatibility.  | [#12967](https://github.com/kubevirt/kubevirt/pull/12967) | [xpivarc](https://github.com/xpivarc) |
| Add dynamic pod interface name feature gate  | [#13078](https://github.com/kubevirt/kubevirt/pull/13078) | [qinqon](https://github.com/qinqon) |
| virtctl: virtctl create vm can now use the Access Credentials API to add credentials to a new VM  | [#13072](https://github.com/kubevirt/kubevirt/pull/13072) | [0xFelix](https://github.com/0xFelix) |
| fix the cpu model issue for s390x.  | [#13050](https://github.com/kubevirt/kubevirt/pull/13050) | [vamsikrishna-siddu](https://github.com/vamsikrishna-siddu) |
| Add kubevirt_vmi_status_addresses metric  | [#12802](https://github.com/kubevirt/kubevirt/pull/12802) | [machadovilaca](https://github.com/machadovilaca) |
| BugFix: Stop creating tokenSecretRef when no volumes to export  | [#13027](https://github.com/kubevirt/kubevirt/pull/13027) | [awels](https://github.com/awels) |
| Relaxed check on modify VM spec during VM snapshot to only check disks/volumes  | [#13001](https://github.com/kubevirt/kubevirt/pull/13001) | [awels](https://github.com/awels) |
| Updated common-instancetypes bundles to v1.2.0  | [#13082](https://github.com/kubevirt/kubevirt/pull/13082) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| vmsnapshot: Enable status subresource for snapshot.kubevirt.io api group  | [#12601](https://github.com/kubevirt/kubevirt/pull/12601) | [mhenriks](https://github.com/mhenriks) |
| Support Dynamic Primary Pod NIC Name  | [#13018](https://github.com/kubevirt/kubevirt/pull/13018) | [orelmisan](https://github.com/orelmisan) |
| virtctl: The flags `--volume-clone-pvc`, `--volume-datasource` and `--volume-blank` are deprecated in favor of the `--volume-import` flag.  | [#13019](https://github.com/kubevirt/kubevirt/pull/13019) | [0xFelix](https://github.com/0xFelix) |
| Network hotplug feature is declared as GA.  | [#13059](https://github.com/kubevirt/kubevirt/pull/13059) | [EdDev](https://github.com/EdDev) |
| network binding plugin: Introduce a new `managedTap` `domainAttachmentType`  | [#13024](https://github.com/kubevirt/kubevirt/pull/13024) | [EdDev](https://github.com/EdDev) |
| Network binding plugins feature is declared as Beta.  | [#13060](https://github.com/kubevirt/kubevirt/pull/13060) | [EdDev](https://github.com/EdDev) |
| Add 'machine_type' label for kubevirt_vm_info metric  | [#13045](https://github.com/kubevirt/kubevirt/pull/13045) | [dasionov](https://github.com/dasionov) |
| Removed the ManualRecoveryRequired field from the VolumeMigrationState and convert it to the VM condition ManualRecoveryRequired  | [#13030](https://github.com/kubevirt/kubevirt/pull/13030) | [alicefr](https://github.com/alicefr) |
| virtctl: Users can specify a sysprep volume in VMs created with virtctl create vm  | [#13053](https://github.com/kubevirt/kubevirt/pull/13053) | [0xFelix](https://github.com/0xFelix) |
| virtctl expose: Drop flag to set deprecated LoadBalancerIP option  | [#12855](https://github.com/kubevirt/kubevirt/pull/12855) | [0xFelix](https://github.com/0xFelix) |
| virtctl: Allow creating a basic cloud-init config with virtctl create vm  | [#13008](https://github.com/kubevirt/kubevirt/pull/13008) | [0xFelix](https://github.com/0xFelix) |
| fix: Proxies configured in kubeconfig are used in client-go for asynchronous subresources like VNC or Console  | [#12829](https://github.com/kubevirt/kubevirt/pull/12829) | [0xFelix](https://github.com/0xFelix) |
| Bugfix: Fix disk expansion logic by checking usable size instead of requested capacity  | [#12733](https://github.com/kubevirt/kubevirt/pull/12733) | [alromeros](https://github.com/alromeros) |
| Update code-generators to 1.31.1  | [#13052](https://github.com/kubevirt/kubevirt/pull/13052) | [fossedihelm](https://github.com/fossedihelm) |
| Build KubeVirt with go v1.22.8  | [#12882](https://github.com/kubevirt/kubevirt/pull/12882) | [brianmcarey](https://github.com/brianmcarey) |
| BugFix: Allow VMExport to work with VM columes that have dots in its name  | [#13040](https://github.com/kubevirt/kubevirt/pull/13040) | [awels](https://github.com/awels) |
| Update k8s dependencies to 0.31.0<br>Dropped `computeResourceOverhead.claims` and `supportContainerResources.claims` fields in KV config  | [#12729](https://github.com/kubevirt/kubevirt/pull/12729) | [fossedihelm](https://github.com/fossedihelm) |
| Fixed additional broken amd64 image in some image manifests  | [#12867](https://github.com/kubevirt/kubevirt/pull/12867) | [jschintag](https://github.com/jschintag) |
| Deprecate the DockerSELinuxMCS FeatureGate  | [#12940](https://github.com/kubevirt/kubevirt/pull/12940) | [Barakmor1](https://github.com/Barakmor1) |
| The `GPU` feature gate is now deprecated with the feature state graduated to `GA` and thus enabled by default  | [#12943](https://github.com/kubevirt/kubevirt/pull/12943) | [Barakmor1](https://github.com/Barakmor1) |
| Add a 'outdated' label to kubevirt_vmi_info metric  | [#12992](https://github.com/kubevirt/kubevirt/pull/12992) | [machadovilaca](https://github.com/machadovilaca) |
| VM admitter: improve validation of vm spec datavolumetemplate  | [#12933](https://github.com/kubevirt/kubevirt/pull/12933) | [ShellyKa13](https://github.com/ShellyKa13) |
| The `PreferredEfi` preference is now only applied when a user has not already enabled either `EFI` or `BIOS` within the underlying `VirtualMachine`.  | [#12986](https://github.com/kubevirt/kubevirt/pull/12986) | [lyarwood](https://github.com/lyarwood) |
| Integrate kwok with sig-scale tests  | [#12117](https://github.com/kubevirt/kubevirt/pull/12117) | [Sreeja1725](https://github.com/Sreeja1725) |
| Update kubevirt_rest_client_request_latency_seconds to count list calls if made using query params  | [#12716](https://github.com/kubevirt/kubevirt/pull/12716) | [Sreeja1725](https://github.com/Sreeja1725) |
| Mark Running field as deprecated  | [#12578](https://github.com/kubevirt/kubevirt/pull/12578) | [dasionov](https://github.com/dasionov) |
| The `CommonInstancetypesDeploymentGate` feature gate and underlying feature are graduated to GA and now always enabled by default. A single new `KubeVirt` configurable is also introduced to allow cluster admins a way of explicitly disabling deployment when required.  | [#12753](https://github.com/kubevirt/kubevirt/pull/12753) | [lyarwood](https://github.com/lyarwood) |
| Add kubevirt_vmsnapshot_succeeded_timestamp_seconds metric  | [#12645](https://github.com/kubevirt/kubevirt/pull/12645) | [avlitman](https://github.com/avlitman) |
| add s390x support for kubevirt builder  | [#11097](https://github.com/kubevirt/kubevirt/pull/11097) | [vamsikrishna-siddu](https://github.com/vamsikrishna-siddu) |
| Rename kubevirt_vm_resource_requests 'vmi' label to 'name'  | [#12910](https://github.com/kubevirt/kubevirt/pull/12910) | [machadovilaca](https://github.com/machadovilaca) |
| Reduce default CompletionTimeoutPerGiB from 800s to 150s  | [#12848](https://github.com/kubevirt/kubevirt/pull/12848) | [iholder101](https://github.com/iholder101) |
| bugfix: fix possible miss update of datavolumename on vmrestore restores  | [#12861](https://github.com/kubevirt/kubevirt/pull/12861) | [ShellyKa13](https://github.com/ShellyKa13) |
| Increase periodicity in domainstats migration metrics  | [#12441](https://github.com/kubevirt/kubevirt/pull/12441) | [machadovilaca](https://github.com/machadovilaca) |
| Add kubevirt_vm_info metric  | [#12718](https://github.com/kubevirt/kubevirt/pull/12718) | [machadovilaca](https://github.com/machadovilaca) |
| MaxCpuSockets won't block creation of VMs with more Sockets than MaxCpuSockets declare  | [#12599](https://github.com/kubevirt/kubevirt/pull/12599) | [xpivarc](https://github.com/xpivarc) |
| BugFix: Fail to create VMExport via virtctl vmexport create  | [#12857](https://github.com/kubevirt/kubevirt/pull/12857) | [akalenyu](https://github.com/akalenyu) |
| Add the volume migration state in the VM status  | [#12355](https://github.com/kubevirt/kubevirt/pull/12355) | [alicefr](https://github.com/alicefr) |
| Concurrent addvolume/removevolume using virtctl no longer fail if concurrent modifications happen  | [#12726](https://github.com/kubevirt/kubevirt/pull/12726) | [awels](https://github.com/awels) |
| bugfix: In case of err in vmrestore, leave VM without RestoreInProgress annotation allowing it to be started  | [#12835](https://github.com/kubevirt/kubevirt/pull/12835) | [ShellyKa13](https://github.com/ShellyKa13) |
| bug-fix: Ensure PDB associated with a VMI is deleted when it Reaches Succeeded or Failed phase  | [#12809](https://github.com/kubevirt/kubevirt/pull/12809) | [dasionov](https://github.com/dasionov) |
| BugFix: can't create export pod on OpenShift  | [#12813](https://github.com/kubevirt/kubevirt/pull/12813) | [akalenyu](https://github.com/akalenyu) |
| virtctl: Created VMs can infer an instancetype or preference from PVC, Registry and Snapshot sources now.  | [#12786](https://github.com/kubevirt/kubevirt/pull/12786) | [0xFelix](https://github.com/0xFelix) |
| bugfix: vmrestore create DVs before creation/update of restored VM  | [#12764](https://github.com/kubevirt/kubevirt/pull/12764) | [ShellyKa13](https://github.com/ShellyKa13) |
| Continue changes to Ginkgo V2 Serial runs  | [#10562](https://github.com/kubevirt/kubevirt/pull/10562) | [dhiller](https://github.com/dhiller) |
| enable initial e2e tests for s390x.  | [#12516](https://github.com/kubevirt/kubevirt/pull/12516) | [vamsikrishna-siddu](https://github.com/vamsikrishna-siddu) |
| A new `PreferredEfi` field has been added to preferences to express the preferred `EFI` configuration for a given `VirtualMachine`.  | [#12739](https://github.com/kubevirt/kubevirt/pull/12739) | [lyarwood](https://github.com/lyarwood) |
| Add evictable label to kubevirt_vmi_info  | [#12737](https://github.com/kubevirt/kubevirt/pull/12737) | [machadovilaca](https://github.com/machadovilaca) |
| The `NUMA` feature gate is now deprecated with the feature state graduated to `GA` and thus enabled by default  | [#12232](https://github.com/kubevirt/kubevirt/pull/12232) | [lyarwood](https://github.com/lyarwood) |
| vmsnapshot: when checking if a VM is running, ignore runStrategy  | [#12582](https://github.com/kubevirt/kubevirt/pull/12582) | [mhenriks](https://github.com/mhenriks) |
| Add kubevirt_vm_resource_requests for CPU resource  | [#12625](https://github.com/kubevirt/kubevirt/pull/12625) | [machadovilaca](https://github.com/machadovilaca) |
| vmexport: enable status subresource for VirtualMachineExport  | [#12605](https://github.com/kubevirt/kubevirt/pull/12605) | [mhenriks](https://github.com/mhenriks) |
| replace `Update()` with `Patch()` for `test VirtualMachineInstancesPerNode`  | [#12616](https://github.com/kubevirt/kubevirt/pull/12616) | [orenc1](https://github.com/orenc1) |
| Optionally create data source using virtctl image upload.  | [#12557](https://github.com/kubevirt/kubevirt/pull/12557) | [codingben](https://github.com/codingben) |
| virt-api: skip clone auth check when DataVolume already exists  | [#12547](https://github.com/kubevirt/kubevirt/pull/12547) | [mhenriks](https://github.com/mhenriks) |
| Bridge binding: Static routes to subnets containing the pod's NIC IP address are passed to the VM.  | [#12613](https://github.com/kubevirt/kubevirt/pull/12613) | [orelmisan](https://github.com/orelmisan) |
| [tests] introduce a decorator for Periodic_only tests  | [#12594](https://github.com/kubevirt/kubevirt/pull/12594) | [tiraboschi](https://github.com/tiraboschi) |
| Add kubevirt_vm_resource_requests metric for memory resource  | [#12593](https://github.com/kubevirt/kubevirt/pull/12593) | [machadovilaca](https://github.com/machadovilaca) |
| grpc from go.mod is now correctly shipped in release images  | [#12617](https://github.com/kubevirt/kubevirt/pull/12617) | [Acedus](https://github.com/Acedus) |
| BugFix: "Cannot allocate memory" warnings for containerdisk VMs  | [#12638](https://github.com/kubevirt/kubevirt/pull/12638) | [akalenyu](https://github.com/akalenyu) |
| Add new condition for VMIStorageLiveMigratable  | [#12395](https://github.com/kubevirt/kubevirt/pull/12395) | [alicefr](https://github.com/alicefr) |
| Add timeout to validation webhooks  | [#12419](https://github.com/kubevirt/kubevirt/pull/12419) | [nunnatsa](https://github.com/nunnatsa) |
| Fixed issue emitting created secret events when not actually creating secrets during VMExport setup  | [#12592](https://github.com/kubevirt/kubevirt/pull/12592) | [awels](https://github.com/awels) |
| Build KubeVirt with go v1.22.6  | [#12584](https://github.com/kubevirt/kubevirt/pull/12584) | [brianmcarey](https://github.com/brianmcarey) |
| Advise users to use RunStrategy in virt-api messages  | [#12575](https://github.com/kubevirt/kubevirt/pull/12575) | [Barakmor1](https://github.com/Barakmor1) |
| tests/vm_tests.go: replace Update() with Patch()  | [#12466](https://github.com/kubevirt/kubevirt/pull/12466) | [orenc1](https://github.com/orenc1) |
| Updated common-instancetypes bundles to v1.1.0  | [#12548](https://github.com/kubevirt/kubevirt/pull/12548) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| Enable live-migration and node labels on s390x  | [#12476](https://github.com/kubevirt/kubevirt/pull/12476) | [jschintag](https://github.com/jschintag) |
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
