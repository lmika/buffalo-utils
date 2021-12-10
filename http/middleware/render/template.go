package render

import (
	"context"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"sync"
)

// Render is a template renderer.  It uses the Go HTML template package and designed to be used
// as middleware for the standard http.Handler.
type Render struct {
	templateFS fs.FS
	master string

	cacheMutex  *sync.RWMutex
	templateSet *template.Template
}

// New creates a new template renderer.  Templates are read from the provided FS.
func New(tmplFS fs.FS) *Render {
	cfg := &Render{
		templateFS: tmplFS,
		master: "",

		cacheMutex:  new(sync.RWMutex),
		templateSet: template.New("/"),
	}

	_ = fs.WalkDir(tmplFS, ".", func(path string, d fs.DirEntry, err error) error {
		if !strings.HasSuffix(path, ".html") {
			return nil
		}

		tmpl, err := cfg.parseTemplate(path)
		if err != nil {
			log.Printf("template %v: %v", path, err)
			return nil
		}

		if _, err := cfg.templateSet.AddParseTree(path, tmpl.Tree); err != nil {
			log.Printf("template %v: %v", path, err)
		}

		return nil
	})
	return cfg
}

// Use enables use of the renderer with the passed in handler.
func (tc *Render) Use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rc := &renderContext{
			render: tc,
			values: make(map[string]interface{}),
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), renderContextKey, rc)))
	})
}

func (tc *Render) template(name string) (*template.Template, error) {
	return tc.templateSet.Lookup(name), nil
}

func (tc *Render) parseTemplate(name string) (*template.Template, error) {
	f, err := tc.templateFS.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	tmplBytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(name).Parse(string(tmplBytes))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}

type renderContext struct {
	render *Render
	values map[string]interface{}
}

type renderContextKeyType struct{}

var renderContextKey = renderContextKeyType{}
