package i18n

import (
	"os"
	"strings"
)

// DetectSystemLocale はシステムのロケール設定を検出します
func DetectSystemLocale() string {
	// 1. LC_ALL環境変数をチェック（全ての言語設定を上書き、最優先）
	if lcAll := os.Getenv("LC_ALL"); lcAll != "" {
		return parseLocale(lcAll)
	}

	// 2. LC_MESSAGES環境変数をチェック（メッセージ専用）
	if lcMessages := os.Getenv("LC_MESSAGES"); lcMessages != "" {
		return parseLocale(lcMessages)
	}

	// 3. LANG環境変数をチェック（最も一般的）
	if lang := os.Getenv("LANG"); lang != "" {
		return parseLocale(lang)
	}

	// 4. LC_CTYPE環境変数をチェック（文字分類）
	if lcCtype := os.Getenv("LC_CTYPE"); lcCtype != "" {
		return parseLocale(lcCtype)
	}

	// 5. デフォルトは英語
	return "en"
}

// parseLocale はロケール文字列から言語コードを抽出します
// 例: "ja_JP.UTF-8" -> "ja", "en_US.UTF-8" -> "en"
func parseLocale(locale string) string {
	if locale == "" || locale == "C" || locale == "POSIX" {
		return "en"
	}

	// "_"で分割して言語部分を取得
	parts := strings.Split(locale, "_")
	if len(parts) == 0 {
		return "en"
	}

	lang := strings.ToLower(parts[0])

	// サポートしている言語かチェック
	switch lang {
	case "ja", "japanese":
		return "ja"
	case "en", "english":
		return "en"
	default:
		// サポートしていない言語の場合は英語にフォールバック
		return "en"
	}
}

// GetSupportedLocales はサポートしている言語のリストを返します
func GetSupportedLocales() []string {
	return []string{"ja", "en"}
}

// IsSupported は指定された言語がサポートされているかチェックします
func IsSupported(locale string) bool {
	supported := GetSupportedLocales()
	for _, lang := range supported {
		if lang == locale {
			return true
		}
	}
	return false
}