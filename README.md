# No-Tel-in

**Telemetry & Sentry Management Scanner**

> *"Ain't no telling what data they be lyin about collectin"*

No-Tel-in is a privacy-focused tool that discovers, analyzes, and manages telemetry and data collection from applications on your macOS system. Because transparency shouldn't be optional.

## Terminal Demo

<div align="center">
<img src="https://raw.githubusercontent.com/notelin/notelin/main/docs/demo.gif" alt="No-Tel-in Demo" width="800"/>
</div>

### Interactive Demo (HTML)

<details>
<summary>🎬 Click to see live terminal demo</summary>

```html
<!DOCTYPE html>
<html>
<head>
    <style>
        .terminal {
            background: #1e1e1e;
            color: #d4d4d4;
            font-family: 'Monaco', 'Menlo', monospace;
            padding: 20px;
            border-radius: 8px;
            margin: 20px 0;
            overflow: auto;
            box-shadow: 0 4px 12px rgba(0,0,0,0.3);
        }
        .prompt { color: #98c379; }
        .command { color: #61dafb; }
        .output { color: #d4d4d4; }
        .highlight { color: #e06c75; }
        .success { color: #98c379; }
        .warning { color: #e5c07b; }
        .header { color: #c678dd; font-weight: bold; }
        .logo { color: #ff6b9d; }
        .blink { animation: blink 1s infinite; }
        @keyframes blink { 0%, 50% { opacity: 1; } 51%, 100% { opacity: 0; } }
    </style>
</head>
<body>

<div class="terminal">
<div class="prompt">$ </div><div class="command">notelin</div><br/>

<div class="logo">
┌─ No-Tel-in Privacy Scanner ───────────────────────────────────┐
│                                                               │
│ ███╗   ██╗ ██████╗       ████████╗███████╗██╗         ██╗███╗ │
│ ████╗  ██║██╔═══██╗      ╚══██╔══╝██╔════╝██║         ██║████╗│
│ ██╔██╗ ██║██║   ██║ █████╗ ██║   █████╗  ██║   █████╗██║██╔██│
│ ██║╚██╗██║██║   ██║ ╚════╝ ██║   ██╔══╝  ██║   ╚════╝██║██║╚█│
│ ██║ ╚████║╚██████╔╝        ██║   ███████╗███████╗      ██║██║ ╚│
│ ╚═╝  ╚═══╝ ╚═════╝         ╚═╝   ╚══════╝╚══════╝      ╚═╝╚═╝  │
│                                                               │
│         📡 Telemetry & Sentry Management Scanner             │
│           "Ain't no telling what data they be lyin about"    │
│                            v1.0.0                            │
└───────────────────────────────────────────────────────────────┘
</div>

<div class="success">=> 🕵️ Start Investigation</div>
<div class="output">   ⚙️ Settings</div> 
<div class="output">   📚 Help</div>
<div class="output">   ❌ Exit</div>

<br/>
<div class="output">🔍 Discovering applications...</div>
<div class="output">[████████████████████████████████████████] 100%</div>
<div class="output">Found 12 applications • 6 with data collection potential</div>

<br/>
<div class="output">
┌─ Suspicious Apps Detected ────────────────────────────────────┐
│                                                               │
│ <div class="highlight">🚨 HIGH SUS</div>                                                   │
│ ├─ 🌐 Nicegram Desktop            <div class="highlight">"They definitely lyin"</div>      │
│ └─ 🎵 Spotify                     <div class="highlight">"Data hungry"</div>              │
│                                                               │
│ <div class="warning">⚠️  MEDIUM SUS</div>                                                │
│ ├─ 💻 Blackbox Terminal           <div class="warning">"Electron = telemetry"</div>     │
│ ├─ 🎨 Adobe Creative Cloud       <div class="warning">"Always watching"</div>          │
│ └─ 💬 Zoom                       <div class="warning">"Meeting your data"</div>        │
│                                                               │
│ <div class="success">✅ LOW SUS</div>                                                    │
│ ├─ 📝 VS Code                    <div class="success">"Open but still tracking"</div>   │
│ └─ 🗂️  The Unarchiver            <div class="success">"Actually clean"</div>           │
│                                                               │
│ [↑↓] Investigate [ENTER] Deep Dive [ESC] Back [Q] Quit      │
└───────────────────────────────────────────────────────────────┘
</div>

<br/>
<div class="success">Selected: 💻 Blackbox Terminal</div>
<div class="output">🔍 Deep diving into application data...</div>

<br/>
<div class="output">
┌─ Application Analysis Results ────────────────────────────────┐
│                                                               │
│ 💻 Application: Blackbox Terminal                            │
│ 📊 Risk Level: <div class="warning">⚠️  MEDIUM</div>                                     │
│ 📡 Telemetry: <div class="highlight">DETECTED</div>                                       │
│ 🗂️  Data Collection: <div class="highlight">Active</div>                                   │
│                                                               │
│ Found Evidence:                                               │
│ • Chromium-based cache (24 files, 2.3MB)                   │
│ • Session tracking (recent directories)                      │
│ • Environment variable logging                               │
│ • Network state persistence                                  │
│                                                               │
│ Privacy Score: <div class="warning">67/100</div> (Higher = Worse)                       │
└───────────────────────────────────────────────────────────────┘
</div>

<br/>
<div class="output">Available Actions:</div>
<div class="success">=> 🧹 Clean Application Data</div>
<div class="output">   📊 Export Detailed Report</div>
<div class="output">   ⚙️ Configure Privacy Settings</div>
<div class="output">   🗑️ Uninstall Application</div>
<div class="output">   ← Back to App List</div>

<br/>
<div class="output">🧹 Cleaning application data...</div>
<div class="output">[████████████████████████████████████████] 100%</div>

<div class="success">
✅ Application data cleaned successfully!
• Cleared 24 cache files (2.3 MB freed)
• Removed recent directory history  
• Cleared session storage
• Environment logs preserved (app needs them)
</div>

<div class="prompt">$ </div><div class="blink">_</div>

</div>

</body>
</html>
```

