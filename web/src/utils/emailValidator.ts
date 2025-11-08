/**
 * Email格式验证工具函数
 * 验证email格式是否符合标准
 */
export const validateEmail = (email: string): boolean => {
  // 基本的email格式正则表达式
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email);
};

/**
 * 获取email验证的错误信息
 */
export const getEmailValidationMessage = (email: string): string => {
  if (!email || email.trim().length === 0) {
    return "邮箱不能为空";
  }

  if (!validateEmail(email)) {
    return "请输入有效的邮箱地址";
  }

  return "";
};
