package gime

import (
	"time"
)

type CommitTimer interface {
	Time() float64
}

type Committer interface {
	Date() time.Time
	Author() string
	Message() string
}

type CommitWithTimer interface {
	Committer
	CommitTimer
}

type Aggregation map[string]float64

func TotalHours(times []CommitTimer) float64 {
	var total float64 = 0.00

	for _, t := range times {
		total += t.Time()
	}

	return total
}

func TotalHoursByDay(times []CommitWithTimer) Aggregation {
	format := "2006-01-02"
	agg := Aggregation{}

	for _, t := range times {
		agg[t.Date().Format(format)] += t.Time()
	}

	return agg
}
