package main

import (
	"errors"
	"net/http"
)

type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port must be positive")
		}

		options.port = &port
		return nil
	}
}

func NewServer(addr string, opts ...Option) (*http.Server, error) {
	options := &options{}

	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}

		opt(options)
	}

	return nil, nil
}
