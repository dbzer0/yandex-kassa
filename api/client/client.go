package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const ua = "yandex-kassa-go-api/1.0"

var ErrInvalidRequest = errors.New("invalid request")

// APIClient определяет транспортный уровень коммуникаций с API.
type APIClient struct {
	HTTP           *http.Client
	APIURL         string
	ShopID, Secret string
}

func (c *APIClient) get(ctx context.Context, uri string) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.APIURL, uri), nil)
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(c.ShopID, c.Secret)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("User-Agent", ua)

	return c.HTTP.Do(request.WithContext(ctx))
}

func (c *APIClient) post(ctx context.Context, uri string, idempKey string, body []byte) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", c.APIURL, uri), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Idempotence-Key", idempKey)
	request.SetBasicAuth(c.ShopID, c.Secret)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Set("User-Agent", ua)

	return c.HTTP.Do(request.WithContext(ctx))
}

func (c *APIClient) Create(ctx context.Context, idempKey string, body *[]byte) (io.ReadCloser, error) {
	response, err := c.post(ctx, "payments", idempKey, *body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return response.Body, ErrInvalidRequest
	}
	return response.Body, nil
}

func (c *APIClient) Find(ctx context.Context, paymentId string) (io.ReadCloser, error) {
	response, err := c.get(ctx, fmt.Sprintf("payments/%s", paymentId))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return response.Body, ErrInvalidRequest
	}
	return response.Body, nil
}

func (c *APIClient) Cancel(ctx context.Context, idempKey string, paymentId string) (io.ReadCloser, error) {
	response, err := c.post(ctx, fmt.Sprintf("payments/%s/cancel", paymentId), idempKey, []byte("{}"))
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
