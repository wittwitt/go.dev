package redis

import (
	"fmt"
	"time"
)

type Config struct {
	Name string
}

type Client struct {
	cnn string
}

func NewClient(cfg *Config) (*Client, error) {
	if cfg.Name == "abc" {
		return nil, fmt.Errorf("name not allowed abc")
	}
	fmt.Println("new redis clinet")
	return &Client{cnn: cfg.Name}, nil
}

func (p *Client) SaveName(string) {

}

func (p *Client) GetName(string) string {
	return time.Now().String()
}
