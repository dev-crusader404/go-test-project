package client

import (
	"net/http"
	"time"
)

var client *rClient

func init() {
	client = &rClient{
		c: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns: 10,
			},
			Timeout: 60 * time.Second,
		},
	}
}

type rClient struct {
	c *http.Client
}

type RestClient interface {
	Do(*http.Request) (*http.Response, error)
}

func (cl *rClient) Do(r *http.Request) (*http.Response, error) {
	return cl.c.Do(r)
}

func GetClient() *rClient {
	return client
}
