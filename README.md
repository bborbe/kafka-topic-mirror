# kafka-topic-mirror

Copies messages from a source Kafka topic to a target topic (optionally across clusters), then exits when the source is drained. Broker-agnostic.

## Run

```
make run
```

## Flags

| Flag | Env | Description |
|---|---|---|
| `-source-kafka-brokers` | `SOURCE_KAFKA_BROKERS` | source brokers (comma-separated) |
| `-source-kafka-group` | `SOURCE_KAFKA_GROUP` | consumer group |
| `-source-topic` | `SOURCE_TOPIC` | topic to read from |
| `-target-kafka-brokers` | `TARGET_KAFKA_BROKERS` | target brokers (comma-separated) |
| `-target-topic` | `TARGET_TOPIC` | topic to write to |
| `-batch-size` | `BATCH_SIZE` | batch consume size (default 1) |
| `-sentry-dsn` | `SENTRY_DSN` | optional Sentry DSN |

## Build

`make buca` builds and publishes `docker.io/bborbe/kafka-topic-mirror:vX.Y.Z` (git-tag semver).
