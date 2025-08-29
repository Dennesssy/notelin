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

## 🎬 Live Demo

**See No-Tel-in in action discovering and disabling Gemini Code telemetry:**

<div align="center">

### [🖥️ View Interactive Terminal Demo](https://dennesssy.github.io/notelin/terminal-demo.html)

*Click above to see a realistic workflow of discovering and disabling app telemetry*

</div>

<details>
<summary>📖 What the demo shows</summary>

**Complete workflow demonstration:**
- 🚀 **App Launch** → Beautiful terminal interface with ASCII art
- 🔍 **Discovery Phase** → Scanning system and finding 23 apps with data collection
- 🎯 **Risk Assessment** → Apps categorized by privacy risk (High/Medium/Low)
- 🤖 **Deep Analysis** → Detailed investigation of Gemini Code's telemetry
- 🛡️ **Privacy Action** → Disabling analytics, Firebase, and Sentry tracking
- ✅ **Results** → Privacy score improved from 85/100 to 15/100

**Real evidence found:**
- Google Analytics integration
- Firebase Remote Config (A/B testing)
- Crash reporting to Sentry  
- Code completion analytics
- Session duration logging
- Feature usage metrics

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