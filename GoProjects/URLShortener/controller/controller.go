package controller

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
	"github.com/sujalkbarnwal/golang/GoProjects/URLShortener/model"
	"github.com/sujalkbarnwal/golang/GoProjects/URLShortener/utils"
)

const connectionString = "mongodb+srv://<atlas_username>:<atlas_password>@cluster0.x2hpi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "urlobject"
const colName = "shortlongdata"

// Most Important
var collection *mongo.Collection

// Connect with MongoDB
func init() {
	// Client options
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	// Collection instance
	fmt.Println("Collection instance is ready")
}

// MongoDB helpers - file
func insertShortenUrlRecord(urlObj model.UrlObject) {
	_, err := collection.InsertOne(context.TODO(), urlObj)
	if err != nil {
		log.Fatal(err)
	}
}

func checkShortenUrl(shortCode string) (model.UrlObject, error) {
	var urlObj model.UrlObject
	filter := bson.M{"shortcode": shortCode}
	err := collection.FindOne(context.TODO(), filter).Decode(&urlObj)
	return urlObj, err
}

// Controller - file
func GetHomePage(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	// Execute the template and write to the response
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func CreateShortenUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	longUrl := r.FormValue("url")

	// Generate a shortcode
	shortCode := utils.GenerateShortCode(6)

	// Create the UrlObject
	urlObj := model.UrlObject{
		LongUrl:   longUrl,
		ShortCode: shortCode,
	}

	// Insert into MongoDB
	insertShortenUrlRecord(urlObj)

	// Redirect back to the homepage with the shortened URL
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, map[string]interface{}{
		"ShortenedURL": fmt.Sprintf("http://localhost:8080/%s", shortCode),
	})
}

func ReturnOriginalUrl(w http.ResponseWriter, r *http.Request) {
	shortCode := mux.Vars(r)["shortcode"] // Get the shortcode from the URL

	// Check if the shortcode exists
	urlObj, err := checkShortenUrl(shortCode)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Redirect to the original long URL
	http.Redirect(w, r, urlObj.LongUrl, http.StatusSeeOther)
}
