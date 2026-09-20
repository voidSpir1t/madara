package tool

import (
	"log"

	"github.com/voidSpir1t/madara/internal/storage"
	"github.com/voidSpir1t/madara/internal/tool/filetool"
	"github.com/voidSpir1t/madara/internal/tool/sshtool"

	"google.golang.org/adk/v2/tool"
)

type ToolRegistry struct {
	tools map[string]tool.Tool
}

func (tr *ToolRegistry) Get(name string) tool.Tool {
	tool := tr.tools[name]
	return tool
}

func (tr *ToolRegistry) Add(name string, tool tool.Tool) {
	tr.tools[name] = tool
}

func NewToolRegistry(storage storage.Storage) (ToolRegistry, error) {
	tr := ToolRegistry{
		tools: make(map[string]tool.Tool),
	}

	sshTool, err := sshtool.NewSshTool()
	if err != nil {
		log.Fatalf("Failed to create run_ssh_command tool: %v", err)
	}

	fileTool, err := filetool.NewFileTool(storage)
	if err != nil {
		log.Fatalf("Failed to create run_ssh_command tool: %v", err)
	}

	tr.Add("sshTool", sshTool)
	tr.Add("fileTool", fileTool)

	return tr, nil
}
