package gof_go

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"strings"
	"testing"
)

func TestNewBufferedReader(t *testing.T) {
	r := NewBufferedReader(strings.NewReader("this is a line"), 1024)

	b := make([]byte, 10, 10)

	read, err := r.Read(b)
	require.NoError(t, err)
	assert.Equal(t, 10, read)
	assert.Equal(t, "this is a ", string(b[0:read]))

	read, err = r.Read(b)
	require.NoError(t, err)
	assert.Equal(t, 4, read)
	assert.Equal(t, "line", string(b[0:read]))

	read, err = r.Read(b)
	assert.Equal(t, 0, read)
	assert.Equal(t, io.EOF, err)
}
