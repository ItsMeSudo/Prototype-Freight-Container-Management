package main

import (
	"backend/initFlag"
	"backend/restApiV1"
	"backend/restApiV2"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {
	initFlag.InitFlag()

	fmt.Println("REST API Port:", initFlag.RestAPIPort)
	fmt.Println("REST API CORS Mode:", initFlag.RestAPICorsMode)
	fmt.Println("JSON Server Host:", initFlag.JsonServerHost)
	fmt.Println("JSON Server Port:", initFlag.JsonServerPort)
	fmt.Println("Frontend Port:", initFlag.FrontendPort)
	fmt.Println("Standalone Mode:", initFlag.Standalone)

	initFlag.JsonServerFullPath = "http://" + initFlag.JsonServerHost + ":" + initFlag.JsonServerPort

	go GoGarbageCollector()
	if initFlag.Standalone == 0 {
		go StartFrontendHost()
		StartRestApiRouter()
	} else if initFlag.Standalone == 1 {
		StartFrontendHost()
	} else if initFlag.Standalone == 2 {
		StartRestApiRouter()
	} else {
		log.Println("Only 0, 1, 2 accepted")
	}
}

func StartRestApiRouter() {
	log.Println("RestApi Started on port: " + initFlag.RestAPIPort)
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "content-type", "Authorization", "content", "accept", "Accept"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS "})
	origins := handlers.AllowedOrigins([]string{initFlag.RestAPICorsMode})

	router.HandleFunc("/api/v2/getall", restApiV2.GetAllContainers).Methods("POST")
	router.HandleFunc("/api/v2/importcsv", restApiV2.HandleCSVFormUpload).Methods("POST")
	router.HandleFunc("/api/v2/getstat", restApiV2.GetBlockInfoPost).Methods("POST")
	router.HandleFunc("/api/v2/importjson", restApiV2.InsertJsonFileData).Methods("POST")

	router.HandleFunc("/api/containers", restApiV1.InsertJsonData).Methods("POST")
	router.HandleFunc("/api/containers/import", restApiV1.HandleCSVBinaryUpload).Methods("POST")
	router.HandleFunc("/api/blocks/stat", restApiV1.GetBlockInfoPost).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+initFlag.RestAPIPort, handlers.CORS(headers, methods, origins)(router)))

}

func StartFrontendHost() {
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)
	http.HandleFunc("/blockstat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})

	http.HandleFunc("/import", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})
	log.Println("Frontend Started on port: " + initFlag.FrontendPort)
	log.Fatal(http.ListenAndServe(":"+initFlag.FrontendPort, nil))
}

func GoGarbageCollector() {
	for {
		time.Sleep(500 * time.Second)
		runtime.GC()
	}
}
