package gcstorage_test

import (
	"os"
	"testing"

	"github.com/Nivl/go-filestorage/implementations/gcstorage"
	"github.com/Nivl/go-filestorage/testfilestorage"
)

func TestGCStorageUploadHappyPath(t *testing.T) {
	if os.Getenv("TEST_GCP") != "true" {
		t.Skip("Not testing GCP")
	}

	apiKey := os.Getenv("GCP_API_KEY")
	bucket := os.Getenv("GCP_BUCKET")

	storage, err := gcstorage.New(apiKey)
	if err != nil {
		t.Fatal(err)
	}
	storage.SetBucket(bucket)
	testfilestorage.StorageHappyPathTest(t, storage, nil)
	testfilestorage.StorageUnexistingReadTest(t, storage)
}
