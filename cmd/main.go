package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/bradhe/stopwatch"

	"github.com/gorilla/mux"
	"github.com/wowlsh93/hyperledger-fabric-400/hyperledger"
)

type Message struct {
	key   string
	value string
}

// Start & Test
func main() {
	hyperledger.StartFabric()
	watch := stopwatch.Start()

	for i := 0; i < 1000200; i++ {
		hyperledger.WriteTrans(strconv.Itoa(i), strconv.Itoa(i))

		i++
	}

	for {
		result := hyperledger.GetTrans("1000000")
		if result != "" {
			break
		}
	}
	watch.Stop()
	fmt.Printf("seconds elapsed: %v \n", watch.Milliseconds())

	//result1 := hyperledger.GetTrans("1")
	//result2 := hyperledger.GetTrans("2")
	//result3 := hyperledger.GetTrans("3")
	//
	//fmt.Printf("key1 : %s \n", result1)
	//fmt.Printf("key2 : %s \n", result2)
	//fmt.Printf("key3 : %s \n", result3)

	//web_server_run()
}

func web_server_run() error {
	mux := makeMuxRouter()
	httpPort := "8080"
	log.Println("HTTP Server Listening on port :", httpPort)
	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", getLetter).Methods("GET")
	muxRouter.HandleFunc("/", writeLetter).Methods("POST")
	return muxRouter
}

func getLetter(w http.ResponseWriter, r *http.Request) {

	letters := hyperledger.GetTrans("1")
	bytes, err := json.MarshalIndent(letters, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func writeLetter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&msg); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	result := hyperledger.WriteTrans(msg.key, msg.value)
	respondWithJSON(w, r, http.StatusCreated, result)

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
