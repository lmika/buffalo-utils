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
		inv.HTML(r, rw, http.StatusOK, "index.html")

		assert.Equal(t, http.StatusOK, rw.Result().StatusCode)
		assert.Equal(t, "text/html; charset=utf-8", rw.Header().Get("Content-type"))
		assert.Equal(t, `Template: Hello - World`, rw.Body.String())
	})
}

func TestTemplate_WithFrame(t *testing.T) {
	t.Run("should render templates in frame if specified", func(t *testing.T) {
		rw := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "https://www.example.com/", nil)

		rnd := render.New(fstest.MapFS{
			"index.html": &fstest.MapFile{
				Data: []byte(`{{.alpha}} - {{.bravo}}`),
			},
			"frame.html": &fstest.MapFile{
				Data: []byte(`Frame: [{{.Content}}]`),
			},
		}, render.WithFrame("frame.html"))

		inv := rnd.NewInv()
		inv.Set("alpha", "Hello")
		inv.Set("bravo", "World")
		inv.HTML(r, rw, http.StatusOK, "index.html")

		assert.Equal(t, http.StatusOK, rw.Result().StatusCode)
		assert.Equal(t, "text/html; charset=utf-8", rw.Header().Get("Content-type"))
		assert.Equal(t, `Frame: [Hello - World]`, rw.Body.String())
	})

	t.Run("should support multiple frames templates in frame if specified", func(t *testing.T) {
		rw := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "https://www.example.com/", nil)

		rnd := render.New(fstest.MapFS{
			"index.html": &fstest.MapFile{
				Data: []byte(`{{.alpha}} - {{.bravo}}`),
			},
			"frame.html": &fstest.MapFile{
				Data: []byte(`Frame: [{{.Content}}]`),
			},
		}, render.WithFrame("frame.html"), render.WithFrame("frame.html"))

		inv := rnd.NewInv()
		inv.Set("alpha", "Hello")
		inv.Set("bravo", "World")
		inv.HTML(r, rw, http.StatusOK, "index.html")

		assert.Equal(t, http.StatusOK, rw.Result().StatusCode)
		assert.Equal(t, "text/html; charset=utf-8", rw.Header().Get("Content-type"))
		assert.Equal(t, `Frame: [Frame: [Hello - World]]`, rw.Body.String())
	})

	t.Run("should render frames in reverse order", func(t *testing.T) {
		rw := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "https://www.example.com/", nil)

		rnd := render.New(fstest.MapFS{
			"index.html": &fstest.MapFile{
				Data: []byte(`{{.alpha}} - {{.bravo}}`),
			},
			"outer.html": &fstest.MapFile{
				Data: []byte(`Outer: [{{.Content}}]`),
			},
			"inner.html": &fstest.MapFile{
				Data: []byte(`Inner: [{{.Content}}]`),
			},
		}, render.WithFrame("outer.html"), render.WithFrame("inner.html"))

		inv := rnd.NewInv()
		inv.Set("alpha", "Hello")
		inv.Set("bravo", "World")
		inv.HTML(r, rw, http.StatusOK, "index.html")

		assert.Equal(t, http.StatusOK, rw.Result().StatusCode)
		assert.Equal(t, "text/html; charset=utf-8", rw.Header().Get("Content-type"))
		assert.Equal(t, `Outer: [Inner: [Hello - World]]`, rw.Body.String())
	})
}
