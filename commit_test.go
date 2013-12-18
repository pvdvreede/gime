package gime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Commit_NewCommit(t *testing.T) {
	run := func(date string) *Commit {
		return NewCommit("test", "test", date)
	}

	Convey("It returns with the correct time object", t, func() {
		So(run("Wed Dec 18 12:13:31 2013 +1100").Date().Hour(), ShouldEqual, 12)
		So(run("Wed Dec 18 12:13:31 2013 +1100").Date().Minute(), ShouldEqual, 13)
		So(run("Wed Dec 18 12:13:31 2013 +1100").Date().Year(), ShouldEqual, 2013)
		So(run("Wed Dec 18 12:13:31 2013 +1100").Date().Day(), ShouldEqual, 18)
	})

	Convey("It panics if the commit date is not parsable", t, func() {
		So(func() { run("This isnt a date!") }, ShouldPanic)
	})

}
