package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	startTime                 time.Time
	suicideHealthAfterSeconds int
)

const (
	suicideHealthAfterSecondsKey = "SUICIDE_HEALTH_AFTER_SECONDS"
)

func main() {
	var err error
	startTime = time.Now()

	// Parse the SUICIDE_HEALTH_AFTER_SECONDS environment variable
	suicideSeconds := os.Getenv(suicideHealthAfterSecondsKey)
	if suicideSeconds == "" {
		panic(fmt.Sprintf("Missing env %s", suicideHealthAfterSecondsKey))
	}
	suicideHealthAfterSeconds, err = strconv.Atoi(suicideSeconds)
	if err != nil {
		log.Fatalf("Invalid %s value: %v", suicideHealthAfterSecondsKey, err)
	}

	http.HandleFunc("/test/health", healthHandler)
	http.HandleFunc("/test/app", appHandler)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if suicideHealthAfterSeconds > 0 && time.Since(startTime).Seconds() > float64(suicideHealthAfterSeconds) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Funny Message</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f0f0f0;
        }
        .message {
            font-size: 24px;
            text-align: center;
            padding: 20px;
            background-color: #ffffff;
            border-radius: 10px;
            box-shadow: 0 0 10px rgba(0,0,0,0.1);
        }
    </style>
</head>
<body>
    <div class="message">
        <p>Why don't scientists trust atoms?</p>
        <p>Because they make up everything! 😄</p>
    </div>
</body>
</html>
`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}
