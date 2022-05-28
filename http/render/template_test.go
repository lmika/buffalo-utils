package render_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/lmika/gopkgs/http/render"
	"github.com/stretchr/testify/assert"
)

func TestTemplate(t *testing.T) {
	t.Run("should render template successfully", func(t *testing.T) {
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
		inv.HTML(rw, r, http.StatusOK, "index.html")

		assert.Equal(t, http.StatusOK, rw.Result().StatusCode)
		assert.Equal(t, "text/html; charset=utf-8", rw.Header().Get("Content-type"))
		assert.Equal(t, `Template: Hello - World`, rw.Body.String())
	})
}

func TestTemplate_Use(t *testing.T) {
	t.Run("should render template using middleware", func(t *testing.T) {
		var handler http.Handler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			render.Set(r, "alpha", "Hello")
			render.Set(r, "bravo", "World")
			render.HTML(rw, r, http.StatusOK, "index.html")
		})

		rw := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "https://www.example.com/", nil)

		rnd := render.New(fstest.MapFS{
			"index.html": &fstest.MapFile{
				Data: []byte(`Template: {{.alpha}} - {{.bravo}}`),
			},
		})
		handler = rnd.Use(handler)

		handler.ServeHTTP(rw, r)

		assert.Equal(t, http.StatusOK, rw.Result().StatusCode)
		assert.Equal(t, "text/html; charset=utf-8", rw.Header().Get("Content-type"))
		assert.Equal(t, `Template: Hello - World`, rw.Body.String())
	})
}
