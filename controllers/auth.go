package controllers

import (
    "encoding/json"
    "net/http"
    "e-ticketing/database"
    "e-ticketing/models"
    "e-ticketing/utils"
    "golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
        return
    }

    // Get user from database
    db := database.GetDB()
    var user models.User
    if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
        http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
        return
    }

    // Check password
    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
        http.Error(w, `{"error": "Invalid credentials"}`, http.StatusUnauthorized)
        return
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(user.ID, user.Username)
    if err != nil {
        http.Error(w, `{"error": "Error generating token"}`, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "token":   token,
        "message": "Login successful",
    })
}