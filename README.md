# Mathematical Expression Parser / Evaluator in Go (Golang)
> By *mhaxanali*, a remaster of [math-evaluate](https://github.com/mhaxanali/math-evaluate) in Go.

## Flow of Program

```
┌──────────────────────────┐
│        START             │
└─────────────┬────────────┘
              │
              ▼
┌──────────────────────────┐
│ Take raw input string    │
│ (e.g. "3 + -2 * 5")      │
└─────────────┬────────────┘
              │
              ▼
┌──────────────────────────┐
│ Tokenize input           │
│ → numbers                │
│ → operators (+ - * /)    │
│ → parentheses            │
└─────────────┬────────────┘
              │
              ▼
┌──────────────────────────┐
│ Normalize tokens         │
│ (fix unary minus,        │
│ merge negative numbers)  │
└─────────────┬────────────┘
              │
              ▼
┌──────────────────────────┐
│ Validate syntax          │
│ - check characters       │
│ - check balanced ()      │
│ - check operator usage   │
└─────────────┬────────────┘
              │
              ▼
     ┌─────────────────┐
     │ parentheses?    │───Yes──────┐
     └───────┬─────────┘            │
             │No                    ▼
             ▼               ┌────────────────────────┐
┌──────────────────────────┐ │ Evaluate inner-most () │
│ Proceed to operator      │ │ recursively            │
│ precedence evaluation    │ └─────────────┬──────────┘
└─────────────┬────────────┘               │
              │                            ▼
              ▼                   replace (...) with result
┌──────────────────────────┐              │
│ Handle * and / first     │ <────────────┘
│ left → right             │
└─────────────┬────────────┘
              │
              ▼
┌──────────────────────────┐
│ Handle + and - next      │
│ left → right             │
└─────────────┬────────────┘
              │
              ▼
┌──────────────────────────┐
│ Single token left?       │
└───────┬────────┬─────────┘
        │Yes     │No (error)
        ▼        ▼
┌───────────────┐ ┌────────────────────────┐
│ Return result │ │ Throw "invalid expr"   │
└───────────────┘ └────────────────────────┘
```

## Why?
I was learning Go and decided to remake one of my python projects in Go and this is just that.

## Usage
- Clone the repo
    ```bash
    git clone https://github.com/mhaxanali/go-math-eval
    ```
- Navigate to the directory
    ```bash
    cd go-math-eval
    ```
- Build the binary file
    ```bash
    go build math-eval/main.go
    ```
- Run the program with expression as argument.
    ```bash
    ./main.exe <expression>
    ```