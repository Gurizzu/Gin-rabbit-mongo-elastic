package elasticsearch

import (
	elastic "github.com/elastic/go-elasticsearch/v7"
	"log"
)

var cfg elastic.Config = elastic.Config{
	Addresses: []string{
		"http://35.213.138.186:9200/",
	},
	Username: "elastic",
	Password: "inipassword",
}

func ElasticConn() (esClient *elastic.Client) {
	esClient, err := elastic.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return
}
