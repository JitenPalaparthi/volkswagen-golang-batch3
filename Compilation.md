## How compiler compiles


Java Source      --> javac -->  Byte Code --> JVM(JIT, Runtime) --> Runs application on a target machine

Closure Source   --> closurec -->  Byte Code --> JVM(JIT, Runtime) --> Runs application on a target machine

Scale Source     --> scalac -->  Byte Code --> JVM(JIT, Runtime) --> Runs application on a target machine

Kotlin Source    --> kotlinc -->  Byte Code --> JVM(JIT, Runtime) --> Runs application on a target machine

CShape Source    --> csc -->  IL --> .NetRuntime(JIT, Runtime) --> Runs application on a target machine

C++ Source --> C++ Compiler --> Assembler --> Linker --> (Binary/Exe)
                            --------------------------
                                    LLVM(Clang)/GCC

Rust Source --> RustC --> LLVM FE(IR) --> Based on BE --> Assembler(Machine Code .o) --> Linker --> Binary/Exe
                          ---------------------------------------------------------------------
                                            LLVM Tool Chain

Go Source --> GoC --> SSA (Single Static Assignment, In Mem IR) --> Assembler(Plan 9 Assember) -> Linker --> Binary/Exe

- No monomorphization in Go/ Rather there is type eraser
- No Macros, no preprocessor
- Excellent Code optimization
- Constant Evaluation and Dead Code elimination
- Most of the Std is pre compiled (as User do not change std)
- Compiler does the escape analysis

                
What Compiler Does, generally

- There is a lexical analysis
- Tokenization
- AST
- Optimized code
- Dead code elimination
- Constant Evaluation
- Monomorphization (All generics are converted to concreate code)
- Preprocessing (Macros Code generation etc..)
- It checks and tell your code is legal or not (Validate the code)