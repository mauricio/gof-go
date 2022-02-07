package gof_go

import (
	"bytes"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net"
	"testing"
	"time"
)

type mockAddress struct {
	network string
	address string
}

func (m *mockAddress) Network() string {
	return m.network
}

func (m *mockAddress) String() string {
	return m.address
}

// mockConnection represents a connection used for testing only
type mockConnection struct {
	closed  bool
	buffer  *bytes.Buffer
	address *mockAddress
}

func (m *mockConnection) Read(b []byte) (n int, err error) {
	return m.buffer.Read(b)
}

func (m *mockConnection) Write(b []byte) (n int, err error) {
	return m.buffer.Write(b)
}

func (m *mockConnection) LocalAddr() net.Addr {
	return m.address
}

func (m *mockConnection) RemoteAddr() net.Addr {
	return m.address
}

func (m *mockConnection) SetDeadline(t time.Time) error {
	return nil
}

func (m *mockConnection) SetReadDeadline(t time.Time) error {
	return nil
}

func (m *mockConnection) SetWriteDeadline(t time.Time) error {
	return nil
}

func (m *mockConnection) Close() error {
	m.closed = true
	return nil
}

type socketClient struct {
	address string
	factory func(ctx context.Context, network, address string) (net.Conn, error)
}

func (s *socketClient) ping(ctx context.Context) error {
	c, err := s.factory(ctx, "tcp4", s.address)
	if err != nil {
		return err
	}

	defer func() {
		if err := c.Close(); err != nil {
			fmt.Printf("failed to close socket: %v\n", err)
		}
	}()

	if _, err := c.Write([]byte("PING")); err != nil {
		return err
	}

	return nil
}

func TestSocketClient(t *testing.T) {
	connection := &mockConnection{
		buffer: bytes.NewBuffer(make([]byte, 0, 1024)),
	}

	c := &socketClient{
		address: "example.com:40",
		factory: func(ctx context.Context, network, address string) (net.Conn, error) {
			connection.address = &mockAddress{
				address: address,
				network: network,
			}

			return connection, nil
		},
	}

	require.NoError(t, c.ping(context.Background()))
	assert.True(t, connection.closed)
	assert.Equal(t, "PING", connection.buffer.String())
}
