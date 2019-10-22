package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// APIClient определяет транспортный уровень коммуникаций с API.
type APIClient struct {
	HTTP           *http.Client
	APIURL         string
	ShopID, Secret string
}

func (c *APIClient) requestCloser(ctx context.Context, method, uriStr string, body *[]byte, idempKey *string) (*http.Response, error) {
	var request *http.Request

	u := fmt.Sprintf("%s/%s", c.APIURL, uriStr)

	var err error
	if body == nil {
		request, err = http.NewRequest(method, u, nil)
	} else {
		request, err = http.NewRequest(method, u, bytes.NewReader(*body))
	}
	if err != nil {
		return nil, err
	}

	request.SetBasicAuth(c.ShopID, c.Secret)
	request.Header.Add("Content-Type", "application/json")
	if idempKey != nil {
		request.Header.Add("Idempotence-Key", *idempKey)
	}

	return c.HTTP.Do(request.WithContext(ctx))
}

func (c *APIClient) PaymentCreate(ctx context.Context, idempKey string, body *[]byte) (io.ReadCloser, error) {
	response, err := c.requestCloser(ctx, http.MethodPost, "payments", body, &idempKey)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}

func (c *APIClient) PaymentFind(ctx context.Context, paymentId string, body *[]byte) (io.ReadCloser, error) {
	uri := fmt.Sprintf("payments/%s", paymentId)
	response, err := c.requestCloser(ctx, http.MethodGet, uri, body, nil)
	if err != nil {
		return nil, err
	}

	return response.Body, nil
}
