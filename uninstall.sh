#!/bin/bash

# Konst アンインストールスクリプト

echo "🗑️  Konst をアンインストールしています..."

# インストール先を探す
KONST_PATH=$(which konst 2>/dev/null)

if [ -z "$KONST_PATH" ]; then
    echo "❌ Konst が見つかりません"
    exit 1
fi

INSTALL_DIR=$(dirname "$KONST_PATH")

# 確認
echo "以下のファイルを削除します:"
echo "  - $KONST_PATH"
echo "  - $INSTALL_DIR/messages/"
echo ""
read -p "続行しますか？ (y/N): " -n 1 -r
echo ""

if [[ $REPLY =~ ^[Yy]$ ]]; then
    # 削除
    rm -f "$KONST_PATH"
    rm -rf "$INSTALL_DIR/messages"
    
    echo "✅ アンインストール完了!"
else
    echo "❌ アンインストールをキャンセルしました"
fi