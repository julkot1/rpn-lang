
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

Currently, only the `I64` (64-bit integer) type is implemented. Future plans include additional types.

**To-Do**: Expand type support.

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
| `FUNCTION(vtable_type, vtable, ...)` | Vtable type, vtable variable, n type rules | Specifies the function's vtable for type checking. |
| `TYPE(...)`   | STC types                    | Defines which types are allowed for this function, creating type-checking rules. |

---

### Example: Creating a C Binding with Decorators
To illustrate how to create a binding, let’s walk through an example.

1. **Include the Header**: First, include `decorators.h` in your code.
2. **Use the `STC` Macro**: Apply the `STC` macro to specify relevant decorators.

```c
#include "decorators.h"

// Example function declaration in a .h file
STC(CODE(255), NAME("add"), FUNCTION(STC_bin_function, add_funcs, TYPE(STC_I64_TYPE, STC_I64_TYPE)))
void stc_add(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
```

#### Explanation of the Example
- **`CODE(4)`**: Links this function to the Stack Code operator with token code `255` (function call).
- **`NAME("add")`**: Binds the function to the literal name `"add"` in Stack Code.
- **`FUNCTION(STC_bin_function, add_funcs, TYPE(STC_I64_TYPE, STC_I64_TYPE))`**: Adds function to `add_funcs` vtable
- **`TYPE(STC_I64_TYPE, STC_I64_TYPE)`**: Specifies that the function only accepts `I64` types.

### Notes
- Decorators simplify type checking and operator assignment, making it easier to extend Stack Code with additional functionality.
- The flexibility of the `STC` macro allows for complex rules and seamless C integration.

---







