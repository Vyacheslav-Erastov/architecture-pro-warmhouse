package mq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Command struct {
	DeviceID int         `json:"device_id"`
	Type     string      `json:"command_type"`
	Params   interface{} `json:"parameters"`
}

type Producer struct {
	ch *amqp.Channel
}

func NewProducer(conn *amqp.Connection) (*Producer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		"device_commands_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	return &Producer{ch: ch}, nil
}

func (p *Producer) SendCommand(cmd Command) error {
	body, _ := json.Marshal(cmd)
	return p.ch.Publish(
		"", "device_commands_queue", false, false,
		amqp.Publishing{ContentType: "application/json", Body: body},
	)
}
