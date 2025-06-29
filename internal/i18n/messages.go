package i18n

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// MessageKey はメッセージキーの型
type MessageKey string

// メッセージキーの定数定義
const (
	MsgMode                MessageKey = "mode"
	MsgOutputDirectory     MessageKey = "output_directory"
	MsgFilesToBeGenerated  MessageKey = "files_to_be_generated"
	MsgGenerated           MessageKey = "generated"
	MsgValidationSuccess   MessageKey = "validation_success"
	MsgInputMustBeDir      MessageKey = "input_must_be_directory"
	MsgOutputMustBeDir     MessageKey = "output_must_be_directory"
	MsgCmdArgError         MessageKey = "command_argument_error"
	MsgValidationError     MessageKey = "validation_error"
	MsgDryRunError         MessageKey = "dry_run_error"
	MsgInputPathError      MessageKey = "input_path_error"
	MsgProcessingError     MessageKey = "processing_error"
	MsgFileError           MessageKey = "file_error"
	MsgOutputRequired      MessageKey = "output_required"
	MsgExecutablePathError MessageKey = "executable_path_error"
)

// Messages は言語別のメッセージを管理する構造体
type Messages struct {
	messages map[string]string
	locale   string
}

// defaultMessages はデフォルト（英語）のメッセージ
var defaultMessages = map[MessageKey]string{
	MsgMode:                "Mode",
	MsgOutputDirectory:     "Output directory",
	MsgFilesToBeGenerated:  "Files to be generated",
	MsgGenerated:           "Generated",
	MsgValidationSuccess:   "Validation successful: No issues found in JSON definitions",
	MsgInputMustBeDir:      "input must be a directory",
	MsgOutputMustBeDir:     "output must be a directory",
	MsgCmdArgError:         "Command line argument error",
	MsgValidationError:     "Validation error",
	MsgDryRunError:         "Dry-run error",
	MsgInputPathError:      "Input path error",
	MsgProcessingError:     "Processing error",
	MsgFileError:           "file",
	MsgOutputRequired:      "please specify output filename with -o option",
	MsgExecutablePathError: "executable path error",
}

var globalMessages *Messages

// Init は指定された言語でメッセージシステムを初期化します
func Init(locale string) error {
	globalMessages = &Messages{
		locale: locale,
	}
	
	// デフォルトメッセージをマップにコピー
	globalMessages.messages = make(map[string]string)
	for key, value := range defaultMessages {
		globalMessages.messages[string(key)] = value
	}
	
	// 指定された言語のメッセージファイルを読み込み
	if locale != "en" {
		if err := globalMessages.loadMessages(locale); err != nil {
			// ファイルが見つからない場合は英語のデフォルトを使用
			return nil
		}
	}
	
	return nil
}

// loadMessages は指定された言語のメッセージファイルを読み込みます
func (m *Messages) loadMessages(locale string) error {
	// 実行ファイルの場所を取得
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	
	// メッセージファイルのパスを構築
	messagesPath := filepath.Join(filepath.Dir(exePath), "messages", locale+".json")
	
	// ファイルが存在しない場合は現在のディレクトリからも探す
	if _, err := os.Stat(messagesPath); os.IsNotExist(err) {
		messagesPath = filepath.Join("messages", locale+".json")
	}
	
	// ファイルを読み込み
	data, err := os.ReadFile(messagesPath)
	if err != nil {
		return err
	}
	
	// JSONをパース
	var localeMessages map[string]string
	if err := json.Unmarshal(data, &localeMessages); err != nil {
		return err
	}
	
	// メッセージを更新
	for key, value := range localeMessages {
		m.messages[key] = value
	}
	
	return nil
}

// T は指定されたキーのメッセージを取得します（翻訳関数）
func T(key MessageKey) string {
	if globalMessages == nil {
		// 初期化されていない場合はデフォルト値を返す
		if msg, ok := defaultMessages[key]; ok {
			return msg
		}
		return string(key)
	}
	
	if msg, ok := globalMessages.messages[string(key)]; ok {
		return msg
	}
	
	// メッセージが見つからない場合はデフォルト値を返す
	if msg, ok := defaultMessages[key]; ok {
		return msg
	}
	
	return string(key)
}

// GetLocale は現在の言語設定を返します
func GetLocale() string {
	if globalMessages == nil {
		return "en"
	}
	return globalMessages.locale
}