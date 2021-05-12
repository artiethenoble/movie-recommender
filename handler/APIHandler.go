package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"movie-recommender/models"
	"net/http"
)

func Suggest(w http.ResponseWriter, r *http.Request) {

	input, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiInput := models.Request{}

	log.Println("in: ", string(input))

	err = json.Unmarshal(input, &apiInput)
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error unmarshalling into input struct"))
		return
	}

	log.Println("input: ", apiInput)

    // TODO: call suggestor and get moview

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successful Response"))

	return
}
