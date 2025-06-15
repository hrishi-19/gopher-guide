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

---
## 📁 04_more_types

### 1. `pointers.go`  
Demonstrates the use of pointers in Go:  
- `&` operator to get the address of a variable.  
- `*` operator to dereference and manipulate values via pointers.  
- Useful for directly modifying values and optimizing memory usage.

### 2. `struct.go`  
Introduces structs (custom types composed of fields):  
- Defining and instantiating a struct using `type`.  
- Accessing and updating struct fields.  
- Using pointers to structs to modify values without copying.

### 3. `struct_literals.go`  
Shows struct literals and shorthand syntax:  
- Initializing structs with values: full, named, and default.  
- Using `&` with struct literals to get a pointer.  
- Demonstrates value vs pointer behavior during initialization.
---
## 📁 05_arrays

### 1. `arrays.go`  
Introduces arrays in Go:  
- Fixed-size collections of elements.  
- Declaration using `var arr [n]T` and initialization using literals.  
- Demonstrates indexing, comparison, and printing of arrays.  
- Highlights default values and equality checks.

### 2. `slices.go`  
Explores Go slices (dynamic views over arrays):  
- Slicing an existing array with `[low:high]` syntax.  
- Slice literals, length and capacity.  
- Modifying a slice affects the underlying array.  
- Re-slicing operations like trimming and extending.

### 3. `slice_append.go`  
Focuses on slice growth and `append`:  
- Appending to nil and non-nil slices.  
- Demonstrates automatic capacity expansion.  
- Multiple element append in a single call.  
- Uses a helper `printSlice` function to show slice state.

---
## 📁 06_maps

### 1. `maps.go`  
Introduces maps in Go:  
- Maps are key-value pairs, like dictionaries in other languages.  
- Declared using `make(map[KeyType]ValueType)`.  
- Example shows storing `Vertex` structs with string keys.  
- Accessing map elements using keys.

### 2. `map_operation.go`  
Demonstrates common map operations:  
- Inserting and updating values.  
- Deleting keys with `delete(map, key)`.  
- Checking for key existence using the `value, ok := map[key]` idiom.  
- Shows the default zero value if a key is not present.
