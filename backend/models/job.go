package models

import "time"

type Job struct {
    ID       string        `json:"id"`
    Name     string        `json:"name"`
    Duration time.Duration `json:"duration"`
    Status   string        `json:"status"`
}