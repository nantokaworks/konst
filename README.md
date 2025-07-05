# Konst

> JSON定義から Go・TypeScript のコードを自動生成するツール

**Konst** は、JSON で定義された定数・列挙型の情報から Go と TypeScript のコードを自動生成するツールです。  
API通信での型安全性を向上させ、多言語間での定数管理を一元化できます。

## ✨ 特長

| 機能 | 説明 |
|---|---|
| 🔄 **多言語出力** | JSON定義から Go・TypeScript の型安全なコードを生成 |
| 🛡️ **enum型完全サポート** | バリデーション・パーサー関数付きのenum自動生成 |
| 🌐 **API通信に最適** | protobuf文字列通信でのenum値検証に最適 |
| 🔧 **開発支援機能** | バリデーション・ドライラン・ウォッチモード |
| 📁 **ディレクトリ処理** | 複数ファイル一括処理とindex.ts自動生成 |
| 🗣️ **i18n対応** | 日本語・英語のヘルプメッセージとシステムロケール自動検出 |
| 📦 **パッケージ分離** | 異なるgoPackageを持つファイルを自動的に別ディレクトリに配置 |

## 📋 JSON 定義フォーマット

### 💻 基本的な定数定義

```json
{
  "version": "1.0",
  "goPackage": "constants",
  "definitions": {
    "MaxRetries": {
      "type": "int",
      "value": 3
    },
    "ApiTimeout": {
      "type": "int64",
      "value": 30000,
      "tsMode": "number"
    }
  }
}
```

### 🏷️ enum型定義（API通信に最適）

```json
{
  "version": "1.0", 
  "goPackage": "enums",
  "definitions": {
    "UserStatus": {
      "type": "enum",
      "values": ["active", "inactive", "pending"],
      "default": "pending"
    },
    "Priority": {
      "type": "enum",
      "values": ["low", "medium", "high"]
    }
  }
}
```

### 📝 フィールド仕様

<details>
<summary><strong>📌 共通フィールド</strong></summary>

| フィールド | 必須 | 説明 | 例 |
|---|---|---|---|
| `version` | ✅ | JSON定義フォーマットのバージョン | `"1.0"` |
| `goPackage` | ✅ | 生成されるGoパッケージ名 | `"constants"` |

</details>

<details>
<summary><strong>🔢 定数型フィールド</strong></summary>

| フィールド | 必須 | 説明 | 例 |
|---|---|---|---|
| `type` | ✅ | 定義の型 | `"int"`, `"string"`, `"bool"` |
| `value` | ✅ | 実際のリテラル値 | `42`, `"hello"` |
| `tsMode` | ❌ | TypeScript用出力指定 | `"number"`, `"bigint"` |

</details>

<details>
<summary><strong>🏷️ enum型フィールド</strong></summary>

| フィールド | 必須 | 説明 | 例 |
|---|---|---|---|
| `type` | ✅ | `"enum"` | `"enum"` |
| `values` | ✅ | 文字列配列（選択肢） | `["active", "inactive"]` |
| `default` | ❌ | デフォルト値 | `"active"` |

</details>

### 🗂️ サポートする型

#### 基本型
| 型 | Go出力 | TypeScript出力 | 備考 |
|---|---|---|---|
| `int` | `int` | `number` | |
| `int32` | `int32` | `number` | |
| `int64` | `int64` | `bigint` | tsMode:"number"で変更可 |
| `string` | `string` | `string` | |
| `bool` | `bool` | `boolean` | |
| `date` | `time.Time` | `Date` | 各種モード指定可 |

#### 🆕 enum型（v0.3.0）
| 型 | Go出力 | TypeScript出力 |
|---|---|---|
| `enum` | カスタム型 + バリデーション関数群 | const object + 型 + 関数群 |

#### 配列型
各型に `[]` を付けて配列型として定義：
- `int[]`, `string[]`, `bool[]`, `date[]`, `enum[]` など

## 🚀 インストール

```bash
go install github.com/yourusername/konst@latest
```

> ※ `github.com/yourusername/konst` は実際のリポジトリパスに置き換えてください

## 📖 使い方

Konst は **出力ディレクトリ（`-o`）** と **出力モード（`-m`）** の指定が必要です。  
入力は単一ファイルまたはディレクトリを指定できます。

### 📦 Go パッケージ分離機能

Go出力では、異なる `goPackage` を持つJSONファイルが自動的に別々のディレクトリに配置されます：

```bash
# 入力ディレクトリ構成
definitions/
├── config.json    # "goPackage": "config"
└── database.json  # "goPackage": "database"

# 出力結果
generated/
├── config/
│   └── config.go     # package config
└── database/
    └── database.go   # package database
```

これにより、同じディレクトリに異なるパッケージが混在することによるGoのビルドエラーを防げます。

### 🔥 基本的な使い方

```bash
# TypeScript生成
konst -i constants.json -o generated/ -m ts

# Go生成  
konst -i constants.json -o generated/ -m go

# ディレクトリ一括処理
konst -i definitions/ -o generated/ -m ts -f
```

### 🛠️ 開発支援機能

