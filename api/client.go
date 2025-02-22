package api

import (
	"bufio"
	"fmt"
	"github.com/ahhcash/llm-cli/defaults"
	"github.com/ahhcash/llm-cli/llm"
	"os"
)

func Prompt(llmType string, prompt string, stream bool, tokens int, model string, system string, clear bool) {
	newLLM, err := llm.NewLLM(llmType)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	_, err = newLLM.Prompt(prompt, stream, tokens, model, system, clear)
	if err != nil {
		fmt.Println("Error fetching completion:", err)
		os.Exit(1)
	}
}

func Chat(llmType string, clr bool) {
	newLLM, err := llm.NewLLM(llmType)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	model, err := defaults.GetDefaultModel(llmType)
	if err != nil {
		fmt.Println("Error getting default model:", err)
		os.Exit(1)
	}
	fmt.Println("Entering chat mode. Type 'exit' to exit.")
	firstPrompt := true
	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		if text == "exit\n" {
			break
		}
		_, err := newLLM.Prompt(text, true, 1024, model, "", firstPrompt && clr)
		if err != nil {
			fmt.Println("Error fetching completion:", err)
			continue
		}
		firstPrompt = false
		fmt.Println()
	}
}
