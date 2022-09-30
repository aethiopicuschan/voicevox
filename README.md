# aethiopicuschan/voicevox

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/voicevox.svg)](https://pkg.go.dev/github.com/aethiopicuschan/voicevox)

Golang client for [voicevox_engine](https://github.com/VOICEVOX/voicevox_engine) API.

## Getting Started

```sh
go get github.com/aethiopicuschan/voicevox
```

## Usage

```go
package main

import "github.com/aethiopicuschan/voicevox"

func main() {
  client := voicevox.NewClient("http", "127.0.0.1:50021")
  ...
}
```
