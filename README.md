# Konst

**Konst** は、JSON で定義された定数、列挙型、オブジェクト（構造体／インターフェース）の情報から、Go および TypeScript のコードを自動生成するツールです。AIで生成されたコードを中心に調整中...  

---

## 特長

- **多言語出力**  
  JSON 定義から Go と TypeScript のコードを生成します。

- **JSON 定義**  
  各定義は `type` と `value` を利用して記述します。  
  - **`type`**: 定義の型（例: `"int"`, `"int64"`, `"date"` など）を示します。  
  - **`value`**: 定義内容そのものを保持します。  
    定数の場合は型（例: `"int"`, `"int64"` など）やリテラル値（`value` キー）、TypeScript 用変換指定（`mode` キー）を指定。

- **カスタムテンプレート**  
  内蔵テンプレートに加え、カスタムテンプレートディレクトリを指定することで出力コードのフォーマットを自由にカスタマイズ可能です。  
  テンプレートファイル名はシンプルに `go.tmpl` および `ts.tmpl` とします。

---

## JSON 定義フォーマット例

以下は、Konst で使用する JSON 定義ファイルのサンプルです。  
この例では、定数の定義が含まれています。

```json
{
  "version": "1.0",
  "goPackage": "nantoka",
  "definitions": {
    "MaxItems": {
      "type": "int",
      "value": 100
    },
    "LargeNumber": {
      "type": "int64",
      "value": 9223372036854775807,
      "mode": "number"
    },
    "DateAt": {
      "type": "date",
      "value": "2025-04-04T12:34:56Z"
    },
    "DateStringAt": {
      "type": "date",
      "value": "2025-04-04T12:34:56Z",
      "mode": "string"
    }
  }
}
```

### 各フィールドの説明

- **version**  
  JSON 定義フォーマットのバージョン。

- **definitions**  
  各定義はキー名で識別され、以下の情報を持ちます:
  - **type**: 定義の型（例: `"int"`、`"int64"`, `"date"` など）
  - **value**: 実際のリテラル値  
  - **mode** (オプション): TypeScript 用の出力指定（例: `"number"`、`"string"` など）

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