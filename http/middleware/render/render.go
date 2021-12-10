package render

import (
	"bytes"
	"html/template"
	"net/http"
)

// HTML renders the template with the given name and status code as the response.  This should
// usually be the last call of the handler.
func HTML(r *http.Request, w http.ResponseWriter, status int, templateName string) {
	rc, ok := r.Context().Value(renderContextKey).(*renderContext)
	if !ok {
		return
	}

	// Render the content template
	tmpl, err := rc.render.template(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bw := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(bw, templateName, rc.values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the master template
	var targetBw *bytes.Buffer
	if rc.render.master != "" {
		masterTmpl, err := rc.render.template(rc.render.master)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		targetBw = new(bytes.Buffer)
		if err := masterTmpl.ExecuteTemplate(targetBw, rc.render.master, struct {
			Content template.HTML
		}{
			Content: template.HTML(bw.String()),
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		targetBw = bw
	}

	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(status)
	targetBw.WriteTo(w)
}
