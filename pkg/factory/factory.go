// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package factory

import (
	"github.com/IBM/sarama"
	libkafka "github.com/bborbe/kafka"
	"github.com/bborbe/log"
	"github.com/bborbe/run"

	"github.com/bborbe/kafka-topic-mirror/pkg"
)

func CreateConsumerRun(
	saramaClient sarama.Client,
	sourceTopic libkafka.Topic,
	sourceGroup libkafka.Group,
	targetTopic libkafka.Topic,
	syncProducer libkafka.SyncProducer,
	batchSize libkafka.BatchSize,
	trigger run.Fire,
) run.Func {
	consumer := libkafka.NewOffsetConsumerHighwaterMarksBatch(
		saramaClient,
		sourceTopic,
		libkafka.NewSaramaOffsetManager(
			saramaClient,
			sourceGroup,
			libkafka.OffsetOldest,
			libkafka.OffsetNewest,
		),
		pkg.NewMessageHandler(
			pkg.NewConverter(
				targetTopic,
			),
			syncProducer,
			log.DefaultSamplerFactory,
		),
		batchSize,
		trigger,
		log.DefaultSamplerFactory,
	)
	return consumer.Consume
}
