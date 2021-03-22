//Stream repository can use any kind of stream who use the io.ReadWriter interface,
//for example a file
package stream

import (
	"io"

	"github.com/dannywolfmx/iwb/app/domain/entity"
)

//camera is the struct repositiory
type camera struct {
	cacheViewport entity.Position
	stream        io.Writer
}

//NewCameraRepo will return a camera reposity from the path file
func NewCameraRepo(stream io.ReadWriter) *camera {
	//Generate a new viewport
	viewport := entity.NewPosition(0, 0)
	//try to load data from the stream to get a cache
	if err := load(stream, &viewport); err != nil {
		//Set a default viewport to save a new cache in the stream
		viewport = entity.NewPosition(0, 0)
	}
	//Set the load viewport
	return &camera{
		stream:        stream,
		cacheViewport: viewport,
	}
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
