# Agent Development Guidelines

## Build/Test Commands
- Build: `go build ./...`
- Run tests: `go test ./...`
- Run single test: `go test -run <TestName> ./<package>`
- Run tests with verbose: `go test -v ./...`
- Test coverage: `go test -cover ./...`

## Code Style Guidelines

### Imports
- Group imports: stdlib, third-party, internal (separated by blank lines)
- Use absolute imports with full module paths
- Avoid unused imports

### Formatting
- Use `gofmt` for code formatting
- Use standard Go conventions for naming (PascalCase for exported, camelCase for unexported)
- Keep lines under 120 characters where possible

### Types & Naming
- Use descriptive names for functions and variables
- Exported types should have clear, documented purposes
- Constants should be UPPERCASE_SNAKE_CASE
- Error types should be prefixed with "Err"

### Error Handling
- Always handle errors explicitly
- Wrap errors with context using `fmt.Errorf("operation: %w", err)`
- Use specific error types for API responses (see `api/errors.go`)
- Return early on errors to reduce nesting

### API Structure
- APIs accept `context.Context` as first parameter
- Use pointer types for optional request fields
- Validate authentication before making requests
- Return structured types, not raw JSON

### Security
- Never log or commit private keys, API secrets, or sensitive data
- Use secure storage for credentials in production
- Validate all inputs before API calls