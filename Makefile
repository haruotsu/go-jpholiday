.PHONY: help
help: ## ヘルプを表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: setup
setup: ## プロジェクトのセットアップ
	@echo "Setting up project structure..."
	@mkdir -p holiday fetcher data cmd/update-holidays .github/workflows
	@go mod init github.com/haruotsu/go-jpholiday 2>/dev/null || true
	@go mod tidy
	@echo "Setup complete!"

.PHONY: test
test: ## テストを実行
	go test -v ./...

.PHONY: test-coverage
test-coverage: ## カバレッジ付きでテストを実行
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

.PHONY: test-race
test-race: ## 競合状態の検出付きでテストを実行
	go test -race ./...

.PHONY: bench
bench: ## ベンチマークテストを実行
	go test -bench=. -benchmem ./...

.PHONY: lint
lint: ## 静的解析を実行
	@which golangci-lint > /dev/null 2>&1 || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	golangci-lint run ./...

.PHONY: fmt
fmt: ## コードフォーマット
	go fmt ./...
	gofmt -s -w .

.PHONY: vet
vet: ## go vetを実行
	go vet ./...

.PHONY: security
security: ## セキュリティチェック
	@which gosec > /dev/null 2>&1 || (echo "Installing gosec..." && go install github.com/securego/gosec/v2/cmd/gosec@latest)
	gosec ./...

.PHONY: check
check: fmt vet lint test ## 全ての品質チェックを実行

.PHONY: build
build: ## ビルド
	go build -o bin/update-holidays cmd/update-holidays/main.go

.PHONY: install
install: ## ローカルにインストール
	go install ./cmd/update-holidays

.PHONY: run
run: ## 祝日データ更新コマンドを実行
	@if [ -z "$(GOOGLE_API_KEY)" ]; then \
		echo "Error: GOOGLE_API_KEY is not set"; \
		exit 1; \
	fi
	go run cmd/update-holidays/main.go

.PHONY: run-dry
run-dry: ## ドライラン（実際に更新しない）
	@if [ -z "$(GOOGLE_API_KEY)" ]; then \
		echo "Error: GOOGLE_API_KEY is not set"; \
		exit 1; \
	fi
	go run cmd/update-holidays/main.go -dry-run

.PHONY: clean
clean: ## ビルド成果物とテストキャッシュをクリーンアップ
	rm -rf bin/ coverage.* *.test
	go clean -testcache

.PHONY: deps
deps: ## 依存関係を更新
	go mod download
	go mod tidy

.PHONY: update-deps
update-deps: ## 依存関係を最新バージョンに更新
	go get -u ./...
	go mod tidy

.PHONY: docs
docs: ## ドキュメントをブラウザで開く
	@which godoc > /dev/null 2>&1 || (echo "Installing godoc..." && go install golang.org/x/tools/cmd/godoc@latest)
	@echo "Opening documentation at http://localhost:6060/pkg/github.com/haruotsu/go-jpholiday/"
	@godoc -http=:6060

.PHONY: profile
profile: ## CPUプロファイリング付きでテストを実行
	go test -cpuprofile=cpu.prof -memprofile=mem.prof -bench=.
	go tool pprof cpu.prof

.PHONY: ci
ci: deps check test-coverage ## CI環境での実行

.PHONY: tag
tag: ## tagprでリリース
	@which tagpr > /dev/null 2>&1 || (echo "Installing tagpr..." && go install github.com/Songmu/tagpr@latest)
	tagpr

.DEFAULT_GOAL := help
