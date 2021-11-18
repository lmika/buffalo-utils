package render

import (
	"net/http"
)

func Set(r *http.Request, name string, value interface{}) {
	rc, ok := r.Context().Value(renderContextKey).(*renderContext)
	if !ok {
		return
	}
	rc.values[name] = value
}
