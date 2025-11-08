import { useState } from "react";
import userApi from "../api/userApi";

const UserPasswordUpdate = () => {
  const [newPassword, setNewPassword] = useState("");
  const [newPassword2, setNewPassword2] = useState("");

  const handleUpdatePasswordBtnClick = () => {
    if (newPassword !== newPassword2) {
      alert("两次密码不一致，请重新输入");
      setNewPassword("");
      setNewPassword2("");
      return;
    }
    userApi
      .updatePassword({
        password: newPassword,
      })
      .then((response) => {
        console.log("updatePassword response", response);
        if (response.data.success) {
          alert("密码更新成功");
        } else {
          alert(response.data.message);
        }
      })
      .finally(() => {
        setNewPassword("");
        setNewPassword2("");
      });
  };

  return (
    <details>
      <summary>更新密码</summary>
      <fieldset>
        <input
          type="password"
          onChange={(e) => setNewPassword(e.target.value)}
          value={newPassword}
          placeholder="请输入新密码"
        />
        <input
          type="password"
          onChange={(e) => setNewPassword2(e.target.value)}
          value={newPassword2}
          placeholder="请再次输入新密码"
        />
      </fieldset>
      <button onClick={handleUpdatePasswordBtnClick}>提交</button>
    </details>
  );
};

export default UserPasswordUpdate;
