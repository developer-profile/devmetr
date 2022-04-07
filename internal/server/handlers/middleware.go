package handlers

import (
	"net/http"

	"github.com/developer-profile/devmetr/internal/server/validator"

	"github.com/gorilla/mux"
)

func (m *metricHandler) MiddlewareCheckHost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			vars := mux.Vars(r)
			mType := vars["mtype"]
			if !validator.IsMType(mType) {
				http.Error(w, ``, http.StatusNotImplemented)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
