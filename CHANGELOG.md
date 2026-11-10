# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards compatible manner, and
* PATCH version when you make backwards compatible bug fixes.

## Unreleased

- chore: update github.com/bborbe/kafka to v1.29.13, github.com/bborbe/log to v1.7.8, github.com/bborbe/metrics to v0.6.11, github.com/bborbe/run to v1.11.11, github.com/bborbe/service to v1.11.3, github.com/bborbe/time to v1.27.27, github.com/onsi/ginkgo/v2 to v2.34.0

## v0.2.2

- chore: restate the do-not-merge contract in the gate workflow

## v0.2.1

- fix: keep the rootless BuildKit build container startable. `allowPrivilegeEscalation: false` and `capabilities.drop: ["ALL"]` on that container stop `rootlesskit` from installing its uid/gid map (`newuidmap: operation not permitted`), so every cluster build died at startup; both are removed and the reason is recorded in `hack/build-job.yaml`. Also hardened the clone path: the token is scoped to its origin host, deleted on the error path as well as the success path, and refused over a non-https URL; `REF` can no longer be read as a git option; and `__CONTEXT__` is listed in the template's token table

## v0.2.0

- feat: build the image on a Kubernetes cluster with `make build BUILD_TARGET=<cluster> REF=<sha> VERSION=<tag>`, keeping the default local arm unchanged. `hack/build-cluster.sh`, `hack/cluster-alias.sh` and `hack/build-job.yaml` are vendored from `bborbe/agent-task-executor`; `check-version-tag` and `check-go-mod` stay local-arm prerequisites, since they describe the local HEAD rather than the commit the cluster builds

## v0.1.12

- chore: update github.com/bborbe/errors to v1.6.6, github.com/bborbe/kafka to v1.29.11, github.com/bborbe/log to v1.7.6, github.com/bborbe/metrics to v0.6.9, github.com/bborbe/run to v1.11.10, github.com/bborbe/sentry to v1.10.7, github.com/bborbe/service to v1.11.2, github.com/bborbe/time to v1.27.25

## v0.1.11

- chore: update Go to 1.27.3 and github.com/bborbe/errors to v1.6.3, github.com/bborbe/kafka to v1.28.4, github.com/bborbe/log to v1.7.4, github.com/bborbe/metrics to v0.6.7, github.com/bborbe/run to v1.11.2, github.com/bborbe/sentry to v1.10.5, github.com/bborbe/service to v1.11.0, github.com/bborbe/time to v1.27.17

## v0.1.10

- fix: bump `osv-scanner` to v2.6.0 so the Linux vulnerability gates stop failing. v2.3.1 pins `golang.org/x/tools` v0.38.0, whose SSA builder aborts with `unexpected expr: *ast.KeyValueExpr` on the promoted-field composite-literal key Go 1.27 permits in the Linux stdlib, so a repo on the old pin passes locally on darwin and fails only in Linux CI. Where `golang.org/x/net` is older than v0.60.0 it is raised to v0.60.0, clearing `GO-2026-6603/6610/6611/6612/6617`, which fail both `vulncheck` and `trivy`.
## v0.1.9

- chore: update Go to 1.27.2 and github.com/bborbe/errors to v1.6.2, github.com/bborbe/log to v1.7.3, github.com/bborbe/metrics to v0.6.4, github.com/bborbe/sentry to v1.10.3, github.com/bborbe/time to v1.27.15, github.com/onsi/ginkgo/v2 to v2.33.1

## v0.1.8

- chore: update github.com/IBM/sarama to v1.61.1, github.com/bborbe/kafka to v1.26.0, github.com/bborbe/log to v1.7.1, github.com/bborbe/metrics to v0.6.3, github.com/bborbe/run to v1.11.0, github.com/bborbe/sentry to v1.10.2, github.com/bborbe/service to v1.10.14, github.com/onsi/ginkgo/v2 to v2.33.0, github.com/onsi/gomega to v1.44.0

## v0.1.7

- chore: update Go to 1.27.1 and github.com/bborbe/errors to v1.6.1, github.com/bborbe/kafka to v1.25.15, github.com/bborbe/metrics to v0.6.2, github.com/bborbe/sentry to v1.10.1, github.com/bborbe/service to v1.10.13, github.com/bborbe/time to v1.27.14

## v0.1.6

- fix: `make build` refuses to stamp a version onto a tree that is not that version's tag (`check-version-tag`, escape hatch `ALLOW_UNTAGGED_BUILD=1`). `VERSION` defaults to the newest tag repo-wide, so an operator-run build from an untagged or older tree silently republishes under the newest tag. The guard compares `git describe --exact-match HEAD` against `$(VERSION)` and exits non-zero on mismatch.

## v0.1.5

- ci: add ci.yml running `make precommit` -- the `test` required status check had no workflow producing it, permanently blocking every PR

- chore: update github.com/bborbe/errors to v1.6.0, github.com/bborbe/kafka to v1.25.11, github.com/bborbe/metrics to v0.6.1, github.com/bborbe/run to v1.10.2, github.com/bborbe/sentry to v1.10.0, github.com/bborbe/service to v1.10.10, github.com/bborbe/time to v1.27.11, github.com/onsi/gomega to v1.43.0

## v0.1.4

- chore: update github.com/IBM/sarama to v1.60.2, github.com/bborbe/errors to v1.5.21, github.com/bborbe/log to v1.6.25, github.com/bborbe/metrics to v0.5.15, github.com/bborbe/sentry to v1.9.27

## v0.1.3

- chore: Bump errcheck to v1.20.0 and golangci-lint to v2.13.1 for Go 1.27 support
## v0.1.2

- update Go to 1.26.6 and update dependencies (fixes GO-2026-6179, GO-2026-6180, CVE-2026-56864, CVE-2026-56865, GO-2026-5026, GO-2026-5972, GO-2026-6090, GO-2026-6218)

## v0.1.1

- docs: add a License section to the README

## v0.1.0

- Initial release — extracted from `bborbe/trading` (`strimzi/topic-mirror`) as a standalone public repo
- Copies messages from a source Kafka topic to a target topic (across clusters), broker-agnostic
- Decoupled from `trading/lib`: build-info via public `github.com/bborbe/metrics`, sync-producer via public `github.com/bborbe/kafka`
- Publish-only build → `docker.io/bborbe/kafka-topic-mirror:vX.Y.Z`
