package docbot

import (
	"context"
	"fmt"
	"os"

	_ "embed"

	"github.com/diveagents/dive"
	"github.com/diveagents/dive/agent"
	"github.com/diveagents/dive/llm"
	"github.com/diveagents/dive/llm/providers/openai"
	"github.com/diveagents/dive/slogger"
	"github.com/diveagents/dive/toolkit"
)

// use go embed to embed a local context.txt
//
//go:embed context.txt
var llmcontext string

type DocBot struct {
	path string
}

func New(path string) *DocBot {
	return &DocBot{path}
}

func (b *DocBot) DocumentCode(ctx context.Context) (*dive.Response, error) {

	var model llm.LLM
	var err error

	model = openai.New(
		openai.WithEndpoint(os.Getenv("OPENAI_ENDPOINT")),
	)

	researcher, err := agent.New(agent.Options{
		Name:         "Risor Documentation Assistant",
		Goal:         "Document Risor code",
		Model:        model,
		Instructions: llmcontext,
		Logger:       slogger.New(slogger.LevelFromString("debug")),
		Tools: []llm.Tool{
			toolkit.NewFileReadTool(toolkit.FileReadToolOptions{MaxSize: 100000}),
			toolkit.NewFileWriteTool(toolkit.FileWriteToolOptions{}),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to create research agent: %w", err)
	}

	return researcher.CreateResponse(ctx, dive.WithInput(fmt.Sprintf("document the Risor code in %s. Update the file.", b.path)))
}
