# HISTORY

Konstの開発履歴とバージョン別変更点を記録しています。

## v0.3.6 (2025-07-05)

### 🔧 修正
- **Go出力の改善**: 同一ディレクトリ内で異なる`goPackage`を持つJSONファイルの処理を改善
  - 各`goPackage`ごとに自動的にサブディレクトリを作成
  - パッケージ名の競合によるビルドエラーを解決
  - 例: `config.json` (goPackage: "config") → `output/config/config.go`
  - 例: `database.json` (goPackage: "database") → `output/database/database.go`

### 🔄 影響範囲
- `internal/process/file.go`: Go出力時のパス構築ロジック改善
- `internal/process/file_resolver.go`: 依存関係解決時のパス構築ロジック改善

### 📝 使用方法
```bash
# 異なるgoPackageを持つファイルがある場合
konst -i ./definitions -o ./output -m go -f
# → output/config/config.go (package config)
# → output/database/database.go (package database)
```

## v0.3.5 (2025-06-29)

### 🌐 主要機能追加
- **完全i18n対応**: ヘルプメッセージの多言語サポート
  - 日本語・英語のヘルプ表示をサポート
  - `--locale=ja` フラグで日本語ヘルプを表示
  - `--locale=en` フラグで英語ヘルプを表示
  - 環境変数 `KONST_LOCALE` による言語設定
  - システムロケールの自動検出機能

### 🔧 内部改善
- **ヘルプメッセージシステム**: `InitHelpMessagesWithLocale` 関数で動的言語切り替え
- **コマンドライン解析**: `--locale` フラグを事前解析してヘルプ表示に反映
- **新規ファイル追加**:
  - `internal/i18n/help.go`: ヘルプメッセージ管理システム
- **言語判定の優先順位**: `--locale` > `KONST_LOCALE` > システムロケール > デフォルト(`en`)

### 📝 使用方法
```bash
# 日本語ヘルプを表示
konst --help --locale=ja

# 英語ヘルプを表示  
konst --help --locale=en

# 環境変数で日本語設定
KONST_LOCALE=ja konst --help
```

## v0.3.4 (2025-06-28)

### 🚀 主要機能追加
- **依存関係解決システム**: `{{variable}}` 形式の変数参照と自動展開機能
  - 例: `"{{BaseRetries}} * 2"` → `"3 * 2"` → `6`
  - 循環依存の自動検出とエラーハンドリング
  - 基本的な四則演算のサポート (`+`, `-`, `*`, `/`)
- **数式評価エンジン**: 文字列内の数式を実際の値に計算変換

### 🐛 バグ修正  
- **TypeScript生成の重大なバグ修正**: Go構造体の生の値が出力される問題を解決
  - 修正前: `export const Value = {int 42 [] [] };`
  - 修正後: `export const Value = 42;`
- **formatTSConstValue関数**: Definition型の適切な処理でTypeScript値の正確な生成

### 🔧 内部改善
- **依存関係グラフ**: 全JSONファイルを先読みして依存順にソート処理
- **新規ファイル追加**:
  - `internal/utils/dependency_resolver.go`: 依存関係解決エンジン
  - `internal/process/file_resolver.go`: 解決済み定義での処理ロジック
- **処理フロー改善**: ディレクトリ処理時の依存関係を事前解決

## v0.3.3 (2025-06-28)

### ✨ 新機能
- **テンプレート文字列サポート**: `%parameter%` 形式でパラメータ置換が可能なテンプレート型を追加
- **キャメルケース変換**: テンプレート引数名を自動的にキャメルケースに変換 (`twitch_id` → `twitchId`)
- **テスト用Taskfile**: 開発・テスト用のタスクを追加
  - `task test:validate` - JSON定義の検証
  - `task test:dry-run` - 生成予定ファイルの確認  
  - `task test:gen` - 完全テスト（検証→生成→確認）
  - `task clean` - 生成ファイルの削除

### 🔄 改善
- **ディレクトリ専用化**: 単一ファイル処理を完全削除し、ディレクトリ指定のみに変更
- **階層的処理**: example構造を整理し、階層的なJSONファイル処理をサポート
- **使用方法の統一**: 全ての機能でディレクトリ指定を必須化

### 📁 リファクタリング
- `example/multi/` 構造を `example/` 直下に統合
- Taskfileからmulti用タスクを削除（冗長なため）
- Usage表示を `<inputDirectory>` に変更

## v0.3.0 (2025-06-28)

### 🚀 主要機能追加
- **enum型完全サポート**: 
  - Go: カスタム型、定数定義、バリデーション関数、パーサー関数
  - TypeScript: const assertion、型ガード、パーサー関数
  - デフォルト値サポート
- **開発支援機能**:
  - `--validate` フラグ: JSON定義の検証のみ実行
  - `--dry-run` フラグ: 生成予定ファイル一覧の表示
  - `--watch` フラグ: ファイル変更監視（実験的機能）

### 🔧 内部改善
- テンプレート関数の拡張 (`hasEnum`, `toTitle`)
- import文の最適化（必要な場合のみstrings, errorsをimport）
- enum用バリデーション・パーサー関数の自動生成

## v0.2.0 (2025-06-28)

### 🗂️ アーキテクチャ変更
- 単一ファイル処理を削除（ディレクトリ処理に統一）
- ファイル処理ロジックの簡素化

## v0.1.2 (2025-04-14)

### ✨ 機能追加
- **出力モード選択**: `-m` フラグでGo/TypeScript出力を選択可能
- **バージョン表示**: `-v` フラグでバージョン情報を表示

### 📝 ドキュメント
- exampleファイルの更新
- templateファイルの手直し

## v0.1.x (2025-04-05 - 2025-04-14)

### 🎯 初期機能実装
- **基本的なコード生成**:
  - Go: package名指定、基本型サポート
  - TypeScript: 基本型変換
- **型サポート**: int, string, bool, date, array型など
- **ディレクトリ入力対応**: 複数JSONファイルの一括処理
- **テンプレートシステム**: カスタマイズ可能なコード生成

### 🔄 設計方針変更
- enum, object型の対応を一旦中止（後のv0.3.0で実装）
- シンプルな型システムに集中

## v0.0.x (2025-04-04)

### 🎉 プロジェクト開始
- 初期プロジェクト構造の作成
- 基本的なJSON→Go/TypeScriptコード生成の実装
- GoのPackage名指定機能

---

## 開発の方向性

Konstは、JSON定義からGo/TypeScriptの定数・型定義を生成するツールとして開発されています。

**設計思想**:
- シンプルで使いやすいCLI
- 型安全なコード生成
- 開発効率の向上
- protobuf代替としての軽量な型共有

**今後の予定**:
- より高度なテンプレートカスタマイズ
- パフォーマンス最適化
- 複雑な数式サポートの拡張