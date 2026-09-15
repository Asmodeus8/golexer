# GoLexer

A compact lexer/tokenizer written in Go.

It recognizes identifiers, integers, quoted strings and operators while tracking source line and column positions. Illegal/unclosed tokens are surfaced explicitly.

```bash
go run . example.txt
```

The reusable lexer lives in `lexer/lexer.go`; `main.go` provides a file-tokenization CLI. Implemented independently by Adewale Babalola as a compiler-tooling exercise.
