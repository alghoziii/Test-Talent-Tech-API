package middleware

import (
    "context"
    "net/http"
    "strings"
    "e-ticketing/utils"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, `{"error": "Authorization header required"}`, http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            http.Error(w, `{"error": "Bearer token required"}`, http.StatusUnauthorized)
            return
        }

        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
            return
        }

        // Add claims to context
        ctx := context.WithValue(r.Context(), "claims", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}