package gof_go

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBuilder_Build(t *testing.T) {

	request, err := NewBuilder("https://example.com/").
		AddHeader("User-Agent", "Golang patterns").
		Build()

	assert.NoError(t, err)

	assert.Equal(t, "Golang patterns", request.Header.Get("User-Agent"))
	assert.Equal(t, http.MethodGet, request.Method)
	assert.Equal(t, "https://example.com/", request.URL.String())
	
}
