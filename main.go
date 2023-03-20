package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	// "github.com/gorilla/mux"
)

type User struct{
	Firstname 		string `json:"firstname"`
	Lastname 		string `json:"lastname"`
	Email 			string `json:"email"`
	Cellnumber 		string `json:"cellnumber"`
	ID 				int `json:"id"`
}

// type Account struct{

// }

func handleGetUser(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/user" {
		http.Error(w, "Not found.", 404)
		return 
	}
	if r.Method != "GET" {
		http.Error(w, "Request method not supported.", 404)
        return
	}

	data := User {Firstname: "Lihle", Lastname : "Mbangela", Email : "mbangelalihle@gmail.com", Cellnumber : "+27672966361", ID: 1,}
	
	userData, error := json.Marshal(data);
	if error != nil {
		http.Error(w, error.Error(), 500)
        return
	}
	
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(userData)
	
	// json.NewEncoder(w).Encode(userData)
	// w.Header().Set("Content-Type", "application/json")
	w.Write(userData)
	return
}

func main(){
	//listen here;
	// router := mux.NewRouter()


	serveLandingPage := http.FileServer(http.Dir("./public"))
	http.Handle("/",serveLandingPage)
	// router.Handle("/", serveLandingPage).Methods("GET")
	// http.HandleFunc("/user", handleGetUser)

	http.HandleFunc("/user", handleGetUser)

    fmt.Println("Listening on port 8080")
	if error := http.ListenAndServe(":8080", nil);
	error != nil {
		log.Fatal(error)

	}
}