</details>

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

### Pre-built Binaries
Download from [GitHub Releases](https://github.com/notelin/notelin/releases)

## Quick Start

```bash
# Interactive mode (recommended)
notelin

# Quick system scan
notelin scan

# Generate detailed report
notelin report --output privacy-assessment.json
```

## Example Analysis

```bash
$ notelin scan

🔍 Scanning system for suspicious applications...
Found 47 applications • 12 with data collection potential

🚨 HIGH RISK APPLICATIONS:
• Nicegram Desktop - Firebase remote config, extensive analytics
• Spotify - Usage tracking, demographic profiling

⚠️  MEDIUM RISK APPLICATIONS:  
• Blackbox Terminal - Electron framework, session tracking
• Adobe Creative Cloud - User behavior analytics
• Zoom - Meeting analytics, system monitoring

✅ LOW RISK APPLICATIONS:
• VS Code - Telemetry with opt-out available
• The Unarchiver - Clean, no telemetry detected
```

## Privacy Assessment Criteria

### 🚨 High Risk (70-100 points)
- Remote configuration capabilities
- Extensive analytics without opt-out
- Known privacy violations
- Multiple telemetry mechanisms

### ⚠️ Medium Risk (30-69 points)
- Standard analytics with opt-out
- Electron-based architecture
- Session/usage tracking
- Crash reporting

### ✅ Low Risk (0-29 points)
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

We welcome contributions! Please read our [Contributing Guide](docs/CONTRIBUTING.md) for details.

### Development Setup
```bash
git clone https://github.com/notelin/notelin.git
cd notelin
go mod download
make build
make test
```

## Requirements

- macOS 10.15 or later (Linux support coming soon)
- [Gum](https://github.com/charmbracelet/gum) for TUI components
  ```bash
  brew install gum
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

## Roadmap

- [x] macOS application analysis
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
A: No. All analysis happens locally on your device. We practice what we preach.

**Q: Can it break my applications?**
A: The cleanup features are designed to be safe, but backup important data before cleaning.

**Q: What about Apple's own telemetry?**
A: No-Tel-in focuses on third-party applications. Apple's system telemetry requires different analysis tools.

**Q: How accurate are the privacy assessments?**
A: We use evidence-based analysis of actual data collection patterns. Risk scores are based on detected telemetry mechanisms, not speculation.

## Real-World Impact

> "No-Tel-in helped me discover that my 'privacy-focused' messaging app was actually sending usage data to 12 different analytics services. I switched to Signal the same day." 
> 
> — Privacy researcher

> "As a developer, I was shocked to see how much data my own tools were collecting. No-Tel-in made me more conscious about the software I build."
>
> — Software engineer

## License

MIT License - see [LICENSE](LICENSE) for details.

## Support & Community

- 🐛 [Report Issues](https://github.com/notelin/notelin/issues)
- 💬 [Discussions](https://github.com/notelin/notelin/discussions)  
- 📧 Email: privacy@notelin.org
- 🐦 Twitter: [@notelintool](https://twitter.com/notelintool)

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=notelin/notelin&type=Date)](https://star-history.com/#notelin/notelin&Date)

---

**"Your privacy, your choice, your data"**

*Because you have the right to know what's running on your machine.*

---

<div align="center">

**🔐 No-Tel-in • Transparency shouldn't be optional**

[![GitHub Stars](https://img.shields.io/github/stars/notelin/notelin.svg)](https://github.com/notelin/notelin/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/notelin/notelin)](https://goreportcard.com/report/github.com/notelin/notelin)
[![codecov](https://codecov.io/gh/notelin/notelin/branch/main/graph/badge.svg)](https://codecov.io/gh/notelin/notelin)

Made with ❤️ for digital privacy

</div>

---

*No-Tel-in is not affiliated with any of the applications it analyzes. All trademarks belong to their respective owners.*