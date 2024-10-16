package http

import (
	"io"
	gohttp "net/http"
	"strings"
)

type NoopClient struct{}

func (c NoopClient) Do(*gohttp.Request) (*gohttp.Response, error) {

	return &gohttp.Response{
		Status:     "200",
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader("")),
	}, nil
}
