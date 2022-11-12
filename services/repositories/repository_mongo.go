package repositories

import (
	"context"
	"fmt"

	"log"
	"time"

	"github.com/pedrofernandezmz/Arq-Software2/dtos"
	model "github.com/pedrofernandezmz/Arq-Software2/models"

	e "github.com/pedrofernandezmz/Arq-Software2/utils/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryMongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
func NewMongoDB(host string, port int, collection string) *RepositoryMongoDB {
	credential := options.Credential{
		Username: "root",
		Password: "root",
	}
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port)).SetAuth(credential))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	titles, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Sprintf("[MongoDB] Available databases: %s", titles))

	return &RepositoryMongoDB{
		Client:     client,
		Database:   client.Database("avisos"),
		Collection: collection,
	}
}

func (repo *RepositoryMongoDB) Get(id string) (dtos.ItemDTO, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.ItemDTO{}, e.NewBadRequestApiError(fmt.Sprintf("error getting item %s invalid id", id))
	}
	result := repo.Database.Collection(repo.Collection).FindOne(context.TODO(), bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dtos.ItemDTO{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	var item model.Item
	if err := result.Decode(&item); err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}
	return dtos.ItemDTO{
		Id:          id,
		Titulo:      item.Titulo,
		Descripcion: item.Descripcion,
		Direccion:	 item.Direccion,
		Ciudad:      item.Ciudad,
		Provincia:   item.Provincia,
		Imagen:      item.Imagen,
		Imagen2:     item.Imagen2,
		Vendedor:    item.Vendedor,
	}, nil
}

func (repo *RepositoryMongoDB) Insert(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	result, err := repo.Database.Collection(repo.Collection).InsertOne(context.TODO(), model.Item{
		Titulo:      item.Titulo,
		Descripcion: item.Descripcion,
		Direccion:	 item.Direccion,
		Ciudad:      item.Ciudad,
		Provincia:   item.Provincia,
		Imagen:      item.Imagen,
		Imagen2:     item.Imagen2,
		Vendedor:    item.Vendedor,
	})
	item.Id = fmt.Sprintf(result.InsertedID.(primitive.ObjectID).Hex())
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	body := item.Id
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s ", body)
	if err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
	}

	return item, nil
}

func (repo *RepositoryMongoDB) Update(item dtos.ItemDTO) (dtos.ItemDTO, e.ApiError) {
	_, err := repo.Database.Collection(repo.Collection).UpdateByID(context.TODO(), fmt.Sprintf("%v", item.Id), model.Item{
		Titulo:      item.Titulo,
		Descripcion: item.Descripcion,
		Direccion:	 item.Direccion,
		Ciudad:      item.Ciudad,
		Provincia:   item.Provincia,
		Imagen:      item.Imagen,
		Imagen2:     item.Imagen2,
		Vendedor:    item.Vendedor,
	})
	if err != nil {
		return dtos.ItemDTO{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
	}
	return item, nil
}

func (repo *RepositoryMongoDB) Delete(id string) e.ApiError {
	_, err := repo.Database.Collection(repo.Collection).DeleteOne(context.TODO(), bson.M{"_id": fmt.Sprintf("%s", id)})
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting item %s", id), err)
	}
	return nil
}