# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## プロジェクト概要

**Konst** は JSON 定義から Go と TypeScript のコードを自動生成するツールです。定数、列挙型、オブジェクト（構造体/インターフェース）の定義を JSON で記述し、対応するコードを生成します。

## 主要なコマンド

### ビルドとテスト

```bash
# プロジェクトのビルド
go build .

# 実行（単一ファイル処理）
go run . -i ./example/konst.json -o gen/konst.ts -f
go run . -i ./example/konst.json -o gen/konst.go -f

# Task タスクランナーを使用
task gen:ts    # TypeScript 生成
task gen:go    # Go コード生成  
task gen:all   # 両方生成

# ディレクトリ処理
task gen:dir:ts   # ディレクトリから TypeScript 生成
task gen:dir:go   # ディレクトリから Go 生成
task gen:dir:all  # ディレクトリから両方生成

# テスト実行
go test ./...
```

### コマンドラインオプション

```bash
-i string    # 入力JSONファイルまたはディレクトリ
-o string    # 出力ファイル名またはディレクトリ（必須）
-t string    # カスタムテンプレートディレクトリ
-f           # 既存ファイルを強制上書き
-m string    # 出力モード（go, ts）- ディレクトリ処理時
-indent int  # インデント数（デフォルト2）
-v, -version # バージョン表示
```

## アーキテクチャ

### ディレクトリ構成

- `main.go` - エントリポイント、コマンドライン引数処理
- `internal/process/` - ファイル・ディレクトリ処理ロジック
  - `file.go` - 単一JSONファイル処理
  - `directory.go` - ディレクトリ再帰処理
- `internal/template/` - テンプレート処理とコード生成
- `internal/types/` - 型定義とスキーマ定義
- `internal/utils/` - ユーティリティ関数
- `example/` - サンプル定義ファイル

### 処理フロー

1. **入力判定**: 単一ファイルかディレクトリかを判定 (`main.go:36-49`)
2. **出力形式決定**: 拡張子またはモードフラグから Go/TypeScript を判定 (`main.go:52-62`)
3. **処理実行**: 
   - 単一ファイル: `process.ProcessFile()`
   - ディレクトリ: `process.ProcessDirectory()`
4. **コード生成**: テンプレートエンジンによる出力生成

### JSON スキーマ形式

```json
{
  "version": "1.0",
  "goPackage": "package_name",
  "definitions": {
    "ConstName": {
      "type": "int|int64|string|bool|date|float|uint...",
      "value": "actual_value",
      "mode": "number|string|bigint"  // TypeScript出力制御
    }
  }
}
```

### 型マッピング

- `int` → Go: `int`, TS: `number`
- `int64` → Go: `int64`, TS: `bigint` (mode: "number" で `number`)
- `string` → Go: `string`, TS: `string`
- `bool` → Go: `bool`, TS: `boolean`
- `date` → Go: `time.Time`, TS: `Date` (mode で変更可能)
- 配列型もサポート: `int[]`, `string[]` など

## カスタムテンプレート

- テンプレートファイル: `go.tmpl`, `ts.tmpl`
- 場所: `-t` オプション、`KONST_TEMPLATES` 環境変数、または実行ファイル同階層の `templates/`
- Go テンプレート記法を使用