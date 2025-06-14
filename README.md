# Go Concepts Learning Repository

This repository is a structured guide for learning core Go (Golang) concepts, organized by topic. Each folder contains example `.go` files that demonstrate key features of the language.

---

## 📁 01_basics

### 1. `variables.go`
Demonstrates variable declaration:
- Using `var` keyword with explicit type
- Using short declaration syntax (`:=`)

### 2. `data_types.go`
Covers basic Go data types:
- `int`, `float64`, `bool`, `string`

### 3. `constants.go`
Shows how to declare and use constants:
- Immutable values declared with `const`

### 4. `comments.go`
Illustrates Go's commenting styles:
- Single-line comments (`//`)
- Multi-line comments (`/* ... */`)

---

## 📁 02_control_flow

### 1. `if_else.go`
Covers conditional statements:
- `if`, `else if`, and `else` for control branching based on conditions.

### 2. `switch.go`
Demonstrates the use of `switch` statements:
- Cleaner alternative to multiple `if-else` checks.
- Supports multiple values in a single `case`.

### 3. `loops.go`
Shows different looping patterns in Go:
- Traditional `for` loop (like `for i := 0; i < n; i++`)
- While-like `for` loop
- Infinite `for` loop with `break` control

---
## 📁 03_functions

### 1. `basic_function.go`  
Demonstrates how to define and call a basic function:  
- Function with a parameter.  
- No return value.

### 2. `function_with_return.go`  
Shows how to define a function that returns  value:  
- Return type specified after parameters.  
- Value returned using `return` keyword.

### 3. `variadic_function.go`  
Covers functions that accept a variable number of arguments:  
- Uses `...` syntax for variadic parameters.  
- Aggregates input into a slice.

### 4. `named_return_values.go`  
Demonstrates named return values:  
- Return variables are declared in the function signature.  

### 5. `defer_panic_recover.go`  
Covers Go’s error handling via:  
- `defer`: delays execution until the surrounding function returns.  
- `panic`: aborts the current control flow.  
- `recover`: captures a panic to prevent crash.
