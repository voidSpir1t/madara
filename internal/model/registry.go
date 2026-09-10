package model

import (
	"context"

	"github.com/voidSpir1t/madara/internal/config"
	adkModel "google.golang.org/adk/v2/model"
	"google.golang.org/adk/v2/model/openaimodel"
	"github.com/openai/openai-go/v3/option"
)

type ModelRegistry struct {
	models map[string]adkModel.LLM
}

func NewModelRegistry(cfg []config.ModelConfig) (ModelRegistry, error) {
	var models = make(map[string]adkModel.LLM)

	for _, modelcfg := range cfg {
		switch modelcfg.Provider {
		case "openai":
			m, err := registOpenaiModel(modelcfg)
			if err != nil {
				panic(err)
			}
			models[modelcfg.ModelName] = m
		default:
		}

	}
	modelReg := ModelRegistry{models: models}

	return modelReg, nil
}

func (mr *ModelRegistry) GetModelNames() []string {
	mn := make([]string, 0, len(mr.models))
	for name := range mr.models {
		mn = append(mn, name)
	}
	return mn
}

func (mr *ModelRegistry) GetModel(modelName string) adkModel.LLM {
	return mr.models[modelName]
}

func (mr *ModelRegistry) GetModelMap() map[string]adkModel.LLM {
	return mr.models
}

func registOpenaiModel(modelcfg config.ModelConfig) (adkModel.LLM, error){
	model, err := openaimodel.NewModel(context.Background(), modelcfg.ModelName, &openaimodel.ClientConfig{
				APIKey:  modelcfg.APIKey,
				BaseURL: modelcfg.BaseURL,
				Options: []option.RequestOption{option.WithHeader("User-Agent", "qagent/1.8.10"),
					option.WithHeader("X-Codegen-Client-Name", "qagent")},
	})
	if err != nil {
		return nil , err
	}

	return model, nil
}
