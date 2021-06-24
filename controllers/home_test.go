package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHome(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetHome)
	handler.ServeHTTP(recorder, req)

	require.Equal(t, recorder.Code, http.StatusOK)
	require.Equal(t, recorder.Body.String(), "Hello, World!")
}
