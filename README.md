# Monkey Programming Language Compiler

This project implements a compiler for the Monkey programming language, written in Go. The Monkey language is a simple, interpreted programming language designed for educational purposes.

## Components

### 1. Main (main.go)

The entry point of the application. It sets up a REPL (Read-Eval-Print Loop) for interactive use of the Monkey language.

### 2. Token (token/token.go)

Defines the structure of tokens and their types. Tokens are the smallest units of meaning in the language, such as keywords, identifiers, and operators.

### 3. Lexer (lexer/lexer.go)

The lexical analyzer that breaks down the input source code into a series of tokens. It implements the following key functions:

- `New(input string)`: Creates a new Lexer instance.
- `NextToken()`: Reads the next token from the input.
- `readChar()`: Advances the lexer to the next character.
- `readIdentifier()`: Reads an identifier from the input.
- `readNumber()`: Reads a number from the input.

### 4. REPL (repl/repl.go)

Implements a Read-Eval-Print Loop for interactive use of the Monkey language. It repeatedly:

1. Prompts the user for input.
2. Lexically analyzes the input using the Lexer.
3. Prints out the resulting tokens.

### 5. Lexer Tests (lexer/lexer_test.go)

Contains unit tests for the Lexer to ensure it correctly tokenizes various inputs.

## Notes (notes.txt)

Contains detailed explanations of the compiler's phases, including:

1. Lexical Analysis
2. Syntax Analysis
3. Semantic Analysis
4. Intermediate Representation Generation
5. Optimization
6. Code Generation
7. Code Linking and Assembly

## Current State

As of now, the project implements the lexical analysis.
