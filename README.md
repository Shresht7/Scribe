# `Scribe`

A library to programmatically generate long-form text from structured data. 

<h2 align="center">⚠️ WORK IN PROGRESS ⚠️</h2>

## Features

- Programmatic generation of long-form text from structured data
- Extensible with custom templates

---

## Installation

```bash
go get github.com/Shresht7/Scribe
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/Shresht7/Scribe"
)

func main() {
    doc := scribe.NewDocument()
    doc.AddParagraph("Hello World!")
    fmt.Println(doc)
}
```

---

## Markdown Package

The `markdown` package provides a set of functions to generate Markdown documents.

### Usage

```go
package main

import (
    "fmt"

    "github.com/Shresht7/Scribe/markdown"
)

func main() {
    doc := markdown.NewDocument()
    doc.AddParagraph("Hello World!")
    fmt.Println(doc)
}
```


---

## 📄 License

This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details