| 機能 | コマンド | 説明 |
|---|---|---|
| 🔍 **バリデーション** | `konst --validate -i constants.json` | JSON検証のみ実行 |
| 👀 **ドライラン** | `konst --dry-run -i definitions/ -o generated/ -m ts` | 生成予定ファイル確認 |
| 👁️ **ウォッチモード** | `konst --watch -i definitions/ -o generated/ -m ts` | ファイル変更監視（実験的） |

### ⚙️ コマンドオプション

| オプション | 必須 | 説明 | 例 |
|---|---|---|---|
| `-i` | ❌ | 入力ファイル/ディレクトリ | `-i constants.json` |
| `-o` | ✅ | 出力ディレクトリ | `-o generated/` |
| `-m` | ✅ | 出力モード（go/ts） | `-m ts` |
| `-f` | ❌ | 強制上書き | `-f` |
| `--validate` | ❌ | バリデーションのみ | `--validate` |
| `--dry-run` | ❌ | 生成予定ファイル表示 | `--dry-run` |
| `--watch` | ❌ | ファイル監視（実験的） | `--watch` |
| `-t` | ❌ | カスタムテンプレートDir | `-t ./templates` |
| `--indent` | ❌ | インデント数 | `--indent 4` |
| `--naming` | ❌ | ファイル命名規則 | `--naming kebab` |
| `--locale` | ❌ | 🌐 言語設定（ja/en） | `--locale ja` |

### 📛 ファイル命名規則

`--naming` オプションで出力ファイル・ディレクトリの命名規則を指定できます：

- `kebab`: kebab-case（例: `user-status.ts`）
- `camel`: camelCase（例: `userStatus.ts`）  
- `snake`: snake_case（例: `user_status.go`）

デフォルト: TypeScript は `kebab-case`、Go は `snake_case`

```bash
# TypeScript でデフォルト（kebab-case）
konst -i ./definitions/user_status.json -o generated/ -m ts
# → generated/user-status.ts

# Go で camelCase を指定
konst -i ./definitions/user-status.json -o generated/ -m go --naming camel
# → generated/userStatus.go
```

### 🌐 言語設定

`--locale` オプションでヘルプメッセージの言語を指定できます：

```bash
# 日本語でヘルプを表示
konst --help --locale=ja

# 英語でヘルプを表示
konst --help --locale=en

# 環境変数での設定
export KONST_LOCALE=ja
konst --help

# システムロケール自動検出（デフォルト）
konst --help
```

**言語判定の優先順位:**
1. `--locale` フラグ
2. `KONST_LOCALE` 環境変数  
3. システムロケール（`LC_ALL`, `LC_MESSAGES`, `LANG`, `LC_CTYPE`）
4. デフォルト（英語）

### 🎨 カスタムテンプレート

```bash
# テンプレートファイルを go.tmpl、ts.tmpl として配置
konst -i constants.json -o generated/ -m ts -t ./custom-templates/
```

## 💡 生成されるコード例

### 🏷️ enum型の生成例

<details>
<summary><strong>📋 JSON定義</strong></summary>

```json
{
  "UserStatus": {
    "type": "enum",
    "values": ["active", "inactive", "pending"],
    "default": "pending"
  }
}
```

</details>

<details>
<summary><strong>🐹 Go出力例</strong></summary>

```go
type UserStatus string

const (
    UserStatusActive   UserStatus = "active"
    UserStatusInactive UserStatus = "inactive"
    UserStatusPending  UserStatus = "pending"
)

// バリデーション関数
func IsValidUserStatus(value string) bool { /* ... */ }

// パーサー関数（エラーハンドリング付き）
func ParseUserStatus(value string) (UserStatus, error) { /* ... */ }

// 全ての値を取得
func GetAllUserStatusValues() []UserStatus { /* ... */ }

// デフォルト値を取得
func GetDefaultUserStatus() UserStatus { /* ... */ }
```

</details>

<details>
<summary><strong>🔷 TypeScript出力例</strong></summary>

```typescript
export const UserStatus = {
    Active: "active",
    Inactive: "inactive", 
    Pending: "pending"
} as const;

export type UserStatusType = typeof UserStatus[keyof typeof UserStatus];

// 型ガード関数
export function isValidUserStatus(value: string): value is UserStatusType { /* ... */ }

// パーサー関数（例外投げる版）
export function parseUserStatus(value: string): UserStatusType { /* ... */ }

// パーサー関数（undefinedを返す版）
export function parseUserStatusSafe(value: string): UserStatusType | undefined { /* ... */ }

// 全ての値を取得
export function getAllUserStatusValues(): UserStatusType[] { /* ... */ }

// デフォルト値を取得
export function getDefaultUserStatus(): UserStatusType { /* ... */ }
```

</details>

### 🌐 API通信での活用例

```typescript
// ✅ API受信時の安全な検証
function handleUserData(data: any) {
  if (!isValidUserStatus(data.status)) {
    throw new Error('Invalid status from API');
  }
  const status = parseUserStatus(data.status); // 型安全 ✨
}

// ✅ protobuf通信での型安全性
const user: User = {
  status: UserStatus.Active // コンパイル時型チェック ✨
};
```

## 📄 ライセンス

このプロジェクトは MIT ライセンスの下で公開されています。