package transcoder

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestTranscode(t *testing.T) {
	url := "http://localhost:8080/api/v1/transcoder/"
	// Make a GET request
	resp, err := http.Get(url + "test")

	//assert that there is an error
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	//read body and check if it is a string
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.IsType(t, "", string(body))
	
	// Non existing file should return 404 (not found)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
