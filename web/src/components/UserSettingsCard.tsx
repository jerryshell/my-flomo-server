import { useState, useEffect } from "react";
import userApi from "../api/userApi";
import pluginApi from "../api/pluginApi";
import api from "../api/api";
import { useAlert } from "../hooks/useModal";
import Alert from "./Alert";

const UserSettingsCard = () => {
  // 密码更新状态
  const [newPassword, setNewPassword] = useState("");
  const [newPassword2, setNewPassword2] = useState("");
  const [passwordLoading, setPasswordLoading] = useState(false);

  // 用户偏好设置状态
  const [dailyReviewEnabled, setDailyReviewEnabled] = useState(false);
  const [preferencesLoading, setPreferencesLoading] = useState(false);

  // Telegram Bot 配置状态
  const [telegramChatId, setTelegramChatId] = useState("");
  const [telegramBotToken, setTelegramBotToken] = useState("");
  const [testing, setTesting] = useState(false);

  // 插件令牌状态
  const [pluginToken, setPluginToken] = useState("");
  const [tokenLoading, setTokenLoading] = useState(false);

  const { alertOpen, alertConfig, showAlert, hideAlert } = useAlert();

  // 初始化数据
  useEffect(() => {
    // 获取插件令牌
    pluginApi.getToken().then((response) => {
      if (response.data.success) {
        setPluginToken(response.data.data);
      }
    });

    // 获取用户设置
    userApi
      .getSettings()
      .then((response) => {
        if (response.data.success) {
          const settings = response.data.data;
          setDailyReviewEnabled(settings.dailyReviewEnabled || false);
          setTelegramChatId(settings.telegramChatId || "");
          setTelegramBotToken(settings.telegramBotToken || "");
        }
      })
      .catch((error) => {
        console.error("获取用户设置失败", error);
        // 如果获取失败，使用默认值
        setDailyReviewEnabled(false);
        setTelegramChatId("");
        setTelegramBotToken("");
      });
  }, []);

  // 更新密码
  const handleUpdatePassword = () => {
    if (newPassword !== newPassword2) {
      showAlert({
        message: "两次密码不一致，请重新输入",
        type: "warning",
        duration: 3000,
      });
      setNewPassword("");
      setNewPassword2("");
      return;
    }

    if (newPassword.length < 6) {
      showAlert({
        message: "密码长度不能少于6位",
        type: "warning",
        duration: 3000,
      });
      return;
    }

    setPasswordLoading(true);
    userApi
      .updatePassword({ password: newPassword })
      .then((response) => {
        if (response.data.success) {
          showAlert({
            message: "密码更新成功",
            type: "success",
            duration: 2000,
          });
          setNewPassword("");
          setNewPassword2("");
        } else {
          showAlert({
            message: response.data.message,
            type: "error",
            duration: 3000,
          });
        }
      })
      .catch((error) => {
        console.error("密码更新失败", error);
        showAlert({
          message: "密码更新失败，请重试",
          type: "error",
          duration: 3000,
        });
      })
      .finally(() => {
        setPasswordLoading(false);
      });
  };

  // 更新偏好设置
  const handleUpdatePreferences = () => {
    setPreferencesLoading(true);
    userApi
      .updateSettings({
        dailyReviewEnabled,
        telegramChatId,
        telegramBotToken,
      })
      .then((response) => {
        if (response.data.success) {
          showAlert({
            message: "偏好设置已保存！",
            type: "success",
            duration: 2000,
          });
        } else {
          showAlert({
            message: response.data.message,
            type: "error",
            duration: 3000,
          });
        }
      })
      .catch((error) => {
        console.error("偏好设置更新失败", error);
        showAlert({
          message: "偏好设置更新失败，请重试",
          type: "error",
          duration: 3000,
        });
      })
      .finally(() => {
        setPreferencesLoading(false);
      });
  };

  // 测试每日回顾
  const handleTestDailyReview = async () => {
    setTesting(true);
    try {
      const response = await userApi.triggerDailyReview();

      if (response.data.success) {
        showAlert({
          message: "每日回顾测试已触发！请检查您的 Telegram 消息",
          type: "success",
          duration: 3000,
        });
      } else {
        showAlert({
          message: response.data.message,
          type: "error",
          duration: 3000,
        });
      }
    } catch (error) {
      console.error("触发每日回顾失败", error);
      showAlert({
        message: "触发每日回顾失败，请检查配置后重试",
        type: "error",
        duration: 3000,
      });
    } finally {
      setTesting(false);
    }
  };

  // 生成插件令牌
  const handleGenerateToken = () => {
    setTokenLoading(true);
    pluginApi
      .createToken()
      .then((response) => {
        if (response.data.success) {
          setPluginToken(response.data.data);
          showAlert({
            message: "令牌生成成功",
            type: "success",
            duration: 2000,
          });
        } else {
          showAlert({
            message: response.data.message,
            type: "error",
            duration: 3000,
          });
        }
      })
      .catch((error) => {
        console.error("生成令牌失败", error);
        showAlert({
          message: "生成令牌失败，请稍后重试",
          type: "error",
          duration: 3000,
        });
      })
      .finally(() => {
        setTokenLoading(false);
      });
  };

  return (
    <>
      <div className="space-y-6">
        {/* 密码更新部分 */}
        <section>
          <h3 className="font-semibold text-lg mb-4 pb-2 border-b border-base-300">
            密码设置
          </h3>
          <div className="space-y-3">
            <div className="form-control">
              <label className="label">
                <span className="label-text">新密码</span>
              </label>
              <input
                type="password"
                className="input input-bordered input-sm"
                value={newPassword}
                onChange={(e) => setNewPassword(e.target.value)}
                placeholder="请输入新密码"
              />
            </div>
            <div className="form-control">
              <label className="label">
                <span className="label-text">确认新密码</span>
              </label>
              <input
                type="password"
                className="input input-bordered input-sm"
                value={newPassword2}
                onChange={(e) => setNewPassword2(e.target.value)}
                placeholder="请再次输入新密码"
              />
            </div>
            <button
              className="btn btn-primary btn-sm w-full"
              onClick={handleUpdatePassword}
              disabled={!newPassword || !newPassword2 || passwordLoading}
            >
              {passwordLoading ? (
                <span className="loading loading-spinner"></span>
              ) : (
                "更新密码"
              )}
            </button>
          </div>
        </section>

        {/* 用户偏好设置 */}
        <section>
          <h3 className="font-semibold text-lg mb-4 pb-2 border-b border-base-300">
            偏好设置
          </h3>
          <div className="space-y-4">
            {/* 开启每日回顾开关 */}
            <div className="form-control">
              <label className="label cursor-pointer justify-start gap-3">
                <input
                  type="checkbox"
                  className="checkbox checkbox-primary checkbox-sm"
                  checked={dailyReviewEnabled}
                  onChange={(e) => setDailyReviewEnabled(e.target.checked)}
                />
                <span className="label-text">开启每日回顾</span>
              </label>
              {dailyReviewEnabled && (
                <div className="text-xs text-base-content/50 mt-1 ml-9">
                  开启后会在指定时间随机发送一条 Memo 到邮箱
                </div>
              )}
            </div>

            {/* Telegram Bot 配置 */}
            <div
              className={`space-y-3 transition-all duration-300 ${
                dailyReviewEnabled
                  ? "opacity-100 max-h-96"
                  : "opacity-50 max-h-0 overflow-hidden"
              }`}
            >
              <div className="form-control">
                <label className="label">
                  <span className="label-text font-medium">
                    Telegram Chat ID
                  </span>
                </label>
                <input
                  type="text"
                  className="input input-bordered input-sm"
                  value={telegramChatId}
                  onChange={(e) => setTelegramChatId(e.target.value)}
                  placeholder="请输入您的 Telegram Chat ID"
                  disabled={!dailyReviewEnabled}
                />
                <div className="text-xs text-base-content/50 mt-1">
                  获取方法：向 @userinfobot 发送消息获取您的 Chat ID
                </div>
              </div>

              <div className="form-control">
                <label className="label">
                  <span className="label-text font-medium">
                    Telegram Bot Token
                  </span>
                </label>
                <input
                  type="password"
                  className="input input-bordered input-sm"
                  value={telegramBotToken}
                  onChange={(e) => setTelegramBotToken(e.target.value)}
                  placeholder="请输入您的 Telegram Bot Token"
                  disabled={!dailyReviewEnabled}
                />
                <div className="text-xs text-base-content/50 mt-1">
                  获取方法：通过 @BotFather 创建 Bot 获取 Token
                </div>
              </div>
            </div>

            <button
              className="btn btn-primary btn-sm w-full"
              onClick={handleUpdatePreferences}
              disabled={preferencesLoading}
            >
              {preferencesLoading ? (
                <span className="loading loading-spinner"></span>
              ) : (
                "保存偏好设置"
              )}
            </button>

            {/* 测试按钮 */}
            <div className="form-control">
              <div className="text-xs text-base-content/50 mb-2 text-center">
                💡 测试前请先保存设置以确保配置生效
              </div>
              <button
                className="btn btn-outline btn-sm w-full"
                onClick={handleTestDailyReview}
                disabled={
                  testing ||
                  !dailyReviewEnabled ||
                  !telegramChatId ||
                  !telegramBotToken
                }
              >
                {testing ? (
                  <span className="loading loading-spinner"></span>
                ) : (
                  "测试每日回顾"
                )}
              </button>
              <div className="text-xs text-base-content/50 mt-1">
                点击测试按钮立刻触发每日回顾
              </div>
            </div>
          </div>
        </section>

        {/* 插件接口 */}
        <section>
          <h3 className="font-semibold text-lg mb-4 pb-2 border-b border-base-300">
            插件接口
          </h3>
          <div className="space-y-4">
            {pluginToken ? (
              <div className="space-y-3">
                {/* 新增记录接口 */}
                <div className="bg-base-200 p-4 rounded-lg border border-base-300">
                  <div className="flex items-center justify-between mb-2">
                    <p className="text-sm font-medium text-base-content">
                      新增记录接口
                    </p>
                    <button
                      className="btn btn-xs btn-ghost text-base-content/60 hover:text-base-content"
                      onClick={() => {
                        navigator.clipboard.writeText(
                          `${api.defaults.baseURL}/plugin/createMemo/${pluginToken}`
                        );
                        showAlert({
                          message: "接口地址已复制到剪贴板",
                          type: "success",
                          duration: 2000,
                        });
                      }}
                      title="复制接口地址"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="h-4 w-4"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
                        />
                      </svg>
                      复制
                    </button>
                  </div>
                  <div className="bg-base-300/50 p-3 rounded">
                    <code className="text-xs break-all font-mono text-base-content">
                      {api.defaults.baseURL}/plugin/createMemo/{pluginToken}
                    </code>
                  </div>
                  <p className="text-xs text-base-content/60 mt-2">
                    POST 请求，用于通过插件创建新的 Memo 记录
                  </p>
                </div>

                {/* 随机获取接口 */}
                <div className="bg-base-200 p-4 rounded-lg border border-base-300">
                  <div className="flex items-center justify-between mb-2">
                    <p className="text-sm font-medium text-base-content">
                      随机获取接口
                    </p>
                    <button
                      className="btn btn-xs btn-ghost text-base-content/60 hover:text-base-content"
                      onClick={() => {
                        navigator.clipboard.writeText(
                          `${api.defaults.baseURL}/plugin/randomMemo/${pluginToken}`
                        );
                        showAlert({
                          message: "接口地址已复制到剪贴板",
                          type: "success",
                          duration: 2000,
                        });
                      }}
                      title="复制接口地址"
                    >
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className="h-4 w-4"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                      >
                        <path
                          strokeLinecap="round"
                          strokeLinejoin="round"
                          strokeWidth={2}
                          d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
                        />
                      </svg>
                      复制
                    </button>
                  </div>
                  <div className="bg-base-300/50 p-3 rounded">
                    <code className="text-xs break-all font-mono text-base-content">
                      {api.defaults.baseURL}/plugin/randomMemo/{pluginToken}
                    </code>
                  </div>
                  <p className="text-xs text-base-content/60 mt-2">
                    GET 请求，用于通过插件随机获取一条 Memo 记录
                  </p>
                </div>

                {/* 令牌信息 */}
                <div className="bg-base-200/50 p-3 rounded border border-base-300/50">
                  <div className="flex items-center justify-between">
                    <span className="text-xs text-base-content/60">
                      当前令牌:
                    </span>
                    <div className="flex items-center gap-2">
                      <code className="text-xs font-mono bg-base-300/30 px-2 py-1 rounded">
                        {pluginToken.substring(0, 8)}...
                      </code>
                      <button
                        className="btn btn-xs btn-ghost text-base-content/60 hover:text-base-content"
                        onClick={() => {
                          navigator.clipboard.writeText(pluginToken);
                          showAlert({
                            message: "插件令牌已复制到剪贴板",
                            type: "success",
                            duration: 2000,
                          });
                        }}
                        title="复制完整令牌"
                      >
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          className="h-3 w-3"
                          fill="none"
                          viewBox="0 0 24 24"
                          stroke="currentColor"
                        >
                          <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth={2}
                            d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
                          />
                        </svg>
                      </button>
                    </div>
                  </div>
                </div>
              </div>
            ) : (
              <div className="text-center py-6">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-12 w-12 mx-auto text-base-content/30 mb-3"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={1}
                    d="M15 7a2 2 0 012 2m4 0a6 6 0 01-7.743 5.743L11 17H9v2H7v2H4a1 1 0 01-1-1v-2.586a1 1 0 01.293-.707l5.964-5.964A6 6 0 1121 9z"
                  />
                </svg>
                <p className="text-sm text-base-content/60 mb-4">
                  尚未生成插件令牌，点击下方按钮创建
                </p>
              </div>
            )}

            <button
              className="btn btn-outline btn-sm w-full gap-2"
              onClick={handleGenerateToken}
              disabled={tokenLoading}
            >
              {tokenLoading ? (
                <span className="loading loading-spinner"></span>
              ) : (
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-4 w-4"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                  />
                </svg>
              )}
              {pluginToken ? "重新生成令牌" : "生成插件令牌"}
            </button>
          </div>
        </section>
      </div>

      <Alert
        isOpen={alertOpen}
        onClose={hideAlert}
        message={alertConfig.message}
        type={alertConfig.type}
        duration={alertConfig.duration}
      />
    </>
  );
};

export default UserSettingsCard;
