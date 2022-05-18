package golang

import (
	ga "github.com/devstream-io/devstream/internal/pkg/plugin/githubactions"
)

// Options is the struct for configurations of the githubactions plugin.
type Options struct {
	Owner    string       `validate:"required"`
	Org      string       `validate:"required"`
	Repo     string       `validate:"required"`
	Branch   string       `validate:"required"`
	Language *ga.Language `validate:"required"`
	Build    *Build
	Test     *Test   `validate:"required"`
	Docker   *Docker `validate:"required"`
}
