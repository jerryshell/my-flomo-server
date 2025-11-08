import { useState } from "react";
import ServerStatus from "./ServerStatus";
import ServerDetails from "./ServerDetails";
import { useHealthCheck } from "../hooks/useHealthCheck";

const ServerStatusCheck = () => {
  const { serverStatus, healthData } = useHealthCheck();
  const [showDetails, setShowDetails] = useState(false);

  return (
    <div>
      <ServerStatus
        status={serverStatus}
        healthData={healthData}
        onToggleDetails={() => setShowDetails(!showDetails)}
        showDetails={showDetails}
      />
      {showDetails && healthData && <ServerDetails healthData={healthData} />}
    </div>
  );
};

export default ServerStatusCheck;
