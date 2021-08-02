package main

import (
	"encoding/json"
	"example.com/go-swagger-test/entity"
	"example.com/go-swagger-test/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	_ "example.com/go-swagger-test/docs" // necessary
)

//go:generate swagger generate spec -o ./swagger.json --scan-models

var entityService service.EntityService

func init() {
	entityService = service.NewEntityService()
}

func main() {
	initServer()
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	b := string(bytes)
	fmt.Printf("Body: %s", b)
	_, err := fmt.Fprintf(w, "got value: %d", rand.Int())
	fmt.Printf("got err: %v", err)
}

func initServer() {
	// init main server
	mainRouter := mux.NewRouter()
	router := mainRouter.PathPrefix("/api/v1/entity-service").Subrouter()
	router.Methods("GET").Path("/entities").HandlerFunc(getAllHandler)
	router.Methods("GET").Path("/entities/{id}").HandlerFunc(getByIdHandler)
	router.Methods("POST").Path("/entities").HandlerFunc(create)
	router.Methods("PUT").Path("/entities/{id}").HandlerFunc(update)

	// init swagger server
	swaggerHandler := http.StripPrefix(
		"/swaggerui/",
		http.FileServer(http.Dir("./swaggerui")),
	)
	mainRouter.PathPrefix("/swaggerui").Handler(swaggerHandler)

	// init server with cors handling
	corsHandler := cors.Default().Handler(mainRouter)
	srv := &http.Server{
		Handler: corsHandler,
		Addr:    "127.0.0.1:8082",
	}

	log.Fatal(srv.ListenAndServe())
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	toJson(entityService.GetAll(), w)
}

func getByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	qVal := r.URL.Query().Get("searchVal")

	toJson(entityService.GetById(id, qVal), w)
}

func create(w http.ResponseWriter, r *http.Request) {
	reqBody := &entity.EntityRequest{}
	json.NewDecoder(r.Body).Decode(reqBody)

	toJson(entityService.Create(reqBody), w)
}

func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody := &entity.EntityRequest{}
	json.NewDecoder(r.Body).Decode(reqBody)

	toJson(entityService.Update(id, reqBody), w)
}

// --- serialization ---
func toJson(v interface{}, w http.ResponseWriter) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	w.Header().Set("Content-Type", "application/json")

	encoder.Encode(v)
}
