package fsstorage

import (
	"github.com/Nivl/go-filestorage"
)

// Makes sure Creator is a logger.Creator
var _ filestorage.Creator = (*Creator)(nil)

// NewCreator returns a filestorage creator that will use the provided keys
// to create a new cloudinary driver for each single logger
func NewCreator(defaultBucket string) *Creator {
	return &Creator{
		defaultBucket: defaultBucket,
	}
}

// Creator creates new filestorage
type Creator struct {
	defaultBucket string
}

// New returns a new le client
func (c *Creator) New() (filestorage.FileStorage, error) {
	fs, err := New()
	if err != nil {
		return nil, err
	}
	return fs, fs.SetBucket(c.defaultBucket)
}
