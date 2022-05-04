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
