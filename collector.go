package gime

import (
	"fmt"
	"regexp"
)

type GitLogCollector interface {
	Collect() []CommitWithTimer
}

type ShellCommander interface {
	Exec() []byte
}

func NewShellCollector(cmd ShellCommander) *ShellCollector {
	return &ShellCollector{cmd}
}

type ShellCollector struct {
	command ShellCommander
}

func (self *ShellCollector) Collect() []CommitWithTimer {
	result := self.command.Exec()
	splitter := regexp.MustCompile(`(^|\n)commit\s`)

	splitCommits := splitter.Split(string(result), -1)
	comms := []CommitWithTimer{}
	for i, commit := range splitCommits {
		if commit == "" {
			continue
		}
		fmt.Println(i)
		fmt.Println(commit)
		//comms = append(comms, parseCommit(commit))
	}

	return comms
}

func parseCommit(commitStr string) CommitWithTimer {
	fmt.Println(commitStr)
	authorReg := regexp.MustCompile(`^Author:\s+(?P<author>.*)$`)
	author := authorReg.FindStringSubmatch(commitStr)
	fmt.Println(author)
	//return CommitTime{}
	return NewCommitTime(NewCommit(author[1], "", ""))
}
