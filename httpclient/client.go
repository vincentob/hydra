package httpclient

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/dantin-s/hydra/json"
)

// DoGetJsonWithHeader send http request with method GET and customHeaders
// Parameters:
//   url              string             // Required
//   customerHeaders  map[string]string  // Custom headers that would be added to request header
//   result           interface{}        // You can specify the response data struct, if not nil,
//                                       // will parse the response data to result interface
func DoGetJson(url string, customHeaders map[string]string, result interface{}) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new http request failed")
	}

	for k, v := range customHeaders {
		req.Header.Add(k, v)
	}

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http do failed")
	}

	return ParseResult(resp, result)
}

// DoPostJson do post method with struct data.
// It return raw http response and error.
func DoPostJson(url string, customHeaders map[string]string, body, result interface{}) (resp *http.Response, err error) {
	return DoJson(http.MethodPost, url, customHeaders, body, result)
}

func DoPutJson(url string, customHeaders map[string]string, body, result interface{}) (resp *http.Response, err error) {
	return DoJson(http.MethodPut, url, customHeaders, body, result)
}

func DoJson(method string, url string, customHeaders map[string]string, body, result interface{}) (resp *http.Response, err error) {
	var bodyData string
	switch body.(type) {
	case string:
		logrus.Debug("do DoJson with string body")
		bodyData = body.(string)
	case nil:
		logrus.Debug("do DoJson with nil body")
	default:
		logrus.Debug("do DoJson with struct body")
		bodyData, err = json.FastJJ.MarshalToString(body)
	}

	req, err := http.NewRequest(method, url, bytes.NewBufferString(bodyData))
	if err != nil {
		return nil, errors.Wrap(err, "new http request failed")
	}

	for k, v := range customHeaders {
		req.Header.Add(k, v)
	}
	req.Header.Add("Content-Type", ContentTypeApplicationJson)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http do failed")
	}

	return ParseResult(resp, result)
}

func ParseResult(resp *http.Response, result interface{}) (*http.Response, error) {
	if result == nil {
		return resp, nil
	}

	var err error
	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
	default:
		reader = resp.Body
	}
	defer reader.Close()
	defer resp.Body.Close()

	respBodyData, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "read resp body failed")
	}

	err = json.FastJJ.Unmarshal(respBodyData, result)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal bodyData to result struct failed")
	}

	return resp, nil
}
