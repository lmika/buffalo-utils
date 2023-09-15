package render

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"html/template"
	"net/http"
)

type Inv struct {
	config            *Config
	extraFrames       []string
	ignoreGlobalFrame bool
	values            map[string]any
	frameArgs         map[string]any
}

func (inv *Inv) Set(name string, value any) {
	inv.values[name] = value
}

func (inv *Inv) SetFrameArg(name string, value any) {
	if inv.frameArgs == nil {
		inv.frameArgs = make(map[string]any)
	}
	inv.frameArgs[name] = value
}

// UseFrame adds the use of the frame to the pending frame stack.
func (inv *Inv) UseFrame(name string) {
	inv.extraFrames = append(inv.extraFrames, name)
}

// SetFrame will replace any pending frame with the given name.
func (inv *Inv) SetFrame(name string) {
	inv.ignoreGlobalFrame = true
	inv.extraFrames = []string{name}
}

func (inv *Inv) HTML(w http.ResponseWriter, r *http.Request, status int, templateName string) {
	// Render the content template
	tmpl, err := inv.config.template(templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if tmpl == nil {
		http.Error(w, "template missing: "+templateName, http.StatusInternalServerError)
		return
	}

	bw := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(bw, templateName, inv.values); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Render any invocation frame templates
	for i := len(inv.extraFrames) - 1; i >= 0; i-- {
		frameTemplateName := inv.extraFrames[i]
		frameOutput, err := inv.renderFrameTemplate(frameTemplateName, bw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bw = frameOutput
	}

	// Render any global frame templates
	if !inv.ignoreGlobalFrame {
		for i := len(inv.config.frameTemplates) - 1; i >= 0; i-- {
			frameTemplateName := inv.config.frameTemplates[i]
			frameOutput, err := inv.renderFrameTemplate(frameTemplateName, bw)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			bw = frameOutput
		}
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	bw.WriteTo(w)
}

func (inv *Inv) renderFrameTemplate(frameTemplateName string, subframeOutput *bytes.Buffer) (*bytes.Buffer, error) {
	frameTemplateData := map[string]any{
		"Content": template.HTML(subframeOutput.String()),
	}
	for k, v := range inv.frameArgs {
		frameTemplateData[k] = v
	}

	frameTemplate, err := inv.config.template(frameTemplateName)
	if err != nil {
		return nil, err
	}

	frameOutput := new(bytes.Buffer)
	if err := frameTemplate.ExecuteTemplate(frameOutput, frameTemplateName, frameTemplateData); err != nil {
		return nil, err
	}

	return frameOutput, nil
}

func (inv *Inv) JSON(w http.ResponseWriter, r *http.Request, status int, data any) {
	bts, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(bts)
}

func (inv *Inv) XML(w http.ResponseWriter, r *http.Request, status int, data any) {
	bts, err := xml.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/xml; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(xml.Header))
	w.Write(bts)
}
