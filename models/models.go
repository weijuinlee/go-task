package models

import "encoding/json"

// Patrol schema of the patrol table
type Patrol struct {
	ID       int64           `json:"id"`
	GraphID  int             `json:"graphID"`
	MapVerID string          `json:"mapVerID"`
	Name     string          `json:"name"`
	Points   json.RawMessage `json:"points"`
}
