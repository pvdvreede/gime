package gime

import (
	"regexp"
	"strconv"
)

func NewCommitTime(commit *Commit) *CommitTime {
	return &CommitTime{
		commit,
		parseMessageForTime(commit.Message()),
	}
}

// Takes the types commit message and will return the number of hours
// that have been worked. This allows hrs or mins as the postfix to the amount
// and will convert the number to hours.
func parseMessageForTime(msg string) float64 {

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

type CommitTime struct {
	*Commit
	time float64
}

func (self *CommitTime) Time() float64 {
	return self.time
}
