package filetool

import (
	"github.com/voidSpir1t/madara/internal/storage"
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/functiontool"
)

type FileArgs struct {
	Key string `json:"key"`
}

func NewFileTool(storage storage.Storage) (tool.Tool, error) {
	return functiontool.New[FileArgs, map[string]any](
		functiontool.Config{
			Name:        "get_file",
			Description: "Get a file from the file storage and return a temporary download URL.",
		},
		func(ctx agent.Context, args FileArgs) (map[string]any, error) {
			url, err := storage.PresignGet(
				ctx,
				args.Key,
			)
			if err != nil {
				return nil, err
			}

			return map[string]any{
				"status": "success",
				"url":    url,
			}, nil
		},
	)
}
