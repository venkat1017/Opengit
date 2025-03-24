# OpenGit 🚀

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://golang.org/doc/go1.20)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A lightweight, educational Git implementation written in Go. OpenGit helps you understand Git's internal workings by implementing core Git functionality from scratch.

## 🌟 Features

- **Core Git Operations**
  - Repository initialization
  - Staging files
  - Committing changes
  - Branching and tagging
  - Basic version control operations

- **Git Plumbing Commands**
  - `hash-object`: Compute object ID
  - `cat-file`: Display object contents
  - `ls-tree`: List tree contents
  - `rev-parse`: Parse revisions

- **Git Porcelain Commands**
  - `init`: Create new repository
  - `add`: Stage files
  - `commit`: Record changes
  - `status`: Show working tree status
  - `log`: Display commit history
  - `checkout`: Switch branches/restore files
  - And more!

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/yourusername/opengit.git
cd opengit

# Build the project
go build

# Initialize a new repository
./opengit init

# Add some files
./opengit add file1.txt

# Create a commit
./opengit commit -m "Initial commit"

# Check status
./opengit status
```

## 📚 Command Reference

### Basic Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `init` | Create a new repository | `opengit init [path]` |
| `add` | Add files to staging | `opengit add <files...>` |
| `commit` | Record changes | `opengit commit -m "message"` |
| `status` | Show working tree status | `opengit status` |
| `log` | Show commit history | `opengit log` |

### Advanced Commands

| Command | Description | Usage |
|---------|-------------|-------|
| `checkout` | Switch branches/restore | `opengit checkout <commit-ish>` |
| `tag` | Create a new tag | `opengit tag <tagname> [commit]` |
| `cat-file` | Show object contents | `opengit cat-file -p <object>` |
| `ls-tree` | List tree contents | `opengit ls-tree <tree-ish>` |

## 🛠️ Project Structure

```
opengit/
├── commands/     # Command implementations
├── objects/      # Git object handling
├── refs/        # Reference management
├── repo/        # Repository operations
├── index/       # Index (staging) management
└── main.go      # Entry point
```

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🎯 Educational Purpose

OpenGit is designed as an educational tool to help developers understand Git's internal mechanisms. While it implements many core Git features, it's not intended for production use. For real version control needs, please use the official Git implementation.

## ⭐ Show Your Support

If you find this project helpful, please consider giving it a star! It helps others discover this educational resource.

## 📬 Contact

Your Name - [@yourtwitterhandle](https://twitter.com/yourtwitterhandle)

Project Link: [https://github.com/yourusername/opengit](https://github.com/yourusername/opengit)