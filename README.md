
# Reverse Polish Notation Language (Stack Code)
Stack Code is a playful and experimental language built with Go and LLVM (version 14). While it’s still a work in progress, the aim is to explore stack-oriented programming in a dynamic, C-compatible environment. Note: This language is not yet fully developed and serves more as a fun exploration!

## Language Features

-  **Dynamically Typed**: Types are checked and managed at runtime.

-  **Stack-Oriented**: Operates primarily through a stack-based paradigm. [Learn more about stack-oriented programming](https://en.wikipedia.org/wiki/Stack-oriented_programming).

-  **glibc Based**: Built on top of the GNU C Library.

-  **Easy C Binding**: Simplifies integration with C functions.

---

## Running Stack Code

**To-Do**: Instructions on setting up and running the interpreter will be available soon.  

---


## Language Fundamentals

  

### Stack Operations

The stack is the primary data space for operations. Each item on the stack is 64 bits wide and includes 16 bits of type information.

**Convention**: `[before operation] -> [after operation]`

#### Basic Operations

| **Name**  |  **Syntax** | **Operation** |
| :---         |     :---:      |         :---: |
| `push` | `a` | `[] -> [a]` |
| `dup` | `\|>` | `[a] -> [a a]` |
| `pop` | `pop` | `[a b] -> [b]` |
| `swap` | `\|s` | `[a b] -> [b a]` |
| `rotate` | `\|r` | `[a b c] -> [c b a]` |

---



### Type System

| **Type**  | **Equivalent C Type** |
| :--- | :---: |
| `I64` | `long` |
| `F64` | `double` |
| `I8`  | `char` |

---

### Number Types

#### Integer Types

There are two integer types in this system: `I8` and `I64`.

1. **`I8`**: Represents a single character.
   - Format: `'x'`, where `x` is a single ASCII character.

2. **`I64`**: Represents a 64-bit integer, which can be defined in multiple numeric formats:
   - **Decimal**: A standard integer, which may optionally include a `+` or `-` sign (e.g., `123`, `-456`).
   - **Hexadecimal**: Prefixed with `0x` or `0X` (e.g., `0x1A`, `0X2B`).
   - **Octal**: Prefixed with `0o` or `0O` (e.g., `0o17`, `0O24`).
   - **Binary**: Prefixed with `0b` or `0B` (e.g., `0b1010`, `0B1101`).

| **Format**      | **Example**  | **Description**                     |
| :-------------- | :----------- | :----------------------------------- |
| Decimal         | `123`        | Positive or negative decimal number  |
| Hexadecimal     | `0x1A`       | Prefixed with `0x` or `0X`           |
| Octal           | `0o17`       | Prefixed with `0o` or `0O`           |
| Binary          | `0b1010`     | Prefixed with `0b` or `0B`           |

#### Floating-Point Types

The `F64` type represents a 64-bit floating-point number equivalent to `double` in C. Floating-point literals typically include a decimal point or an exponent (e.g., `1.23`, `4.56e-2`).




---

### Stack Persistence Modifier

The `!` suffix ensures that values used in operations remain on the stack after execution.
#### Example
```stack
5 4 !+
```
**Before**  `!+` operation:
```
[4 5]
```

**After**  `!+` operation:

```
[9 4 5]
```
---

  

### Control Flow
#### Conditional Execution (`if`)

The `if` statement pops a value from the stack and executes the block if the value is non-zero (true).
```stack
3 3 == if {
	5 print
}
```

**Output**:

```

5

```

---

### Loops

#### Repetition (`repeat`)

The `repeat` statement pops a value `n` from the stack and performs the enclosed block `n` times.

```stack
0
5 repeat {
	1 + !print
}
```
**Output**:

```
1
2
3
4
5
```
---
## C Library Integration

### STCClib (Stack Code C Library)
STCClib, located in the `lib/` directory, provides essential operations and bindings that expand Stack Code’s capabilities. This library includes built-in functions for arithmetic, I/O, logical operations, and more. It also defines crucial bindings to handle stack operations seamlessly in C.

---

### Using Decorators for C Binding
Decorators in STCClib simplify the process of linking C functions to Stack Code. These are macros that define key metadata and rules about how functions should interact with the Stack Code environment.

To use decorators, you need to import `decorators.h`. Then, when defining a function in your `.h` file, you can apply the `STC` macro to specify various attributes.

---

### Decorator Macros
Here's a breakdown of the available macros and how to use them:

| **Macro**     | **Arguments**                | **Description**                                             |
|---------------|-------------------------------|-------------------------------------------------------------|
| `STC(...)`    | Any sequence of decorators   | The main macro used to apply a set of decorators to a function. |
| `CODE(x)`     | STC token code (`int32`)     | Associates the function with a built-in operator in Stack Code. |
| `NAME(x)`     | Function name (`string`)     | Binds the function to a specific literal name in Stack Code. |


---

### Example: Creating a C Binding with Decorators
To illustrate how to create a binding, let’s walk through an example.

1. **Include the Header**: First, include `decorators.h` in your code.
2. **Use the `STC` Macro**: Apply the `STC` macro to specify relevant decorators.

```c
#include "decorators.h"

// Example function declaration in a .h file
STC(CODE(255), NAME("add"))
void stc_add(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
```

#### Explanation of the Example
- **`CODE(255)`**: Links this function to the Stack Code operator with token code `255` (function call).
- **`NAME("add")`**: Binds the function to the literal name `"add"` in Stack Code.


### Notes
- Decorators simplify type checking and operator assignment, making it easier to extend Stack Code with additional functionality.
- The flexibility of the `STC` macro allows for complex rules and seamless C integration.

---







