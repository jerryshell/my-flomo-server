import { KeyboardEvent, useState } from "react";
import authApi from "../api/authApi";
import LoginResponse from "../interfaces/LoginResponse";
import { useRecoilState, useSetRecoilState } from "recoil";
import { atoms } from "../atoms/atoms";
import { validateEmail, getEmailValidationMessage } from "../utils/emailValidator";

const Logging = () => (
  <button disabled>
    <svg width="1em" height="1em" viewBox="0 0 24 24">
      <path
        fill="currentColor"
        d="M12 2A10 10 0 1 0 22 12A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8A8 8 0 0 1 12 20Z"
        opacity=".5"
      />
      <path
        fill="currentColor"
        d="M20 12h2A10 10 0 0 0 12 2V4A8 8 0 0 1 20 12Z"
      >
        <animateTransform
          attributeName="transform"
          dur="1s"
          from="0 12 12"
          repeatCount="indefinite"
          to="360 12 12"
          type="rotate"
        />
      </path>
    </svg>
  </button>
);

const LoginPage = () => {
  const [email, setEmail] = useRecoilState(atoms.email);
  const setToken = useSetRecoilState(atoms.token);

  const [password, setPassword] = useState("");
  const [logging, setLogging] = useState(false);
  const [emailError, setEmailError] = useState("");

  const handleLoginClick = () => {
    // 验证email格式
    const emailValidationMessage = getEmailValidationMessage(email);
    if (emailValidationMessage) {
      setEmailError(emailValidationMessage);
      return;
    }
    
    if (email.length <= 0 || password.length <= 0) {
      alert("邮箱和密码不能为空");
      return;
    }
    
    const postData = {
      email,
      password,
    };
    setLogging(true);
    authApi
      .login(postData)
      .then((response) => {
        console.log("login response", response);
        const success = response.data.success;
        if (!success) {
          alert(response.data.message);
          return;
        }
        const data = response.data.data;
        handleLoginSuccess(data);
      })
      .catch((error) => {
        console.error("login error", error);
        alert("登录失败");
      })
      .finally(() => {
        setLogging(false);
      });
  };

  const handleLoginSuccess = (loginResponse: LoginResponse) => {
    setEmail(loginResponse.email);
    setToken(loginResponse.token);
    localStorage.setItem("email", loginResponse.email);
    localStorage.setItem("token", loginResponse.token);
    localStorage.setItem("expiresAt", loginResponse.expiresAt);
  };

  const handleEmailChange = (value: string) => {
    setEmail(value);
    // 实时验证email格式
    if (value.trim().length > 0) {
      const validationMessage = getEmailValidationMessage(value);
      setEmailError(validationMessage);
    } else {
      setEmailError("");
    }
  };

  const handleKeyUp = (e: KeyboardEvent) => {
    if (e.key === "Enter") {
      handleLoginClick();
    }
  };

  return (
    <fieldset onKeyUp={handleKeyUp}>
      <legend>不存在的账号将自动注册</legend>
      <input
        type="email"
        placeholder="邮箱"
        value={email}
        onChange={(e) => handleEmailChange(e.target.value)}
        style={emailError ? { borderColor: "red" } : {}}
      />
      {emailError && (
        <div style={{ color: "red", fontSize: "12px", marginTop: "5px" }}>
          {emailError}
        </div>
      )}
      <input
        type="password"
        placeholder="密码"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      {logging ? <Logging /> : <button onClick={handleLoginClick}>登录</button>}
    </fieldset>
  );
};

export default LoginPage;
