import { useState } from "react";
import userApi from "../api/userApi";
import { useAlert } from "../hooks/useModal";
import Alert from "./Alert";

const UserPasswordUpdate = () => {
  const [newPassword, setNewPassword] = useState("");
  const [newPassword2, setNewPassword2] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const { alertOpen, alertConfig, showAlert, hideAlert } = useAlert();

  const handleUpdatePasswordBtnClick = () => {
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

    setIsLoading(true);
    userApi
      .updatePassword({
        password: newPassword,
      })
      .then((response) => {
        console.log("updatePassword response", response);
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
        setIsLoading(false);
      });
  };

  return (
    <>
      <div className="space-y-3">
        <h4 className="font-semibold">更新密码</h4>

        <div className="form-control">
          <label className="label">
            <span className="label-text">新密码</span>
          </label>
          <input
            type="password"
            className="input input-bordered"
            onChange={(e) => setNewPassword(e.target.value)}
            value={newPassword}
            placeholder="请输入新密码"
          />
        </div>

        <div className="form-control">
          <label className="label">
            <span className="label-text">确认新密码</span>
          </label>
          <input
            type="password"
            className="input input-bordered"
            onChange={(e) => setNewPassword2(e.target.value)}
            value={newPassword2}
            placeholder="请再次输入新密码"
          />
        </div>

        <button
          className="btn btn-primary btn-sm w-full"
          onClick={handleUpdatePasswordBtnClick}
          disabled={!newPassword || !newPassword2 || isLoading}
        >
          {isLoading ? (
            <span className="loading loading-spinner"></span>
          ) : (
            "更新密码"
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

export default UserPasswordUpdate;
