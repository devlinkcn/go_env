package strcat

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	at := assert.New(t)

	body := "{\"strings\":[\"prefix\",\"suffix\"]}"
	r, err := http.NewRequest("POST", "/strcat", strings.NewReader(body))
	at.Nil(err)

	w := httptest.NewRecorder()

	handler := New()
	handler.ServeHTTP(w, r)

	at.Equal(w.Code, http.StatusOK)
	at.Equal(w.Body.String(), "{\"string\":\"prefixsuffix\"}\n")
}
