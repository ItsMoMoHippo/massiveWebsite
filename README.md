# Massive Attack Clash DB Website
This repository contains the Go source code for a web app that serves a dynamically rendered website using Templ, TailwindCSS, and Go's standard HTTP tools.

### 📦 Prerequisites

Make sure you have following installed:
- **Go** - Dowload via your package manager or the [official Go Website](https://golang.google.cn/)
- **Templ** - Used for generating and rendering the HTML components 
 Install via:
 ```bash
 go install github.com/a-h/templ/cmd/templ@latest
 ```
- Project dependencies - Install with:
```bash
go mod tidy
```

> [!IMPORTANT]
> Missing these tools will result in compile-time errors

### 🔨 How to run
1. Generate Templ files (HTML to Go components):
```bash
make generate
```
or 
```bash
templ generate
```

2. Run the project
 - To build and run seperately:
```bash
make build
``` 
or
 ```bash
 go build .
 ./main.exe
 ```
 - Or to run in one step:
```bash
make run
```
or
 ```bash
 go run .
 ```

> [!NOTE]
> You can switch between PostgreSQL and SQLite easily for development vs. production — just change the DB config.

> [!NOTE]
> Use [gofmt](https://go.dev/blog/gofmt) (or [gofumpt](https://github.com/mvdan/gofumpt)) for consistent formatting

### 📊 Progress
- [x] static webpage
- [x] Templ to inject information into a HTML page
- [x] connecting to the database
- [x] using HTMX to easily switch out contents in the HTML
- [x] TailwindCSS for looks
- [ ] maybe a little JS/web Components for little effects
- [ ] go live
