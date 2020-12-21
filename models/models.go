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

// Robot schema of the robot table
type Robot struct {
	RobotID string `json:"robotID"`
	Name    string `json:"name"`
}

// Task schema of the robot table
type Task struct {
	ID          int64           `json:"id"`
	TaskDetails json.RawMessage `json:"taskDetails"`
}

// Graph schema of the graph table
type Graph struct {
	ID       int64           `json:"id"`
	MapVerID string          `json:"mapVerID"`
	Scale    float32         `json:"scale"`
	Name     string          `json:"name"`
	Location string          `json:"location"`
	Level    int64           `json:"level"`
	Lanes    json.RawMessage `json:"lanes"`
	Vertices json.RawMessage `json:"vertices"`
}

// GraphNonDetailed schema of the graph table with less details
type GraphNonDetailed struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}
