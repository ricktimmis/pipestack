package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/ricktimmis/pipestack/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
