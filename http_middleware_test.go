package gof_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOkHandler(t *testing.T) {
	ts := httptest.NewServer(MidddlewareToHandler(LoggingMiddleware, http.HandlerFunc(OkHandler)))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	require.NoError(t, err)

	ok, err := io.ReadAll(res.Body)
	require.NoError(t, err)
	require.NoError(t, res.Body.Close())

	assert.Equal(t, "OK", string(ok))
}
