package git_pr

import (
	"github.com/jawher/mow.cli"
	"github.com/sjanota/git-pr/pkg/git"
	"github.com/sjanota/git-pr/pkg/ops"
)

type root struct {
	repo    git.Repo
	details *bool
}

func (r *root) Configure(app *cli.Cli) {
	r.details = app.BoolOpt("details d", false, "Adds more details to output")
	app.Action = r.action
}

func (r root) action() {
	err := ops.List(r.repo, *r.details)
	if err != nil {
		panic(err)
	}
}
