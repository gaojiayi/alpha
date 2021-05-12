package common

import (
	"alpha-core/src/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Databae struct {
	Mongo *mongo.Client
}

var DB *Databae

//初始化
func init() {
	DB = &Databae{
		Mongo: getConn(),
	}
}

func getConn() *mongo.Client {
	uri := config.DB_URL
	ctx, cancel := context.WithTimeout(context.Background(), config.DB_WAIT_TIME*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(config.DB_MAX_CONNECTION)) // 连接池
	if err != nil {
		log.Fatal(err)
	}
	return client
}


// TEST
func main() {
	collection := getConn().Database("alpha").Collection("messages")
	collection.Find(context.TODO(),bson.D{})
	//_, err := collection.InsertOne(context.TODO(), bson.M{"name": "pi", "value": 3.14159})
	//if err != nil {
	//	fmt.Println(err)
	//}
}
