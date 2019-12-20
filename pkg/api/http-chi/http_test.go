package httpchi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("empty payload", func(t *testing.T) {
		response := `{"error": "EOF"}`
		_, body := testRequest(t, ts, "POST", "/", nil)
		assert.Equal(t, body, response)
	})

	t.Run("correct payload", func(t *testing.T) {
		payload, err := json.Marshal(addRequest{
			URL:      "http://test.com",
			Describe: "",
		})
		assert.Nil(t, err)
		response := `{"error": "Not found subscribe to event METHOD_ADD"}`
		_, body := testRequest(t, ts, "POST", "/", bytes.NewReader(payload))
		assert.Equal(t, body, response)
	})

}

func TestGet(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("correct payload", func(t *testing.T) {
		response := `{"error": "Not found subscribe to event METHOD_GET"}`
		_, body := testRequest(t, ts, "GET", "/hash", nil)
		assert.Equal(t, body, response)
	})
}

func TestList(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("correct payload", func(t *testing.T) {
		response := `{"error": "Not found subscribe to event METHOD_LIST"}`
		_, body := testRequest(t, ts, "GET", "/links", nil)
		assert.Equal(t, body, response)
	})
}

func TestDelete(t *testing.T) {
	server := &API{}

	r := chi.NewRouter()
	r.Mount("/", server.Routes())

	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("correct payload", func(t *testing.T) {
		payload, err := json.Marshal(deleteRequest{
			Hash: "hash",
		})
		assert.Nil(t, err)
		response := `{"error": "Not found subscribe to event METHOD_DELETE"}`
		_, body := testRequest(t, ts, "DELETE", "/", bytes.NewReader(payload))
		assert.Equal(t, body, response)
	})
}

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) { // nolint unparam
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		assert.Nil(t, err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		assert.Nil(t, err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		assert.Nil(t, err)
		return nil, ""
	}

	defer resp.Body.Close()
	return resp, string(respBody)
}
