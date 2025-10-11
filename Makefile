.PHONY: lint fmt test test-unit

lint:
	@echo "🚀 Running linter..."
	@golangci-lint run
	@echo "✅ Linting complete"

fmt:
	@echo "🔧 Formatting code..."
	@golangci-lint fmt
	@echo "✅ Code formatted"

test:
	@echo "👾 Running all tests..."
	@make test-unit

test-unit:
	@echo "👾 Running unit tests..."
	@go test -v ./test/unit/...
