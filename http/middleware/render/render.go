package render

import (
	"net/http"
)

func HTML(r *http.Request, w http.ResponseWriter, status int, templateName string) {
	rc, ok := r.Context().Value(renderContextKey).(*Inv)
	if !ok {
		return
	}

	rc.HTML(r, w, status, templateName)
}
