package gof_go

import (
	"errors"
	"io"
)

var (
	_ io.Reader = &bufferedReader{}
)

func NewBufferedReader(wrapped io.Reader, length int) io.Reader {
	if length <= 0 {
		length = 1024
	}

	return &bufferedReader{
		currentIndex:  0,
		lastIndex:     0,
		buffer:        make([]byte, length, length),
		wrappedReader: wrapped,
		err:           nil,
	}
}

type bufferedReader struct {
	currentIndex  int
	lastIndex     int
	buffer        []byte
	wrappedReader io.Reader
	err           error
}

func (b *bufferedReader) Read(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, errors.New("an empty slice was provided to Read")
	}

	availableBytes := b.lastIndex - b.currentIndex

	if availableBytes == 0 {
		if b.err != nil {
			return 0, b.err
		}

		if read, err := b.wrappedReader.Read(b.buffer); err == nil || err == io.EOF {
			b.err = err
			b.currentIndex = 0
			b.lastIndex = read
			availableBytes = read

			if availableBytes == 0 {
				return 0, b.err
			}
		} else {
			b.err = err
			return 0, err
		}
	}

	expectedBytes := len(p)

	bytesToRead := availableBytes
	if availableBytes > expectedBytes {
		bytesToRead = expectedBytes
	}

	copy(p, b.buffer[b.currentIndex:b.currentIndex+bytesToRead])
	b.currentIndex += bytesToRead

	return bytesToRead, b.err
}
