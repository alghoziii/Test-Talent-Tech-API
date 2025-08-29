package controllers

import (
    "encoding/json"
    "net/http"
    "e-ticketing/database"
    "e-ticketing/models"
)

func CreateTerminal(w http.ResponseWriter, r *http.Request) {
    var req struct {
        Code     string `json:"code"`
        Name     string `json:"name"`
        Location string `json:"location"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, `{"error": "Invalid request"}`, http.StatusBadRequest)
        return
    }

    // Create terminal
    terminal := models.Terminal{
        Code:     req.Code,
        Name:     req.Name,
        Location: req.Location,
    }

    db := database.GetDB()
    if err := db.Create(&terminal).Error; err != nil {
        http.Error(w, `{"error": "Error creating terminal"}`, http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Terminal created successfully",
        "terminal": terminal,
    })
}