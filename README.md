# Julian Ramirez AI

**One command. Any agent. Any OS. The Julian Ramirez AI ecosystem -- configured and ready.**

![Release](https://img.shields.io/github/v/release/julianramirezreyes/julian-ai) ![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg) ![Go 1.24+](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white) ![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## What It Does

This is NOT an AI agent installer. Most agents are easy to install. This is an **ecosystem configurator** -- it takes whatever AI coding agent(s) you use and supercharges them with the Julian Ramirez AI stack: persistent memory, Spec-Driven Development workflow, curated coding skills, MCP servers, an AI provider switcher, a teaching-oriented persona with security-first permissions, and per-phase model assignment so each SDD step can run on a different model.

**Before**: "I installed Claude Code / OpenCode / Cursor, but it's just a chatbot that writes code."

**After**: Your agent now has memory, skills, workflow, MCP tools, and a persona that actually teaches you.

### 8 Supported Agents


| Agent               | Delegation Model             | Key Feature                            |
| ------------------- | ---------------------------- | -------------------------------------- |
| **Claude Code**     | Full (Task tool)             | Sub-agents, output styles              |
| **OpenCode**        | Full (multi-mode overlay)    | Per-phase model routing                |
| **Gemini CLI**      | Full (experimental)          | Custom agents in `~/.gemini/agents/`   |
| **Cursor**          | Full (native subagents)      | 9 SDD agents in `~/.cursor/agents/`    |
| **VS Code Copilot** | Full (runSubagent)           | Parallel execution                     |
| **Codex**           | Solo-agent                   | CLI-native, TOML config                |
| **Windsurf**        | Solo-agent                   | Plan Mode, Code Mode, native workflows |
| **Antigravity**     | Solo-agent + Mission Control | Built-in Browser/Terminal sub-agents   |


> **Note**: This project is based on [Gentleman-Programming/gentle-ai](https://github.com/Gentleman-Programming/gentle-ai) (open source). All features from the original are included with the Julian Ramirez persona.

---

## Quick Start

### macOS / Linux

```
curl -fsSL https://raw.githubusercontent.com/julianramirezreyes/julian-ai/main/scripts/install.sh | bash -s -- --method binary
```

### Windows (PowerShell)

```
irm https://raw.githubusercontent.com/julianramirezreyes/julian-ai/main/scripts/install.ps1 | iex
```

### Update path

```
# Opción temporal (para esta sesión):
$env:PATH += ";C:\Users\Admin\AppData\Local\gentle-ai\bin"
```

```
# O permanentemente:
[Environment]::SetEnvironmentVariable('PATH', $env:PATH + ';C:\Users\Admin\AppData\Local\gentle-ai\bin', 'User')
```

This downloads the latest release for your platform and launches the interactive TUI. No Go toolchain required.

### After install: project-level setup

Once your agents are configured, open your AI agent in a project and run these two commands to register the project context:


| Command          | What it does                                                                | When to re-run                                                                 |
| ---------------- | --------------------------------------------------------------------------- | ------------------------------------------------------------------------------ |
| `/sdd-init`      | Detects stack, testing capabilities, activates Strict TDD Mode if available | When your project adds/removes test frameworks, or first time in a new project |
| `skill-registry` | Scans installed skills and project conventions, builds the registry         | After installing/removing skills, or first time in a new project               |


These are **not required** for basic usage. The SDD orchestrator runs `/sdd-init` automatically if it detects no context. But if something changed in your project (new test runner, new dependencies), re-running them manually ensures the agents have up-to-date context.

---

## Install

### Go install (any platform with Go 1.24+)

```
go install github.com/julianramirezreyes/julian-ai/cmd/julian-ai@latest
```

### Windows (PowerShell)

```
# Option 1: PowerShell installer (downloads binary from GitHub Releases)
irm https://raw.githubusercontent.com/julianramirezreyes/julian-ai/main/scripts/install.ps1 | iex

# Option 2: Go install (requires Go 1.24+)
go install github.com/julianramirezreyes/julian-ai/cmd/julian-ai@latest
```

### From releases

Download the binary for your platform from [GitHub Releases](https://github.com/julianramirezreyes/julian-ai/releases).

---

## Para compilar y hacer release

### En tu máquina Linux

1 Compilar linux

```
go build -o julian-ai ./cmd/gentle-ai
```

Compilar windows

```
GOOS=windows GOARCH=amd64 go build -o julian-ai.exe ./cmd/gentle-ai
```

2 Crear el archivo comprimido

```
tar -czf julian-ai_1.0.2_linux_amd64.tar.gz julian-ai
zip julian-ai_1.0.2_windows_amd64.zip julian-ai.exe
```

3 Commit y push

4 Crear la release

 Opción A: con gh (si lo tenés)

```
gh release create v1.0.2 --title "v1.0.1" --notes "Updated tagline to Julian Ramirez AI"
gh release upload v1.0.2 julian-ai_1.0.1_linux_amd64.tar.gz
gh release upload v1.0.2 julian-ai_1.0.1_windows_amd64.zip
```

 Opción B: manualmente desde GitHub UI en el apartado de releases

---

## Documentation


| Topic                                              | Description                                                             |
| -------------------------------------------------- | ----------------------------------------------------------------------- |
| [Intended Usage](docs/intended-usage.md)           | How julian-ai is meant to be used — the mental model                    |
| [Agents](docs/agents.md)                           | Supported agents, feature matrix, config paths, and per-agent notes     |
| [Components, Skills & Presets](docs/components.md) | All components, GGA behavior, skill catalog, and preset definitions     |
| [Usage](docs/usage.md)                             | Persona modes, interactive TUI, CLI flags, and dependency management    |
| [Platforms](docs/platforms.md)                     | Supported platforms, Windows notes, security verification, config paths |
| [Architecture & Development](docs/architecture.md) | Codebase layout, testing, and relationship to Gentleman.Dots            |


---

## Agent Persona

The agent persona is defined in these embedded files (built into the binary):


| Agent                  | Persona File                                       |
| ---------------------- | -------------------------------------------------- |
| Generic (all agents)   | `internal/assets/generic/persona-gentleman.md`     |
| Claude Code            | `internal/assets/claude/persona-gentleman.md`      |
| OpenCode               | `internal/assets/opencode/persona-gentleman.md`    |
| Output Style (Claude)  | `internal/assets/claude/output-style-gentleman.md` |
| Skills Index           | `AGENTS.md`                                        |
| Crear tu propio agente | internal/components/persona/inject.go              |


To modify the Julian Ramirez persona, edit these files and rebuild the binary.

---

Otros archivos de interes


|                               |                                 |
| ----------------------------- | ------------------------------- |
| File                          | Link                            |
| Script de instalacion linux   | scripts/install.sh              |
| Script de instalacion windows | scripts/install.ps1             |
| Julian-ai menu                | internal/tui/styles/styles.go   |
| Julian-ai persona screen      | internal/tui/screens/persona.go |


![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

---

# Modo educativo

Ya no es el comando julian-ai usa el comando gentleman-ai

## Quick Start

### macOS / Linux

```bash
curl -fsSL https://raw.githubusercontent.com/Gentleman-Programming/gentle-ai/main/scripts/install.sh | bash
```

### Windows

**Opción Recomendada (Scoop):**
Si todavía no tenés Scoop, instalalo primero (desde PowerShell):
```powershell
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
irm get.scoop.sh | iex
```

Una vez que tengas Scoop, agregá el bucket e instalá:
```powershell
scoop bucket add gentleman https://github.com/Gentleman-Programming/scoop-bucket
scoop install gentle-ai
```

**Alternativa (Script directo):**
```powershell
irm https://raw.githubusercontent.com/Gentleman-Programming/gentle-ai/main/scripts/install.ps1 | iex
```
*(Nota: Confirmá que se guardó en tu `$env:PATH` para luego poder correrlo desde cualquier lado).*

**Para ejecutar:**
Una vez instalado, simplemente llamá a tu agente desde la consola:
```powershell
gentleman-ai
```

### After install: project-level setup

Once your agents are configured, open your AI agent in a project and run these two commands to register the project context:

| Command | What it does | When to re-run |
|---------|-------------|----------------|
| `/sdd-init` | Detects stack, testing capabilities, activates Strict TDD Mode if available | When your project adds/removes test frameworks, or first time in a new project |
| `skill-registry` | Scans installed skills and project conventions, builds the registry | After installing/removing skills, or first time in a new project |

These are **not required** for basic usage. The SDD orchestrator runs `/sdd-init` automatically if it detects no context. But if something changed in your project (new test runner, new dependencies), re-running them manually ensures the agents have up-to-date context.

---

## Install

### Homebrew (macOS / Linux)

```bash
brew tap Gentleman-Programming/homebrew-tap
brew install gentle-ai
```

### Go install (any platform with Go 1.24+)

```bash
go install github.com/gentleman-programming/gentle-ai/cmd/gentle-ai@latest
```

### Scoop (Windows)

Si todavía no lo tenés, asegurate de instalar Scoop primero:
```powershell
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
irm get.scoop.sh | iex
```

Luego avanzá instalando el agente:
```powershell
scoop bucket add gentleman https://github.com/Gentleman-Programming/scoop-bucket
scoop install gentle-ai
```

**Migrating from PowerShell installer to Scoop?** Remove the old binary first:

```powershell
Remove-Item "$env:LOCALAPPDATA\gentle-ai" -Recurse -Force
# Then install via Scoop as shown above
```

### Windows (PowerShell — alternative)

```powershell
# Option 1: PowerShell installer (downloads binary from GitHub Releases)
irm https://raw.githubusercontent.com/Gentleman-Programming/gentle-ai/main/scripts/install.ps1 | iex

# Option 2: Go install (requires Go 1.24+)
go install github.com/gentleman-programming/gentle-ai/cmd/gentle-ai@latest
```

### From releases

Download the binary for your platform from [GitHub Releases](https://github.com/Gentleman-Programming/gentle-ai/releases).

---

## Backups

Every install, sync, and upgrade automatically snapshots your config files. Backups are **compressed** (tar.gz), **deduplicated** (identical configs are not re-backed up), and **auto-pruned** (keeps the 5 most recent). Pin important backups via the TUI (`p` key) to protect them from pruning.

See [Backup & Rollback Guide](docs/rollback.md) for details.

---

## Documentation

| Topic | Description |
|-------|-------------|
| [Intended Usage](docs/intended-usage.md) | How gentle-ai is meant to be used — the mental model |
| [Agents](docs/agents.md) | Supported agents, feature matrix, config paths, and per-agent notes |
| [Components, Skills & Presets](docs/components.md) | All components, GGA behavior, skill catalog, and preset definitions |
| [Usage](docs/usage.md) | Persona modes, interactive TUI, CLI flags, and dependency management |
| [Backup & Rollback](docs/rollback.md) | Backup retention, compression, dedup, pinning, and restore |
| [Platforms](docs/platforms.md) | Supported platforms, Windows notes, security verification, config paths |
| [Architecture & Development](docs/architecture.md) | Codebase layout, testing, and relationship to Gentleman.Dots |