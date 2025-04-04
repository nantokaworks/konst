# Konst

**Konst** は、JSON で定義された定数、列挙型、オブジェクト（構造体／インターフェース）の情報から、Go および TypeScript のコードを自動生成するツールです。AIで生成されたコードを中心に調整中...  

---

## 特長

- **多言語出力**  
  JSON 定義から Go と TypeScript のコードを生成します。

- **JSON 定義**  
  各定義は `descriptor` と `content` を利用して記述します。  
  - **`descriptor`**: 定義の種類（`const`、`enum`、`object` など）を示します。  
  - **`content`**: 定義内容そのものを保持します。  
    定数の場合は型（例: `"int"`, `"int64"` など）やリテラル値（`value` キー）、TypeScript 用変換指定（`tsMode` キー）を指定。  
    列挙型は `values`、オブジェクトは `fields` を使って定義します。

- **カスタムテンプレート**  
  内蔵テンプレートに加え、カスタムテンプレートディレクトリを指定することで出力コードのフォーマットを自由にカスタマイズ可能です。  
  テンプレートファイル名はシンプルに `go.tmpl` および `ts.tmpl` とします。

---

## JSON 定義フォーマット例

以下は、Konst で使用する JSON 定義ファイルのサンプルです。  
この例では、定数、列挙型、オブジェクトの各定義が含まれています。

```json
{
  "version": "1.0",
  "definitions": {
    "MaxItems": {
      "descriptor": "const",
      "content": {
        "type": "int",
        "value": 100
      }
    },
    "LargeNumber": {
      "descriptor": "const",
      "content": {
        "type": "int64",
        "value": 9223372036854775807,
        "tsMode": "number"
      }
    },
    "DateAt": {
      "descriptor": "const",
      "content": {
        "type": "date",
        "value": "2025-04-04T12:34:56Z"
      }
    },
    "DateStringAt": {
      "descriptor": "const",
      "content": {
        "type": "date",
        "value": "2025-04-04T12:34:56Z",
        "tsMode": "string"
      }
    },
    "Status": {
      "descriptor": "enum",
      "content": {
        "values": {
          "Inactive": "inactive",
          "Active": "active",
          "Pending": "pending"
        }
      }
    },
    "User": {
      "descriptor": "object",
      "content": {
        "fields": {
          "Id": { "type": "string" },
          "Name": { "type": "string" },
          "Age": { "type": "int" }
        }
      }
    }
  }
}
```

### 各フィールドの説明

- **version**  
  JSON 定義フォーマットのバージョン。将来的な拡張に備えています。

- **definitions**  
  各定義はキー名で識別され、次の情報を持ちます:
  - **descriptor**: 定義の種類  
    - `"const"`: 定数  
    - `"enum"`: 列挙型  
    - `"object"`: オブジェクト（構造体／インターフェース）
  - **content**: 定義内容を保持  
    - 定数の場合:
      - `"type"`: 定数の型（例: `"int"`, `"int64"`, `"string"`, `"string[]"`, `"date"` など）
      - `"value"`: 実際のリテラル値
      - `"tsMode"` (オプション): たとえば `"int64"` の場合、TypeScript で number として出力する場合に `"number"` を指定。`"date"` の場合、TypeScript で string として出力する場合に `"string"` を指定
    - 列挙型の場合:
      - `"values"`: 各メンバーの名前と値のマッピング
    - オブジェクトの場合:
      - `"fields"`: 各フィールドの名前と型情報

---

## インストール

Konst は Go モジュールとして管理されているので、以下のようにインストールできます:

```bash
go install github.com/yourusername/konst@latest
```

※ `github.com/yourusername/konst` は実際のリポジトリパスに置き換えてください。

---

## 使い方

Konst は、出力先ファイル名（`-o` オプション）が必須です。  
入力ファイルは `-i` オプションで指定しますが、指定がなければデフォルトで `constants.json` が使用されます。  
また、カスタムテンプレートディレクトリは `-t` オプションまたは環境変数 `KONST_TEMPLATES` で指定可能です。

### 基本例

```bash
konst constants.json -o=output/konst.ts
```

または

```bash
konst -i=constants.json -o=output/konst.go
```

### カスタムテンプレートの利用

カスタムテンプレートを使用する場合は、テンプレートディレクトリを指定してください。  
テンプレートファイルは以下のように配置します:

- **Go 用テンプレート**: `go.tmpl`
- **TypeScript 用テンプレート**: `ts.tmpl`

例:

```bash
konst -i=constants.json -o=output/konst.go -t=/path/to/custom/templates
```

または環境変数を利用:

```bash
export KONST_TEMPLATES=/path/to/custom/templates
konst -i=constants.json -o=output/konst.ts
```

---

## テンプレート例

exampleを参照してください

---

## ライセンス

このプロジェクトは MIT ライセンスの下で公開されています。