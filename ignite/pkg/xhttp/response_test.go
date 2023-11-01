package xhttp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	errors "github.com/ignite/cli/ignite/pkg/xerrors"
)

func TestResponseJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]interface{}{"a": 1}
	require.NoError(t, ResponseJSON(w, http.StatusCreated, data))
	resp := w.Result()

	require.Equal(t, http.StatusCreated, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	body, _ := io.ReadAll(resp.Body)
	dataJSON, _ := json.Marshal(data)
	require.Equal(t, dataJSON, body)
}

func TestNewErrorResponse(t *testing.T) {
	require.Equal(t, ErrorResponseBody{
		Error: ErrorResponse{
			Message: "error",
		},
	}, NewErrorResponse(errors.New("error")))
}
