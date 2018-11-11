package github

import (
	"errors"
	"fmt"
	std_url "net/url"
	"strings"
)

type URL struct {
	Owner          string
	RepositoryName string
}

func ParseURL(s string) (*URL, error) {
	url, err := std_url.Parse(s)
	if err != nil {
		return nil, err
	}

	if url.Host != "github.com" {
		return nil, errors.New("not GitHub URL")
	}

	path := strings.Split(url.Path, "/")

	// There is a leading slash in url.Path
	if len(path) < 3 || path[1] == "" && path[2] == "" {
		return nil, fmt.Errorf("invalid remote path %s", url.Path)
	}

	return &URL{path[1], path[2]}, nil
}