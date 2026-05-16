package bw

import "time"

type LockStatus string

const (
	LockStatusUnauthenticated LockStatus = "unauthenticated"
	LockStatusUnlocked        LockStatus = "unlocked"
	LockStatusLocked          LockStatus = "locked"
)

type clientStatus struct {
	ServerUrl string     `json:"serverUrl"`
	LastSync  *time.Time `json:"lastSync"`
	UserEmail *string    `json:"userEmail"`
	UserId    *string    `json:"userId"`
	Status    LockStatus `json:"status"`
}
