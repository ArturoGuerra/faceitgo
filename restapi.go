package faceitgo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	req, err := http.NewRequest("GET", BASE_URL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	return c.Do(req)
}

func (c *RESTClient) getJSON(path string, v interface{}) error {
	resp, err := c.get(path)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
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
	req, err := http.NewRequest("GET", BASE_URL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.PrivToken)
	return c.Do(req)
}

func (c *RESTClient) postPriv(path string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", BASE_URL+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.PrivToken)
	return c.Do(req)
}
