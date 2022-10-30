package faceitgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	BASE_DATA_URL    = "https://open.faceit.com/data/v4"
	BASE_CHAT_URL    = "https://open.faceit.com/chat/v1"
	BASE_PRIVATE_URL = "https://api.faceit.com"
)

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

func httpErrorHandler(resp *http.Response) (*http.Response, error) {
	switch resp.StatusCode {
	case http.StatusOK:
		return resp, nil
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized()
	case http.StatusNotFound:
		return nil, ErrNotFound()
	case http.StatusTooManyRequests:
		return nil, ErrTooManyRequests()
	case http.StatusInternalServerError:
		return nil, ErrInternalServerError()
	default:
		return nil, ErrInternalServerError()
	}
}

func New(cfg *RESTConfig) *RESTClient {
	return &RESTClient{
		Client: &http.Client{},
		Config: *cfg,
	}
}

func (c *RESTClient) getJSON(path string, v interface{}, query url.Values) error {
	if query != nil {
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	resp, err := c.get(path)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
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

func (c *RESTClient) get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BASE_DATA_URL, path), nil)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.Token))

	resp, err := c.Do(req)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}

	return httpErrorHandler(resp)
}

func (c *RESTClient) getChat(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BASE_CHAT_URL, path), nil)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.Token))
	resp, err := c.Do(req)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}

	return httpErrorHandler(resp)

}

func (c *RESTClient) postChat(path string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", BASE_CHAT_URL, path), bytes.NewBuffer(body))
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.Token))
	resp, err := c.Do(req)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}

	return httpErrorHandler(resp)
}

func (c *RESTClient) getPriv(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", BASE_PRIVATE_URL, path), nil)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.PrivToken))
	resp, err := c.Do(req)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}

	return httpErrorHandler(resp)
}

func (c *RESTClient) deletePriv(path string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", BASE_PRIVATE_URL, path), nil)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.PrivToken))
	resp, err := c.Do(req)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}

	return httpErrorHandler(resp)
}

func (c *RESTClient) postPriv(path string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", BASE_PRIVATE_URL, path), bytes.NewBuffer(body))
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.PrivToken))
	resp, err := c.Do(req)
	if err != nil {
		return nil, ErrInternalServerError().Wrap(err)
	}

	return httpErrorHandler(resp)
}
