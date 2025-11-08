import { useEffect, useState } from "react";
import healthApi from "../api/healthApi";

const ServerStatusCheck = () => {
  const [serverStatus, setServerStatus] = useState("检查中");
  const [serverVersion, setServerVersion] = useState("");

  useEffect(() => {
    healthApi
      .health()
      .then((response) => {
        console.log("health response", response);
        const success = response.data.success;
        if (success) {
          setServerStatus("在线");
          setServerVersion(response.data.data.commit);
        } else {
          setServerStatus("离线");
        }
      })
      .catch((e) => {
        console.error("health error", e);
        setServerStatus("离线");
      });
  }, []);

  return (
    <>
      <p>服务器状态：{serverStatus}</p>
      <p>
        服务器版本：
        <a
          target="_blank"
          href={`https://github.com/jerryshell/my-flomo-server/tree/${serverVersion}`}
        >
          {serverVersion}
        </a>
      </p>
    </>
  );
};

export default ServerStatusCheck;
