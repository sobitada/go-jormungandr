package dto

import "time"

type LeaderAssignment struct {
	CreationTime time.Time `json:"created_at_time"`
	ScheduleTime time.Time `json:"scheduled_at_time"`
	FinishingTime time.Time `json:"finished_at_time"`
	LeaderID uint32 `json:"enclave_leader_id"`
}
