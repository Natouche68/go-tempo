# Go Tempo

A pomodoro timer written in Go, right in your terminal.

_Go Tempo is written with [Bubbletea](https://github.com/charmbracelet/bubbletea) and [Lipgloss](https://github.com/charmbracelet/lipgloss)._

## Installation

There is compiled binaries for Windows and Linux in the releases, but you can also install it using Go :

```sh
go install github.com/Natouche68/go-tempo@latest
```

## Usage

To launch go-tempo, run :

```sh
go-tempo
```

Use the space bar to pause and resume the timer. Type `r` to reset the timer. To quit, simply type `q` or `ctrl+c`.

## Compiling

When development, you can simply run :

```sh
go run .
```

When you want to build the project, run :

```sh
# For Windows :
GOOS=windows go build -o bin/go-tempo.exe main.go

# For Linux :
GOOS=linux go build -o bin/go-tempo main.go
```

> Don't forget to build to the bin directory, so the compiled binaries are not commited to Github.
