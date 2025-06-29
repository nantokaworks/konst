#!/bin/bash

# Konst インストールスクリプト

set -e

echo "🚀 Konst をインストールしています..."

# ビルド
echo "📦 ビルド中..."
go build -ldflags="-s -w" -o konst .

# インストール先を決定
if [ -n "$GOPATH" ]; then
    INSTALL_DIR="$GOPATH/bin"
else
    INSTALL_DIR="$HOME/go/bin"
fi

# ディレクトリ作成
mkdir -p "$INSTALL_DIR"
mkdir -p "$INSTALL_DIR/messages"

# 実行ファイルをコピー
echo "📋 実行ファイルをインストール中..."
cp konst "$INSTALL_DIR/"

# メッセージファイルをコピー
echo "🌐 メッセージファイルをインストール中..."
cp messages/*.json "$INSTALL_DIR/messages/"

# 実行権限を付与
chmod +x "$INSTALL_DIR/konst"

# パスの確認
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo ""
    echo "⚠️  警告: $INSTALL_DIR がPATHに含まれていません"
    echo "以下のコマンドを ~/.bashrc, ~/.zshrc などに追加してください:"
    echo ""
    echo "export PATH=\"\$PATH:$INSTALL_DIR\""
    echo ""
fi

echo "✅ インストール完了!"
echo "📍 インストール先: $INSTALL_DIR/konst"
echo ""
echo "使い方: konst -i ./example -o ./output -m ts"