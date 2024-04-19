package middleware

import (
	"go_rest_mongo/utils"
	"net/http"
)

func RoleValidatorMiddleware(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := utils.GetAuthorization(r); err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		h.ServeHTTP(w, r)
	})
}
