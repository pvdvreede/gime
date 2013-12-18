package gime

import (
	"time"
)

func NewCommit(author, message, date string) *Commit {
	return &Commit{
		message: message,
		author:  author,
		date:    parseTime(date),
	}
}

func parseTime(dateTime string) time.Time {
	format := "Mon Jan 2 15:04:05 2006 -0700"
	parsedTime, err := time.Parse(format, dateTime)

	if err != nil {
		panic(err)
	}

	return parsedTime
}

type Commit struct {
	message string
	date    time.Time
	author  string
}

func (self *Commit) Message() string {
	return self.message
}

func (self *Commit) Date() time.Time {
	return self.date
}

func (self *Commit) Author() string {
	return self.author
}
