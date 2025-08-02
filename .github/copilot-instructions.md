# Copilot Instructions

## General Best Practices

- Be descriptive: Use clear and descriptive names for variables, functions, and other entities.
- Keep it simple: Aim for simplicity in design and implementation. Avoid unnecessary complexity.
- DRY: Don't Repeat Yourself. Reuse code where possible instead of duplicating it.
- KISS: Keep It Simple, Stupid. Don't make things more complicated than they need to be.
- YAGNI: You Aren't Gonna Need It. Don't add functionality until it's necessary.

## Golang Best Practices

- Use idiomatic Go: prefer clear, concise, and readable code.
- Follow Go naming conventions (e.g., `CamelCase` for exported names, `camelCase` for local variables).
- Use `gofmt` or `goimports` for formatting.
- Handle errors explicitly; avoid ignoring returned errors.
- Prefer short variable declarations (`:=`) where appropriate.
- Keep functions small and focused.
- Use slices and maps idiomatically.
- Avoid global variables unless necessary.
- Use context for cancellation and timeouts in concurrent code.
- Write table-driven tests for unit testing.

