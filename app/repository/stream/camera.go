//Stream repository can use any kind of stream who use the io.ReadWriter interface,
//for example a file
//
package stream

import (
	"encoding/gob"
	"io"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

type camera struct {
	cacheViewport entity.Position
	stream        io.Writer
}

//NewCameraRepo will return a camera reposity from the path file
func NewCameraRepo(stream io.ReadWriter) *camera {
	//try to read the stream to save a cache
	cam := &camera{
		stream: stream,
	}

	//Set a default viewport
	viewport := entity.NewPosition(0, 0)

	if err := load(stream, &viewport); err != nil {

		//Set a default viewport to save a new cache in the stream
		cam.SetViewport(entity.NewPosition(0, 0))

		//Return the default cam
		return cam
	}

	//Set the load viewport
	cam.cacheViewport = viewport

	return cam
}

//Viewport
//SetViewport will set the actual position of the camera
func (c *camera) SetViewport(viewport entity.Position) error {
	//Update the cache
	c.cacheViewport = viewport
	//save the new viewport in the stream

	return save(c.stream, &c.cacheViewport)
}

//GetViewport will set the actual position of the camera
//Return a viewport, and chunkLocation
func (c *camera) GetViewport() (entity.Position, error) {
	return c.cacheViewport, nil
}

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
