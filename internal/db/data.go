package db

import "time"

type Data struct {
	ChallengeID   string    `json:"challenge_id"`
	TimeCompleted time.Time `json:"time"`
}
