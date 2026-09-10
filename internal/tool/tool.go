package tool


import (
    "context"
	"log"
	"time"
    "github.com/voidSpir1t/madara/internal/tool/ssh"
    "google.golang.org/adk/v2/tool/functiontool"
    adktool "google.golang.org/adk/v2/tool"
    adkagent "google.golang.org/adk/v2/agent"
)

type SSHCommandArgs struct {
	Command  string `json:"command"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ToolRegistry struct {
    tools map[string]adktool.Tool
}

func (tr *ToolRegistry) Get(name string) (adktool.Tool) {
    tool := tr.tools[name]
    return tool
}

func (tr *ToolRegistry) Add(name string, tool adktool.Tool) {
    tr.tools[name] = tool
}

func NewToolRegistry() (ToolRegistry, error){
    tr := ToolRegistry{
        tools:  make(map[string]adktool.Tool),
    }

    sshTool, err := functiontool.New[SSHCommandArgs, map[string]any](
		functiontool.Config{
			Name: "run_ssh_command",
			Description: "Runs a command on a remote host over SSH (password authentication, port 22) and returns the combined stdout/stderr. " +
				"Host, user and password are optional;",
		},
		func(ctx adkagent.Context, args SSHCommandArgs) (map[string]any, error) {
			host, user, password := args.Host, args.User, args.Password
			runCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
			defer cancel()
			output, err := ssh.RunSSHCommand(runCtx, host, user, password, args.Command)
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
	if err != nil {
		log.Fatalf("Failed to create run_ssh_command tool: %v", err)
	}
    
    tr.Add("sshTool", sshTool)

    return tr, nil
}