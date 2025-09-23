# go-jpholiday

Google Calendar APIを使用して日本の祝日を判定するGoライブラリです。

[![Test](https://github.com/haruotsu/go-jpholiday/actions/workflows/test.yml/badge.svg)](https://github.com/haruotsu/go-jpholiday/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/haruotsu/go-jpholiday.svg)](https://pkg.go.dev/github.com/haruotsu/go-jpholiday)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## インストール

```bash
go get github.com/haruotsu/go-jpholiday
```

## 使い方

```go
package main

import (
    "fmt"
    "time"
    "github.com/haruotsu/go-jpholiday/holiday"
)

func main() {

    // 特定の日付が祝日かどうか判定
    date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
    if holiday.IsHoliday(date) {
        fmt.Printf("%s は祝日です: %s\n", date.Format("2006-01-02"), holiday.GetHolidayName(date))
    }
    // 出力: 2024-01-01 は祝日です: 元日

    // 年の全祝日を取得
    holidays := holiday.GetHolidaysInYear(2024)
    fmt.Printf("2024年の祝日数: %d\n", len(holidays))
    for _, h := range holidays[:3] { // 最初の3つだけ表示
        fmt.Printf("  %s: %s\n", h.Date.Format("2006-01-02"), h.Name)
    }

    // 期間内の祝日を取得
    start := time.Date(2024, 4, 29, 0, 0, 0, 0, time.Local)
    end := time.Date(2024, 5, 5, 0, 0, 0, 0, time.Local)
    gwHolidays := holiday.GetHolidaysInRange(start, end)
    fmt.Printf("\nゴールデンウィーク期間の祝日:\n")
    for _, h := range gwHolidays {
        fmt.Printf("  %s: %s\n", h.Date.Format("2006-01-02"), h.Name)
    }
}
```

## API

### 主要な関数

#### 祝日判定
- `IsHoliday(date time.Time) bool` - 指定した日付が祝日かどうかを判定
- `GetHolidayName(date time.Time) string` - 指定した日付の祝日名を取得（祝日でない場合は空文字列）

#### 祝日リストの取得
- `GetHolidaysInYear(year int) []Holiday` - 指定した年の全祝日を取得
- `GetHolidaysInRange(start, end time.Time) []Holiday` - 指定した期間の祝日を取得

#### キャッシュ管理
- `LoadCache(filePath string) (*HolidayCache, error)` - キャッシュファイルから祝日データを読み込み
- `SaveCache(filePath string, cache *HolidayCache) error` - 祝日データをキャッシュファイルに保存
- `SetCache(cache *HolidayCache)` - 使用するキャッシュを設定
- `IsStale(cache *HolidayCache, maxAge time.Duration) bool` - キャッシュが古いかどうかを判定

### データ構造

```go
// Holiday represents a Japanese holiday
type Holiday struct {
    Date        time.Time `json:"date"`
    Name        string    `json:"name"`
    Description string    `json:"description,omitempty"`
}

// HolidayCache represents cached holiday data
type HolidayCache struct {
    LastUpdated time.Time          `json:"last_updated"`
    Holidays    map[string]Holiday `json:"holidays"` // Key: "YYYY-MM-DD"
}
```

## CLIツール

祝日データを更新するためのCLIツールが含まれています。

### インストール

```bash
go install github.com/haruotsu/go-jpholiday/cmd/update-holidays@latest
```

### 使い方

```bash
# 環境変数でAPIキーを設定
export GOOGLE_API_KEY=your-google-calendar-api-key

# 現在年と翌年の祝日データを取得
update-holidays

# 特定の年範囲を指定
update-holidays -start-year 2024 -end-year 2025

# ドライラン（実際には更新しない）
update-holidays -dry-run

# デバッグモード
update-holidays -debug

# ヘルプを表示
update-holidays -help
```

### オプション

- `-start-year`: 取得開始年（デフォルト: 現在年）
- `-end-year`: 取得終了年（デフォルト: 現在年+1）
- `-cache-file`: キャッシュファイルのパス（デフォルト: data/holidays.json）
- `-dry-run`: 実際には更新せず、取得する祝日を表示
- `-debug`: デバッグ情報を表示
- `-help, -h`: ヘルプを表示
- `-version, -v`: バージョン情報を表示

## 開発

### セットアップ

```bash
# プロジェクトのセットアップ
make setup

# テスト実行
make test

# 静的解析とフォーマット
make lint
make fmt

# 全チェック実行
make check

# カバレッジ付きテスト
make test-coverage
```

### 祝日データの更新

Google Calendar APIキーを設定後、以下のコマンドで祝日データを更新できます：

```bash
export GOOGLE_API_KEY=your-api-key
go run cmd/update-holidays/main.go
# または
make run
```

GitHub Actionsで毎月1日に自動更新されます。

## コントリビューション

プルリクエストを歓迎します！以下の手順で貢献してください：

1. このリポジトリをフォーク
2. フィーチャーブランチを作成 (`git checkout -b feature/amazing-feature`)
3. テストを追加・更新
4. 変更をコミット (`git commit -m 'Add amazing feature'`)
5. ブランチにプッシュ (`git push origin feature/amazing-feature`)
6. プルリクエストを作成

詳細は[DEVELOPMENT.md](DEVELOPMENT.md)をご覧ください。

## ライセンス

MIT License - 詳細は[LICENSE](LICENSE)をご覧ください。

## 作者

[@haruotsu](https://github.com/haruotsu)

---

**注意**: このライブラリはGoogle Calendar APIに依存しています。商用利用の場合は、Google Calendar APIの利用規約とレート制限をご確認ください。
