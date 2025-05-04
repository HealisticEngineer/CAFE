package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HealisticEngineer/CAFE/db"
	"github.com/HealisticEngineer/CAFE/utils"
)

type ConfigResponse struct {
	ClientName       string            `json:"client_name"`
	AccessGroups     []string          `json:"access_groups"`
	SoftwarePackages []string          `json:"software_packages"`
	CustomConfigs    map[string]string `json:"custom_configs"`
}

// ConfigHandler handles the /config endpoint
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	if !utils.IPAllowed(r.RemoteAddr) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	clientName := r.Header.Get("X-Client-Name")
	if clientName == "" {
		http.Error(w, "Missing X-Client-Name header", http.StatusBadRequest)
		return
	}

	query := `SELECT AccessGroups, SoftwarePackages, CustomKey, CustomValue FROM Configuration WHERE ClientName = @p1`
	rows, err := db.DB.Query(query, clientName)
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var accessGroupsStr, packagesStr string
	custom := make(map[string]string)
	first := true
	for rows.Next() {
		var key, val sql.NullString
		err := rows.Scan(&accessGroupsStr, &packagesStr, &key, &val)
		if err != nil {
			http.Error(w, "Failed to read row", http.StatusInternalServerError)
			return
		}
		if key.Valid && val.Valid {
			custom[key.String] = val.String
		}
		first = false
	}

	if first {
		http.Error(w, "No configuration found", http.StatusNotFound)
		return
	}

	resp := ConfigResponse{
		ClientName:       clientName,
		AccessGroups:     []string{},
		SoftwarePackages: []string{},
		CustomConfigs:    custom,
	}
	json.Unmarshal([]byte(accessGroupsStr), &resp.AccessGroups)
	json.Unmarshal([]byte(packagesStr), &resp.SoftwarePackages)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
