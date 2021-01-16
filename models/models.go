package models

import "encoding/json"

// Patrol schema of the patrol table
type Patrol struct {
	ID       int64           `json:"id"`
	GraphID  int64           `json:"graphID"`
	MapVerID string          `json:"mapVerID"`
	Name     string          `json:"name"`
	Points   json.RawMessage `json:"points"`
}

// Robot schema of the robot table
type Robot struct {
	ID      int64  `json:"id"`
	RobotID string `json:"robotID"`
	Name    string `json:"name"`
}

// Task schema of the task table
type Task struct {
	ID          int64           `json:"id"`
	Type        int64           `json:"type"`
	TaskDetails json.RawMessage `json:"taskDetails"`
}

// Graph schema of the graph table
type Graph struct {
	ID           int64           `json:"id"`
	MapVerID     string          `json:"mapVerID"`
	CollectionID int64           `json:"collectionID"`
	Scale        float32         `json:"scale"`
	Name         string          `json:"name"`
	Location     string          `json:"location"`
	Level        string          `json:"level"`
	Lanes        json.RawMessage `json:"lanes"`
	Vertices     json.RawMessage `json:"vertices"`
}

// GraphNonDetailed schema of the graph table with less details
type GraphNonDetailed struct {
	ID           int64  `json:"id"`
	CollectionID int64  `json:"collectionID"`
	Name         string `json:"name"`
	Location     string `json:"location"`
}

// Collection schema of the graph table with less details
type Collection struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
