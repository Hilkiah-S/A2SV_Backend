## Run
1. Open a terminal in `task_3/library_management`.
2. Run:
```bash
go run .
```

## Structure
- `main.go`: starts the app.
- `controllers/library_controller.go`: console input/output.
- `models/book.go`: `Book` struct.
- `models/member.go`: `Member` struct.
- `services/library_service.go`: `LibraryManager` interface + implementation.
- `docs/documentation.md`: this file.
- `go.mod`: module definition.

## Usage
- Members preloaded: `1: Alice`, `2: Bob`.
- Menu options let you:
  - Add / Remove books
  - Borrow / Return books
  - List available books
  - List books borrowed by a member

## Notes
- `Status` is "Available" or "Borrowed".
- Errors shown for missing items or invalid actions.

