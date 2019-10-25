package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var ErrInvalidRequest = errors.New("invalid request")

type simpleError struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
	Parameter   *string `json:"parameter,omitempty"`
}

func (c *APIClient) errorMessageExtractor(r io.ReadCloser) *string {
	var se simpleError
	if err := json.NewDecoder(r).Decode(&se); err != nil {
		return nil
	}
	if se.Parameter != nil {
		s := fmt.Sprintf("%s: %s", *se.Description, *se.Parameter)
		return &s
	}
	return se.Description
}

func (c *APIClient) errorWrap(r io.ReadCloser) error {
	if errDetails := c.errorMessageExtractor(r); errDetails != nil {
		return errors.New(*errDetails)
	}
	return ErrInvalidRequest
}
