package render

import (
	"bytes"
	"html/template"
	"net/http"
)

type Inv struct {
	config *Config
	values map[string]interface{}
}

func (inv *Inv) Set(name string, value interface{}) {
	inv.values[name] = value
}

func (inv *Inv) HTML(r *http.Request, w http.ResponseWriter, status int, templateName string) {
	// Render the content template
	tmpl, err := inv.config.template(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bw := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(bw, templateName, inv.values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render any frame templates
	for i := len(inv.config.frameTemplates) - 1; i >= 0; i-- {
		frameTemplateName := inv.config.frameTemplates[i]
		frameOutput, err := inv.renderFrameTemplate(frameTemplateName, bw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		bw = frameOutput
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	bw.WriteTo(w)
}

func (inv *Inv) renderFrameTemplate(frameTemplateName string, subframeOutput *bytes.Buffer) (*bytes.Buffer, error) {
	frameTemplate, err := inv.config.template(frameTemplateName)
	if err != nil {
		return nil, err
	}

	frameOutput := new(bytes.Buffer)
	if err := frameTemplate.ExecuteTemplate(frameOutput, frameTemplateName, struct {
		Content template.HTML
	}{
		Content: template.HTML(subframeOutput.String()),
	}); err != nil {
		return nil, err
	}

	return frameOutput, nil
}
