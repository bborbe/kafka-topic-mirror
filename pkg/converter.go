// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkg

import (
	"context"

	"github.com/IBM/sarama"
	libkafka "github.com/bborbe/kafka"
)

type Converter interface {
	Convert(ctx context.Context, msg sarama.ConsumerMessage) (*sarama.ProducerMessage, error)
}

type ConverterFunc func(ctx context.Context, msg sarama.ConsumerMessage) (*sarama.ProducerMessage, error)

func (c ConverterFunc) Convert(
	ctx context.Context,
	msg sarama.ConsumerMessage,
) (*sarama.ProducerMessage, error) {
	return c(ctx, msg)
}

func NewConverter(targetTopic libkafka.Topic) Converter {
	return ConverterFunc(
		func(ctx context.Context, msg sarama.ConsumerMessage) (*sarama.ProducerMessage, error) {
			headers := make([]sarama.RecordHeader, len(msg.Headers))
			for i, header := range msg.Headers {
				headers[i] = *header
			}
			return &sarama.ProducerMessage{
				Headers: headers,
				Topic:   targetTopic.String(),
				Key:     sarama.ByteEncoder(msg.Key),
				Value:   sarama.ByteEncoder(msg.Value),
			}, nil
		},
	)
}
