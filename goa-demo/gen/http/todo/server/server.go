// Code generated by goa v3.19.1, DO NOT EDIT.
//
// todo HTTP server
//
// Command:
// $ goa gen demo/design

package server

import (
	"context"
	todo "demo/gen/todo"
	"net/http"
	"path"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the todo service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	Create             http.Handler
	GenHTTPOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the todo service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *todo.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemGenHTTPOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemGenHTTPOpenapiJSON == nil {
		fileSystemGenHTTPOpenapiJSON = http.Dir(".")
	}
	fileSystemGenHTTPOpenapiJSON = appendPrefix(fileSystemGenHTTPOpenapiJSON, "/gen/http")
	return &Server{
		Mounts: []*MountPoint{
			{"Create", "POST", "/task/create"},
			{"Serve ./gen/http/openapi.json", "GET", "/openapi.json"},
		},
		Create:             NewCreateHandler(e.Create, mux, decoder, encoder, errhandler, formatter),
		GenHTTPOpenapiJSON: http.FileServer(fileSystemGenHTTPOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "todo" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Create = m(s.Create)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return todo.MethodNames[:] }

// Mount configures the mux to serve the todo endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCreateHandler(mux, h.Create)
	MountGenHTTPOpenapiJSON(mux, h.GenHTTPOpenapiJSON)
}

// Mount configures the mux to serve the todo endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountCreateHandler configures the mux to serve the "todo" service "create"
// endpoint.
func MountCreateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/task/create", f)
}

// NewCreateHandler creates a HTTP handler which loads the HTTP request and
// calls the "todo" service "create" endpoint.
func NewCreateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateRequest(mux, decoder)
		encodeResponse = EncodeCreateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create")
		ctx = context.WithValue(ctx, goa.ServiceKey, "todo")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// appendFS is a custom implementation of fs.FS that appends a specified prefix
// to the file paths before delegating the Open call to the underlying fs.FS.
type appendFS struct {
	prefix string
	fs     http.FileSystem
}

// Open opens the named file, appending the prefix to the file path before
// passing it to the underlying fs.FS.
func (s appendFS) Open(name string) (http.File, error) {
	switch name {
	}
	return s.fs.Open(path.Join(s.prefix, name))
}

// appendPrefix returns a new fs.FS that appends the specified prefix to file paths
// before delegating to the provided embed.FS.
func appendPrefix(fsys http.FileSystem, prefix string) http.FileSystem {
	return appendFS{prefix: prefix, fs: fsys}
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}