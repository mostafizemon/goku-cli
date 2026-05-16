# Goku CLI

Goku CLI is a small command-line tool for converting configuration files between
JSON and YAML formats.

## Features

- Convert JSON to YAML
- Convert YAML/YML to JSON
- Detect input format from file extension
- Validate unsupported formats
- Print converted output directly to the terminal
- Built with Cobra for a clean CLI experience

## Supported Formats

| Input | Output |
| --- | --- |
| `.json` | `yaml` / `yml` |
| `.yaml` / `.yml` | `json` |

The input and output formats must be different.

## Project Structure

```text
goku-cli/
├── main.go                         # Application entry point
├── cmd/
│   └── root.go                     # CLI command, flags, and validation
├── internal/
│   └── converter/
│       └── converter.go            # JSON/YAML conversion logic
├── config.json                     # Example JSON file
├── config.yaml                     # Example YAML file
├── go.mod                          # Go module definition
└── go.sum                          # Dependency checksums
```

## Requirements

- Go 1.26.1 or later

## Installation

Clone the repository:

```bash
git clone https://github.com/mostafizemon/goku-cli.git
cd goku-cli
```

Download dependencies:

```bash
go mod download
```

Build the binary:

```bash
go build -o goku
```

## Usage

```bash
./goku -i <input-file> -o <output-format>
```

### Options

| Flag | Short | Description |
| --- | --- | --- |
| `--input` | `-i` | Input file path. Must be `.json`, `.yaml`, or `.yml` |
| `--output` | `-o` | Output format. Must be `json`, `yaml`, or `yml` |

## Examples

Convert JSON to YAML:

```bash
./goku -i config.json -o yaml
```

Convert YAML to JSON:

```bash
./goku -i config.yaml -o json
```

You can also pass the output format with a leading dot:

```bash
./goku -i config.json -o .yaml
```

## How It Works

1. The application starts from `main.go`.
2. `main()` calls `cmd.Execute()`.
3. Cobra parses the CLI flags from `cmd/root.go`.
4. The input file and output format are validated.
5. The input format is detected from the file extension.
6. The input file is read using `os.ReadFile`.
7. `internal/converter.Convert()` converts the file content.
8. The converted result is printed to standard output.

The converter uses an intermediate Go value:

```go
var intermediate interface{}
```

Input data is first unmarshaled into this generic value, then marshaled into the
requested output format.

## Conversion Flow

```text
Input file
   ↓
Read bytes
   ↓
Unmarshal JSON/YAML into Go data
   ↓
Normalize map keys when needed
   ↓
Marshal Go data into JSON/YAML
   ↓
Print result to terminal
```

## Error Handling

Goku CLI returns clear errors when:

- The input file is missing
- The output format is missing
- The input file extension is unsupported
- The output format is unsupported
- The input and output formats are the same
- The input file cannot be read
- The JSON or YAML content is invalid

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI command and flag handling
- [yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) - YAML parsing and encoding
- [encoding/json](https://pkg.go.dev/encoding/json) - JSON parsing and encoding

## License

This project is licensed under the terms of the [LICENSE](LICENSE) file.
