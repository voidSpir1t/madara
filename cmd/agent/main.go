package main

import (
	"log"
	"fmt"
	"os"
	"context"

	adkagent "google.golang.org/adk/v2/agent"
	"github.com/voidSpir1t/madara/internal/config"
	"github.com/voidSpir1t/madara/internal/model"
	"github.com/voidSpir1t/madara/internal/agent"
	"github.com/voidSpir1t/madara/internal/tool"
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

	tr, err := tool.NewToolRegistry()

	a, err := agent.NewRootAgent(mr.GetModel("glm5.3"), tr)

	lcfg := &launcher.Config{
        AgentLoader: adkagent.NewSingleLoader(a),
    }

	l := full.NewLauncher()
    if err = l.Execute(ctx, lcfg, os.Args[1:]); err != nil {
        log.Fatalf("Run failed: %v\n\n%s", err, l.CommandLineSyntax())
    }
}
