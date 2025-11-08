import { useState, useEffect } from "react";
import userApi from "../api/userApi";
import { useAlert } from "../hooks/useModal";
import Alert from "./Alert";

const UserSettings = () => {
  const [dailyReviewEnabled, setDailyReviewEnabled] = useState(false);
  const [loading, setLoading] = useState(false);
  const { alertOpen, alertConfig, showAlert, hideAlert } = useAlert();

  // 这里需要从API获取当前用户的设置，但由于目前没有获取设置的API，
  // 我们先使用默认值，后续可以添加获取设置的API
  useEffect(() => {
    // 暂时使用默认值，后续可以添加获取用户设置的API
    setDailyReviewEnabled(false);
  }, []);

  const handleUpdateSettings = () => {
    setLoading(true);

    // 这里需要调用更新用户设置的API，但由于目前没有更新设置的API，
    // 我们先模拟一个API调用，后续可以添加更新用户设置的API
    setTimeout(() => {
      showAlert({
        message: "设置已保存！",
        type: "success",
        duration: 2000,
      });
      setLoading(false);
    }, 1000);
  };

  return (
    <>
      <div className="space-y-3">
        <h4 className="font-semibold">用户设置</h4>

        <div className="form-control">
          <label className="label cursor-pointer justify-start gap-3">
            <input
              type="checkbox"
              className="checkbox checkbox-primary"
              checked={dailyReviewEnabled}
              onChange={(e) => setDailyReviewEnabled(e.target.checked)}
            />
            <span className="label-text">开启每日回顾邮件</span>
          </label>
          <div className="text-xs text-base-content/50 mt-1">
            每日回顾功能会在指定时间随机发送一条 Memo 到邮箱
          </div>
        </div>

        <button
          className="btn btn-primary btn-sm"
          onClick={handleUpdateSettings}
          disabled={loading}
        >
          {loading ? (
            <span className="loading loading-spinner"></span>
          ) : (
            "保存设置"
          )}
        </button>
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

export default UserSettings;
