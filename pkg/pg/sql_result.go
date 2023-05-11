package pg

import (
	"fmt"
	"time"
)

type (
	// QueryStatus represents if the query run was a success or failure.
	QueryStatus string

	// QueryStart is the timestamp the query started.
	QueryStart time.Time

	// QueryTime represents the amount of time the query took.
	QueryTime float64
)

// SQLResult tracks sql response and time taken.
type SQLResult struct {
	Timestamp QueryStart  `json:"timestamp"`
	Message   string      `json:"message"`
	TimeTaken QueryTime   `json:"time_taken"`
	Status    QueryStatus `json:"status"`
}

const (
	Success QueryStatus = "success"
	Failure QueryStatus = "failed"
)

// MarshalJSON to format the QueryStart timestamp.
func (t QueryStart) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("15:04:05"))

	return []byte(stamp), nil
}

// MarshalJSON to format the QueryTime.
func (t QueryTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%.3fms\"", t)

	return []byte(stamp), nil
}
