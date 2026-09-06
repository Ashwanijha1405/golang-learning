# ⚙️ Project Config Parser (.env Parser)

A lightweight, robust configuration and `.env` file parser written in **Go**.

This project demonstrates how to process structured configuration text by combining Go's **`bufio`**, **`strings`**, and **`regexp`** standard library packages to parse key-value pairs, handle quoted strings, strip inline comments, ignore empty lines, and handle edge cases gracefully.

> **Learn → Build → Break → Debug → Understand**

---

## 🚀 What This Project Does

The program takes a multi-line environment configuration string:

```ini
# Application Configuration
APP_NAME = "My Cool App"
APP_VERSION="1,0,2-beta" #Version with quotes
PORT=8000
DEBUG_MODE="true"
# Database Settings
DB_HOST=localhost
DB_USER=admin
DB_PASSWORD = "p@S$W Ord With Sp@ces!" #Quoted password
API_ENDPOINT = https://api.example.com/v1

# An empty Value 
EMPTY_KEY= 
ANOTHER_KEY_NO_VALUE =
```

And converts it into a Go map (`map[string]string`):

```text
APP_NAME             → "My Cool App"
APP_VERSION          → "1,0,2-beta"
PORT                 → "8000"
DEBUG_MODE           → "true"
DB_HOST              → "localhost"
DB_USER              → "admin"
DB_PASSWORD          → "p@S$W Ord With Sp@ces!"
API_ENDPOINT         → "https://api.example.com/v1"
EMPTY_KEY            → ""
ANOTHER_KEY_NO_VALUE → ""
```

---

## 🌟 Key Features

- 🧹 **Whitespace Tolerant:** Handles flexible spacing around keys, `=`, and values (e.g., `KEY = value` and `KEY=value`).
- 💬 **Comment Support:** Ignores full-line comments (`# Comment`) and strips inline comments (`KEY=value # comment`).
- 🔤 **Quoted String Handling:** Supports both single quotes (`'...'`) and double quotes (`"..."`), preserving spaces and special characters inside quotes.
- 🔗 **Raw / Unquoted Values:** Parses unquoted values (numbers, URLs, booleans, hosts).
- 📭 **Empty Values:** Safely handles empty keys (`EMPTY_KEY=` or `EMPTY_KEY= `).
- 🛡️ **Validation & Error Handling:** Identifies invalid lines without crashing and checks for scanner stream errors.

---

# 🔄 System Architecture & Execution Flow

```text
                      Raw Configuration String
                                 │
                                 ▼
                        strings.NewReader()
                                 │
                                 ▼
                         bufio.NewScanner()
                                 │
                 ┌───────────────┴───────────────┐
                 │       Line-by-Line Loop       │
                 │         scanner.Scan()        │
                 └───────────────┬───────────────┘
                                 │
                                 ▼
                        strings.TrimSpace()
                                 │
                  ┌──────────────┴──────────────┐
                  ▼                             ▼
       Is Empty or Starts with #           Has Content
                  │                             │
               [SKIP]                           ▼
                                    re.FindStringSubmatch()
                                                │
                                  ┌─────────────┴─────────────┐
                                  ▼                           ▼
                             No Match                    Match Found
                                  │                           │
                        [Log Invalid Line]                    ▼
                                                      Extract Key & Value
                                                    (Group 2 / 3 / 4 check)
                                                              │
                                                              ▼
                                                     Store in config map
                                                              │
                                                              ▼
                                                    Return (map, error)
```

---

# 🧠 Core Go Concepts & Theory

---

## 1. 📖 Streaming Input with `bufio.Scanner` & `strings.NewReader`

### Theory
Instead of splitting a giant string into memory all at once with `strings.Split()`, Go provides `bufio.Scanner` for memory-efficient, line-by-line reading.

```go
scanner := bufio.NewScanner(strings.NewReader(content))
for scanner.Scan() {
    line := scanner.Text()
    // process line
}
if err := scanner.Err(); err != nil {
    return nil, err
}
```

### Visual Representation

```text
 Raw String: "APP=Go\nPORT=8080\n"
         │
         ▼ strings.NewReader()
 [ Reader Stream ]
         │
         ▼ bufio.NewScanner()
 ┌────────────────────────┐
 │ 1. Scan() → "APP=Go"   │ → scanner.Text()
 │ 2. Scan() → "PORT=8080"│ → scanner.Text()
 │ 3. Scan() → false (EOF)│
 └────────────────────────┘
```

- **`strings.NewReader(content)`**: Converts a `string` into an `io.Reader` compatible stream.
- **`scanner.Scan()`**: Advances the scanner to the next token (by default, next line). Returns `false` on EOF or error.
- **`scanner.Text()`**: Returns the current line as a Go `string`.
- **`scanner.Err()`**: Checks if any non-EOF error occurred during scanning.

---

## 2. ✂️ Line Cleaning with `strings` Package

Before applying regex matching, the line is sanitized:

```go
trimmedLine := strings.TrimSpace(line)
if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
    continue
}
```

