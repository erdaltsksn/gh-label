# GitHub Label

[![PkgGoDev](https://pkg.go.dev/badge/github.com/erdaltsksn/gh-label)](https://pkg.go.dev/github.com/erdaltsksn/gh-label)
![Go](https://github.com/erdaltsksn/gh-label/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/gh-label)](https://goreportcard.com/report/github.com/erdaltsksn/gh-label)

Github Label AKA `gh-label` is a cli app to manage GitHub Issue Labels.

## Features

- Cross-Platform
- Secure token management using key managers
- Working on public and private repositories
- Predefined [label list](https://github.com/erdaltsksn/gh-label/tree/main/labels)

## Requirements

- GitHub Personal Access Token with `repo` scope

## Getting Started

### 1. Create a new personal access token

This is required. You can generate a new one from [here](https://github.com/settings/tokens/new).

### 2. Install and configure the application

```sh
brew install erdaltsksn/tap/gh-label
gh-label config --token <GITHUB_TOKEN>
```

### 3. Generate the labels using a predefined list

```sh
gh-label generate --repo erdaltsksn/playground --list "insane"
```

## Installation

### Using Homebrew

```sh
brew install erdaltsksn/tap/gh-label
```

### Using Go Modules

```sh
go get github.com/erdaltsksn/gh-label
```

## Updating / Upgrading

### Using Homebrew

```sh
brew upgrade erdaltsksn/tap/gh-label
```

### Using Go Modules

```sh
go get -u github.com/erdaltsksn/gh-label
```

## Usage

You may find the documentation for [each command](docs/gh-label.md) inside the
[docs](docs) folder.

### Getting Help

```sh
gh-label --help
gh-label [command] --help
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](docs/CONTRIBUTING.md) for more information.

## Security Policy

If you discover a security vulnerability within this project, please follow our
[Security Policy Guide](docs/SECURITY.md).

## Code of Conduct

This project adheres to the Contributor Covenant [Code of Conduct](docs/CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
