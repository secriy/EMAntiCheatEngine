package model

import "time"

// Log is Log Model
type Log struct {
	ID        uint64
	IP        string
	Player    string
	Content   string
	CreatedAt time.Time
}
