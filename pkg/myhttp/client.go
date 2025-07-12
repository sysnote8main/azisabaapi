package myhttp

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type BearerHttpClient struct {
	client      *http.Client
	bearerToken string
}

func NewBearerHttpClient(bearerToken string) BearerHttpClient {
	return BearerHttpClient{
		client:      &http.Client{Timeout: 5 * time.Second},
		bearerToken: bearerToken,
	}
}

func (c *BearerHttpClient) do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func (c *BearerHttpClient) addAuth(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+c.bearerToken)
}

func (c *BearerHttpClient) newRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	c.addAuth(req)
	return req, nil
}

func (c *BearerHttpClient) Get(url string, body io.Reader) (*[]byte, error) {
	req, err := c.newRequest(http.MethodGet, url, body)
	if err != nil {
		return nil, err
	}

	res, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is not ok (code:%d)", res.StatusCode)
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (c *BearerHttpClient) SetBearerToken(newToken string) {
	c.bearerToken = newToken
}
