package ops

import (
	"fmt"
	"github.com/sjanota/git-hub/pkg/git"
	"strings"
)

const (
	statusCommentHeaderPadding = "└──"
	statusCommentPadding       = "   "
)

func Status(repo git.Repo) error {
	prs, err := repo.ListPullRequests()
	if err != nil {
		return err
	}

	currentPr, err := getPullRequestForCurrentBranch(repo)
	if _, ok := err.(git.NoPullRequestForBranch); err != nil && !ok {
		return err
	}

	for _, pr := range prs {
		if currentPr != nil && pr.Number == currentPr.Number {
			fmt.Printf("* %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		} else {
			fmt.Printf("  %-6v %-32s %s\n", pr.Number, pr.HeadRef, pr.Title)
		}

		if pr.Comment != "" {
			lines := strings.Split(pr.Comment, "\n")
			fmt.Println(statusCommentHeaderPadding, lines[0])
			for _, line := range lines[1:] {
				fmt.Println(statusCommentPadding, line)
			}
		}
	}

	return nil
}
