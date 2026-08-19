# Text Processor

Process and transform text data with ease - find, replace, and analyze.

[![Go Version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)
[![CI](https://github.com/Qyroxen/text-processor/actions/workflows/ci.yml/badge.svg)](https://github.com/Qyroxen/text-processor/actions/workflows/ci.yml)

> Process and transform text data with ease - find, replace, and analyze.

## What is it?

Text Processor is a command-line tool built with Go that helps developers process and transform text data with ease - find, replace, and analyze. It's designed to be fast, reliable, and easy to use.

## Why?

Every developer needs text processor — but existing tools are either too complex, too slow, or require cloud dependencies. We built Text Processor to be:
- **Fast** — Written in Go for maximum performance
- **Offline** — No cloud dependencies, your data stays on your machine
- **Simple** — Clean CLI interface with sensible defaults
- **Extensible** — Easy to customize and integrate into your workflow

## Features

- **Regular expression support** — Regular expression support
- **Batch processing** — Batch processing
- **Text analysis** — Text analysis
- **Format conversion** — Format conversion
- **Encoding detection** — Encoding detection
- **CLI interface** — CLI interface

## Quick Start

### Prerequisites

- Go 1.23 or later

### Install

```bash
# Install with go install
go install github.com/Qyroxen/text-processor@latest

# Or build from source
git clone https://github.com/Qyroxen/text-processor.git
cd text-processor
go build -o text-processor .
```

### Usage

```bash
# Basic usage
.text-processor --help

# Example
./text-processor replace --input file.txt --pattern 'old' --replacement 'new'
```

## Output

```
Text Processor v1.0.0

Scanning...

✓ Analysis complete
✓ Results ready

{
  "status": "success",
  "results": [...]
}
```

## Configuration

Create a `.config.yaml` file in your project root:

```yaml
# Configuration options
verbose: true
output: json
timeout: 30s
```

## CLI Flags

```
text processor [command]

Flags:
  --path string      Target path (default ".")
  --format string    Output format: json, text (default "text")
  --verbose          Enable verbose output
  --config string    Config file path
  --output string    Output file path
```

## Examples

### Basic Example

```bash
.text-processor --path ./src
```

### Advanced Example

```bash
.text-processor --path ./src --format json --output report.json --verbose
```

### CI/CD Integration

```yaml
# .github/workflows/ci.yml
- name: Run Text Processor
  run: |
    go install github.com/Qyroxen/text-processor@latest
    text-processor --path . --format json --output report.json
```

## Documentation

- [Getting Started](docs/getting-started.md)
- [Configuration](docs/configuration.md)
- [API Reference](docs/api-reference.md)
- [Examples](examples/)
- [Contributing](CONTRIBUTING.md)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

**Qyroxen** - [GitHub](https://github.com/Qyroxen)

---

**Found this useful?** Give it a ⭐ on GitHub!
