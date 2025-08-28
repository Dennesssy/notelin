# No-Tel-in

**Telemetry & Sentry Management Scanner**

> *"Ain't no telling what data they be lyin about collectin"*

No-Tel-in is a privacy-focused tool that discovers, analyzes, and manages telemetry and data collection from applications on your macOS system. Because transparency shouldn't be optional.

## Features

🕵️ **Application Discovery** - Finds non-Apple applications with data collection potential
📊 **Privacy Risk Assessment** - Categorizes apps by their data collection practices  
🔍 **Deep Analysis** - Identifies specific telemetry, analytics, and remote configuration
🧹 **Data Cleanup** - Safely removes tracking data and cached information
📱 **Interactive TUI** - Beautiful terminal interface powered by Gum
📈 **Detailed Reports** - Export comprehensive privacy assessments

## What No-Tel-in Detects

- **Firebase Remote Config** - Apps that can be controlled remotely
- **Analytics & Telemetry** - Usage tracking and behavioral analysis  
- **Crash Reporting** - Sentry and similar error monitoring
- **Session Tracking** - Workflow and usage pattern recording
- **Environment Logging** - System configuration exposure
- **Electron Framework** - Chromium-based data collection mechanisms

## Installation

### Via Homebrew (Recommended)
```bash
brew install notelin/tap/notelin
```

### From Source
```bash
git clone https://github.com/notelin/notelin.git
cd notelin
go build -o notelin cmd/notelin/main.go
```

## Usage

### Interactive Mode
```bash
notelin
```
Launches the full TUI interface for comprehensive privacy analysis.

### Quick Scan
```bash  
notelin scan
```
Performs a rapid system scan and displays results.

### Generate Report
```bash
notelin report
```
Creates a detailed privacy assessment report.

## Example Output

```
┌─ Suspicious Apps Detected ────────────────────────────────────┐
│                                                               │
│ 🚨 HIGH SUS                                                   │
│ ├─ 🌐 Nicegram Desktop            "They definitely lyin"      │
│ └─ 🎵 Spotify                     "Data hungry"              │
│                                                               │
│ ⚠️  MEDIUM SUS                                                │
│ ├─ 💻 Blackbox Terminal           "Electron = telemetry"     │
│ ├─ 🎨 Adobe Creative Cloud       "Always watching"          │
│ └─ 💬 Zoom                       "Meeting your data"        │
│                                                               │
│ ✅ LOW SUS                                                    │
│ ├─ 📝 VS Code                    "Open but still tracking"   │
│ └─ 🗂️  The Unarchiver            "Actually clean"           │
└───────────────────────────────────────────────────────────────┘
```

## Privacy Assessment Criteria

### 🚨 High Risk
- Remote configuration capabilities
- Extensive analytics without opt-out
- Known privacy violations
- Multiple telemetry mechanisms

### ⚠️ Medium Risk  
- Standard analytics with opt-out
- Electron-based architecture
- Session/usage tracking
- Crash reporting

### ✅ Low Risk
- Minimal or no telemetry
- Open source with transparency
- Privacy-focused design
- App Store restrictions compliant

## Supported Applications

No-Tel-in automatically detects and analyzes:
- **Messaging Apps**: Telegram clients, Discord, Slack
- **Development Tools**: VS Code, terminal emulators, IDEs
- **Creative Software**: Adobe Suite, design tools
- **Media Players**: Spotify, streaming applications
- **Communication**: Zoom, video conferencing tools
- **Utilities**: Any app with potential data collection

## Privacy Philosophy

No-Tel-in operates on these principles:

1. **Your Data, Your Choice** - You should know what data is collected
2. **Transparency First** - No hidden analysis or reporting  
3. **Local Operation** - All scanning happens on your device
4. **Actionable Results** - Not just detection, but remediation
5. **Privacy Preservation** - Tool itself collects no data

## Contributing

We welcome contributions! Please read our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup
```bash
git clone https://github.com/notelin/notelin.git
cd notelin
go mod download
go run cmd/notelin/main.go
```

## Architecture

```
notelin/
├── cmd/notelin/          # Main CLI application
├── pkg/scanner/          # Application discovery & analysis  
├── pkg/analyzer/         # Privacy assessment engine
├── pkg/ui/               # Terminal user interface
└── scripts/              # Security analysis scripts
```

## Requirements

- macOS 10.15 or later
- [Gum](https://github.com/charmbracelet/gum) for TUI components

## Roadmap

- [ ] Linux support
- [ ] Windows support  
- [ ] Browser extension analysis
- [ ] Network traffic monitoring
- [ ] Automated privacy policy analysis
- [ ] Integration with security tools

## FAQ

**Q: Is No-Tel-in safe to use?**
A: Yes. No-Tel-in only reads data, never modifies applications without explicit consent.

**Q: Does No-Tel-in collect any data?**  
A: No. All analysis happens locally on your device.

**Q: Can it break my applications?**
A: The cleanup features are designed to be safe, but backup important data before cleaning.

**Q: What about Apple's own telemetry?**
A: No-Tel-in focuses on third-party applications. Apple's telemetry is system-level and requires different tools.

## License

MIT License - see [LICENSE](LICENSE) for details.

## Support

- 🐛 [Report Issues](https://github.com/notelin/notelin/issues)
- 💬 [Discussions](https://github.com/notelin/notelin/discussions)  
- 📧 Email: privacy@notelin.org

---

**"Because you have the right to know what's running on your machine"**

---

*No-Tel-in is not affiliated with any of the applications it analyzes. All trademarks belong to their respective owners.*