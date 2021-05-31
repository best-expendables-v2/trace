package trace

import (
	"net/http"
	"testing"
)

func TestRequestIDToHeader(t *testing.T) {
	header := http.Header{}
	RequestIDToHeader(header, "REQUEST-ID")
	if requestID := header.Get(headerKey); requestID != "REQUEST-ID" {
		t.Fatal()
	}
}

func TestRequestIDFromHeader(t *testing.T) {
	header := http.Header{}
	header.Set(headerKey, "REQUEST-ID")
	if requestID := RequestIDFromHeader(header); requestID != "REQUEST-ID" {
		t.Fatal()
	}
}
