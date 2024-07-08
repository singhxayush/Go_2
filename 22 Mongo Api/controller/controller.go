package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"myMongoApi/model"
)

const connectionStringKey = "mongodb://localhost:27017/"
const dbName = "netflix"
const collName = "watchlist"

var collection *mongo.Collection

func init() { // init() function is a special function in Go that gets executed automatically when a package is loaded. You use it to perform essential setup tasks before your code can run.
	// In this case, ⁠init() is used to establish the connection to the MongoDB database and set up the collection that your code will interact with.
	// Client option
	clientOptions := options.Client().ApplyURI(connectionStringKey)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB ✅")
	/*
		Docs: https://pkg.go.dev/context#Context

		- In literal terms: the circumstances that form the setting for an event,
		statement, or idea, and in terms of which it can be fully understood.

		- In Go, a context.Context is like a backpack that carries important information about a request or operation.

		- Package context defines the Context type, which carries deadlines, cancellation signals,
		and other request-scoped values across API boundaries and between processes.

		- Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
		The chain of function calls between them must propagate the Context, optionally replacing it with a
		derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue.

		- When a Context is canceled, all Contexts derived from it are also canceled.
	*/

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance is ready")
}

/// MongoDB helpers

// Insert one record
func insertOneMovie(movie model.Netfix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Interted 1 movie in DB with ID:", inserted.InsertedID)
}

// Update one record
func updateOneMovie(movieId string) {
	/// P F U C - primitive, filter, update, collection.

	id, err := primitive.ObjectIDFromHex(movieId) // converts a hexadecimal string representation of a MongoDB ObjectID into an actual primitive.ObjectID
	// whereas primitive.ObjectID is a Go struct "a custom data type" specifically designed to represent MongoDB's ObjectID type.

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}                                               // creates a filter to target the specific document you want to update.
	update := bson.M{"$set": bson.M{"watched": true}}                         // defines the update operation to be performed on the matched document.
	result, err := collection.UpdateOne(context.Background(), filter, update) // executes the update operation on the MongoDB collection.
	// collection is variable representing the MongoDB collection.
	// filter we created earlier, specifying which document to update.
	// update operation we defined, "setting the watched field to true"

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count:", result.ModifiedCount)
	/*
	   /Key differences b/n bson.M and bson.D
	   1.	Ordering: ⁠bson.M doesn't guarantee field order, while ⁠bson.D does.  This matters if you need to maintain a specific field order in your BSON documents, especially for applications like data analysis or database operations that rely on ordered data.
	   2.	Data Structure: ⁠bson.M uses a Go map, which is unordered, while ⁠bson.D uses a slice of ⁠bson.E structures, which enforces order.
	   3.	Performance: ⁠bson.M might be slightly faster for simple document creation as it relies on the underlying Go map implementation. However, ⁠bson.D might offer better performance for larger or complex documents, particularly when you need to access fields by their position or manipulate the document order.

	   /When to use which
	     - bson.M:  Use when:
	       ▪	Field order isn't crucial.
	       ▪	You're building simple documents quickly.
	     - bson.D: Use when:
	       ▪	You require specific field order for BSON documents.
	       ▪	You need to perform operations based on the order of fields.
	       ▪	You're working with larger or complex documents where order might impact
	*/
}

// Delete a record
func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleted, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete count:", deleted.DeletedCount)
}

// Delete all record
func deleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("NUmber of movies delete: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all record
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

/// Actual Controller Section

func CreateOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/w-xxx-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST") // This is the name of the HTTP header being set. It's part of the CORS (Cross-Origin Resource Sharing: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) mechanism.

	var movie model.Netfix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/w-xxx-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/w-xxx-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
