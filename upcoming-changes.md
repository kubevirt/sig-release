---
title: Noteworthy changes for the next KubeVirt release
---

# Noteworthy changes for the next KubeVirt release

This list contains the noteworthy changes made after the latest KubeVirt release. The community expects these changes to be included in the next Kubevirt release.

> [!WARNING]
> **Please be aware that any of these might be excluded from the next release.**

| Upcoming changes | PR                                                                   | Author                                          |
|------------------|----------------------------------------------------------------------|-------------------------------------------------|
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
| Release passt CNI image, instead the CNI binary itself.  | [#14887](https://github.com/kubevirt/kubevirt/pull/14887) | [oshoval](https://github.com/oshoval) |
| Changed the time unit conversion in the kubevirt_vmi_vcpu_seconds_total metric from microseconds to nanoseconds.  | [#13898](https://github.com/kubevirt/kubevirt/pull/13898) | [brandboat](https://github.com/brandboat) |
| Add virtctl objectgraph command  | [#14935](https://github.com/kubevirt/kubevirt/pull/14935) | [alromeros](https://github.com/alromeros) |
| A few dynamic annotations are synced from VMs template to VMIs and to virt-launcher pods  | [#14744](https://github.com/kubevirt/kubevirt/pull/14744) | [tiraboschi](https://github.com/tiraboschi) |
| Allow virtio bus for hotplugged disks  | [#14907](https://github.com/kubevirt/kubevirt/pull/14907) | [mhenriks](https://github.com/mhenriks) |
| Allocate more PCI ports for hotplug  | [#14754](https://github.com/kubevirt/kubevirt/pull/14754) | [mhenriks](https://github.com/mhenriks) |
| Feature: Support for defining panic devices in VirtualMachineInstances to catch crash signals from the guest.  | [#13103](https://github.com/kubevirt/kubevirt/pull/13103) | [varunrsekar](https://github.com/varunrsekar) |
| BugFix: Can't LiveMigrate Windows VM after Storage Migration from HPP to OCS  | [#14961](https://github.com/kubevirt/kubevirt/pull/14961) | [akalenyu](https://github.com/akalenyu) |
| Added CRC to ADOPTERS document.  | [#14956](https://github.com/kubevirt/kubevirt/pull/14956) | [RobertoMachorro](https://github.com/RobertoMachorro) |
| The migration controller in virt-handler has been re-architected, migrations should be more stable  | [#14705](https://github.com/kubevirt/kubevirt/pull/14705) | [jean-edouard](https://github.com/jean-edouard) |
| KubeVirt doesn't use PDBs anymore  | [#13764](https://github.com/kubevirt/kubevirt/pull/13764) | [xpivarc](https://github.com/xpivarc) |
| VirtualMachinePool now supports a `.ScaleInStrategy.Proactive.SelectionPolicy.BasePolicy` field to control scale-down behavior. The new `"DescendingOrder"` strategy deletes VMs by descending ordinal index, offering predictable downscale behavior. Defaults to `"random"` if not specified.  | [#14801](https://github.com/kubevirt/kubevirt/pull/14801) | [arsiesys](https://github.com/arsiesys) |
| Integrate NIC hotplug with LiveUpdate rollout strategy  | [#14259](https://github.com/kubevirt/kubevirt/pull/14259) | [orelmisan](https://github.com/orelmisan) |
| Add Video Configuration Field for VMs to Enable Explicit Video Device Selection  | [#14673](https://github.com/kubevirt/kubevirt/pull/14673) | [dasionov](https://github.com/dasionov) |
| Windows offline activation with ACPI MSDM table  | [#14681](https://github.com/kubevirt/kubevirt/pull/14681) | [victortoso](https://github.com/victortoso) |
| Add VolumeRestorePolicies and VolumeRestoreOverrides to VMRestores  | [#14723](https://github.com/kubevirt/kubevirt/pull/14723) | [SkalaNetworks](https://github.com/SkalaNetworks) |
| Add support for Secure Execution VMs on IBM Z  | [#14040](https://github.com/kubevirt/kubevirt/pull/14040) | [jschintag](https://github.com/jschintag) |
| Declarative Volume Hotplug with CD-ROM Inject/Eject  | [#13847](https://github.com/kubevirt/kubevirt/pull/13847) | [mhenriks](https://github.com/mhenriks) |
| Add Object Graph subresource  | [#14807](https://github.com/kubevirt/kubevirt/pull/14807) | [alromeros](https://github.com/alromeros) |
| Failed post-copy migrations now always end in VMI failure  | [#14793](https://github.com/kubevirt/kubevirt/pull/14793) | [jean-edouard](https://github.com/jean-edouard) |
| virt-handler: Reduce Get() calls for KSM handling  | [#14632](https://github.com/kubevirt/kubevirt/pull/14632) | [iholder101](https://github.com/iholder101) |
| Bugfix: Update backend-storage logic so it works with PVCs with non-standard naming convention  | [#14658](https://github.com/kubevirt/kubevirt/pull/14658) | [alromeros](https://github.com/alromeros) |
| Fix network setup when emulation is enabled  | [#14827](https://github.com/kubevirt/kubevirt/pull/14827) | [orelmisan](https://github.com/orelmisan) |
| Move cgroup v1 to maintenance mode  | [#14538](https://github.com/kubevirt/kubevirt/pull/14538) | [iholder101](https://github.com/iholder101) |
| Adding Isovalent to Adopters  | [#14823](https://github.com/kubevirt/kubevirt/pull/14823) | [xmulligan](https://github.com/xmulligan) |
| Expose CONTAINER_NAME on hook sidecars.  | [#14768](https://github.com/kubevirt/kubevirt/pull/14768) | [oshoval](https://github.com/oshoval) |
| Add maxUnavailable support to VirtualMachinePool<br>Adds support for controlling the maximum number of unavailable VMIs during pool updates. This helps maintain service availability by limiting how many VMs can be updated simultaneously.<br>- Defaults to 100% of total replicas for backward compatibility<br>- Can be configured as percentage (e.g. "25%") or absolute number<br>- Prevents updates if current unavailable VMs exceed the limit<br>- Similar to Kubernetes Deployment's maxUnavailable behavior  | [#14183](https://github.com/kubevirt/kubevirt/pull/14183) | [aqilbeig](https://github.com/aqilbeig) |
| Bugfix: Fix online expansion by requeuing VMIs on PVC size change  | [#14695](https://github.com/kubevirt/kubevirt/pull/14695) | [alromeros](https://github.com/alromeros) |
| Clean absent interfaces and their relative networks from stopped VMs.  | [#14738](https://github.com/kubevirt/kubevirt/pull/14738) | [oshoval](https://github.com/oshoval) |
| virt-Freeze: skip freeze if domain is not in running state  | [#14737](https://github.com/kubevirt/kubevirt/pull/14737) | [ShellyKa13](https://github.com/ShellyKa13) |
| CPU hotplug with net multi-queue is now allowed  | [#14728](https://github.com/kubevirt/kubevirt/pull/14728) | [orelmisan](https://github.com/orelmisan) |
| VirtualMachineInstanceMigrations can now express that they are source or target migrations  | [#14616](https://github.com/kubevirt/kubevirt/pull/14616) | [awels](https://github.com/awels) |
| virtctl vnc command now supports user provided VNC clients.  | [#14619](https://github.com/kubevirt/kubevirt/pull/14619) | [cloud-j-luna](https://github.com/cloud-j-luna) |
| bug-fix: persist VM's firmware UUID for existing VMs  | [#14130](https://github.com/kubevirt/kubevirt/pull/14130) | [dasionov](https://github.com/dasionov) |
| ARM: CPU pinning doesn't panic now  | [#14640](https://github.com/kubevirt/kubevirt/pull/14640) | [xpivarc](https://github.com/xpivarc) |
| Build KubeVirt with go v1.23.9  | [#14664](https://github.com/kubevirt/kubevirt/pull/14664) | [brianmcarey](https://github.com/brianmcarey) |
| Enabled watchdog validation on watchdog device models  | [#14599](https://github.com/kubevirt/kubevirt/pull/14599) | [HarshithaMS005](https://github.com/HarshithaMS005) |
| Dirty rate is reported as part of a new `GetDomainDirtyRateStats()` gRPC method and by a Prometheus metric: `kubevirt_vmi_dirty_rate_bytes_per_second`.  | [#13806](https://github.com/kubevirt/kubevirt/pull/13806) | [iholder101](https://github.com/iholder101) |
| Added support for custom JSON patches in VirtualMachineClones.<br>A new property `patches` was added to the VirtualMachineClones CRD.  | [#14617](https://github.com/kubevirt/kubevirt/pull/14617) | [SkalaNetworks](https://github.com/SkalaNetworks) |
| Label backend PVC to support CDI WebhookPvcRendering  | [#14637](https://github.com/kubevirt/kubevirt/pull/14637) | [alromeros](https://github.com/alromeros) |
| add CloudCasa by Catalogic to integrations in the adopters.md  | [#14440](https://github.com/kubevirt/kubevirt/pull/14440) | [pstaniec-catalogicsoftware](https://github.com/pstaniec-catalogicsoftware) |
| The "RestartRequired" condition is not set on VM objects for live-updatable network fields  | [#14602](https://github.com/kubevirt/kubevirt/pull/14602) | [orelmisan](https://github.com/orelmisan) |
| Implement container disk functionality using ImageVolume, protected by the ImageVolume feature gate.  | [#14267](https://github.com/kubevirt/kubevirt/pull/14267) | [Barakmor1](https://github.com/Barakmor1) |
| Enable vhost-user mode for passt network binding plugin  | [#14539](https://github.com/kubevirt/kubevirt/pull/14539) | [nirdothan](https://github.com/nirdothan) |
| Enable node-labeller for ARM64 clusters, supporting machine-type labels.  | [#14520](https://github.com/kubevirt/kubevirt/pull/14520) | [dasionov](https://github.com/dasionov) |
| Trigger VMCannotBeEvicted only for running VMIs  | [#14203](https://github.com/kubevirt/kubevirt/pull/14203) | [machadovilaca](https://github.com/machadovilaca) |
| The 64-Bit PCI hole can now be disabled by adding the kubevirt.io/disablePCIHole annotation to VirtualMachineInstances. This allows legacy OSes such as Windows XP or Server 2003 to boot on KubeVirt using the Q35 machine type.  | [#14449](https://github.com/kubevirt/kubevirt/pull/14449) | [0xFelix](https://github.com/0xFelix) |
| hotplug volume: Boot from hotpluggable disk  | [#13297](https://github.com/kubevirt/kubevirt/pull/13297) | [mhenriks](https://github.com/mhenriks) |
| Network conformance tests are now marked using the `Conformance` decorator. Use `--ginkgo.label-filter '(sig-network && conformance)` to select them.  | [#14509](https://github.com/kubevirt/kubevirt/pull/14509) | [phoracek](https://github.com/phoracek) |
| Bug fix: MaxSockets is limited so maximum of vcpus doesn't go over 512.  | [#14338](https://github.com/kubevirt/kubevirt/pull/14338) | [dasionov](https://github.com/dasionov) |
| Handle lowercase instancetypes/preference keys in VM monitoring  | [#14327](https://github.com/kubevirt/kubevirt/pull/14327) | [machadovilaca](https://github.com/machadovilaca) |
| Ensure stricter check for valid machine type when validating VMI  | [#14437](https://github.com/kubevirt/kubevirt/pull/14437) | [jschintag](https://github.com/jschintag) |
| VirtHandlerRESTErrorsHigh, VirtOperatorRESTErrorsHigh, VirtAPIRESTErrorsHigh and VirtControllerRESTErrorsHigh alerts removed.  | [#13911](https://github.com/kubevirt/kubevirt/pull/13911) | [avlitman](https://github.com/avlitman) |
| Enable Watchdog device support on s390x using the Diag288 device model.  | [#14277](https://github.com/kubevirt/kubevirt/pull/14277) | [HarshithaMS005](https://github.com/HarshithaMS005) |
| guest console log: make virt-tail a proper sidecar  | [#13422](https://github.com/kubevirt/kubevirt/pull/13422) | [mhenriks](https://github.com/mhenriks) |
| Added kubevirt_vmi_migrations_in_unset_phase, instead of including it in kubevirt_vmi_migration_failed.  | [#14426](https://github.com/kubevirt/kubevirt/pull/14426) | [avlitman](https://github.com/avlitman) |
| To use nfs-csi, the env variable KUBEVIRT_NFS_DIR has to be set to a location on the host for NFS data  | [#14428](https://github.com/kubevirt/kubevirt/pull/14428) | [jean-edouard](https://github.com/jean-edouard) |
| Bugfix: Truncate volume names in export pod  | [#13951](https://github.com/kubevirt/kubevirt/pull/13951) | [alromeros](https://github.com/alromeros) |
| supplementalPool added to the description of the ioThreadsPolicy possible values  | [#14405](https://github.com/kubevirt/kubevirt/pull/14405) | [jpeimer](https://github.com/jpeimer) |
| handle nil pointer dereference in cellToCell  | [#14145](https://github.com/kubevirt/kubevirt/pull/14145) | [ayushpatil2122](https://github.com/ayushpatil2122) |
| VMRestore: Keep VM RunStrategy as before the restore  | [#14281](https://github.com/kubevirt/kubevirt/pull/14281) | [ShellyKa13](https://github.com/ShellyKa13) |
| Updated common-instancetypes bundles to v1.3.1  | [#14374](https://github.com/kubevirt/kubevirt/pull/14374) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| A request to create a VirtualMachines that references a non-existent  instance type or preference are no longer rejected. The VirtualMachine will be created but will fail to start until the missing resources are created in the cluster.  | [#14219](https://github.com/kubevirt/kubevirt/pull/14219) | [lyarwood](https://github.com/lyarwood) |
| Don't expose as VMI status the implicit qemu domain pause at the end of live migration  | [#14288](https://github.com/kubevirt/kubevirt/pull/14288) | [qinqon](https://github.com/qinqon) |
| Fixed persistent reservation support for multipathd by improving socket access and multipath files in pr-helper  | [#14309](https://github.com/kubevirt/kubevirt/pull/14309) | [alicefr](https://github.com/alicefr) |
| fix: disks-images-provider is pointing to wrong alpine image for s390x.  | [#14325](https://github.com/kubevirt/kubevirt/pull/14325) | [vamsikrishna-siddu](https://github.com/vamsikrishna-siddu) |
| The `v1alpha{1,2}` versions of the `instancetype.kubevirt.io` API group are no longer served or supported.  | [#14048](https://github.com/kubevirt/kubevirt/pull/14048) | [lyarwood](https://github.com/lyarwood) |
| A new `Enabled` attribute has been added to the `TPM` device allowing users to explicitly disable the device regardless of any referenced preference.  | [#14316](https://github.com/kubevirt/kubevirt/pull/14316) | [lyarwood](https://github.com/lyarwood) |
| Cleanup: Fix unit tests on a sane, non-host-cgroup-sharing development setup  | [#14328](https://github.com/kubevirt/kubevirt/pull/14328) | [akalenyu](https://github.com/akalenyu) |
| Add interface name label to kubevirt_vmi_status_addresses  | [#14108](https://github.com/kubevirt/kubevirt/pull/14108) | [machadovilaca](https://github.com/machadovilaca) |
| The `InstancetypeReferencePolicy` feature has graduated to GA and no longer requires the associated feature gate to be enabled.  | [#14050](https://github.com/kubevirt/kubevirt/pull/14050) | [lyarwood](https://github.com/lyarwood) |
| Register k8s client-go latency metrics on init  | [#14286](https://github.com/kubevirt/kubevirt/pull/14286) | [machadovilaca](https://github.com/machadovilaca) |
| Update module github.com/containers/common to v0.60.4  | [#14304](https://github.com/kubevirt/kubevirt/pull/14304) | [jean-edouard](https://github.com/jean-edouard) |
| VM Persistent State GA  | [#14065](https://github.com/kubevirt/kubevirt/pull/14065) | [jean-edouard](https://github.com/jean-edouard) |
| VMSnapshot: add QuiesceFailed indication to snapshot if freeze failed  | [#14096](https://github.com/kubevirt/kubevirt/pull/14096) | [ShellyKa13](https://github.com/ShellyKa13) |
| Update module golang.org/x/oauth2 to v0.27.0  | [#14215](https://github.com/kubevirt/kubevirt/pull/14215) | [dominikholler](https://github.com/dominikholler) |
| Default VM Rollout Strategy is now LiveUpdate. Important: to preserve previous behavior, rolloutStrategy needs to be set to "Stage" in the KubeVirt CR.  | [#14068](https://github.com/kubevirt/kubevirt/pull/14068) | [jean-edouard](https://github.com/jean-edouard) |
| Update module golang.org/x/net to v0.36.0  | [#14222](https://github.com/kubevirt/kubevirt/pull/14222) | [dominikholler](https://github.com/dominikholler) |
| Update golang.org/x/crypto to v0.35.0  | [#14218](https://github.com/kubevirt/kubevirt/pull/14218) | [dominikholler](https://github.com/dominikholler) |
| Update module github.com/opencontainers/runc to v1.1.14  | [#14217](https://github.com/kubevirt/kubevirt/pull/14217) | [dominikholler](https://github.com/dominikholler) |
| Large number of migrations should no longer lead to active migrations timing out  | [#14141](https://github.com/kubevirt/kubevirt/pull/14141) | [jean-edouard](https://github.com/jean-edouard) |
| Ensure launcher pods are finalized and deleted before removing the VMI finalizer when the VMI is marked for deletion.  | [#13870](https://github.com/kubevirt/kubevirt/pull/13870) | [dasionov](https://github.com/dasionov) |
| libvirt: 10.10.0-7, qemu: 9.1.0-15  | [#14101](https://github.com/kubevirt/kubevirt/pull/14101) | [qinqon](https://github.com/qinqon) |
| Add entrypoint to the pr-helper for creating the symlink to the multipath socket  | [#14071](https://github.com/kubevirt/kubevirt/pull/14071) | [alicefr](https://github.com/alicefr) |
| Support live migration to a named node  | [#12725](https://github.com/kubevirt/kubevirt/pull/12725) | [tiraboschi](https://github.com/tiraboschi) |
| Add v1.5.0 perf and scale benchmarks data  | [#13888](https://github.com/kubevirt/kubevirt/pull/13888) | [Sreeja1725](https://github.com/Sreeja1725) |
| The virtctl port-forward/ssh/scp syntax was changed to type/name[/namespace]. It now supports resources with dots in their name properly.  | [#13939](https://github.com/kubevirt/kubevirt/pull/13939) | [0xFelix](https://github.com/0xFelix) |
| virt-launcher now uses bash to retrieve disk info and verify container-disk files, requiring bash to be included in the launcher image  | [#13807](https://github.com/kubevirt/kubevirt/pull/13807) | [Barakmor1](https://github.com/Barakmor1) |
| Network interfaces state can be set to `down` or `up` in order to set the link state accordingly when VM is running. Hot plugging of interface in these states is also supported.  | [#13744](https://github.com/kubevirt/kubevirt/pull/13744) | [nirdothan](https://github.com/nirdothan) |
| Interrupted migrations will now be reconciled on next VM start.  | [#13536](https://github.com/kubevirt/kubevirt/pull/13536) | [jean-edouard](https://github.com/jean-edouard) |
| bug-fix: add machine type to `NodeSelector` to prevent breaking changes on unsupported nodes  | [#13690](https://github.com/kubevirt/kubevirt/pull/13690) | [dasionov](https://github.com/dasionov) |
| The node-restriction Validating Admission Policy will return consistent reasons on failures  | [#13940](https://github.com/kubevirt/kubevirt/pull/13940) | [tiraboschi](https://github.com/tiraboschi) |
| Instance type and preference runtime data is now stored under `Status.{Instancetype,Preference}Ref` and is no longer mutated into the core VirtualMachine` `Spec`.  | [#13916](https://github.com/kubevirt/kubevirt/pull/13916) | [lyarwood](https://github.com/lyarwood) |
| VMClone: Remove webhook that checks Snapshot Source  | [#13831](https://github.com/kubevirt/kubevirt/pull/13831) | [ShellyKa13](https://github.com/ShellyKa13) |
| GA ClusterProfiler FG and add a config to enable it  | [#13815](https://github.com/kubevirt/kubevirt/pull/13815) | [acardace](https://github.com/acardace) |
| Updated common-instancetypes bundles to v1.3.0  | [#13928](https://github.com/kubevirt/kubevirt/pull/13928) | [kubevirt-bot](https://github.com/kubevirt-bot) |
| Fetch non-cluster instance type and preferences with namespace key  | [#13805](https://github.com/kubevirt/kubevirt/pull/13805) | [machadovilaca](https://github.com/machadovilaca) |


_This page is updated daily._
