package main

import "net/http"

type Client struct {
	Request  *http.Request
	Response *http.Response
}

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	c.Request = r
	return c.Response, nil
}
