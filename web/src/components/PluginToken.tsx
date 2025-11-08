import { useEffect, useState } from "react";
import pluginApi from "../api/pluginApi";
import api from "../api/api";

const PluginToken = () => {
  const [pluginToken, setPluginToken] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [isLoading, setIsLoading] = useState(false);

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
    setIsLoading(true);
    pluginApi
      .createToken()
      .then((response) => {
        const success = response.data.success;
        if (success) {
          setPluginToken(response.data.data);
          setErrorMessage("");
        } else {
          setPluginToken("");
          setErrorMessage(response.data.message);
        }
      })
      .catch((error) => {
        console.error("createToken error", error);
        setErrorMessage("生成令牌失败，请稍后重试");
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  return (
    <div className="space-y-3">
      <h4 className="font-semibold">插件接口</h4>

      {pluginToken ? (
        <div className="space-y-2">
          <div className="bg-base-200 p-3 rounded">
            <p className="text-sm font-medium">新增记录:</p>
            <code className="text-xs break-all">
              {api.defaults.baseURL}/plugin/createMemo/{pluginToken}
            </code>
          </div>
          <div className="bg-base-200 p-3 rounded">
            <p className="text-sm font-medium">随机获取:</p>
            <code className="text-xs break-all">
              {api.defaults.baseURL}/plugin/randomMemo/{pluginToken}
            </code>
          </div>
        </div>
      ) : (
        <p className="text-error text-sm">{errorMessage}</p>
      )}

      <button
        className="btn btn-outline btn-sm w-full"
        onClick={handleGeneratePluginTokenBtnClick}
        disabled={isLoading}
      >
        {isLoading ? (
          <span className="loading loading-spinner"></span>
        ) : (
          "重新生成令牌"
        )}
      </button>
    </div>
  );
};

export default PluginToken;
