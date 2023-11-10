package rabbit

import (
	"bytes"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

const EXCHANGE_NAME = "logs_rpc"
const QueueToSend = "rpc_queue"
const KeyOfQueueToSend = "rpc_send_key"

type EmailRequest struct {
	From    string
	To      string
	Subject string
	Content string
}

type RabbitMQClient struct {
	Conn *amqp.Channel
}

func Serialize(msg EmailRequest) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}
