package gime

import (
	"regexp"
	"strconv"
)

func NewCommitTime(message string) *CommitTime {
	return &CommitTime{
		message: message,
		time:    -1,
	}
}

type CommitTime struct {
	message string
	time    float64
}

// Takes the types commit message and will return the number of hours
// that have been worked. This allows hrs or mins as the postfix to the amount
// and will convert the number to hours.
func (self *CommitTime) Time() float64 {
	if self.time != -1 {
		return self.time
	}

	msg := self.message
	reg := regexp.MustCompile(`\s(?P<time>[0-9\.]+)(?P<type>hrs|mins)`)
	if !reg.MatchString(msg) {
		return 0
	}

	matches := reg.FindStringSubmatch(msg)

	time, err := strconv.ParseFloat(matches[1], 64)

	if err != nil {
		panic(err)
	}

	switch matches[2] {
	case "hrs":
		return time
	case "mins":
		return time / 60
	default:
		return 0
	}

}
