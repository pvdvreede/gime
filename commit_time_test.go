package gime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_CommitTime_Time(t *testing.T) {
	run := func(msg string) *CommitTime {
		return NewCommitTime(msg)
	}

	Convey("It returns 0 when there are no hours in the msg", t, func() {
		So(run("This is a message without time.").Time(), ShouldEqual, 0)
	})

	Convey("It returns 5 when there is '5hrs' in the msg", t, func() {
		So(run("Worked hard, 5hrs.").Time(), ShouldEqual, 5)
	})

	Convey("It returns 0.5 when there is '30mins' in the msg", t, func() {
		So(run("Didnt work to hard, only did 30mins here.").Time(), ShouldEqual, 0.5)
	})

	Convey("It returns 0 when there is 'ADhrs' in the msg", t, func() {
		So(run("I dont know what I am saying, blah ADhrs.").Time(), ShouldEqual, 0)
	})

	Convey("It caches the result and will return the same value more than once", t, func() {
		ct := run("Didnt work to hard, only did 30mins here.")
		So(ct.Time(), ShouldEqual, 0.5)
		So(ct.Time(), ShouldEqual, 0.5)
	})

}
