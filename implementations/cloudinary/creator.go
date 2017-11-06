package cloudinary

import (
	"github.com/Nivl/go-filestorage"
)

// Makes sure Creator is a logger.Creator
var _ filestorage.Creator = (*Creator)(nil)

// NewCreator returns a filestorage creator that will use the provided keys
// to create a new cloudinary driver for each single logger
func NewCreator(apiKey, secret, defaultBucket string) *Creator {
	return &Creator{
		apiKey:        apiKey,
		secret:        secret,
		defaultBucket: defaultBucket,
	}
}

// Creator creates new filestorage
type Creator struct {
	apiKey        string
	secret        string
	defaultBucket string
}

// New returns a new le client
func (c *Creator) New() (filestorage.FileStorage, error) {
	fs := New(c.apiKey, c.secret)
	return fs, fs.SetBucket(c.defaultBucket)
}
