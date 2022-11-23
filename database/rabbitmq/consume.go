package main

import (
	"encoding/json"
	"gin_es-rabbit/database/elasticsearch"
	"gin_es-rabbit/models"
	elastic "github.com/elastic/go-elasticsearch/v7"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"strings"
)

func RabbitConn() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:Im4dm1nDud3@35.213.138.186:5672/")
	if err != nil {
		log.Println(err, "Fail while connecting to rabbitmq server")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err, "Fail while connecting to channel")
	}
	return ch
}

var ch *amqp.Channel = RabbitConn()

var es *elastic.Client = elasticsearch.ElasticConn()

func main() {
	queue, err := ch.QueueDeclare(
		"books",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	consume, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println(err)
	}

	var forever chan struct{}
	go func() {
		for d := range consume {
			var res models.Book
			_ = json.Unmarshal(d.Body, &res)
			js, _ := json.Marshal(res)
			resu, err := es.Index(
				"books-es-rabbit",
				strings.NewReader(string(js)),
				es.Index.WithDocumentID(res.ID.Hex()),
				es.Index.WithRefresh("true"),
			)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(resu)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
