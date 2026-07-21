// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/bborbe/errors"
	libkafka "github.com/bborbe/kafka"
	"github.com/bborbe/log"
	"github.com/golang/glog"
)

func NewMessageHandler(
	converter Converter,
	syncProducer libkafka.SyncProducer,
	logSamplerFactory log.SamplerFactory,
) libkafka.MessageHandlerBatch {
	logSampler := logSamplerFactory.Sampler()
	var counter uint64
	return libkafka.MessageHandlerBatchFunc(
		func(ctx context.Context, messages []*sarama.ConsumerMessage) error {
			producerMessages := make([]*sarama.ProducerMessage, 0, len(messages))
			for _, msg := range messages {
				producerMessage, err := converter.Convert(ctx, *msg)
				if err != nil {
					return errors.Wrapf(ctx, err, "convert failed")
				}
				producerMessages = append(producerMessages, producerMessage)
				counter++
			}
			if err := syncProducer.SendMessages(ctx, producerMessages); err != nil {
				return errors.Wrapf(ctx, err, "send messages failed")
			}
			if logSampler.IsSample() {
				glog.Infof("send messages %d completed", counter)
			}
			return nil
		},
	)
}
