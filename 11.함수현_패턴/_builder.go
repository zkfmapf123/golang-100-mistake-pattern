package main

import (
	"errors"
	"net/http"
)

type Config struct {
	port int
}

/*							첫번째 방법 고전적인 Builder 패턴 적용방법 ✅ 								*/
type ConfigBuilder struct {
	port *int
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port != nil {
		cfg.port = *b.port
	}

	if cfg.port < 0 {
		return Config{}, errors.New("port must be positive")
	}

	if cfg.port == 0 {
		return Config{}, errors.New("port must be non-zero")
	}

	return cfg, nil
}

func NewServerUseBuilder(addr string, builder *ConfigBuilder) (*http.Server, error) {
	cBuilder := ConfigBuilder{}
	cBuilder.Port(8080)
	_, err := cBuilder.Build()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
