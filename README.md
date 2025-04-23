# kubeclip

A Golang CLI tool that copies Kubernetes secrets to your clipboard.

## Overview

`kubeclip` is a simple yet powerful tool designed to streamline the process of accessing Kubernetes secrets. It allows you to quickly copy secret values to your clipboard without the need to decode base64 values manually or write complex kubectl commands.

## Features

- Copy Kubernetes secrets directly to your clipboard
- Support for namespace specification
- Auto-completion for namespaces, secret names, and secret keys
- Simple and intuitive command-line interface

## Installation

### Prerequisites

- Go 1.18 or higher
- kubectl configured with access to your Kubernetes cluster
- For auto-completion: fish, bash, or zsh shell

### Building from source

```bash
git clone https://github.com/yourusername/kubeclip.git
cd kubeclip
go build -o secret main.go
```

Move the binary to your PATH:

```bash
sudo mv secret /usr/local/bin/
```

## Usage

### Basic usage

```bash
# Copy a secret from the default namespace
secret secret-name

# Copy a secret from a specific namespace
secret -n namespace-name secret-name

# Copy a specific key from a secret
secret -n namespace-name secret-name key-name
```

### Shell completion

#### Fish shell

To enable auto-completion for fish shell:

```bash
secret -fish > ~/.config/fish/completions/secret.fish
```

Optionally, you can choose your name binary

```bash
secret -fish -p secret-name > ~/.config/fish/completions/secret.fish
```

#### Bash shell

To enable auto-completion for bash shell:

```bash
secret -bash > /etc/bash_completion.d/secret
```

Optionally, you can choose your name binary

```bash
secret -bash -p secret-name > ~/.config/fish/completions/secret.fish
```

#### Zsh shell

To enable auto-completion for zsh shell:

```bash
secret -zsh > ~/.zsh/completion/_secret
```

Optionally, you can choose your name binary

```bash
secret -zsh -p secret-name > ~/.config/fish/completions/secret.fish
```

## Examples

```bash
# Copy the value of 'API_KEY' from 'api-credentials' secret in 'application' namespace
secret -n application api-credentials API_KEY

# List all available keys in a secret
secret -n monitoring prometheus-creds
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
