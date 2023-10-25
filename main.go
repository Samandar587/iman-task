package main

import (

	"time"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)

const secretToken = "L24V5jDsTU"

type DaysLeftResponse struct{	
	DaysLeft int `json:"days_left"`

}

func calculateDays(w http.ResponseWriter, r *http.Request){
	currentTime := time.Now()
	givenTime := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	timeDif := givenTime.Sub(currentTime)

	daysLeft := int(timeDif.Hours()/24)

	res := DaysLeftResponse{DaysLeft: daysLeft}

	jsonRes, err := json.Marshal(res)
	if err != nil{
		http.Error(w, "Failed to marshal Json response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonRes))
}

func tokenAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        token := r.Header.Get("Authorization")

        if token == "" || token != "Bearer "+secretToken {
            http.Error(w, "Token is missing or wrong token is inserted!", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}


func main(){
	
	mux := mux.NewRouter()
	mux.Handle("/days-left", tokenAuthMiddleware(http.HandlerFunc(calculateDays)))

	http.Handle("/", mux)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)

}