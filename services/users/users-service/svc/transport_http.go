// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: d5b3153b9f
// Version Date: Thu Jul 27 18:20:46 UTC 2017

package svc

// This file provides server-side bindings for the HTTP transport.
// It utilizes the transport/http.Server.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	// This service
	pb "github.com/myambition/ambition/services/users/users-service"
)

const contentType = "application/json; charset=utf-8"

var (
	_ = fmt.Sprint
	_ = bytes.Compare
	_ = strconv.Atoi
	_ = httptransport.NewServer
	_ = ioutil.NopCloser
	_ = pb.RegisterUsersServer
	_ = io.Copy
	_ = errors.Wrap
)

// MakeHTTPHandler returns a handler that makes a set of endpoints available
// on predefined paths.
func MakeHTTPHandler(ctx context.Context, endpoints Endpoints, logger log.Logger) http.Handler {
	m := mux.NewRouter()

	return m
}

// ErrorEncoder writes the error to the ResponseWriter, by default a content
// type of application/json, a body of json with key "error" and the value
// error.Error(), and a status code of 500. If the error implements Headerer,
// the provided headers will be applied to the response. If the error
// implements json.Marshaler, and the marshaling succeeds, the JSON encoded
// form of the error will be used. If the error implements StatusCoder, the
// provided StatusCode will be used instead of 500.
func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	body, _ := json.Marshal(errorWrapper{Error: err.Error()})
	if marshaler, ok := err.(json.Marshaler); ok {
		if jsonBody, marshalErr := marshaler.MarshalJSON(); marshalErr == nil {
			body = jsonBody
		}
	}
	w.Header().Set("Content-Type", contentType)
	if headerer, ok := err.(httptransport.Headerer); ok {
		for k := range headerer.Headers() {
			w.Header().Set(k, headerer.Headers().Get(k))
		}
	}
	code := http.StatusInternalServerError
	if sc, ok := err.(httptransport.StatusCoder); ok {
		code = sc.StatusCode()
	}
	w.WriteHeader(code)
	w.Write(body)
}

type errorWrapper struct {
	Error string `json:"error"`
}

// httpError satisfies the Headerer and StatusCoder interfaces in
// package github.com/go-kit/kit/transport/http.
type httpError struct {
	error
	statusCode int
	headers    map[string][]string
}

func (h httpError) StatusCode() int {
	return h.statusCode
}

func (h httpError) Headers() http.Header {
	return h.headers
}

// Server Decode

// EncodeHTTPGenericResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeHTTPGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	return encoder.Encode(response)
}

// Helper functions

func headersToContext(ctx context.Context, r *http.Request) context.Context {
	for k, _ := range r.Header {
		// The key is added both in http format (k) which has had
		// http.CanonicalHeaderKey called on it in transport as well as the
		// strings.ToLower which is the grpc metadata format of the key so
		// that it can be accessed in either format
		ctx = context.WithValue(ctx, k, r.Header.Get(k))
		ctx = context.WithValue(ctx, strings.ToLower(k), r.Header.Get(k))
	}

	return ctx
}

func HTTPDecodeLogger(next httptransport.DecodeRequestFunc, logger log.Logger) httptransport.DecodeRequestFunc {
	return func(ctx context.Context, r *http.Request) (interface{}, error) {
		logger.Log("method", r.Method, "url", r.URL.String())
		rv, err := next(ctx, r)
		if err != nil {
			logger.Log("method", r.Method, "url", r.URL.String(), "err", err)
		}
		return rv, err
	}
}
