// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"os"
	"time"

	"github.com/bborbe/errors"
	libkafka "github.com/bborbe/kafka"
	"github.com/bborbe/metrics"
	"github.com/bborbe/run"
	libsentry "github.com/bborbe/sentry"
	"github.com/bborbe/service"
	libtime "github.com/bborbe/time"
	"github.com/golang/glog"
	_ "github.com/lib/pq"

	"github.com/bborbe/kafka-topic-mirror/pkg/factory"
)

const serviceName = "kafka-topic-mirror"

func main() {
	app := &application{}
	os.Exit(service.Main(context.Background(), app, &app.SentryDSN, &app.SentryProxy))
}

type application struct {
	SentryDSN          string             `required:"true"  arg:"sentry-dsn"           env:"SENTRY_DSN"           usage:"SentryDSN"                                    display:"length"`
	SentryProxy        string             `required:"false" arg:"sentry-proxy"         env:"SENTRY_PROXY"         usage:"Sentry Proxy"`
	SourceKafkaBrokers libkafka.Brokers   `required:"true"  arg:"source-kafka-brokers" env:"SOURCE_KAFKA_BROKERS" usage:"Comma separated list of source Kafka brokers"`
	SourceKafkaGroup   libkafka.Group     `required:"true"  arg:"source-kafka-group"   env:"SOURCE_KAFKA_GROUP"   usage:"kafka group"`
	SourceTopic        string             `required:"true"  arg:"source-topic"         env:"SOURCE_TOPIC"         usage:"topic read from"`
	TargetKafkaBrokers libkafka.Brokers   `required:"true"  arg:"target-kafka-brokers" env:"TARGET_KAFKA_BROKERS" usage:"Comma separated list of target Kafka brokers"`
	TargetTopic        string             `required:"true"  arg:"target-topic"         env:"TARGET_TOPIC"         usage:"topic written to"`
	BatchSize          libkafka.BatchSize `required:"true"  arg:"batch-size"           env:"BATCH_SIZE"           usage:"batch consume size"                                            default:"1"`
	BuildGitVersion    string             `required:"false" arg:"build-git-version"    env:"BUILD_GIT_VERSION"    usage:"Build Git version"                                             default:"dev"`
	BuildGitCommit     string             `required:"false" arg:"build-git-commit"     env:"BUILD_GIT_COMMIT"     usage:"Build Git commit hash"                                         default:"none"`
	BuildDate          *libtime.DateTime  `required:"false" arg:"build-date"           env:"BUILD_DATE"           usage:"Build timestamp (RFC3339)"`
}

func (a *application) Run(ctx context.Context, sentryClient libsentry.Client) error {
	metrics.NewBuildInfoMetrics().SetBuildInfo(a.BuildGitVersion, a.BuildGitCommit, a.BuildDate)

	trigger := run.NewTrigger()

	saramaClientProvider, err := libkafka.NewSaramaClientProviderByType(
		ctx,
		libkafka.SaramaClientProviderTypeReused,
		a.SourceKafkaBrokers,
	)
	if err != nil {
		return errors.Wrapf(ctx, err, "create sarama client failed")
	}
	defer saramaClientProvider.Close()

	rawSyncProducer, err := libkafka.NewSyncProducerWithName(
		ctx,
		a.TargetKafkaBrokers,
		serviceName,
	)
	if err != nil {
		return errors.Wrapf(ctx, err, "create sync producer failed")
	}
	syncProducer := libkafka.NewSyncProducerMetrics(rawSyncProducer)
	defer syncProducer.Close()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
		case <-trigger.Done():
			duration := 10 * time.Second
			glog.V(2).Infof("copy completed => cancel service in %v", duration)
			time.Sleep(duration)
			cancel()
		}
	}()

	saramaClient, err := saramaClientProvider.Client(ctx)
	if err != nil {
		return errors.Wrapf(ctx, err, "create sarama client failed")
	}

	return factory.CreateConsumerRun(
		saramaClient,
		libkafka.Topic(a.SourceTopic),
		libkafka.Group(a.SourceKafkaGroup),
		libkafka.Topic(a.TargetTopic),
		syncProducer,
		a.BatchSize,
		trigger,
	)(ctx)
}
