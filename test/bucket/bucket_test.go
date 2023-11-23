package bucket

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestDownloadInvalidFile(t *testing.T) {
	url := "http://localhost:8080/api/v1/bucket/"
	// Make a GET request

	resp, _ := http.Get(url + "thisImageDoesNotExist")

	// Non existing file should return 500 (internal server error)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestDownloadValidFile(t *testing.T) {
	url := "http://localhost:8080/api/v1/bucket/"
	// Make a GET request

	resp, err := http.Get(url + "test")

	//assert that there is an error
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	//read body and check if it is a string
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.IsType(t, "", string(body))

	//check that the body can be parsed as base64
	_, err = io.ReadAll(base64.NewDecoder(base64.StdEncoding, resp.Body))
	assert.NoError(t, err)

	// Non existing file should return 500 (internal server error)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestListBucket(t *testing.T) {
	url := "http://localhost:8080/api/v1/bucket/"
	// Make a GET request
	resp, err := http.Get(url)

	//assert that there is an error
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	//read body and check if it is a string
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.IsType(t, "", string(body))
	t.Logf(string(body))

	//check that the body can be parsed as base64
	_, err = io.ReadAll(base64.NewDecoder(base64.StdEncoding, resp.Body))
	assert.NoError(t, err)

	// Non existing file should return 500 (internal server error)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
