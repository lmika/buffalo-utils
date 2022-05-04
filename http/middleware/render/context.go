package render

import (
	"bytes"
	"net/http"
)

type Inv struct {
	render *Render
	values map[string]interface{}
}

func (inv *Inv) Set(name string, value interface{}) {
	inv.values[name] = value
}

func (inv *Inv) HTML(r *http.Request, w http.ResponseWriter, status int, templateName string) {
	// Render the content template
	tmpl, err := inv.render.template(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bw := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(bw, templateName, inv.values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render the master template
	/*
		masterTmpl, err := inv.render.template("masters/frame.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		masterBw := new(bytes.Buffer)
		if err := masterTmpl.ExecuteTemplate(masterBw, "masters/frame.html", struct {
			Content template.HTML
		}{
			Content: template.HTML(bw.String()),
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/

	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(status)
	masterBw.WriteTo(w)
}
