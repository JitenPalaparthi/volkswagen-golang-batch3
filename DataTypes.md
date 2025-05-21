# Golang Data Types

Go is a statically typed language. Data types in Go are categorized as follows:

---

## 🧮 Basic/defined Types

| Category | Types |
|----------|-------|
| **Boolean** | `bool` (true or false) |
| **String** | `string` (UTF-8 text) |
| **Numeric** | `int`, `int8`, `int16`, `int32`, `int64`<br>`uint`, `uint8` (alias `byte`), `uint16`, `uint32`, `uint64`, `uintptr`<br>`float32`, `float64`<br>`complex64`, `complex128` |

---

## 🪞 Alias Types

| Alias | Original |
|-------|----------|
| `byte` | `uint8` |
| `rune` | `int32` (used to represent a Unicode code point) |

---

## 📦 Aggregate Types

| Type | Description |
|------|-------------|
| `array` | Fixed-size collection of elements of the same type: `[5]int` |
| `slice` | Dynamically-sized, flexible view into an array: `[]int` |
| `struct` | Collection of named fields: `struct { name string; age int }` |
| `map` | Key-value store: `map[string]int` |

---

## 🧩 Reference Types

| Type | Description |
|------|-------------|
| `pointer` | Holds the memory address of a value: `*int` |
| `function` | Functions are first-class values: `func(int, int) int` |
| `channel` | Used for goroutine communication: `chan int` |
| `interface` | Specifies method set: `interface{}` or custom ones |
| `slice` | Backed by array, reference type (as above) |
| `map` | Also a reference type (as above) |

---

## 🔧 Special Type

| Type | Description |
|------|-------------|
| `interface{}` | The **empty interface** — can hold any value |
| `nil` | Zero value for pointers, interfaces, maps, slices, channels, functions |

---

## 🎯 User-Defined Types

You can create your own types based on existing ones:

```go
type Age int
type Person struct {
    Name string
    Age  int
}