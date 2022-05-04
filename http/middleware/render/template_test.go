package render_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/lmika/gopkgs/http/middleware/render"
	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	t.Run("should generate template successfully", func(t *testing.T) {
		rw := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "https://www.example.com/", nil)

		rnd := render.New(fstest.MapFS{
			"index.html": &fstest.MapFile{
				Data: []byte(`Template: {{.alpha}} - {{.bravo}}`),
			},
		})

		inv := rnd.NewInv()
		inv.Set("alpha", "Hello")
		inv.Set("bravo", "World")
		inv.HTML(r, rw, http.StatusOK, "index.html")

		assert.Equal(t, http.StatusOK, rw.Result().StatusCode)
		assert.Equal(t, "text/html; charset=utf-8", rw.Header().Get("Content-type"))
		assert.Equal(t, `Template: Hello - World`, rw.Body.String())
	})
}
