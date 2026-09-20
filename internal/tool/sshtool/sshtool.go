package sshtool

import (
	"context"
	"time"

	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/functiontool"
)

type SSHCommandArgs struct {
	Command  string `json:"command"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func NewSshTool() (tool.Tool, error) {
	return functiontool.New[SSHCommandArgs, map[string]any](
		functiontool.Config{
			Name: "run_ssh_command",
			Description: "Runs a command on a remote host over SSH (password authentication, port 22) and returns the combined stdout/stderr. " +
				"Host, user and password are optional;",
		},
		func(ctx agent.Context, args SSHCommandArgs) (map[string]any, error) {
			host, user, password := args.Host, args.User, args.Password
			runCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
			defer cancel()
			output, err := RunSSHCommand(runCtx, host, user, password, args.Command)
			if err != nil {
				return map[string]any{
					"status":        "error",
					"error_message": err.Error(),
				}, nil
			}
			return map[string]any{
				"status": "success",
				"host":   host,
				"output": output,
			}, nil
		},
	)
}
