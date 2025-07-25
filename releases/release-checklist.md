# KubeVirt Pre-release Checklist

A KubeVirt release entails a variety of tasks to complete in concert with the release tagging.

| Event or time Before Release (BR) | Tasks to initiate                                            | SIGs responsible                           |
|-----------------------------------|--------------------------------------------------------------|--------------------------------------------|
| Beta 0                            | Prepare unconference for next version development cycle      | All SIGs                                   | 
|                                   | Update common-instancetypes bundle in KubeVirt repo          | SIG-compute                                | 
| Feature Freeze                    | Gather scale & performance data                              | SIG-scale                                  |
|                                   | Prepare schedule for subsequent version                      | SIG-release                                |
| After FF                          | Unconference                                                 | All SIGs                                   |
| RC 0                              | Categorise release notes and create User Guide PR with /hold | Docs                                       |
|                                   | Draft release highlights blog                                | Community                                  |
|                                   | - Check for outstanding feature docs PRs                     | Docs and relevant SIGs                     |
| 1 week BR                         | Collect SIG sign-offs on blog                                | SIGs-compute, network, storage, scale, etc |
|                                   | Publish scale & performance data                             | SIG-scale                                  |
| 1 day BR                          | Ensure docs and blog PRs are approved                        | Docs and relevant SIGs                     |
|                                   | Update release notes for any last-minute bugfixes            | Docs & SIG-release                         |
| GA                                | Email to mailing list                                        | SIG-release                                |
|                                   | Publish/unhold release notes and blog                        | Docs and Community                         |
|                                   | Promote through social media                                 | Community                                  |