```text
Raw Line:       "   # Database Settings    "
                     ↓ strings.TrimSpace()
Trimmed:        "# Database Settings"
                     ↓ strings.HasPrefix("#") == true
Result:         SKIP (Comment)
```

- **`strings.TrimSpace(line)`**: Strips leading and trailing whitespace (`\t`, `\n`, `\r`, ` `).
- **`strings.HasPrefix(trimmedLine, "#")`**: Efficiently checks if the line is a comment line.

---

## 3. 🔍 Regular Expression Breakdown (`regexp`)

The heart of the parser is the regular expression:

```go
re := regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(?:'([^']*)'|"([^"]*)"|([^#\s]*))?(?:\s*#.*)?$`)
```

### Anatomy of the Regex

```text
 ^\s* ([\w.-]+) \s*=\s* (?: '([^']*)' | "([^"]*)" | ([^#\s]*) )? (?:\s*#.*)? $
 │ │      │        │    │ │    │          │             │     │      │     │ │
 │ │      │        │    │ │    │          │             │     │      │     │ └─ End of line
 │ │      │        │    │ │    │          │             │     │      │     └─── Match inline comment (optional)
 │ │      │        │    │ │    │          │             │     │      └───────── Optional value group
 │ │      │        │    │ │    │          │             │     └──────────────── Group 4: Unquoted value
 │ │      │        │    │ │    │          └─────────────┴────────────────────── Group 3: Double-quoted value
 │ │      │        │    │ └────┴─────────────────────────────────────────────── Group 2: Single-quoted value
 │ │      │        └────┴────────────────────────────────────────────────────── Equals sign with optional spaces
 │ │      └──────────────────────────────────────────────────────────────────── Group 1: Key name
 │ └─────────────────────────────────────────────────────────────────────────── Leading whitespace
 └───────────────────────────────────────────────────────────────────────────── Start of line
```

### Token-by-Token Explanation

| Token | Meaning | Purpose | Example Match |
|---|---|---|---|
| `^` | Start anchor | Asserts the start of the string | — |
| `\s*` | 0+ Whitespace | Tolerates leading spaces/tabs | `"   "` |
| `([\w.-]+)` | **Group 1** | Captures alphanumeric, `_`, `.`, `-` as Key | `DB_PASSWORD`, `api.host` |
| `\s*=\s*` | Equals with spaces | Matches the assignment operator | `=`, ` = `, ` =` |
| `(?: ... )?` | Optional non-capturing group | Makes the entire value part optional | Allows empty values |
| `'([^']*)'` | **Group 2** | Captures content between single quotes | `'hello world'` → `hello world` |
| `\|` | OR operator | Alternation between single, double, unquoted | — |
| `"([^"]*)"` | **Group 3** | Captures content between double quotes | `"p@S$W Ord!"` → `p@S$W Ord!` |
| `([^#\s]*)` | **Group 4** | Captures unquoted chars (no `#` or space) | `8000`, `localhost`, `true` |
| `(?:\s*#.*)?` | Optional comment group | Strips inline comments after values | ` # Version with quotes` |
| `$` | End anchor | Asserts the end of the string | — |

---

## 4. 🎯 Capture Group Unpacking & Value Selection

When `re.FindStringSubmatch(trimmedLine)` executes, it returns a slice of strings:

```text
matches[0] → Full matched line
matches[1] → Key (Group 1)
matches[2] → Single-quoted value (Group 2)
matches[3] → Double-quoted value (Group 3)
matches[4] → Unquoted value (Group 4)
```

### Selection Logic

```go
key := matches[1]
var value string 

if matches[2] != "" {
    value = matches[2] // Single quotes matched
} else if matches[3] != "" {
    value = matches[3] // Double quotes matched
} else {
    value = matches[4] // Unquoted or empty value
}

config[key] = value
```

### Capture Group Truth Table

| Line Example | `matches[1]` (Key) | `matches[2]` (Single) | `matches[3]` (Double) | `matches[4]` (Unquoted) | Resolved `value` |
|---|---|---|---|---|---|
| `APP_NAME = "My Cool App"` | `APP_NAME` | `""` | `"My Cool App"` | `""` | `"My Cool App"` |
| `SECRET = 's3cr3t!'` | `SECRET` | `"s3cr3t!"` | `""` | `""` | `"s3cr3t!"` |
| `PORT=8000` | `PORT` | `""` | `""` | `"8000"` | `"8000"` |
| `EMPTY_KEY=` | `EMPTY_KEY` | `""` | `""` | `""` | `""` |
| `EMPTY_QUOTES=""` | `EMPTY_QUOTES` | `""` | `""` | `""` | `""` |

---

## 5. 🗺️ Data Storage with `map[string]string`

Maps in Go provide fast $O(1)$ average time complexity for lookups, insertions, and updates.

```go
config := make(map[string]string)
config[key] = value
```

