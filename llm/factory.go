package llm

import (
	"fmt"
	"github.com/ahhcash/llm-cli/session"
	"github.com/ahhcash/llm-cli/variants/anthropic"
	"github.com/ahhcash/llm-cli/variants/mistral"
	"github.com/ahhcash/llm-cli/variants/openai"
	"net/http"
)

type LLM interface {
	Prompt(prompt string, stream bool, tokens int, model string, system string, clear bool) (string, error)
	MarshalRequest(prompt string, stream bool, tokens int, model string, system string, session *session.Session) (*http.Request, error)
	ParseResponse(resp *http.Response) (string, error)
	ParseStreamingResponse(resp *http.Response) (string, error)
}

func NewLLM(llmType string) (LLM, error) {
	switch llmType {
	case "claude":
		config, err := anthropic.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("error loading config for Claude: %w", err)
		}
		return anthropic.NewClient(config), nil
	case "mistral":
		config, err := mistral.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("error loading config for Mistral: %w", err)
		}
		return mistral.NewClient(config), nil
	case "gpt":
		config, err := openai.LoadConfig()
		if err != nil {
			return nil, fmt.Errorf("error loading config for OpenAI: %w", err)
		}
		return openai.NewClient(config), nil
	default:
		return nil, fmt.Errorf("unknown LLM type: %s", llmType)
	}
}
