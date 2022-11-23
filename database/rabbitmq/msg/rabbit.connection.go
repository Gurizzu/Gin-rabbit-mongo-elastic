package msg

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitUtil struct {
	srv   string
	qname string
	ctx   context.Context
}

func NewRabbitUtil(qname string) *RabbitUtil {
	return &RabbitUtil{
		srv:   "amqp://guest:Im4dm1nDud3@35.213.138.186:5672/",
		qname: qname,
		ctx:   context.Background(),
	}

}

func (this RabbitUtil) Connect() (client *amqp.Connection, err error) {
	client, err = amqp.Dial(this.srv)
	if err != nil {
		log.Println(err, "Fail while connecting to rabbitmq server")
		return
	}
	return
}

func (this RabbitUtil) Send(data interface{}) (err error) {

	conn, err := this.Connect()
	if err != nil {
		log.Println(err)
	}
	defer func(conn *amqp.Connection) {
		_, err := conn.Channel()
		if err != nil {

		}
	}(conn)

	ch, err := conn.Channel()
	queue, err := ch.QueueDeclare(
		this.qname,
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println(err)
	}

	js, _ := json.Marshal(data)
	err = ch.PublishWithContext(
		this.ctx,
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        js,
		})

	if err != nil {
		log.Println(err)
	}
	log.Printf(" Sent %s\n", js)
	return
}
