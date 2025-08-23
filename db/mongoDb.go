package db

import (
	"context"
	"log"

	"github.com/flash-backend/config"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Collection *mongo.Collection
var Database *mongo.Database

func InitMongoDB() {
	log.Println("Initializing MongoDB...")

	client, err := mongo.Connect(options.Client().ApplyURI(config.MONGO_DB_URL))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	log.Println("Connected to MongoDB!")
	Database = client.Database(config.MONGO_DB_DATABASE)
	CreateCollectionAndValidate(Database)
	// log.Println("MongoDB Initialized:", Collection.Name())
}
func CreateCollectionAndValidate(db *mongo.Database) {
	validators := map[string]bson.M{
		"users": {
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"name", "email"},
				"properties": bson.M{
					"name":  bson.M{"bsonType": "string"},
					"email": bson.M{"bsonType": "string", "pattern": "^.+@.+\\..+$"},
					"age":   bson.M{"bsonType": "int", "minimum": 18},
				},
			},
		},
		"products": {
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"title", "price"},
				"properties": bson.M{
					"title": bson.M{"bsonType": "string"},
					"price": bson.M{"bsonType": "double", "minimum": 0},
				},
			},
		},
		"orders": {
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"userId", "productIds"},
				"properties": bson.M{
					"userId":     bson.M{"bsonType": "objectId"},
					"productIds": bson.M{"bsonType": "array"},
					"status":     bson.M{"enum": []string{"pending", "shipped", "delivered"}},
				},
			},
		},
		"payments": {
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"orderId", "amount", "method"},
				"properties": bson.M{
					"orderId": bson.M{"bsonType": "objectId"},
					"amount":  bson.M{"bsonType": "double", "minimum": 0},
					"method":  bson.M{"enum": []string{"card", "bank_transfer", "cash"}},
				},
			},
		},
		"reviews": {
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": []string{"userId", "productId", "rating"},
				"properties": bson.M{
					"userId":    bson.M{"bsonType": "objectId"},
					"productId": bson.M{"bsonType": "objectId"},
					"rating":    bson.M{"bsonType": "int", "minimum": 1, "maximum": 5},
					"comment":   bson.M{"bsonType": "string"},
				},
			},
		},
	}

	// Loop and create collections with validators
	for col, validator := range validators {
		opts := options.CreateCollection().SetValidator(validator)
		err := db.CreateCollection(context.TODO(), col, opts)
		if err != nil {
			// If collection already exists, ignore error
			log.Printf("Collection %s might already exist: %v\n", col, err)
		} else {
			log.Printf("Collection %s created with validator!\n", col)
		}
	}
}
