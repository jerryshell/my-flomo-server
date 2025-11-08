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

  const getStatusColor = () => {
    switch (serverStatus) {
      case "在线":
        return "text-success";
      case "离线":
        return "text-error";
      default:
        return "text-warning";
    }
  };

  return (
    <div className="flex flex-col gap-1 text-sm">
      <div className="flex items-center gap-2">
        <span className="font-medium">服务器状态:</span>
        <span className={`badge badge-sm ${getStatusColor()}`}>
          {serverStatus}
        </span>
      </div>

      <div className="flex items-center gap-2">
        <span className="font-medium">服务器版本:</span>
        <a
          target="_blank"
          rel="noopener noreferrer"
          href={`https://github.com/jerryshell/my-flomo-server/tree/${serverVersion}`}
          className="link link-primary text-xs"
        >
          {serverVersion || "未知"}
        </a>
      </div>
    </div>
  );
};

export default ServerStatusCheck;
