package test

import (
	"github.com/Educado-App/educado-transcoding-service/api/v1/handlers/bucket"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// MockGCPService is a mock of the GCP service
type MockGCPService struct {
	mock.Mock
}

func (m *MockGCPService) DownloadFile(filename string) ([]byte, error) {
	args := m.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}

func DeleteFileTest(t *testing.T) {

}

func DownloadFileTest(t *testing.T) {
	// Initialize Fiber app
	app := fiber.New()
	app.Get("/:fileName", bucket.DownloadFile)

	// Define test cases
	tests := []struct {
		description    string
		route          string
		expectedStatus int
	}{
		{
			description:    "File exists",
			route:          "/existing-file.txt",
			expectedStatus: fiber.StatusOK,
		},
		{
			description:    "File does not exist",
			route:          "/nonexistent-file.txt",
			expectedStatus: fiber.StatusInternalServerError,
		},
	}

	// Run the tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest("GET", test.route, nil)
			resp, err := app.Test(req, -1)

			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatus, resp.StatusCode)
		})
	}
}

func ListFileTest(t *testing.T) {

}

func UploadFileTest(t *testing.T) {

}
