//Stream file will perform save and load method from an stream
package stream

import (
	"encoding/gob"
	"io"
)

//Load the "data" from the stream and set it in the pointer data reference
//data need to be a pointer to decode and set the data from the stream
func load(stream io.Reader, data interface{}) error {
	decoder := gob.NewDecoder(stream)

	err := decoder.Decode(&data)
	//we don't need a if becouse the client will check the returned error
	return err
}

//Save the "data" to the stream using the gob encoder
func save(stream io.Writer, data interface{}) error {
	encoder := gob.NewEncoder(stream)
	return encoder.Encode(data)
}
