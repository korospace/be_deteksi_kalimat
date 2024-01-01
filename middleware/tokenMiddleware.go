package middleware

import (
	"be_deteksi_kalimat/helpers"
	"context"
	"net/http"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token is exist?
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			helpers.Response(w, 401, "Unauthorized", nil)
			return
		}

		// token is valid?
		claims, err := helpers.ValidateToken(accessToken)
		if err != nil {
			helpers.Response(w, 401, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), "tokeninfo", claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
