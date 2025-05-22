# Risor-DocBot

An AI-powered documentation assistant for [Risor](https://github.com/risor-io/risor) code using the [Dive Agents](https://github.com/diveagents/dive) framework.

## Overview

Risor-DocBot is a tool that automatically analyzes and documents Risor code files. It leverages OpenAI's models through the Dive Agents framework to generate comprehensive documentation for Risor scripts and modules.

## Documentation Format

Risor-DocBot uses a documentation format documented in [risor-docgen](https://github.com/rubiojr/risor-docgen).

## Features

- Automatic documentation generation for Risor code files
- Support for documenting functions, objects, methods, and constructors
- Recognition of Risor-specific syntax and patterns
- Integration with the Dive Agents framework for AI-powered documentation

## Installation

### Prerequisites

- Go 1.24 or later
- An OpenAI API key

### Building from source

```bash
git clone https://github.com/rubiojr/risor-docbot.git
cd risor-docbot
go build ./cmd/risor-docbot
```

## Usage

```bash
export OPENAI_API_KEY=your-api-key
export OPENAI_ENDPOINT=https://api.openai.com/v1  # optional, defaults to OpenAI's endpoint
# Or an Azure AI endpoint
# OPENAI_ENDPOINT=https://your-azure-endpoint.openai.azure.com/openai/deployments/gpt-4.1/chat/completions?api-version=2025-01-01-preview


# Basic usage
./risor-docbot path/to/your/risor/file.risor

# Using command line flags
./risor-docbot --file path/to/your/risor/file.risor
./risor-docbot -f path/to/your/risor/file.risor

# Enable verbose output
./risor-docbot --verbose path/to/your/risor/file.risor
```

The tool will analyze the specified Risor file and update it with appropriate documentation comments.

### Command Line Options

- `--file, -f`: Path to the Risor file to document
- `--model, -m`: OpenAI model to use (default: "gpt-4o")
- `--verbose, -v`: Enable verbose output
- `--help, -h`: Show help
- `--version, -v`: Show version information

## How It Works

Risor-DocBot:

1. Reads the specified Risor code file
2. Sends the code to an AI assistant through the Dive Agents framework
3. The AI agent analyzes the code structure, identifying functions, objects, methods, etc.
4. Documentation comments are generated for each code element
5. The file is updated with the newly generated documentation

## Example

The bot takes a `.risor` file like:

```go
// Example file to be automatically documented with docbot.

// global.function
func fn1(arg1) {
}

// @object.constructor(obj)
func new() {
  // @object.class
  obj := {
  }

  // object.method
  obj.m1 = func() {
  }
}
```

And provides a comprehensive documentation for each element:

```go
// Example file to be automatically documented with docbot.

// Function `fn1`
// This is an example global function.
//
// Parameters:
// - arg1: Description of argument 1.
//
// @global.function
func fn1(arg1) {
}

// Constructor for `obj`
// Initializes and returns a new `obj`.
//
// Returns: A new instance of `obj`.
// @object.constructor(obj)
func new() {
  // Class `obj`
  // Represents an example object.
  //
  // @object.class
  obj := {
  }

  // Method `m1` in `obj`
  // This is an example method for `obj`.
  //
  // @object.method
  obj.m1 = func() {
  }
}
```

Which can then be converted to markdown documentation with [risor-docgen](https://github.com/rubiojr/risor-docgen).

## Environment Variables

- `OPENAI_API_KEY` (required): Your OpenAI API key
- `OPENAI_ENDPOINT` (optional): The OpenAI API endpoint to use

## License

[MIT License](LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
