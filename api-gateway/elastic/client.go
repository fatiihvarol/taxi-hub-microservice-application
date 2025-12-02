package elastic

import (
	"log"
	"os"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

var ES *elasticsearch.Client
var ElasticIndex string

func InitElastic() {
	ElasticIndex = os.Getenv("ELASTIC_INDEX")
	if ElasticIndex == "" {
		ElasticIndex = "gateway-logs"
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_URL"),
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Elasticsearch connection error: %v", err)
	}

	ES = client

	log.Println("[Elastic] connected successfully")
}
