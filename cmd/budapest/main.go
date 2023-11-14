package main

// "context"
// "fmt"
// "html/template"
// "log"
// "net/http"

// "github.com/gin-gonic/gin"
// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"

func main() {
	// // Initialize the Gin router
	// router := gin.Default()

	// // Connect to MongoDB
	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	// if err != nil {
	// 	log.Fatalf("Failed to connect to MongoDB: %v", err)
	// }
	// defer client.Disconnect(context.Background())

	// // Set up a MongoDB collection
	// collection := client.Database(dbName).Collection(collectionName)

	// // Define a route to serve the HTML page
	// router.GET("/", func(c *gin.Context) {
	// 	// Retrieve HTML content from MongoDB
	// 	var result HTMLContent
	// 	err := collection.FindOne(context.Background(), nil).Decode(&result)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve HTML content"})
	// 		return
	// 	}

	// 	// Parse the HTML template
	// 	tmpl, err := template.New("html").Parse(result.Content)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse HTML template"})
	// 		return
	// 	}

	// 	// Execute the template and write it as an HTML response
	// 	err = tmpl.Execute(c.Writer, nil)
	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render HTML template"})
	// 		return
	// 	}
	// })

	// // Start the Gin server
	// port := 8080 // Update with your desired port number
	// addr := fmt.Sprintf(":%d", port)
	// log.Printf("Server is running on %s", addr)
	// if err := router.Run(addr); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}
