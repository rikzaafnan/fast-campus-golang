package repository

import "net/http"

type HttpConnector interface {
	Do(req *http.Request) (*http.Response, error)
}
