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

    // 获取用户设置（暂时使用默认值）
    setDailyReviewEnabled(false);
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
    setTimeout(() => {
      showAlert({
        message: "偏好设置已保存！",
        type: "success",
        duration: 2000,
      });
      setPreferencesLoading(false);
    }, 1000);
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
          <div className="space-y-3">
            <div className="form-control">
              <label className="label cursor-pointer justify-start gap-3">
                <input
                  type="checkbox"
                  className="checkbox checkbox-primary checkbox-sm"
                  checked={dailyReviewEnabled}
                  onChange={(e) => setDailyReviewEnabled(e.target.checked)}
                />
                <span className="label-text">开启每日回顾邮件</span>
              </label>
              <div className="text-xs text-base-content/50 mt-1 ml-9">
                每日回顾功能会在指定时间随机发送一条 Memo 到邮箱
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
          </div>
        </section>

        {/* 插件接口 */}
        <section>
          <h3 className="font-semibold text-lg mb-4 pb-2 border-b border-base-300">
            插件接口
          </h3>
          <div className="space-y-3">
            {pluginToken ? (
              <div className="space-y-2">
                <div className="bg-base-200 p-3 rounded">
                  <p className="text-sm font-medium mb-1">新增记录:</p>
                  <code className="text-xs break-all bg-base-300 p-1 rounded">
                    {api.defaults.baseURL}/plugin/createMemo/{pluginToken}
                  </code>
                </div>
                <div className="bg-base-200 p-3 rounded">
                  <p className="text-sm font-medium mb-1">随机获取:</p>
                  <code className="text-xs break-all bg-base-300 p-1 rounded">
                    {api.defaults.baseURL}/plugin/randomMemo/{pluginToken}
                  </code>
                </div>
              </div>
            ) : (
              <p className="text-sm text-base-content/60">
                点击下方按钮生成插件令牌
              </p>
            )}
            <button
              className="btn btn-outline btn-sm w-full"
              onClick={handleGenerateToken}
              disabled={tokenLoading}
            >
              {tokenLoading ? (
                <span className="loading loading-spinner"></span>
              ) : pluginToken ? (
                "重新生成令牌"
              ) : (
                "生成令牌"
              )}
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
