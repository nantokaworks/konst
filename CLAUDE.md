# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 前提

- 日本語でチャットすること。
- 適宜README.mdの更新を行う。
- ユーザーの選択を仰ぐ場合は通知で知らせる。

## プロジェクト概要

**Konst** は JSON 定義から Go と TypeScript のコードを自動生成するツールです。定数、列挙型、オブジェクト（構造体/インターフェース）の定義を JSON で記述し、対応するコードを生成します。

## 主要なコマンド

### ビルドとテスト

```bash
# プロジェクトのビルド
go build .

# システムインストール
go install .

# 基本実行（ディレクトリベース）
konst -i ./example/enum.json -o gen/output -m ts -f
konst -i ./example/multi -o gen/output -m go -f

# Task タスクランナーを使用
task gen:ts       # TypeScript 生成
task gen:go       # Go コード生成  
task gen:all      # 両方生成
task gen:dir:ts   # ディレクトリから TypeScript 生成
task gen:dir:go   # ディレクトリから Go 生成
task gen:dir:all  # ディレクトリから両方生成

# テスト実行
go test ./...
```

### コマンドラインオプション

```bash
-i string      # 入力JSONファイルまたはディレクトリ
-o string      # 出力ディレクトリ（必須、バリデーション時除く）
-m string      # 出力モード（go, ts）必須
-f             # 既存ファイルを強制上書き
-t string      # カスタムテンプレートディレクトリ
--indent int   # インデント数（デフォルト2）
--naming string # ファイル命名規則（kebab, camel, snake）
               # TypeScriptはデフォルトでkebab、Goはデフォルトでsnake
--validate     # JSON定義の検証のみ（コード生成なし）
--dry-run      # 生成予定ファイル一覧表示
--watch        # ファイル変更監視（実験的）
-v, --version  # バージョン表示
```

### ファイル命名規則

`--naming` オプションで出力ファイルとディレクトリの命名規則を指定できます：

- `kebab`: kebab-case（例: `test-file.ts`, `sub-directory/`）
- `camel`: camelCase（例: `testFile.ts`, `subDirectory/`）  
- `snake`: snake_case（例: `test_file.go`, `sub_directory/`）

デフォルトでは、TypeScript は `kebab-case`、Go は `snake_case` を使用します。

例：
```bash
# TypeScript でデフォルト（kebab-case）
konst -i ./example/test_snake_case.json -o gen/output -m ts

# Go で camelCase を指定
konst -i ./example/test-kebab-case.json -o gen/output -m go --naming camel
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

1. **バリデーション/ドライラン**: 特殊モードの場合は早期リターン
2. **入力判定**: 単一ファイルかディレクトリかを判定  
3. **出力形式決定**: `-m` フラグから Go/TypeScript を判定
4. **処理実行**: 常に `process.ProcessDirectory()` を使用（統一化）
5. **コード生成**: テンプレートエンジンによる出力生成

### JSON スキーマ形式

```json
{
  "version": "1.0",
  "goPackage": "package_name", 
  "definitions": {
    "ConstName": {
      "type": "int|string|bool|date|enum...",
      "value": "actual_value",
      "tsMode": "number|string|bigint"
    },
    "EnumName": {
      "type": "enum",
      "values": ["option1", "option2"],
      "default": "option1"
    }
  }
}
```

### 型マッピング

#### 基本型
- `int` → Go: `int`, TS: `number`
- `int64` → Go: `int64`, TS: `bigint` (tsMode: "number" で `number`)
- `string` → Go: `string`, TS: `string`
- `bool` → Go: `bool`, TS: `boolean`
- `date` → Go: `time.Time`, TS: `Date` (各種モード指定可能)

#### enum型（v0.3.0新機能）
- `enum` → Go: カスタム型 + バリデーション関数群
- `enum` → TS: const object + 型 + バリデーション関数群

#### 配列型
- 各型に `[]` 付加: `int[]`, `string[]`, `enum[]` など

## カスタムテンプレート

- テンプレートファイル: `go.tmpl`, `ts.tmpl`
- 場所: `-t` オプション、`KONST_TEMPLATES` 環境変数、または実行ファイル同階層の `templates/`
- Go テンプレート記法を使用

## メモ

- バージョンを上げる際はHISTORY.mdの更新も忘れずに