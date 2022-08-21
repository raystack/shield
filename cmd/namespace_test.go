package cmd_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/odpf/shield/cmd"
	"github.com/stretchr/testify/assert"
)

var expectedNamespaceUsageHelp = heredoc.Doc(`

\x1b[1mUsage\x1b[0m
  shield namespace [flags]

\x1b[1mCore commands\x1b[0m
  create      Create a namespace
  edit        Edit a namespace
  list        List all namespaces
  view        View a namespace

\x1b[1mFlags\x1b[0m
  -h, --host string   Shield API service to connect to

\x1b[1mInherited flags\x1b[0m
  --help   Show help for command

\x1b[1mExamples\x1b[0m
  $ shield namespace create
  $ shield namespace edit
  $ shield namespace view
  $ shield namespace list

`)

func TestClientNamespace(t *testing.T) {
	t.Run("without config file", func(t *testing.T) {
		tests := []struct {
			name        string
			cliConfig   *cmd.Config
			subCommands []string
			want        string
			err         error
		}{
			{
				name: "`namespace` only should show usage help",
				want: expectedNamespaceUsageHelp,
				err:  nil,
			},
			{
				name:        "`namespace` list only should throw error host not found",
				want:        "",
				subCommands: []string{"list"},
				err:         cmd.ErrClientConfigHostNotFound,
			},
			{
				name:        "`namespace` list with host flag should pass",
				want:        "",
				subCommands: []string{"list", "-h", "test"},
				err:         context.DeadlineExceeded,
			},
			{
				name:        "`namespace` create only should throw error host not found",
				want:        "",
				subCommands: []string{"create"},
				err:         cmd.ErrClientConfigHostNotFound,
			},
			{
				name:        "`namespace` create with host flag should throw error missing required flag",
				want:        "",
				subCommands: []string{"create", "-h", "test"},
				err:         errors.New("required flag(s) \"file\" not set"),
			},
			{
				name:        "`namespace` edit without host should throw error host not found",
				want:        "",
				subCommands: []string{"edit", "123"},
				err:         cmd.ErrClientConfigHostNotFound,
			},
			{
				name:        "`namespace` edit with host flag should throw error missing required flag",
				want:        "",
				subCommands: []string{"edit", "123", "-h", "test"},
				err:         errors.New("required flag(s) \"file\" not set"),
			},
			{
				name:        "`namespace` view without host should throw error host not found",
				want:        "",
				subCommands: []string{"view", "123"},
				err:         cmd.ErrClientConfigHostNotFound,
			},
			{
				name:        "`namespace` view with host flag should pass",
				want:        "",
				subCommands: []string{"view", "123", "-h", "test"},
				err:         context.DeadlineExceeded,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tt.cliConfig = &cmd.Config{}
				cli := cmd.New(tt.cliConfig)

				buf := new(bytes.Buffer)
				cli.SetOutput(buf)
				args := append([]string{"namespace"}, tt.subCommands...)
				cli.SetArgs(args)

				err := cli.Execute()
				got := buf.String()

				assert.Equal(t, tt.err, err)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
