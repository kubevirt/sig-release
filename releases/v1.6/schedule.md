# KubeVirt 1.6 Release Schedule

| **When**   | **Milestone**          | **Who**     | **What**                         | **Notes**                                                                                                                               |
|------------|------------------------|-------------|----------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| **March**  |                        |             |                                  |                                                                                                                                         |
| 2025/03/11 |                        | Kubernetes  | Tag 1.33.0-beta.0                |                                                                                                                                         |
| 2025/03/25 |                        | KubeVirt CI | Start the 1.33 provider          | Add on-demand pre-submit                                                                                                                |
| **April**  |                        |             |                                  |                                                                                                                                         |
| 2025/04/23 | Kubernetes Release     | Kubernetes  | 1.33.0 release                   |                                                                                                                                         |
| 2025/04/30 | KubeVirt Alpha 0  | KubeVirt    | Tag v1.6.0-alpha.0               |                                                                                                                                         |
| 2025/04/30 | KubeVirt Alpha 0  | KubeVirt    | Virtualization Enhancement Proposal (VEP) Freeze  |                                                                                                                        |
| **May**    |                        |             |                                  |                                                                                                                                         |
| 2025/05/14 |                        | KubeVirt CI | CI lanes for provider are voting |                                                                                                                                         |
| 2025/05/21 | KubeVirt Beta 0        | KubeVirt    | Tag v1.6.0-beta.0                | Begin [release checklist](../../release-checklist.md)                                                                                   |
| 2025/05/21 | KubeVirt Beta 0        | KubeVirt CI | Mandatory pre-submit             | Non-standard lanes (i.e. migrations, multus, ipv6) are to be bumped by stakeholder                                                      |
| **June**   |                        |             |                                  |                                                                                                                                         |
| 2025/06/25 | KubeVirt Code Freeze   | KubeVirt    | Branch release-1.6               | Review [list of quarantined tests](https://storage.googleapis.com/kubevirt-prow/reports/quarantined-tests/kubevirt/kubevirt/index.html) |
| **July**   |                        |             |                                  |                                                                                                                                         |
| 2025/07/02 | KubeVirt RC 0          | KubeVirt    | Tag v1.6.0-rc.0 on release-1.6   |                                                                                                                                         |
| 2025/07/09 | KubeVirt RC 1          | KubeVirt    | Tag v1.6.0-rc.1 on release-1.6   |                                                                                                                                         |
| 2025/07/16 | **KubeVirt GA**        | KubeVirt    | Tag v1.6.0 on release-1.6        |                                                                                                                                         |
| 2025/07/24 | **CDI GA**             | CDI         | v1.63.0 release                  | [CDI](https://github.com/kubevirt/containerized-data-importer) aims to release approximately a week following KubeVirt                  |

