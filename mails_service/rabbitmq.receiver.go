package main

import (
	"bytes"
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"mailer-service/api"
	"mailer-service/mailer"
	"os"
)

type EmailRequest struct {
	From    string
	To      string
	Subject string
	Content string
}

func ConnectRBM() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5902/")
	if err != nil {
		return nil
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil
	}
	return ch
}

func Deserialize(b []byte) (EmailRequest, error) {
	var msg EmailRequest
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}

const EXCHANGE_NAME1 = "logs_rpc"
const QueueToSend1 = "rpc_queue"
const KeyOfQueueToSend1 = "rpc_send_key"

func StartRabbitConsumer() {
	ch := ConnectRBM() // connect rabbitmq

	ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	q1, _ := ch.QueueDeclare(
		QueueToSend1, // name
		false,        // delete when turn off
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)

	msgs, _ := ch.Consume(
		q1.Name, // queue
		"",      // consumer
		false,   // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)

	var forever chan struct{}

	go func() {
		for d := range msgs {

			response, _ := Deserialize(d.Body)
			SendMailToSMTPServer(response)
			log.Printf(" [.] request from sender is %v", response)
			log.Printf(response.Content)

			err := ch.PublishWithContext(context.Background(),
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte("SUCCESS"),
				})
			if err != nil {
				return
			}

			err = d.Ack(false)
			if err != nil {
				return
			}
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func SendMailToSMTPServer(emailRec EmailRequest) {
	m := mailer.Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        1025,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("ENCRYPTION"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
		FromName:    os.Getenv("FORM_NAME"),
	}

	server, err := api.NewServer(m)
	if err != nil {
		log.Fatal(err)
	}

	err = server.SendEmailInternal(emailRec.From, emailRec.To, emailRec.Subject, emailRec.Content)
	if err != nil {
		return
	}
}
