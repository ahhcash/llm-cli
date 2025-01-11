# LLM CLI 🚀

A command-line interface for getting instant coding help and technical explanations while maintaining conversational context.

## Supported Models
- ✅ Anthropic (Claude)
- ✅ Mistral AI
- ✅ OpenAI (GPT)

## Features

- **Coding Assistant**: Get concise explanations for programming concepts, debugging help, and code reviews
- **Context-Aware**: Conversations persist locally, enabling follow-up questions and detailed discussions
- **Multi-Model**: Switch between different LLMs while maintaining conversation context
- **Multiple Sessions**: Manage parallel conversations with unique contexts

## Installation

### Using Go Install
```bash
go install github.com/ahhcash/llm-cli/cmd/llm@latest
```

### Building from Source
```bash
git clone https://github.com/ahhcash/llm-cli
cd llm-cli
make
```

## Quick Start

1. Set API keys:
```bash
export ANTHROPIC_API_KEY="your-key"
export MISTRAL_API_KEY="your-key"
export OPENAI_API_KEY="your-key"
```

2. Set default model:
```bash
export DEFAULT_COMPLETER="claude"  # or "mistral" or "gpt"
```

## Usage

### Get Coding Help
```bash
llm claude "Explain goroutines in Go"
llm claude "Review my function: func add(a, b int) int { return a + b }"
llm claude "Debug error: fatal: not a git repository"
```

### Interactive Sessions
```bash
llm claude chat     # Start chat session
llm session list    # List all sessions
llm session switch <uuid>  # Switch context
```

### Advanced Usage
```bash
llm claude -S "You are a security expert" "Review my SSH config"
llm claude --model claude-3-opus-20240229 "Complex system design"
```

## Troubleshooting
- Check environment variables if API calls fail
- Ensure Go 1.21+ for building
- Verify config directory permissions

## License
MIT License - See LICENSE file