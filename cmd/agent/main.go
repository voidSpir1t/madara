package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/voidSpir1t/madara/internal/agent"
	"github.com/voidSpir1t/madara/internal/config"
	"github.com/voidSpir1t/madara/internal/model"
	"github.com/voidSpir1t/madara/internal/storage"
	"github.com/voidSpir1t/madara/internal/tool"
	adkagent "google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/cmd/launcher"
	"google.golang.org/adk/v2/cmd/launcher/full"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	mr, err := model.NewModelRegistry(cfg.Models)
	fmt.Println(mr.GetModelNames())

	s, err := storage.NewStorage(cfg.Storage)

	tr, err := tool.NewToolRegistry(s)

	a, err := agent.NewRootAgent(mr.GetModel("glm5.3"), tr)

	lcfg := &launcher.Config{
		AgentLoader: adkagent.NewSingleLoader(a),
	}

	l := full.NewLauncher()
	if err = l.Execute(ctx, lcfg, os.Args[1:]); err != nil {
		log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
	}
}
