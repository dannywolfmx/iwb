//Stream file will perform save and load method from an stream
package stream

import (
	"encoding/gob"
	"io"
)

//Load the "data" from the stream and set it in the pointer data reference
//data need to be a pointer to decode and set the data from the stream
func load(stream io.Reader, data interface{}) error {
	//Generate a new decoder from the stream
	decoder := gob.NewDecoder(stream)
	//Decode the data and return the error (or nil)
	return decoder.Decode(&data)
}

//Save the "data" to the stream using the gob encoder
func save(stream io.Writer, data interface{}) error {
	//Generate a new encoder from the stream
	encoder := gob.NewEncoder(stream)
	//Encode the data and return the error (or nil)
	return encoder.Encode(data)
}
