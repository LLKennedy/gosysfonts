# gosysfonts

[![GoDoc](https://godoc.org/github.com/LLKennedy/gosysfonts?status.svg)](https://godoc.org/github.com/LLKennedy/gosysfonts)
[![Build Status](https://travis-ci.org/disintegration/imaging.svg?branch=master)](https://travis-ci.org/LLKennedy/gosysfonts)
[![Coverage Status](https://coveralls.io/repos/github/LLKennedy/gosysfonts/badge.svg?branch=master)](https://coveralls.io/github/LLKennedy/gosysfonts?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/LLKennedy/gosysfonts)](https://goreportcard.com/report/github.com/LLKennedy/gosysfonts)

Get system TTF fonts by name in Go. Currently only works for windows.

## Installation
`go get "github.com/LLKennedy/gosysfonts"`

## Basic Usage
```
font, err := gosysfonts.New().GetFont("arial")
```

## Testing
On windows, the simplest way to test is to use the powershell script.

`./test.ps1`

To emulate the testing which occurs in build pipelines, run the following:

`go test . -race -coverprofile=coverage.out;`