package gime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

type TestCommitTime struct {
	time float64
	date time.Time
}

func (self *TestCommitTime) Date() time.Time {
	return self.date
}

func (self *TestCommitTime) Time() float64 {
	return self.time
}

func (self *TestCommitTime) Author() string {
	return "Me"
}

func (self *TestCommitTime) Message() string {
	return "Message"
}

func Test_Tally_TotalHours(t *testing.T) {
	run := func(nums ...float64) []CommitTimer {
		comms := []CommitTimer{}

		for _, num := range nums {
			ct := &TestCommitTime{num, time.Now()}
			comms = append(comms, ct)
		}

		return comms
	}

	Convey("It adds up the correct numbers", t, func() {
		So(TotalHours(run(1, 3, 6)), ShouldEqual, 10)
		So(TotalHours(run(1.5, 3, 6.8)), ShouldEqual, 11.3)
	})

}

func Test_Tally_TotalHoursByDay(t *testing.T) {
	run := func(comms ...CommitWithTimer) []CommitWithTimer {
		cwt := []CommitWithTimer{}

		for _, c := range comms {
			cwt = append(cwt, c)
		}

		return cwt
	}

	Convey("It aggregates totals by day", t, func() {
		comms := run(
			&TestCommitTime{
				3,
				time.Date(2013, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			&TestCommitTime{
				7,
				time.Date(2013, time.November, 10, 23, 0, 0, 0, time.UTC),
			},
			&TestCommitTime{
				6,
				time.Date(2013, time.November, 11, 23, 0, 0, 0, time.UTC),
			},
		)

		agg := TotalHoursByDay(comms)

		So(len(agg), ShouldEqual, 2)
		So(agg["2013-11-10"], ShouldEqual, 10)
		So(agg["2013-11-11"], ShouldEqual, 6)
	})
}
