package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/dannywolfmx/iwb/world"
	"github.com/dannywolfmx/iwb/world/file"
	"github.com/gorilla/mux"
)

var data world.PersistantWorld

type chunkUpdatePayload struct {
	X     string
	Y     string
	Value string
}

func getChunk(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	//ParseUint will return a uint64 we will need to convert to uint8,
	xc, err := strconv.ParseUint(vars["x"], 10, 8)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No way")
		return
	}

	//ParseUint will return a uint64 we will need to convert to uint8,
	yc, err := strconv.ParseUint(vars["y"], 10, 8)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No way")
		return
	}

	position := world.Position{X: uint8(xc), Y: uint8(yc)}
	fmt.Fprintf(w, "%v", data.GetChunk(position))
}

func decodeNumber(str string) (val uint8, err error) {
	var intval int64
	if intval, err = strconv.ParseInt(str, 10, 8); err != nil {
		return
	}
	return uint8(intval), nil
}

func decodeRune(str string) rune {
	char, _ := utf8.DecodeRuneInString(str[0:])
	return char
}

func decodePayload(r *http.Request) (world.Position, world.Element, error) {
	var body []byte
	var x, y uint8
	//Use a space like a rune
	element := ' '
	var payload chunkUpdatePayload
	var err error

	position := world.Position{X: x, Y: y}

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		return position, element, err
	}
	if err = json.Unmarshal(body, &payload); err != nil {
		return position, element, err
	}
	if x, err = decodeNumber(payload.X); err != nil {
		return position, element, err
	}
	if y, err = decodeNumber(payload.Y); err != nil {
		return position, element, err
	}
	element = decodeRune(payload.Value)

	position = world.Position{X: x, Y: y}

	return position, element, nil
}

func putChunks(w http.ResponseWriter, r *http.Request) {
	position, element, err := decodePayload(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	//TODO GET A CHUNK LOCATION FROM THE CLIENT
	chunk := data.GetChunk(position)
	chunk.SetElement(position, element)
	err = data.Persist()
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintf(w, "OK")
}

// CreateHTTPServer returns a http.Handler for creating an HTTP REST API.
func CreateHTTPServer() *mux.Router {
	data = file.NewFileWorld()
	router := mux.NewRouter()
	router.HandleFunc("/chunks/{x}/{y}", getChunk).Methods("GET")
	router.HandleFunc("/chunks", putChunks).Methods("PUT")
	return router
}
