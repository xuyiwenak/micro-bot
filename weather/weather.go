package weather

import (
	"log"
	"net/http"
)

func GetWeatherDataAPI() {
	request, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		log.Fatalf("http.NewRequest(%q, %q, nil) failed with %v; want success", "GET", "http://www.example.com", err)
	}
}
