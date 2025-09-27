# go-jpholiday 🎌🌸

Google Calendar APIを使用して日本の祝日を判定するGoライブラリ。

祝日データはGitHub Actionsによって毎月自動更新されるため、手動での作業なしに常に最新の日本の祝日情報を利用できます

[![Test](https://github.com/haruotsu/go-jpholiday/actions/workflows/test.yml/badge.svg)](https://github.com/haruotsu/go-jpholiday/actions/workflows/test.yml)
[![Update Holidays](https://github.com/haruotsu/go-jpholiday/actions/workflows/update-holidays.yml/badge.svg)](https://github.com/haruotsu/go-jpholiday/actions/workflows/update-holidays.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/haruotsu/go-jpholiday.svg)](https://pkg.go.dev/github.com/haruotsu/go-jpholiday)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/haruotsu/go-jpholiday)](https://goreportcard.com/report/github.com/haruotsu/go-jpholiday)

## インストール

```bash
go get github.com/haruotsu/go-jpholiday
```

## クイックスタート

```go
package main

import (
    "fmt"
    "time"
    "github.com/haruotsu/go-jpholiday/holiday"
)

func main() {
    // 日付が祝日かどうかチェック
    date := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
    if holiday.IsHoliday(date) {
        fmt.Printf("%s は祝日です: %s\n",
            date.Format("2006-01-02"),
            holiday.GetHolidayName(date))
    }
    // 出力: 2024-01-01 は祝日です: 元日

    // 年の全祝日を取得
    holidays := holiday.GetHolidaysInYear(2024)
    fmt.Printf("2024年の祝日数: %d\n", len(holidays))

    // 期間内の祝日を取得（例：ゴールデンウィーク）
    start := time.Date(2024, 4, 29, 0, 0, 0, 0, time.Local)
    end := time.Date(2024, 5, 5, 0, 0, 0, 0, time.Local)
    gwHolidays := holiday.GetHolidaysInRange(start, end)
    for _, h := range gwHolidays {
        fmt.Printf("  %s: %s\n",
            h.Date.Format("2006-01-02"),
            h.Name)
    }
}
```

## データ管理

### 祝日データソース

このライブラリは、公式のGoogle Calendar API（カレンダーID: `ja.japanese.official#holiday@group.v.calendar.google.com`）から日本の国民の祝日データを取得します。このカレンダーはGoogleが管理しており、以下のような日本の公式祝日がすべて含まれています：

- 通常の国民の祝日（元日、成人の日等）
- 特別な祝日（天皇即位の日等）
- 振替休日

### データ保存・キャッシュ
- **データ保持期間**: デフォルトでは、11年分の祝日データを取得・保存（現在年±5年）
- **キャッシュ形式**: O(1)の検索性能のため、日付キー（"YYYY-MM-DD"）でJSON形式

### 更新スケジュール

| 更新タイプ | 頻度 | 方法 | 説明 |
|------------|------|------|------|
| **自動** | 毎月（1日） | GitHub Actions | 祝日データを常に最新状態に保持 |
| **オンデマンド** | いつでも | CLIツール | 必要な時に手動更新 |
| **パッケージリリース** | 新バージョンと共に | 埋め込みデータ | 各リリースで事前にデータを更新 |

### データの鮮度

- 祝日データは通常、日本政府によって1-2年前に公表されます
- 特別な祝日（皇室関連行事等）はより短い期間で発表される場合があります
- 毎月の自動更新により、新しく発表された祝日を迅速に取得します

### 自動更新

祝日データは複数の仕組みによって自動的に更新されます：

1. GitHub Actions（毎月）: 毎月1日の00:00 UTCに実行
2. パッケージ更新: 新しいリリースには最新の祝日データが含まれます
3. ランタイムフォールバック: キャッシュデータが存在しない場合、埋め込みのデフォルトデータを使用

### インストール

```bash
go install github.com/haruotsu/go-jpholiday/cmd/update-holidays@latest
```

### 使用方法

```bash
# Google Calendar API キーを設定
export GOOGLE_API_KEY=your-google-calendar-api-key

# デフォルト範囲（現在年±5年）の祝日を取得
update-holidays

# カスタム年範囲を指定
update-holidays -start-year 2024 -end-year 2029

# ドライラン（更新せずにプレビュー）
update-holidays -dry-run

# デバッグ出力を有効化
update-holidays -debug
```

### CLIオプション

| オプション | 説明 | デフォルト |
|-----------|------|-----------|
| `-start-year` | 取得開始年 | 現在年 - 5 |
| `-end-year` | 取得終了年 | 現在年 + 5 |
| `-cache-file` | キャッシュファイルのパス | `data/holidays.json` |
| `-dry-run` | 更新せずに変更をプレビュー | `false` |
| `-debug` | デバッグ出力を有効化 | `false` |
| `-help, -h` | ヘルプメッセージを表示 | - |
| `-version, -v` | バージョン情報を表示 | - |

### セットアップ

```bash
# リポジトリをクローン
git clone https://github.com/haruotsu/go-jpholiday.git
cd go-jpholiday

# 依存関係をインストール
make setup

# テストを実行
make test

# カバレッジ付きで実行
make test-coverage

# リンターを実行
make lint

# コードをフォーマット
make fmt
```

### 祝日データの更新

#### 手動更新

祝日データを手動で更新するには：

```bash
export GOOGLE_API_KEY=your-api-key
make run
# または直接実行:
go run cmd/update-holidays/main.go
```

#### カスタム年範囲

特定の年の祝日データを取得できます：

```bash
# 5年分のデータを取得（2024-2028）
update-holidays -start-year 2024 -end-year 2028
```

**注意**: Google Calendar APIには濫用防止のためのリクエスト制限があります。5年を超えるデータを一度に取得すると、レート制限にかかる可能性があります。

## コントリビューティング

コントリビューションを歓迎します！詳細は[コントリビューティングガイドライン](CONTRIBUTING.md)をご覧ください。

1. リポジトリをフォーク
2. フィーチャーブランチを作成（`git checkout -b feature/amazing-feature`）
3. 必要に応じてテストを追加・更新
4. 変更をコミット（`git commit -m 'Add amazing feature'`）
5. ブランチにプッシュ（`git push origin feature/amazing-feature`）
6. プルリクエストを開く

詳細は[DEVELOPMENT.md](DEVELOPMENT.md)をご覧ください。

## ライセンス

このプロジェクトはMITライセンスの下で公開されています - 詳細は[LICENSE](LICENSE)ファイルをご覧ください。

## 謝辞

- 祝日データは Google Calendar API（日本の祝日）から取得
- 他言語の類似ライブラリからインスピレーションを得ています

## サポート

- [バグ報告](https://github.com/haruotsu/go-jpholiday/issues)
- [機能リクエスト](https://github.com/haruotsu/go-jpholiday/issues)
- [ドキュメントを読む](https://pkg.go.dev/github.com/haruotsu/go-jpholiday)

## 作者

[@haruotsu](https://github.com/haruotsu)

---

**注意**: このライブラリは祝日データの取得にGoogle Calendar APIに依存しています。商用利用の場合は[Google Calendar API利用規約](https://developers.google.com/terms)とレート制限をご確認ください。
