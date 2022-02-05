package gof_go

import (
	"context"
	"io"
	"net/http"
)

func NewBuilder(url string) HTTPBuilder {
	return &builder{
		headers: map[string][]string{},
		url:     url,
		body:    nil,
		method:  http.MethodGet,
		ctx:     context.Background(),
		close:   false,
	}
}

type HTTPBuilder interface {
	AddHeader(name, value string) HTTPBuilder
	Body(r io.Reader) HTTPBuilder
	Method(method string) HTTPBuilder
	Close(close bool) HTTPBuilder
	Build() (*http.Request, error)
}

type builder struct {
	headers map[string][]string
	url     string
	method  string
	body    io.Reader
	close   bool
	ctx     context.Context
}

func (b *builder) Close(close bool) HTTPBuilder {
	b.close = close

	return b
}

func (b *builder) Method(method string) HTTPBuilder {
	b.method = method

	return b
}

func (b *builder) AddHeader(name, value string) HTTPBuilder {
	values, found := b.headers[name]

	if !found {
		values = make([]string, 0, 10)
	}

	b.headers[name] = append(values, value)

	return b
}

func (b *builder) Body(r io.Reader) HTTPBuilder {
	b.body = r

	return b
}

func (b *builder) Build() (*http.Request, error) {
	r, err := http.NewRequestWithContext(b.ctx, b.method, b.url, b.body)
	if err != nil {
		return nil, err
	}

	for key, values := range b.headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}

	r.Close = b.close

	return r, nil
}
