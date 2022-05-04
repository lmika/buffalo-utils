package render

import (
	"net/http"
)

// HTML renders the template with the given name and status code as the response.  This should
// usually be the last call of the handler.
func HTML(r *http.Request, w http.ResponseWriter, status int, templateName string) {
	rc, ok := r.Context().Value(renderContextKey).(*Inv)
	if !ok {
		return
	}

	rc.HTML(r, w, status, templateName)
}
