import { useEffect, useState } from "react";
import pluginApi from "../api/pluginApi";
import api from "../api/api";

const PluginToken = () => {
  const [pluginToken, setPluginToken] = useState("");
  const [errorMessage, setErrorMessage] = useState("");

  const fetchPluginToken = () => {
    pluginApi.getToken().then((response) => {
      const success = response.data.success;
      if (success) {
        setPluginToken(response.data.data);
      } else {
        setPluginToken("");
        setErrorMessage(response.data.message);
      }
    });
  };

  useEffect(() => {
    fetchPluginToken();
  }, []);

  const handleGeneratePluginTokenBtnClick = () => {
    pluginApi.createToken().then((response) => {
      const success = response.data.success;
      if (success) {
        setPluginToken(response.data.data);
        setErrorMessage("");
      } else {
        setPluginToken("");
        setErrorMessage(response.data.message);
      }
    }).catch((error) => {
      console.error("createToken error", error);
      setErrorMessage("生成令牌失败，请稍后重试");
    });
  };

  return (
    <details>
      <summary>插件接口</summary>
      {pluginToken ? (
        <div>
          <p>
            新增记录: {api.defaults.baseURL}/plugin/createMemo/{pluginToken}
          </p>
          <p>
            随机获取: {api.defaults.baseURL}/plugin/randomMemo/{pluginToken}
          </p>
        </div>
      ) : (
        <p>{errorMessage}</p>
      )}
      <button onClick={handleGeneratePluginTokenBtnClick}>重新生成</button>
    </details>
  );
};

export default PluginToken;
