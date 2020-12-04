package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var BaseUrl = "http://localhost:8989"

func TestHomePage(test *testing.T) {
	response, err := http.Get(BaseUrl + "/")
	assert.NoError(test, err, "error not nil")
	assert.Equal(test, 200, response.StatusCode, "StatusCode Not Equal")
}
