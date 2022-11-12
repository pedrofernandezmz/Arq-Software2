package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	logger "github.com/sirupsen/logrus"
	"github.com/stevenferrer/solr-go"
)

type Publi struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Direction   string `json:"direction"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Image       string `json:"image"`
	Image2      string `json:"image2"`
	Seller      string `json:"seller"`
}
type SolrClient struct {
	Client     *solr.JSONClient
	Collection string
}

func NewSolrClient(host string, port int, collection string) *SolrClient {
	logger.Debug(fmt.Sprintf("%s:%d", host, port))
	Client := solr.NewJSONClient(fmt.Sprintf("http://%s:%d", host, port))
	return &SolrClient{
		Client:     Client,
		Collection: collection,
	}
}
func (sc *SolrClient) Update() error {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"COLA2", false, false, false, false, nil,
	)
	failOnError(err, "Failed to declare a queue")
	msgs, error2 := ch.Consume(
		q.Name, "", true, false, false, true, nil)
	failOnError(error2, "Failed to register consumer")

	d := <-msgs
	msg := string(d.Body)
	//solucionar
	url := "http://localhost:8090/items/" + msg
	fmt.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
	json.Marshal(sb)

	json_bytes := []byte(sb)
	var info Publi
	json.Unmarshal(json_bytes, &info)

	//Post en Solr
	collection := "avisos"
	//baseURL := "http://localhost:8983"
	//client := solr.NewJSONClient(baseURL)
	docs := []solr.M{
		{"id": msg, "title": info.Title, "description": info.Description, "direction": info.Direction, "city": info.City, "province": info.Province, "image": info.Image, "image2": info.Image2, "seller": info.Seller},
	}
	buf := &bytes.Buffer{}
	error := json.NewEncoder(buf).Encode(docs)

	ctx := context.Background()

	_, error = sc.Client.Update(ctx, collection, solr.JSON, buf)

	error = sc.Client.Commit(ctx, collection)

	return error
}
