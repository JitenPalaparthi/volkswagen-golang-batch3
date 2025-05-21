## 🔍 Segment Breakdown

| Segment | Description |
|---------|-------------|
| **Code/Text** | Stores compiled machine code (functions, methods). Read-only. |
| **RO Data** | Stores read-only constants, literals, and strings. |
| **Initialized Data** | Stores global/static variables with explicit initialization. |
| **BSS** | Stores uninitialized global/static variables; zero-filled at runtime. |
| **Heap** | Stores dynamically allocated data (via `new`, `make`, slices, maps, etc. not always based on escape analysis). Managed by the garbage collector. |
| **Stack** | Stores function call frames: local vars, return addresses, params. Automatically managed by Go. |
