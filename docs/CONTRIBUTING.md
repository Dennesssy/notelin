# Contributing to No-Tel-in

Thank you for your interest in contributing to No-Tel-in! This guide will help you get started.

## Code of Conduct

We are committed to providing a welcoming and inspiring community for all. Please read and follow our Code of Conduct.

## How to Contribute

### Reporting Issues

Before creating an issue, please check if it already exists. When reporting:

1. **Use a clear, descriptive title**
2. **Describe the expected behavior**
3. **Describe the actual behavior**  
4. **Include steps to reproduce**
5. **Add system information** (macOS version, No-Tel-in version)
6. **Include relevant logs** (sanitize any personal data)

### Suggesting Features

Feature requests are welcome! Please:

1. **Check existing feature requests** first
2. **Describe the use case** clearly
3. **Explain why this would benefit users**
4. **Consider privacy implications**

### Development Workflow

1. **Fork the repository**
2. **Create a feature branch** from `main`
3. **Make your changes**
4. **Test thoroughly** 
5. **Submit a pull request**

## Development Setup

### Prerequisites

- Go 1.21 or later
- macOS 10.15 or later (for testing)
- [Gum](https://github.com/charmbracelet/gum) installed

### Getting Started

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/notelin.git
cd notelin

# Install dependencies
go mod download

# Run tests
go test ./...

# Build and test locally
go build -o notelin cmd/notelin/main.go
./notelin
```

## Project Structure

```
notelin/
├── cmd/notelin/          # Main CLI application entry point
├── pkg/
│   ├── scanner/          # Application discovery and basic analysis
│   ├── analyzer/         # Deep privacy analysis and assessment  
│   └── ui/               # Terminal user interface components
├── scripts/              # Supporting shell scripts for analysis
├── docs/                 # Documentation
├── build/                # Build scripts and configurations
└── homebrew/             # Homebrew formula
```

## Coding Standards

### Go Code Style

- Follow standard Go formatting (`gofmt`)
- Use meaningful variable and function names
- Add comments for exported functions
- Keep functions focused and small
- Handle errors appropriately

### Example

```go
// AnalyzePrivacyRisk assesses the privacy implications of an application
// based on its data collection patterns and returns a risk score (0-100).
func AnalyzePrivacyRisk(app Application) (int, error) {
    if app.Name == "" {
        return 0, fmt.Errorf("application name is required")
    }
    
    score := calculateBaseScore(app)
    score += assessDataSources(app.DataSources)
    
    return clamp(score, 0, 100), nil
}
```

### Privacy Guidelines

When adding new detection capabilities:

1. **Document what you're detecting** and why it's a privacy concern
2. **Avoid false positives** - be specific in pattern matching  
3. **Respect user data** - never log or transmit sensitive information
4. **Provide context** - explain why something is concerning
5. **Offer solutions** - include remediation suggestions when possible

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/scanner
```

### Test Requirements

- **Unit tests** for all public functions
- **Integration tests** for key workflows  
- **Example data** for testing (use sanitized/mock data only)
- **Error case coverage** for edge conditions

### Test Data

Never include real user data in tests. Use:

```go
func TestAnalyzeApplication(t *testing.T) {
    mockApp := scanner.Application{
        Name: "TestApp",
        Path: "/Applications/TestApp.app",
        // ... other mock data
    }
    
    result, err := analyzer.AnalyzeApplication(mockApp)
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

## Documentation

### Code Documentation

- Document all exported functions with Go doc comments
- Include examples in documentation where helpful
- Update README.md for new features
- Add inline comments for complex logic

### User Documentation

- Update README.md for user-facing changes
- Add examples for new commands or features
- Document any new privacy implications
- Keep language accessible to non-technical users

## Privacy & Security

### Privacy First

No-Tel-in is a privacy tool - we must hold ourselves to the highest standards:

- **Never collect user data** - all analysis is local
- **Respect user privacy** during analysis
- **Sanitize any debug output** 
- **Document data handling** clearly
- **Minimize data retention** (e.g., clear temp files)

### Security Considerations  

- **Validate all inputs** from file system operations
- **Use safe path operations** to prevent directory traversal
- **Handle permissions gracefully** 
- **Avoid running as root** unless absolutely necessary
- **Sanitize shell commands** if used

## Pull Request Process

### Before Submitting

1. **Test your changes** thoroughly
2. **Update documentation** as needed
3. **Add tests** for new functionality
4. **Run linting** and fix any issues
5. **Verify no personal data** is committed

### Pull Request Description

Include in your PR description:

- **What changes were made** and why
- **How to test** the changes
- **Any breaking changes**
- **Screenshots** for UI changes  
- **Performance implications** if applicable

### Review Process

1. **Automated checks** must pass
2. **Code review** by maintainers
3. **Testing** by reviewers when needed
4. **Documentation review** for user-facing changes
5. **Privacy impact assessment** for detection changes

## Application Detection Contributions

### Adding New App Detection

When adding support for detecting a new application:

1. **Research the app's data practices** thoroughly
2. **Document your findings** with evidence
3. **Implement detection logic** with clear patterns
4. **Add appropriate risk assessment**
5. **Include remediation suggestions**
6. **Test with the actual application** if possible

### Detection Pattern Guidelines

```go
// Good: Specific, documented pattern
func (s *Scanner) checkSpecificTelemetry(app *Application) {
    // Check for MyApp's specific telemetry file
    telemetryPath := filepath.Join(supportPath, "myapp-analytics.db")
    if s.pathExists(telemetryPath) {
        app.DataSources = append(app.DataSources, DataSource{
            Type:        "analytics",
            Description: "MyApp stores detailed usage analytics in local database",
            Evidence:    []string{telemetryPath},
            OptOut:      true, // MyApp allows opting out in settings
        })
    }
}

// Bad: Too generic, may cause false positives
func (s *Scanner) checkGenericTelemetry(app *Application) {
    // Don't do this - too broad
    if strings.Contains(app.Name, "app") {
        // This would match almost everything
    }
}
```

## Release Process

Releases are handled by maintainers, but contributors should be aware:

1. **Semantic versioning** (MAJOR.MINOR.PATCH)
2. **Release notes** documenting changes
3. **Homebrew formula** updates
4. **Binary releases** for major platforms

## Getting Help

Need help contributing?

- 💬 [GitHub Discussions](https://github.com/notelin/notelin/discussions)
- 🐛 [Open an Issue](https://github.com/notelin/notelin/issues)
- 📧 Email maintainers: dev@notelin.org

## Recognition

Contributors are recognized in:

- **README.md** acknowledgments section
- **Release notes** for significant contributions  
- **GitHub contributors** page

---

Thank you for helping make digital privacy more accessible to everyone!