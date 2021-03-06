package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type fetch struct {
	repo   git.Repo
	remote *string
	all    *bool
}

func (f *fetch) Configure(app *cli.Cli) {
	app.Command("fetch f", "fetch Pull Requests", func(cmd *cli.Cmd) {
		cmd.Spec = "[REMOTE | -a]"
		f.remote = cmd.StringArg("REMOTE", "origin", "Optional remote name to fetch")
		f.all = cmd.BoolOpt("all a", false, "fetch all remotes")

		cmd.Action = f.action
	})
}

func (f *fetch) action() {
	var remotes git.RemotesLister
	if *f.all {
		remotes = git.AllRemotesLister{}
	} else {
		remotes = git.OneRemoteLister{Remote: *f.remote}
	}

	err := ops.FetchPullRequests(f.repo, remotes)
	if err != nil {
		panic(err)
	}
}
