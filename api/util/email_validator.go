package util

import (
	"regexp"
)

// ValidateEmail 验证email格式是否符合标准
func ValidateEmail(email string) bool {
	// 基本的email格式正则表达式
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailRegex, email)
	return matched
}

// GetEmailValidationMessage 获取email验证的错误信息
func GetEmailValidationMessage(email string) string {
	if email == "" {
		return "邮箱不能为空"
	}

	if !ValidateEmail(email) {
		return "请输入有效的邮箱地址"
	}

	return ""
}
