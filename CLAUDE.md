# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Building and Running
- **Build the CLI**: `go build -o heyFrame-cli .`
- **Run directly**: `go run main.go [command]`
- **Run tests**: `go test ./...`
- **Run specific test**: `go test ./[package_path]`
- **Verbose test output**: `go test -v ./...`

### Code Quality
- **Prefer Go 1.24 packages** like slices package
- **Check Go modules**: `go mod tidy`
- **Format code**: `go fmt ./...`
- **Run static analysis**: `go vet ./...`

## Architecture Overview

### Core Command Groups
1. **`account/`** - HeyFrame Account management (login, companies, extensions)
2. **`extension/`** - Extension development tools (build, validate, AI assistance)
3. **`project/`** - HeyFrame project management (creation, configuration, deployment)

### Key Internal Packages
- **`internal/verifier/`** - Code quality tools (PHPStan, ESLint, Twig linting)
- **`internal/account-api/`** - HeyFrame Account API integration
- **`internal/system/`** - System utilities (PHP/Node detection, filesystem)
- **`internal/packagist/`** - Composer/Packagist integration
- **`internal/llm/`** - AI integration (OpenAI, Gemini, OpenRouter)
- **`internal/twigparser/`** - Custom Twig parsing for validation
- **`internal/html/`** - HTML parsing and formatting
- **`internal/git/`** - Git operations
- **`internal/ci/`** - CI/CD configuration generation

### Extension System
The CLI supports three extension types:
- **Platform Plugins**: Standard HeyFrame 6 plugins with `composer.json`
- **HeyFrame Apps**: App system extensions with `manifest.xml`
- **HeyFrame Bundles**: Custom bundle implementations

Extension detection is automatic based on file presence. Each type has specific build, validation, and packaging rules.

### Configuration Files
- **`.heyFrame-cli.yaml`** - Global CLI configuration
- **`.heyFrame-extension.yml`** - Extension-specific settings (schema: `extension/heyFrame-extension-schema.json`)
- **`.heyFrame-project.yml`** - Project-specific settings (schema: `shop/heyFrame-project-schema.json`)

## Development Patterns

### Command Structure
Commands follow Cobra CLI patterns with:
- Main command in `cmd/[group]/[group].go`
- Subcommands in `cmd/[group]/[group]_[subcommand].go`
- Service containers for dependency injection

### Testing Strategy
- Unit tests alongside source files (`*_test.go`)
- Use testify assert for test assertions (`github.com/stretchr/testify/assert`)
- Test data in `testdata/` directories
- Integration tests use real extension samples in `testdata/`
- Prefer assert.ElementsMatch on lists to ignore ordering issues
- Use t.Setenv for environment variables
- Use t.Context() for Context creation in tests

### Error Handling
- Use structured logging via `go.uber.org/zap`
- Context-based logging: `logging.FromContext(ctx)`
- Graceful error reporting to users

### AI Integration
The CLI includes AI-powered features for:
- Twig template upgrades (`extension ai twig-upgrade`)
- Code quality suggestions
- Automated fixes for common issues

LLM providers are configurable (OpenAI, Gemini, OpenRouter) with API key management.

## Extension Development Workflow

### Building Extensions
```bash
# Build extension (auto-detects type)
heyFrame-cli extension build

# Watch mode for development
heyFrame-cli extension admin-watch

# Validate extension
heyFrame-cli extension validate

# Create distribution package
heyFrame-cli extension zip
```

### Project Management
```bash
# Create new project
heyFrame-cli project create

# Build assets
heyFrame-cli project admin-build
heyFrame-cli project storefront-build

# Development servers
heyFrame-cli project admin-watch
heyFrame-cli project storefront-watch
```

## Code Quality Integration

The verifier system provides comprehensive code quality checks:
- **PHP**: PHPStan, PHP-CS-Fixer, Rector
- **JavaScript**: ESLint, Prettier, Stylelint  
- **Twig**: Custom admin Twig linter with auto-fix capabilities
- **Composer**: Dependency validation

Tools are configurable via JSON schemas and run automatically during builds.