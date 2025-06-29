package i18n

import (
	"os"
	"testing"
)

func TestParseLocale(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"日本語（標準）", "ja_JP.UTF-8", "ja"},
		{"日本語（シンプル）", "ja", "ja"},
		{"英語（標準）", "en_US.UTF-8", "en"},
		{"英語（シンプル）", "en", "en"},
		{"英語（イギリス）", "en_GB.UTF-8", "en"},
		{"空文字", "", "en"},
		{"C", "C", "en"},
		{"POSIX", "POSIX", "en"},
		{"サポート外言語", "fr_FR.UTF-8", "en"},
		{"不正な形式", "invalid", "en"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseLocale(tt.input)
			if result != tt.expected {
				t.Errorf("parseLocale(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDetectSystemLocale(t *testing.T) {
	// 元の環境変数を保存
	originalLang := os.Getenv("LANG")
	originalLcAll := os.Getenv("LC_ALL")
	originalLcMessages := os.Getenv("LC_MESSAGES")
	originalLcCtype := os.Getenv("LC_CTYPE")

	// テスト後に環境変数を復元
	defer func() {
		os.Setenv("LANG", originalLang)
		os.Setenv("LC_ALL", originalLcAll)
		os.Setenv("LC_MESSAGES", originalLcMessages)
		os.Setenv("LC_CTYPE", originalLcCtype)
	}()

	tests := []struct {
		name        string
		lang        string
		lcAll       string
		lcMessages  string
		lcCtype     string
		expected    string
	}{
		{
			name:     "LANG日本語",
			lang:     "ja_JP.UTF-8",
			expected: "ja",
		},
		{
			name:     "LANG英語",
			lang:     "en_US.UTF-8",
			expected: "en",
		},
		{
			name:      "LC_ALL優先",
			lang:      "en_US.UTF-8",
			lcAll:     "ja_JP.UTF-8",
			expected:  "ja",
		},
		{
			name:       "LC_MESSAGES",
			lcMessages: "ja_JP.UTF-8",
			expected:   "ja",
		},
		{
			name:     "LC_CTYPE",
			lcCtype:  "ja_JP.UTF-8",
			expected: "ja",
		},
		{
			name:     "全て空の場合",
			expected: "en",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 環境変数をクリア
			os.Unsetenv("LANG")
			os.Unsetenv("LC_ALL")
			os.Unsetenv("LC_MESSAGES")
			os.Unsetenv("LC_CTYPE")

			// テスト用の環境変数を設定
			if tt.lang != "" {
				os.Setenv("LANG", tt.lang)
			}
			if tt.lcAll != "" {
				os.Setenv("LC_ALL", tt.lcAll)
			}
			if tt.lcMessages != "" {
				os.Setenv("LC_MESSAGES", tt.lcMessages)
			}
			if tt.lcCtype != "" {
				os.Setenv("LC_CTYPE", tt.lcCtype)
			}

			result := DetectSystemLocale()
			if result != tt.expected {
				t.Errorf("DetectSystemLocale() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGetSupportedLocales(t *testing.T) {
	locales := GetSupportedLocales()
	expected := []string{"ja", "en"}

	if len(locales) != len(expected) {
		t.Errorf("GetSupportedLocales() returned %d locales, want %d", len(locales), len(expected))
		return
	}

	for i, locale := range locales {
		if locale != expected[i] {
			t.Errorf("GetSupportedLocales()[%d] = %q, want %q", i, locale, expected[i])
		}
	}
}

func TestIsSupported(t *testing.T) {
	tests := []struct {
		name     string
		locale   string
		expected bool
	}{
		{"日本語サポート", "ja", true},
		{"英語サポート", "en", true},
		{"フランス語非サポート", "fr", false},
		{"空文字", "", false},
		{"不正な値", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSupported(tt.locale)
			if result != tt.expected {
				t.Errorf("IsSupported(%q) = %v, want %v", tt.locale, result, tt.expected)
			}
		})
	}
}