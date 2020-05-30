# GitHub Label

[![GoDoc](https://godoc.org/github.com/erdaltsksn/gh-label?status.svg)](https://godoc.org/github.com/erdaltsksn/gh-label)
![Go](https://github.com/erdaltsksn/gh-label/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/erdaltsksn/gh-label)](https://goreportcard.com/report/github.com/erdaltsksn/gh-label)

`gh-label` provides us a cli to manage GitHub issue labels.

## Features

- Cross-Platform
- Secure token management using key managers
- Working on public and private repositories
- Insane label list

## Getting Started

### 1. Create a new personal access token

This is required. You can generate a new one from [here](https://github.com/settings/tokens/new).

### 2. Install and configure the application

```sh
go get github.com/erdaltsksn/gh-label
gh-label config --token <GITHUB_TOKEN>
```

### 3. Generate the labels using a predefined list

```sh
gh-label generate --repo erdaltsksn/playground --list "insane"
```

## Installation

```sh
go get github.com/erdaltsksn/gh-label
gh-label config --token <GITHUB_TOKEN>
```

## Updating

```sh
go get -u github.com/erdaltsksn/gh-label
```

## Usage

### Getting Help

```sh
gh-label --help
gh-label [command] --help
```

### Export labels from a repository to a file

```sh
gh-label export --repo erdaltsksn/playground
```

This will export the labels and write them into a file at the current directory.

**You can export the labels into a file by specifying file path.**

```sh
gh-label export --repo erdaltsksn/playground --out=$HOME/Desktop/mylabels.json
```

### Generate the labels using a predefined list

```sh
gh-label generate --repo erdaltsksn/playground --list "insane"
```

**You may use your custom file as a list.**

```sh
gh-label  generate --repo erdaltsksn/playground --file $HOME/my-labels.json
```

**DANGER:** You may add `--force` parameter if you want to delete all labels
before generating new labels.

```sh
gh-label  generate --repo erdaltsksn/playground --list "insane" --force
```

## Contributing

If you want to contribute to this project and make it better, your help is very
welcome. See [CONTRIBUTING](docs/CONTRIBUTING.md) for more information.

## Disclaimer

In no event shall we be liable to you or any third parties for any special,
punitive, incidental, indirect or consequential damages of any kind, or any
damages whatsoever, including, without limitation, those resulting from loss of
use, data or profits, and on any theory of liability, arising out of or in
connection with the use of this software.
