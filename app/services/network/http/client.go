package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

type Request struct {
	URL     string
	Method  string
	Headers map[string]string
	Data    map[string]string
}

func (request *Request) Config(r Request) *Request {
	request.URL = r.URL
	request.Method = r.Method
	request.Headers = r.Headers
	request.Data = r.Data
	return request
}

func (request *Request) Do(ctx context.Context, dataOutcome interface{}) (*http.Response, error) {
	var resp *http.Response
	var data io.Reader
	sec, _ := time.ParseDuration("10s")
	ctx, cancel := context.WithTimeout(ctx, sec)
	defer cancel()

	if request.Data != nil {
		dataBytes, err := json.Marshal(request.Data)
		if err != nil {
			return resp, err
		}
		data = bytes.NewBuffer(dataBytes)
	}

	req, err := http.NewRequestWithContext(ctx, request.Method, request.URL, data)
	if err != nil {
		return resp, err
	}

	for index, value := range request.Headers {
		req.Header.Add(index, value)
	}

	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}

	if reflect.TypeOf(dataOutcome).Kind() == reflect.Array {
		dataOutcome = body
	} else {
		err = json.Unmarshal(body, &dataOutcome)
		if err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func NewClient() *Request {
	return &Request{}
}
