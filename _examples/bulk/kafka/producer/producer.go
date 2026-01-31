// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package producer

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	sides    = []string{"BUY", "SELL"}
	symbols  = []string{"KBCU", "KBCU", "KBCU", "KJPR", "KJPR", "KSJD", "KXCV", "WRHV", "WTJB", "WMLU"}
	accounts = []string{"ABC123", "ABC123", "ABC123", "LMN456", "LMN456", "STU789"}
)

func init() {
	kafka.DefaultClientID = "go-elasticsearch-kafka-demo"
}

type Producer struct {
	BrokerURL   string
	TopicName   string
	TopicParts  int
	MessageRate int

	writer *kafka.Writer

	startTime     time.Time
	totalMessages int64
	totalErrors   int64
	totalBytes    int64
}

func (p *Producer) Run(ctx context.Context) error {
	var messages []kafka.Message
	p.startTime = time.Now()

	p.writer = kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{p.BrokerURL},
		Topic:   p.TopicName,
	})
	defer p.writer.Close()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for t := range ticker.C {
		for i := 1; i <= p.MessageRate; i++ {
			messages = append(messages, kafka.Message{Value: p.generateMessage(t)})
		}
		if err := p.writer.WriteMessages(ctx, messages...); err != nil {
			messages = messages[:0]
			return err
		}
		messages = messages[:0]
	}

	return nil
}

func (p *Producer) CreateTopic(ctx context.Context) error {
	conn, err := net.Dial("tcp", p.BrokerURL)
	if err != nil {
		return err
	}

	return kafka.NewConn(conn, "", 0).CreateTopics(
		kafka.TopicConfig{
			Topic:             p.TopicName,
			NumPartitions:     p.TopicParts,
			ReplicationFactor: 1,
		})
}

func (p *Producer) generateMessage(t time.Time) []byte {
	var (
		buf bytes.Buffer

		timestamp time.Time
		timeshift time.Duration
		side      string
		quantity  int
		price     int
		amount    int
		symbol    string
		account   string
	)

	timestamp = t
	if timestamp.Minute() == 2 {
		timeshift = -(time.Duration(time.Minute))
	} else {
		timeshift = time.Duration(time.Duration(rand.ExpFloat64()/2.0*100) * time.Second)
	}
	switch {
	case timestamp.Minute()%5 == 0:
		side = "SELL"
	default:
		if timestamp.Second()%3 == 0 {
			side = "SELL"
		} else {
			side = "BUY"
		}
	}
	switch {
	case timestamp.Minute()%3 == 0:
		quantity = rand.Intn(250) + 500
	case timestamp.Second()%6 == 0:
		quantity = rand.Intn(300) + 50
	case timestamp.Second()%12 == 0:
		quantity = rand.Intn(10) + 1
	default:
		quantity = rand.Intn(150) + 10
	}
	if side == "SELL" {
		price = int(100.0 + 15.0*rand.NormFloat64())
	} else {
		price = int(250.0 + 50.0*rand.NormFloat64())
	}
	amount = quantity * price
	if timestamp.Second()%4 == 0 {
		symbol = "KXCV"
	} else {
		symbol = symbols[rand.Intn(len(symbols))]
	}
	if timestamp.Minute()%5 == 0 && timestamp.Second() > 30 {
		account = "STU789"
	} else {
		account = accounts[rand.Intn(len(accounts))]
	}

	fmt.Fprintf(&buf,
		`{"time":"%s", "symbol":"%s", "side":"%s", "quantity":%d, "price":%d, "amount":%d, "account":"%s"}`,
		timestamp.UTC().Add(timeshift).Format(time.RFC3339),
		symbol,
		side,
		quantity,
		price,
		amount,
		account,
	)
	return buf.Bytes()
}

type Stats struct {
	Duration      time.Duration
	TotalMessages int64
	TotalErrors   int64
	TotalBytes    int64
	Throughput    float64
}

func (p *Producer) Stats() Stats {
	if p.writer == nil {
		return Stats{}
	}

	duration := time.Since(p.startTime)
	writerStats := p.writer.Stats()

	p.totalMessages += writerStats.Messages
	p.totalErrors += writerStats.Errors
	p.totalBytes += writerStats.Bytes

	rate := float64(p.totalMessages) / duration.Seconds()

	return Stats{
		Duration:      duration,
		TotalMessages: p.totalMessages,
		TotalErrors:   p.totalErrors,
		TotalBytes:    p.totalBytes,
		Throughput:    rate,
	}
}
