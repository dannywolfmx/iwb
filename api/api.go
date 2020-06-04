package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"unicode/utf8"

	"danirod.es/pkg/iwb/world"
	"github.com/gorilla/mux"
)

var data *world.World

type chunkUpdatePayload struct {
	X     string
	Y     string
	Value string
}

func getChunk(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	var xc, yc int64
	xc, err = strconv.ParseInt(vars["x"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No way")
		return
	}
	yc, err = strconv.ParseInt(vars["y"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No way")
		return
	}
	fmt.Fprintf(w, "%s", data.GetChunk(int32(xc), int32(yc)).GetRunes())
}

func decodeNumber(str string) (val int32, err error) {
	var intval int64
	if intval, err = strconv.ParseInt(str, 10, 64); err != nil {
		return
	}
	return int32(intval), nil
}

func decodeRune(str string) rune {
	char, _ := utf8.DecodeRuneInString(str[0:])
	return char
}

func decodePayload(r *http.Request) (int32, int32, rune, error) {
	var body []byte
	var x, y, value int32
	var payload chunkUpdatePayload
	var err error

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return 0, 0, 0, err
	}
	if err = json.Unmarshal(body, &payload); err != nil {
		return 0, 0, 0, err
	}
	if x, err = decodeNumber(payload.X); err != nil {
		return 0, 0, 0, err
	}
	if y, err = decodeNumber(payload.Y); err != nil {
		return 0, 0, 0, err
	}
	value = decodeRune(payload.Value)
	return x, y, value, nil
}

func putChunks(w http.ResponseWriter, r *http.Request) {
	x, y, value, err := decodePayload(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	chunk := data.GetChunk(int32(x/256), int32(y/256))
	chunk.SetRune(int32(x%256), int32(y%256), value)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "OK")
}

// CreateHTTPServer returns a http.Handler for creating an HTTP REST API.
func CreateHTTPServer() *mux.Router {
	data = world.NewWorld()
	router := mux.NewRouter()
	router.HandleFunc("/chunks/{x}/{y}", getChunk).Methods("GET")
	router.HandleFunc("/chunks", putChunks).Methods("PUT")
	return router
}
