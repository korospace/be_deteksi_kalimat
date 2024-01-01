package middleware

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/helpers"
	"be_deteksi_kalimat/models"
	"context"
	"net/http"
)

func PakarMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token info
		tokeninfo := r.Context().Value("tokeninfo").(*helpers.TokenInfo)

		// get user access id
		var user models.User
		if err := database.DB.First(&user, "id = ?", tokeninfo.ID).Error; err != nil {
			helpers.Response(w, 500, "failed", err)
			return
		}

		if user.UserAccessID != 2 {
			helpers.Response(w, 401, "Access Denied", nil)
			return
		}

		ctx := context.WithValue(r.Context(), "userinfo", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
