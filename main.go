package main

import (
	"fmt"
	"log"
	"net/http"
	"task-management/api"
	"task-management/app"
	_ "task-management/docs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("default")  // Name of the configuration file without extension
	viper.SetConfigType("yaml")     // Configuration file type
	viper.AddConfigPath(".")        // Path to look for the config file
	viper.AddConfigPath("./config") // Add additional config path if needed

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	log.Println("Configuration loaded successfully")
}

// @title Task management api
// @version 1.0
// @description Task management server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8082
// @BasePath /task-service
func main() {
	// Initialize the App instance with MongoDB connection
	application, err := app.New()
	if err != nil {
		log.Fatalf("Error initializing the app: %v", err)
	}

	// Initialize the API with the app context
	apiInstance, err := api.New(application)
	if err != nil {
		log.Fatalf("Error initializing the API: %v", err)
	}

	// Serve the API
	serveAPI(apiInstance)
}

func serveAPI(api *api.API) {
	cors := handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"http://localhost:8081", "*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "auth-token"}), // Set allowed headers as needed
	)
	router := mux.NewRouter()
	router.Use(cors)

	// Initialize API routes under /task-service
	api.Init(router.PathPrefix("/task-service").Subrouter().StrictSlash(true))

	// Get the port from Viper configuration or default to 8081
	port := viper.GetInt("server.port") // Assuming "server.port" is defined in the config file
	if port == 0 {
		port = 8081 // Fallback to default port
	}

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	log.Printf("Serving API at http://127.0.0.1:%d", port)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
