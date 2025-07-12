package aziapi

import (
	"encoding/json"
	"strings"

	"github.com/sysnote8main/azisabaapi/pkg/myhttp"
)

type AziApiClient struct {
	httpClient *myhttp.BearerHttpClient
	baseUrl    string
}

func NewAziApiClient(token string, baseUrl string) AziApiClient {
	httpClient := myhttp.NewBearerHttpClient(token)
	return AziApiClient{
		httpClient: &httpClient,
		baseUrl:    strings.TrimRight(baseUrl, "/"),
	}
}

func (c *AziApiClient) getUrl(subPath string) string {
	return c.baseUrl + subPath
}

func (c *AziApiClient) GetCounts() (*CountsResponse, error) {
	b, err := c.httpClient.Get(c.getUrl("/counts"), nil)
	if err != nil {
		return nil, err
	}
	var data CountsResponse
	if err = json.Unmarshal(*b, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
