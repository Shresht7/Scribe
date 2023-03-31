# `Scribe`

A library to programmatically generate markdown.

<h2 align="center">⚠️ WORK IN PROGRESS ⚠️</h2>

## Features

- Programmatically generate markdown from structured data

---

## Installation

```sh
go get github.com/Shresht7/Scribe/markdown
```

## Usage

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

## 📕 API Reference

### `Blockquote`

```go
BlockQuote("text")
```

> text

### `Bold`

```go
Bold("text")
```

**text**

### `BoldItalic`

```go
BoldItalic("text")
```

***text***

### `Code`

```go
Code("text")
```

`text`

### `CodeBlock`

```go
CodeBlock("fmt.Println(\"Hello World!\")", "go")
```

```go
fmt.Println("Hello World!")
```

### `Details`

```go
Details("description", "This is a simple details block")
```

<details>

<summary>description</summary>

This is a simple details block

</details>

### `FencedBlock`

```go
FencedBlock("This is a fenced block")
```

```
This is a fenced block
```

### `FrontMatter`

```go
FrontMatter("yaml", "title: Hello World")
```

```yaml
title: Hello World
```

### `Heading`

```go
Heading(4, "Sub-Heading")
```

#### Sub-Heading

### `HorizontalRule`

```go
HorizontalRule('-', 3)
```

---

### `Image`

```go
Image("Image of a cat", "[ImageUrl]")
```

![Image of a cat](ImageUrl)

### `Italic`

```go
Italic("text")
```

*text*

### `Link`

```go
Link("Link to GitHub Homepage", "https://github.com")
```

[Link to GitHub Homepage](https://github.com)

### `List`

```go
UnorderedList([]string{"Item 1", "Item 2", "Item 3"})
```

- Item 1
- Item 2
- Item 3

```go
OrderedList([]string{"Item 1", "Item 2", "Item 3"})
```

1. Item 1
2. Item 2
3. Item 3

### `Paragraph`

```go
Paragraph("This is a simple paragraph")
```

This is a simple paragraph

### `Strikethrough`

```go
Strikethrough("text")
```

~~text~~

### `Table`

```go
Table([]string{"Name", "Age"}, [][]string{{"John", "21"}, {"Jane", "22"}})
```

| Name | Age |
| ---- | --- |
| John | 21  |
| Jane | 22  |

### `TaskList`

<!-- TODO: TaskList (map[string]bool) -->

---

## 📄 License

This project is licensed under the [MIT License](LICENSE) - see the [LICENSE](LICENSE) file for details
