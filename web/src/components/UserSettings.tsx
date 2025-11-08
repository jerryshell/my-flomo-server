import { useState, useEffect } from "react";
import userApi from "../api/userApi";

const UserSettings = () => {
  const [dailyReviewEnabled, setDailyReviewEnabled] = useState(false);
  const [loading, setLoading] = useState(false);

  // 这里需要从API获取当前用户的设置，但由于目前没有获取设置的API，
  // 我们先使用默认值，后续可以添加获取设置的API
  useEffect(() => {
    // 暂时使用默认值，后续可以添加获取用户设置的API
    setDailyReviewEnabled(false);
  }, []);

  const handleUpdateSettings = () => {
    setLoading(true);
    userApi
      .updateSettings({
        dailyReviewEnabled: dailyReviewEnabled,
      })
      .then((response) => {
        console.log("updateSettings response", response);
        if (response.data.success) {
          alert("设置更新成功");
        } else {
          alert(response.data.message);
        }
      })
      .catch((error) => {
        console.error("更新设置失败", error);
        alert("更新设置失败，请重试");
      })
      .finally(() => {
        setLoading(false);
      });
  };

  return (
    <details>
      <summary>用户设置</summary>
      <fieldset>
        <div style={{ marginBottom: "15px" }}>
          <label style={{ display: "flex", alignItems: "center", gap: "10px" }}>
            <input
              type="checkbox"
              checked={dailyReviewEnabled}
              onChange={(e) => setDailyReviewEnabled(e.target.checked)}
            />
            <span>开启每日回顾邮件</span>
          </label>
          <div style={{ fontSize: "12px", color: "#666", marginTop: "5px" }}>
            每日回顾功能会在指定时间随机发送一条 Memo 到邮箱
          </div>
        </div>
      </fieldset>
      <button
        onClick={handleUpdateSettings}
        disabled={loading}
        style={{ opacity: loading ? 0.6 : 1 }}
      >
        {loading ? "更新中..." : "保存设置"}
      </button>
    </details>
  );
};

export default UserSettings;
