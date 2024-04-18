package middleware

import (
	"go_rest_mongo/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func RoleValidatorMiddleware(s []string) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := utils.GetAuthorization(r); err != nil {
				utils.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			}
			h.ServeHTTP(w, r)
		})

	}
}
