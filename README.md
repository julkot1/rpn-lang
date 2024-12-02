
# Reverse Polish Notation Language (Stack Code)
Stack Code is a playful and experimental language built with Go and LLVM (version 14). While it’s still a work in progress, the aim is to explore stack-oriented programming in a dynamic, C-compatible environment. Note: This language is not yet fully developed and serves more as a fun exploration!

## Language Features

-  **Dynamically Typed**: Types are checked and managed at runtime.

-  **Stack-Oriented**: Operates primarily through a stack-based paradigm. [Learn more about stack-oriented programming](https://en.wikipedia.org/wiki/Stack-oriented_programming).

-  **glibc Based**: Built on top of the GNU C Library.

-  **Easy C Binding**: Simplifies integration with C functions.

---

## Running Stack Code

Below is a Markdown description for the two `Config` structs you provided, designed for documentation purposes:

---

## Configuration Structs

### Config Type


The `Config` struct defines the configuration options for your application. These options can be specified in a TOML configuration file. Each field is described below:

| Field Name      | TOML Key         | Type    | Description                                                                 |
|------------------|------------------|---------|-----------------------------------------------------------------------------|
| `RootPath`      | `root_path`      | string  | The root directory path for the project.                                   |
| `Destination`   | `destination`    | string  | The output directory where generated files will be stored.                 |
| `CompileLibs`   | `compile_libs`   | bool    | A flag indicating whether to compile libraries.                            |
| `ClangPath`     | `clang_path`     | string  | The file path to the Clang compiler executable.                            |
| `LlcPath`       | `llc_path`       | string  | The file path to the LLVM `llc` tool for compiling LLVM intermediate code. |
| `LLDisPath`     | `llc_dis_path`   | string  | The file path to the LLVM `lldis` tool for disassembling LLVM bytecode.    |
| `LinkerPath`    | `linker_path`    | string  | The file path to the linker executable.                                    |


---

### Commands

#### `stc` (Root Command)
The root command serves as the entry point for the `stc` CLI tool. It includes the following subcommands:

##### 1. `compile`
Compiles given program files into an executable or LLVM IR, based on the provided flags.

**Usage:**
```sh
stc compile [files...] [flags]
```

**Arguments:**
- `files...`: A list of source files to compile.

**Flags:**
| Flag                | Shorthand | Default | Description                                                |
|---------------------|-----------|---------|------------------------------------------------------------|
| `--output`          | `-o`      | `out`   | Specifies the name of the output file.                     |
| `--opt`             | `-O`      | `0`     | Specifies the optimization level (0–3).                    |
| `--emit-llvm`       | `-l`      | `false` | Generates output in LLVM Intermediate Representation (IR). |

**Behavior:**
- Validates the optimization level.
- Compiles each file specified using the `Compile` function, configured via a TOML file (`config.debug.toml`).

##### 2. `libs`
Compiles the default libraries for the `stc` tool.

**Usage:**
```sh
stc libs
```

**Behavior:**
- Invokes the `CompileStcLibs` function to build standard libraries, using the configuration from `config.debug.toml`.

---

### Example Usage

1. Compile source files with default settings:
   ```sh
   stc compile file1.rpn file2.rpn
   ```

2. Specify an output file and optimization level:
   ```sh
   stc compile file1.rpn -o myprogram -O2
   ```

3. Generate LLVM IR:
   ```sh
   stc compile file1.rpn -l
   ```

4. Compile default libraries:
   ```sh
   stc libs
   ```

---

### Notes
- The `config.CreateTOMLConfig` function is used to load configurations from a TOML file.
- Ensure the optimization level is between 0 and 3 to avoid errors.
- This CLI package requires the `cobra` library for command management.

---

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
| `Type` | `long` |

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

### Typeof Operator

In Stack Code, the `typeof` operator is used to push the type of the item currently on the stack onto the stack. This operator can be used to determine the type of a value at runtime.

#### Syntax:
```stack
value typeof
```

- `value`: The item whose type you want to inspect.
- The result is pushed onto the stack as a string representing the type of the item.

#### Example:
```stack
42 typeof
```

**Before**:
```
[42]
```

**After**:
```
["I64"type]
```

In this example, the `typeof` operator inspects the type of `42`, which is an `I64` integer, and pushes the string `"I64"` onto the stack.

#### Handling Types in Stack Code

- The `typeof` operator outputs a string type (e.g., `"I64"`, `"F64"`, `"I8"`) that can be used for conditional checks or debugging.
- This string type can be used in conjunction with other operators for dynamic type handling in Stack Code programs.

---

### Print Operator

The `print` operator outputs the current item on the stack as a string representation, converting it based on its type. It will display the value in a human-readable format.

#### Syntax:
```stack
value print
```

- `value`: The item you want to print to the console.

#### Example:
```stack
42 print
```

**Before**:
```
[42]
```

**After**:
```
[]
```

**Output**:
```
42
```

#### Behavior:
- The `print` operator converts the value on top of the stack into a readable string and outputs it to the console.
- For numeric types like `I64` or `F64`, it will print the number.
- For types like `I8` (which is a character), it will print the ASCII representation of the character.

#### Example with `I8` Type:
```stack
'I' print
```

**Before**:
```
['I']
```

**After**:
```
[]
```

**Output**:
```
I
```

In this case, the `I8` value `'I'` is printed as the character `I`.

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







