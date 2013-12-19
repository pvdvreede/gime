package gime

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestShellCommand struct {
	returnValue []byte
}

func (self *TestShellCommand) Exec() []byte {
	return self.returnValue
}

func Test_ShellCollector_Collect(t *testing.T) {
	logStream := []byte(`
commit 2f67de042f2af45412886babb41f215a53b60b27
Author: pvdvreede <paul@vdvreede.net>
Date:   Wed Dec 18 15:17:50 2013 +1100

    Add godoc badge to Readme.

commit 351d40bba117f7738267e5118c14482ceadf2217
Author: pvdvreede <paul@vdvreede.net>
Date:   Wed Dec 18 14:50:26 2013 +1100

    Add tallying functions for total and aggregate by day.

`)

	cmd := &TestShellCommand{logStream}
	collector := NewShellCollector(cmd)

	Convey("The Log stream is parsed into commit objects", t, func() {
		commits := collector.Collect()

		So(len(commits), ShouldEqual, 2)
		So(commits[0].Message(), ShouldEqual, "Add godoc badge to Readme.")
		So(commits[0].Author(), ShouldEqual, "pvdvreede <paul@vdvreede.net>")
	})
}

func Test_Collector_parseCommit(t *testing.T) {
	commit := `
2f67de042f2af45412886babb41f215a53b60b27
Author: pvdvreede <paul@vdvreede.net>
Date:   Wed Dec 18 15:17:50 2013 +1100

    Add godoc badge to Readme.

`

	Convey("A string commit is parsed properly to a CommitTime object", t, func() {

	})
}
