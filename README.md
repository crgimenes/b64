# b64
[![Build Status](https://travis-ci.org/crgimenes/b64.svg?branch=master)](https://travis-ci.org/crgimenes/b64)
[![Go Report Card](https://goreportcard.com/badge/github.com/crgimenes/b64)](https://goreportcard.com/report/github.com/crgimenes/b64)
[![codecov](https://codecov.io/gh/crgimenes/b64/branch/master/graph/badge.svg)](https://codecov.io/gh/crgimenes/b64)
[![GoDoc](https://godoc.org/github.com/crgimenes/b64?status.png)](https://godoc.org/github.com/crgimenes/b64)
[![Go project version](https://badge.fury.io/go/github.com%2Fcrgimenes%2Fb64.svg)](https://badge.fury.io/go/github.com%2Fcrgimenes%2Fb64)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

A small utility to convert files to and from base64


In my work I need from time to time to convert a file to base64 to send via a REST API using curl and sometimes convert back to the original format. I could do this using uuencode or with a single line in Python but it is more fun and instructive to write my own utility.


## Install

```
go get github.com/go-br/b64
```

## Examples of use

Encodes fileName to base64 and sends the content to screen

```
b64 fileName
```

Decode a base64 file and show the content on the screen

```
b64 -decode fileName
```

Encodes a file to base64 and saves the contents to a file pointed by *-output*

```
b64 -output=fileName.b64 fileName
```

## Contributing

- Fork the repo on GitHub
- Clone the project to your own machine
- Create a *branch* with your modifications `git checkout -b fantastic-feature`.
- Then _commit_ your changes `git commit -m 'Implementation of new fantastic feature'`
- Make a _push_ to your _branch_ `git push origin fantastic-feature`.
- Submit a **Pull Request** so that we can review your changes
