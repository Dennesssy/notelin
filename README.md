<div align="center">

# 🕵️ No-Tel-in

**Telemetry & Sentry Management Scanner**

*"Ain't no telling what data they be lyin about collectin"*

[![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)](https://github.com/notelin/notelin/releases)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![macOS](https://img.shields.io/badge/platform-macOS-lightgrey.svg)](https://www.apple.com/macos/)
[![Go](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)

[🚀 Quick Start](#-quick-start) • [📖 Documentation](#-what-it-does) • [🏗️ Architecture](#️-how-its-built) • [🤝 Contributing](#️-want-to-help-build-this)

</div>

---

## 🎯 What is No-Tel-in?

**Your privacy detective for macOS applications**

Ever wonder what your apps are really doing behind the scenes? **No-Tel-in** is your privacy detective that discovers, analyzes, and manages telemetry and data collection from applications on your macOS system. Because transparency shouldn't be optional.

### ✨ Key Features
- 🔍 **Smart Discovery** → Automatically finds apps collecting your data
- 📊 **Risk Assessment** → Scores privacy risk from 0-100 based on evidence  
- 🧹 **Data Cleanup** → Safely removes tracking data and cached information
- 💻 **Beautiful Interface** → Interactive terminal UI that's actually fun to use
- 📈 **Detailed Reports** → Export comprehensive privacy assessments
- 🛡️ **Zero Telemetry** → This tool practices what it preaches

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

## 🎯 What it does

**Think of it as a privacy audit for your Mac**

- 🎪 **Finds the culprits** → Scans your system for apps doing shady stuff
- 🎭 **Shows their tricks** → Reveals exactly what data they're collecting  
- 🎨 **Gives you control** → Clean up their mess with one click
- 🎬 **Makes it visual** → Beautiful terminal interface (yes, terminals can be pretty!)

**Real talk:** This tool helped me discover my "privacy-focused" messaging app was sending data to 12 different analytics companies. Wild, right?

## 🔍 The sneaky stuff we catch

**Your apps are busier than you think:**

- 🎛️ **Remote puppet strings** → Firebase configs that let companies control your app from afar
- 📈 **Usage stalking** → Analytics tracking every click, swipe, and keystroke
- 💥 **Crash snitching** → Error reports with way more info than needed
- 📝 **Session recording** → Apps remembering everything you do
- 🖥️ **System spying** → Logging your computer's setup and environment
- ⚡ **Electron shenanigans** → Those "native" apps that are actually websites in disguise

## 🚀 Quick Start

### Prerequisites
- **macOS 10.15+** (Big Sur or later recommended)
- **Go 1.21+** for building from source
- **Gum** for terminal UI components
  ```bash
  brew install gum
  ```

### Installation Options

#### 🍺 Homebrew (Recommended)
```bash
# Coming soon - Homebrew formula pending
brew install notelin/tap/notelin
```

#### 🔧 Build from Source
```bash
git clone https://github.com/notelin/notelin.git
cd notelin
make build
```

#### 📦 Pre-built Binaries
Download from [GitHub Releases](https://github.com/notelin/notelin/releases)

### Basic Usage

```bash
# Interactive mode (recommended for first-time users)
notelin

# Quick system scan
notelin scan

# Generate detailed privacy report
notelin report --output privacy-assessment.json

# Get help
notelin --help
```

## 👀 What you'll see

**Here's what happened when I ran it on my Mac:**

```bash
$ notelin scan

🔍 Digging through your apps...
Found 47 applications • 12 are up to suspicious stuff

🚨 DEFINITELY SHADY:
• Nicegram Desktop - "This app has trust issues" 
• Spotify - "Knows more about you than your therapist"

⚠️  KINDA SKETCHY:
• Blackbox Terminal - "Electron apps gonna electron"
• Adobe Creative Cloud - "Creative with your data too"
• Zoom - "Your meetings aren't the only thing being recorded"

✅ PROBABLY FINE:
• VS Code - "Telemetry but at least they tell you"
• The Unarchiver - "Actually does what it says"
```

**Plot twist:** Even VS Code was sending more data than I expected!

## 🎯 How we score the sus-ness

### 🚨 **DEFINITELY SHADY** (70-100 sus points)
**These apps are working overtime collecting your data:**
- Can be remote controlled by companies
- Hoovering up data with no way to opt out  
- Caught red-handed violating privacy before
- Using multiple spy techniques at once

### ⚠️ **KINDA SKETCHY** (30-69 sus points)  
**Problematic but not completely evil:**
- Standard tracking (but at least you can turn it off)
- Built on Electron (so basically a browser in disguise)
- Remembers what you do and when
- Reports crashes with extra details

### ✅ **PROBABLY FINE** (0-29 sus points)
**These apps passed the vibe check:**  
- Barely any tracking or gives you control
- Open source so you can see what they're doing
- Actually respects your privacy
- Plays by App Store rules

## 📱 Apps we investigate

**We check pretty much everything suspicious:**

- 💬 **Chat apps** → Telegram, Discord, Slack (spoiler: they're all chatty about your data)
- 👨‍💻 **Dev tools** → VS Code, terminals, IDEs (even your code editor is snooping)  
- 🎨 **Creative stuff** → Adobe everything, design apps (creative with data collection too)
- 🎵 **Media players** → Spotify, streaming apps (they know your music AND your habits)
- 📹 **Video calls** → Zoom, Teams (meeting recordings aren't the only thing saved)
- 🛠️ **Random utilities** → Basically anything that could be collecting data

## 🤝 Our privacy promise 

**We practice what we preach:**

✅ **Your data stays yours** → Everything happens on your Mac, nowhere else  
✅ **No secrets** → This tool doesn't phone home or collect anything  
✅ **Actually helpful** → We don't just point out problems, we help fix them  
✅ **Open book** → You can see exactly what we're doing (it's all open source)  
✅ **Your choice** → We show you what's happening, you decide what to do about it

## 🛠️ Want to help build this?

**Join the privacy revolution!** We'd love your help making this tool even better.

**Get set up in 30 seconds:**
```bash
git clone https://github.com/notelin/notelin.git
cd notelin && go mod download && make build && make test
```

**What you'll need:**
- macOS 10.15+ (Linux coming soon, Windows... maybe someday)
- [Gum](https://github.com/charmbracelet/gum) for the pretty terminal stuff
  ```bash
  brew install gum  # Takes 2 seconds, makes everything beautiful
  ```

**Ideas for contributions:**
- 🔍 New app detection patterns
- 🎨 UI/UX improvements  
- 📊 Better privacy scoring algorithms
- 🐛 Bug fixes (there are always bugs)
- 📖 Documentation improvements

## 🏗️ How it's built

**Simple but effective architecture:**

```
notelin/
├── cmd/notelin/          # The main show (CLI commands)
├── pkg/scanner/          # Detective work (finds suspicious apps)  
├── pkg/analyzer/         # The judge (scores privacy risk)
├── pkg/ui/               # Pretty terminal magic
└── scripts/              # Extra security tools
```

## 🚀 What's next

**The roadmap (in order of "when we'll actually do this"):**

- [x] ✅ macOS app analysis (done!)
- [ ] 🐧 Linux support (working on it)
- [ ] 🪟 Windows support (maybe... Windows is weird)  
- [ ] 🌐 Browser extension analysis (your extensions are probably sketchy too)
- [ ] 📡 Network traffic monitoring (see what they're actually sending)
- [ ] 📄 Automated privacy policy analysis (because who reads those?)
- [ ] 🛡️ Integration with security tools (play nice with others)

## ❓ Questions people always ask

**Q: Will this break my apps?**  
A: Nah, we're read-only detectives by default. The cleanup features are super careful, but maybe backup your important stuff first (you should do that anyway).

**Q: Are you spying on me while checking for spying?**  
A: Absolutely not! Everything happens on your Mac. We literally cannot see what you're doing – that's the whole point.

**Q: What about Apple's own data collection?**  
A: We focus on third-party apps. Apple's telemetry is a whole different beast that needs specialized tools.

**Q: How do you know this stuff is accurate?**  
A: We look for actual evidence – config files, network requests, tracking pixels. We don't guess; we investigate.

**Q: Can I trust you?**  
A: Don't trust us – trust the code! It's all open source, so you can see exactly what we're doing.

## 🎉 Success stories

**Real people, real discoveries:**

> *"Holy sh*t, my 'privacy-focused' messaging app was sending data to 12 different tracking companies! Switched to Signal immediately."*  
> — Privacy researcher who thought they knew better

> *"I'm a developer and I was SHOCKED at what my own tools were doing. This changed how I think about the software I build."*  
> — Software engineer having an existential crisis

> *"Found out my design app was tracking every single project I worked on. Client confidentiality much?"*  
> — Freelance designer who values NDAs

---

## 📞 Get help or say hi

- 🐛 [Found a bug?](https://github.com/notelin/notelin/issues) (we'll fix it fast)
- 💬 [Want to chat?](https://github.com/notelin/notelin/discussions) (we're friendly)  
- 📧 Email us: privacy@notelin.org (we actually reply)
- 🐦 Twitter: [@notelintool](https://twitter.com/notelintool) (for the memes)

---

<div align="center">

**🔐 No-Tel-in**  
*Because privacy shouldn't be a guessing game*

[![⭐ Star us on GitHub](https://img.shields.io/github/stars/notelin/notelin.svg?style=social)](https://github.com/notelin/notelin/stargazers)
[![📄 MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![🚀 Go Report](https://goreportcard.com/badge/github.com/notelin/notelin)](https://goreportcard.com/report/github.com/notelin/notelin)

*Made with ❤️ by people who care about your privacy*

**"Your data, your rules, your choice"**

</div>