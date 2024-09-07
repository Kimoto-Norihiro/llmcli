package llm_models

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
)

type LLM struct {
	llm llms.Model
}

func NewLLM(ctx context.Context, kind string) (*LLM, error) {
	switch kind {
	case "gpt":
		gpt, err := generateGptModel()
		if err != nil {
			return nil, err
		}
		return &LLM{llm: gpt}, nil
	case "gemini":
		gemini, err := generateGeminiModel(ctx)
		if err != nil {
			return nil, err
		}

		return &LLM{llm: gemini}, nil
	default:
		return nil, errors.New("unknown LLM kind")
	}
}

func (l *LLM) GenerateContent(ctx context.Context, messages []llms.MessageContent, options ...llms.CallOption) (*llms.ContentResponse, error) {
	return l.llm.GenerateContent(ctx, messages, options...)
}

func generateGptModel() (*openai.LLM, error) {
	gptKey := viper.GetString("GPT_KEY")
	if gptKey == "" {
		return nil, fmt.Errorf("GPT_KEY not set")
	}

	gpt, err := openai.New(openai.WithToken(gptKey))
	if err != nil {
		return nil, err
	}

	return gpt, nil
}

func generateGeminiModel(ctx context.Context) (*googleai.GoogleAI, error) {
	geminiKey := viper.GetString("GEMINI_KEY")
	if geminiKey == "" {
		return nil, fmt.Errorf("GEMINI_KEY not set")
	}

	gemini, err := googleai.New(ctx, googleai.WithAPIKey(geminiKey))
	if err != nil {
		return nil, err
	}

	return gemini, nil
}
