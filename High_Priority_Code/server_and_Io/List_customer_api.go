package serverandio

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongo_client *mongo.Collection  // from inti() it is connected with db

type Customer struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Address *Address `json:"address,omitempty"`
	Number  int64    `json:"number,omitempty"`
}

type Address struct {
	Addressline1 string `json:"addressline1,omitempty"`
	Addressline2 string `json:"addressline2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          int    `json:"zip,omitempty"`
}

func ListCustomer(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, _ := mongo_client.Find(ctx, bson.M{})

	var result []Customer
	for cur.Next(ctx) {
		var c Customer
		cur.Decode(&c)
		result = append(result, c)
	}

	json.NewEncoder(w).Encode(result)
}

// ----------- MAIN -----------

func main_() {
	// Mongo connect
	log.Println("Server running on :8080")
	http.HandleFunc("/get", ListCustomer)
	http.ListenAndServe(":8080", nil)

	// http.HandleFunc("/create", createCustomer)
	// http.HandleFunc("/delete", deleteCustomer)
	// http.HandleFunc("/update", updateCustomer)

}

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	mongo_client = client.Database("testdb").Collection("customer")
}

// /////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func createCustomer(w http.ResponseWriter, r *http.Request) {
	var c Customer
	json.NewDecoder(r.Body).Decode(&c)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongo_client.InsertOne(ctx, c)

	w.Write([]byte("created"))
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	numberStr := r.URL.Query().Get("number")
	number, _ := strconv.ParseInt(numberStr, 10, 64)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongo_client.DeleteOne(ctx, bson.M{"number": number})

	w.Write([]byte("deleted"))
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var c Customer
	json.NewDecoder(r.Body).Decode(&c)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mongo_client.UpdateOne(ctx, bson.M{"number": c.Number}, bson.M{"$set": bson.M{"name": c.Name, "age": c.Age}})

	w.Write([]byte("updated"))
}
