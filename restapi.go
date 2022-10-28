package faceitgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const BASE_URL = "https://open.faceit.com/data/v4"

type (
	RESTConfig struct {
		Token     string
		PrivToken string
	}
	RESTClient struct {
		*http.Client
		Config RESTConfig
	}
)

func New(cfg *RESTConfig) *RESTClient {
	return &RESTClient{
		Client: &http.Client{},
		Config: *cfg,
	}
}

func (c *RESTClient) get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BASE_URL, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.Token))
	return c.Do(req)
}

func (c *RESTClient) getJSON(path string, v interface{}, query url.Values) error {
	if query != nil {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	resp, err := c.get(path)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

func (c *RESTClient) getPriv(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BASE_URL, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.PrivToken))
	return c.Do(req)
}

func (c *RESTClient) deletePriv(path string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", BASE_URL, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.PrivToken))
	return c.Do(req)
}

func (c *RESTClient) postPriv(path string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", BASE_URL, path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.PrivToken))
	return c.Do(req)
}
