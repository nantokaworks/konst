package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nantokaworks/konst/internal/types"
)

// ResolveDependencies は定義間の依存関係を解決して値を展開します
func ResolveDependencies(definitions map[string]types.Definition) (map[string]types.Definition, error) {
	resolved := make(map[string]types.Definition)
	processing := make(map[string]bool)

	var resolve func(string) error
	resolve = func(name string) error {
		if _, ok := resolved[name]; ok {
			return nil // すでに解決済み
		}
		if processing[name] {
			return fmt.Errorf("circular dependency detected: %s", name)
		}

		def, exists := definitions[name]
		if !exists {
			return fmt.Errorf("undefined dependency: %s", name)
		}

		processing[name] = true

		// 値が文字列で依存関係を含む場合、展開する
		if strValue, ok := def.Value.(string); ok {
			expandedValue, err := expandDependencies(strValue, definitions, resolve)
			if err != nil {
				return err
			}

			// 展開された値を適切な型に変換
			convertedValue, err := convertToTargetType(expandedValue, def.Type)
			if err != nil {
				return err
			}

			def.Value = convertedValue
		}

		resolved[name] = def
		delete(processing, name)
		return nil
	}

	// すべての定義を解決
	for name := range definitions {
		if err := resolve(name); err != nil {
			return nil, err
		}
	}

	return resolved, nil
}

// expandDependencies は文字列内の依存関係を展開します
func expandDependencies(value string, definitions map[string]types.Definition, resolve func(string) error) (string, error) {
	re := regexp.MustCompile(`\{\{([^}]+)\}\}`)
	
	result := re.ReplaceAllStringFunc(value, func(match string) string {
		// {{name}} から name を抽出
		depName := strings.TrimSpace(match[2 : len(match)-2])
		
		// 依存関係を再帰的に解決
		if err := resolve(depName); err != nil {
			return match // エラーの場合は元の値を返す
		}
		
		// 依存する定義の値を取得
		if dep, exists := definitions[depName]; exists {
			return fmt.Sprintf("%v", dep.Value)
		}
		
		return match // 見つからない場合は元の値を返す
	})
	
	return result, nil
}

// convertToTargetType は文字列を指定された型に変換します
func convertToTargetType(value string, targetType types.DefinitionType) (interface{}, error) {
	switch targetType {
	case types.DefinitionTypeInt, types.DefinitionTypeInt32, types.DefinitionTypeInt64:
		// 数式の評価（簡単な計算のみサポート）
		result, err := evaluateSimpleExpression(value)
		if err != nil {
			return value, nil // 計算できない場合は文字列のまま
		}
		return result, nil
	case types.DefinitionTypeFloat, types.DefinitionTypeFloat32, types.DefinitionTypeFloat64:
		result, err := evaluateSimpleExpression(value)
		if err != nil {
			return value, nil
		}
		return float64(result), nil
	case types.DefinitionTypeString:
		return value, nil
	default:
		return value, nil
	}
}

// evaluateSimpleExpression は簡単な数式を評価します（+ - * / のみサポート）
func evaluateSimpleExpression(expr string) (int, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	
	// 乗算と除算を先に処理
	for strings.Contains(expr, "*") || strings.Contains(expr, "/") {
		re := regexp.MustCompile(`(\d+)\s*([*/])\s*(\d+)`)
		match := re.FindStringSubmatch(expr)
		if len(match) != 4 {
			break
		}
		
		left, _ := strconv.Atoi(match[1])
		op := match[2]
		right, _ := strconv.Atoi(match[3])
		
		var result int
		switch op {
		case "*":
			result = left * right
		case "/":
			if right == 0 {
				return 0, errors.New("division by zero")
			}
			result = left / right
		}
		
		expr = strings.Replace(expr, match[0], strconv.Itoa(result), 1)
	}
	
	// 加算と減算を処理
	for strings.Contains(expr, "+") || strings.Contains(expr, "-") {
		re := regexp.MustCompile(`(\d+)\s*([+-])\s*(\d+)`)
		match := re.FindStringSubmatch(expr)
		if len(match) != 4 {
			break
		}
		
		left, _ := strconv.Atoi(match[1])
		op := match[2]
		right, _ := strconv.Atoi(match[3])
		
		var result int
		switch op {
		case "+":
			result = left + right
		case "-":
			result = left - right
		}
		
		expr = strings.Replace(expr, match[0], strconv.Itoa(result), 1)
	}
	
	// 最終的に数値のみが残っているかチェック
	return strconv.Atoi(expr)
}