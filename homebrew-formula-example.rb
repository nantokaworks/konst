# このファイルは Homebrew Formula の例です
# 実際に使用する場合は homebrew-core にPRを送るか、独自のtapを作成してください

class Konst < Formula
  desc "JSON定義からGo・TypeScriptのコードを自動生成するツール"
  homepage "https://github.com/nantokaworks/konst"
  url "https://github.com/nantokaworks/konst/archive/v0.3.4.tar.gz"
  sha256 "YOUR_SHA256_HERE"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
    
    # メッセージファイルをインストール
    (pkgshare/"messages").install Dir["messages/*.json"]
    
    # メッセージファイルへのシンボリックリンクを作成
    (bin/"messages").install_symlink Dir["#{pkgshare}/messages/*"]
  end

  test do
    system "#{bin}/konst", "--version"
  end
end