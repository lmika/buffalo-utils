package render

import (
	"net/http"
)

func Set(r *http.Request, name string, value interface{}) {
	rc, ok := r.Context().Value(renderContextKey).(*Inv)
	if !ok {
		return
	}

	rc.Set(name, value)
}

func UseFrame(r *http.Request, name string) {
	rc, ok := r.Context().Value(renderContextKey).(*Inv)
	if !ok {
		return
	}

	rc.UseFrame(name)
}

func SetFrameArg(r *http.Request, name string, value interface{}) {
	rc, ok := r.Context().Value(renderContextKey).(*Inv)
	if !ok {
		return
	}

	rc.SetFrameArg(name, value)
}

func HTML(w http.ResponseWriter, r *http.Request, status int, templateName string) {
	rc, ok := r.Context().Value(renderContextKey).(*Inv)
	if !ok {
		return
	}

	rc.HTML(w, r, status, templateName)
}

func JSON(w http.ResponseWriter, r *http.Request, status int, data any) {
	rc, ok := r.Context().Value(renderContextKey).(*Inv)
	if !ok {
		return
	}

	rc.JSON(w, r, status, data)
}
