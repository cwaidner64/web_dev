package main

import (
	//generic imports
	"net/http"
	"log"
	"errors"
	"os"
	"github.com/joho/godotenv"
	//custom imports
	"web/handler"
	"web/middleware"
	"web/utils"


)
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Access-Control-Allow-Origin header to allow requests from your Vue app's origin
		// Replace "http://localhost:5173" with the actual origin of your Vue frontend.
		// For development, you can use "*" to allow all origins, but DO NOT use this in production.
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true") // If you're sending cookies/auth headers

		// Handle preflight requests (OPTIONS method)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var jwtSecret []byte
	// Load environment variables from .env file
	var err error
	err= godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file"+err.Error())
		os.Exit(1)
	}

	db:=utils.GetDB()
	if db==nil{
		log.Println("Error connecting to database"+err.Error())
		os.Exit(1)
	}

	err=os.MkdirAll(os.Getenv("ROOT_DIR"), 0755)
	if err != nil {
		log.Println("Error creating directory"+err.Error())
		os.Exit(1)
	}

	jwtSecret=[]byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		log.Println("JWT_SECRET environment variable is not set")
		os.Exit(1)
	}

	
	//Set up server from environment variables and functions
	port := os.Getenv("PORT")
	server := http.NewServeMux()
	// Test server to see if it is running
	server.HandleFunc("/hello",handler.HelloHandler)  // GET hello statment



	log.Println("Starting server on :"+port+" ...")

	wrappedServer := enableCORS(middleware.RequestLogger(server)) // RequestLogger is already there

	err = http.ListenAndServe(":"+port, wrappedServer)
	if errors.Is(err, http.ErrServerClosed) {
		log.Println("Server closed")
	} else if err != nil {
		log.Println("Error starting server:", err)
		os.Exit(1)
	}
}