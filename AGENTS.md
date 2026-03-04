## Cursor Cloud specific instructions

This is a CodeCrafters "Build Your Own Shell" challenge implemented in Go 1.25. It is a single standalone CLI binary with no external dependencies, no databases, and no Docker.

### Build & Run

- Build: `go build -o /tmp/codecrafters-build-shell-go app/*.go`
- Run: `./your_program.sh` (compiles and runs in one step)
- The shell reads from stdin and writes to stdout; pipe input for non-interactive testing: `echo "hello" | ./your_program.sh`

### Lint

- `go vet ./...` — the only lint check available (no third-party linters configured)

### Testing

- No local test suite exists. Testing is done via the CodeCrafters platform (`codecrafters submit`), which requires an account and network access.
- For local validation, pipe commands into the shell binary and verify output.
