package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"premart/painting"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// MetroClient is
type MetroClient struct {
	URL      string
	painting painting.Painting
}

// NewMetroClient is
func NewMetroClient() *MetroClient {
	return &MetroClient{
		URL:      "https://collectionapi.metmuseum.org/public/collection/v1/objects",
		painting: painting.Painting{},
	}
}

func fetchPainting(w http.ResponseWriter, r *http.Request) {

	objID := rand.Intn(370000)
	dist := NewMetroClient()

	resp, err := http.Get(dist.URL + "/" + strconv.Itoa(objID))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// jsonを読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// jsonを構造体へデコード
	if err := json.Unmarshal(body, &dist.painting); err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%+v", dist.painting)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/random", fetchPainting)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
