package grafana

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Grafana interface {
	// Dashboards
	CreateDashboard(dashboard *Dashboard, overwrite bool) error
	GetDashboard(slug string) (*Dashboard, error)
	DeleteDashboard(slug string) error
}

var (
	// ErrInvalidResponse is thrown when grafana responds with invalid or error response
	ErrInvalidResponse = errors.New("invalid response from Grafana")
	// ErrDoesNotExist is thrown when the resource does not exists
	ErrDoesNotExist = errors.New("the resource does not exist")
	// ErrInvalidArgument is thrown when invalid argument
	ErrInvalidArgument = errors.New("the argument passed is invalid")
	// ErrTimeoutError is thrown when the operation has timed out
	ErrTimeoutError = errors.New("the operation has timed out")
	// ErrConflict is thrown when the resource is conflicting (ie. app already exists)
	ErrConflict = errors.New("conflicting resource")
	// ErrAuthFailed is thrown when Authication Failed
	ErrAuthFailed = errors.New("Invalid API key")
	// ErrPreconditionFailed is thrown when Precondition Failed
	ErrPreconditionFailed = errors.New("Precondition failed")
)

type grafanaClient struct {
	sync.RWMutex
	// the configuration for the client
	url string
	// API Token
	token string
	// the http client use for making requests
	httpClient *http.Client
}

func NewGrafanaClient(config Config) Grafana {
	service := new(grafanaClient)
	service.url = config.URL
	service.token = config.Token
	service.httpClient = &http.Client{
		Timeout: (60 * time.Second),
	}
	return service
}

func (r *grafanaClient) GetDashboard(slug string) (*Dashboard, error) {
	var result GetDashboardResponse
	uri := grafanaAPIDashboard + "/" + slug
	err := r.apiGet(uri, nil, &result)
	// debugln(result.Meta.Slug)
	// b, err := json.MarshalIndent(result.Dashboard, "", "    ")
	// if err != nil {
	// 	debugln(err)
	// }
	// debugln("\n", string(b))
	if err != nil {
		debugln(err)
		return &result.Dashboard, err
	}
	return &result.Dashboard, nil
}

func (r *grafanaClient) CreateDashboard(dashboard *Dashboard, Overwrite bool) error {
	var request PostDashboardRequest
	request.Dashboard = *dashboard
	request.Overwrite = Overwrite
	var result PostDashboardResponse
	uri := grafanaAPIDashboard
	err := r.apiPost(uri, request, &result)
	if err != nil {
		debugln(err)
		debugf("%+v", result)
		return err
	}
	return nil
}

func (r *grafanaClient) DeleteDashboard(slug string) error {
	type Result struct {
		Title string `json:"title"`
	}
	var result Result
	uri := grafanaAPIDashboard + "/" + slug
	err := r.apiDelete(uri, nil, &result)
	if err != nil {
		debugln(err)
	}
	debugln(result.Title)
	return nil
}

func (r *grafanaClient) apiGet(uri string, post, result interface{}) error {
	return r.apiCall("GET", uri, post, result)
}

func (r *grafanaClient) apiPut(uri string, post, result interface{}) error {
	return r.apiCall("PUT", uri, post, result)
}

func (r *grafanaClient) apiPost(uri string, post, result interface{}) error {
	return r.apiCall("POST", uri, post, result)
}

func (r *grafanaClient) apiDelete(uri string, post, result interface{}) error {
	return r.apiCall("DELETE", uri, post, result)
}

func (r *grafanaClient) apiCall(method, uri string, body, result interface{}) error {
	url := fmt.Sprintf("%s/%s", r.url, uri)
	debugf("[http] request: %s, uri: %s, url: %s", method, uri, url)
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}
	request, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		debugln(err)
		return err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", "Bearer "+r.token)
	response, err := r.httpClient.Do(request)
	if err != nil {
		debugln(err)
		return err
	}
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if len(jsonBody) > 0 {
		debugf("apiCall(): %v %v %s returned %v %s\n", request.Method, request.URL.String(), jsonBody, response.Status, oneLogLine(respBody))
	} else {
		debugf("apiCall(): %v %v returned %v %s\n", request.Method, request.URL.String(), response.Status, oneLogLine(respBody))
	}

	switch {
	case response.StatusCode >= 200 && response.StatusCode <= 299:
		if result != nil {
			if err := json.Unmarshal(respBody, result); err != nil {
				debugf("apiCall(): failed to unmarshall the response from marathon, error: %s\n", err)
				return err
			}
		}
		return nil
	case response.StatusCode == 404:
		return ErrDoesNotExist

	case response.StatusCode == 401:
		return ErrAuthFailed

	case response.StatusCode == 409:
		return ErrConflict

	case response.StatusCode == 412:
		return ErrPreconditionFailed

	case response.StatusCode >= 500:
		return ErrInvalidResponse

	default:
		debugf("apiCall(): unknown error: %s", oneLogLine(respBody))
		return ErrInvalidResponse
	}
}