```text
                config map
        ┌──────────────────────┬────────────────────────┐
        │ Key                  │ Value                  │
        ├──────────────────────┼────────────────────────┤
        │ "APP_NAME"           │ "My Cool App"          │
        │ "PORT"               │ "8000"                 │
        │ "DB_HOST"            │ "localhost"            │
        │ "EMPTY_KEY"          │ ""                     │
        └──────────────────────┴────────────────────────┘
```

---

## 6. 🖨️ Formatted Output with `%q`

In `main()`:

```go
for k, v := range config {
    fmt.Printf("%s=%q\n", k, v)
}
```

- `%s` prints raw unquoted string (e.g. `APP_NAME`).
- `%q` prints a safely **Go-quoted string literal** (e.g. `"My Cool App"` or `""`).
- `%q` makes empty strings and values with spaces immediately obvious when inspecting output.

---

# 📝 Step-by-Step Line Parsing Walkthrough

Here is a trace of how the parser evaluates each line in the sample input:

### 1. `APP_NAME = "My Cool App"`
- **Trimmed**: `APP_NAME = "My Cool App"`
- **Matched Groups**:
  - Key: `APP_NAME`
  - Group 3 (Double): `My Cool App`
- **Result**: `config["APP_NAME"] = "My Cool App"`

### 2. `APP_VERSION="1,0,2-beta" #Version with quotes`
- **Trimmed**: `APP_VERSION="1,0,2-beta" #Version with quotes`
- **Matched Groups**:
  - Key: `APP_VERSION`
  - Group 3 (Double): `1,0,2-beta`
  - Inline comment `#Version with quotes` is ignored by `(?:\s*#.*)?`
- **Result**: `config["APP_VERSION"] = "1,0,2-beta"`

### 3. `PORT=8000`
- **Trimmed**: `PORT=8000`
- **Matched Groups**:
  - Key: `PORT`
  - Group 4 (Unquoted): `8000`
- **Result**: `config["PORT"] = "8000"`

### 4. `DB_PASSWORD = "p@S$W Ord With Sp@ces!" #Quoted password`
- **Trimmed**: `DB_PASSWORD = "p@S$W Ord With Sp@ces!" #Quoted password`
- **Matched Groups**:
  - Key: `DB_PASSWORD`
  - Group 3 (Double): `p@S$W Ord With Sp@ces!`
  - Quotes preserve special characters (`@`, `$`, `!`) and spaces.
- **Result**: `config["DB_PASSWORD"] = "p@S$W Ord With Sp@ces!"`

### 5. `API_ENDPOINT = https://api.example.com/v1`
- **Trimmed**: `API_ENDPOINT = https://api.example.com/v1`
- **Matched Groups**:
  - Key: `API_ENDPOINT`
  - Group 4 (Unquoted): `https://api.example.com/v1`
- **Result**: `config["API_ENDPOINT"] = "https://api.example.com/v1"`

### 6. `EMPTY_KEY= `
- **Trimmed**: `EMPTY_KEY=`
- **Matched Groups**:
  - Key: `EMPTY_KEY`
  - Group 2, 3, 4: `""`
- **Result**: `config["EMPTY_KEY"] = ""`

---

# 🛠️ Project Structure

```text
project-config-parser/
├── main.go      # Complete parser implementation and demo
└── readme.md    # Comprehensive documentation and theory
```

---

# ▶️ How to Run

Ensure Go is installed:

```bash
go version
```

Execute the parser:

```bash
go run main.go
```

### Expected Output

```text
APP_NAME="My Cool App"
APP_VERSION="1,0,2-beta"
PORT="8000"
DEBUG_MODE="true"
DB_HOST="localhost"
DB_USER="admin"
DB_PASSWORD="p@S$W Ord With Sp@ces!"
API_ENDPOINT="https://api.example.com/v1"
EMPTY_KEY=""
ANOTHER_KEY_NO_VALUE=""
```

*(Note: Map iteration order in Go is randomized by design, so the keys may appear in varying order across runs).*

---

# 📚 Summary & Key Takeaways

| Go Feature / Function | Role in Project | Why We Use It |
|---|---|---|
| `bufio.NewScanner` | Line-by-line reading | Memory efficient streaming of text data |
| `strings.NewReader` | Stream adaptor | Wraps a `string` to satisfy `io.Reader` |
| `strings.TrimSpace` | Whitespace stripping | Eliminates extraneous indentation and newlines |
| `strings.HasPrefix` | Comment filtering | Quickly skips lines starting with `#` |
| `regexp.MustCompile` | Regex compilation | Pre-compiles pattern for fast repeated matching |
| `re.FindStringSubmatch` | Group extraction | Extracts captured keys, quoted strings, and unquoted values |
| `map[string]string` | Key-value store | Provides direct $O(1)$ associative lookup for config values |
| `fmt.Printf` with `%q` | Safe display | Displays exact strings with quotes, making empty values visible |
| `scanner.Err()` | Stream validation | Catches any I/O or tokenization errors during scanning |

---

<div align="center">

### `Learn → Build → Break → Understand → Rebuild`

</div>
