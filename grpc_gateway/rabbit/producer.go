package rabbit

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (s *RabbitMQClient) SendEmail(email EmailRequest) bool {

	err := s.Conn.ExchangeDeclare(
		EXCHANGE_NAME, // name
		"direct",      // type
		false,         // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return true
	}
	q1, _ := s.Conn.QueueDeclare(
		QueueToSend, // name
		false,       // delete when turn off
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	err = s.Conn.QueueBind(
		q1.Name,          // queue name
		KeyOfQueueToSend, // routing key
		EXCHANGE_NAME,    // exchange
		false,
		nil,
	)
	if err != nil {
		return true
	}

	q2, _ := s.Conn.QueueDeclare( // default exchange must be use with rpc
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)

	corrID := "corrId"
	body, err := Serialize(email)
	if err != nil {
		return true
	}

	err = s.Conn.PublishWithContext(context.Background(),
		EXCHANGE_NAME,    // exchange
		KeyOfQueueToSend, // routing key
		false,            // mandatory
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       q2.Name,
			Body:          body,
		})
	if err != nil {
		return true

	}

	log.Printf(" [x] Sent %s", body)

	msgs, err := s.Conn.Consume(
		q2.Name, // queue
		"",      // consumer
		true,    // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)

	for d := range msgs {
		if corrID == d.CorrelationId {
			res := string(d.Body)
			log.Printf(" [.] Send %v", res)
			if err != nil {
				log.Println(err)
				return false
			}
			break
		}
	}
	return true
}
