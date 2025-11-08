import api from "../api/api";

const CsvExport = () => {
  const handleCsvExportBtnClick = () => {
    const token = localStorage.getItem("token");
    window.open(`${api.defaults.baseURL}/csvExport/token/${token}`);
  };

  return (
    <details>
      <summary>CSV 导出</summary>
      <button onClick={handleCsvExportBtnClick}>CSV 导出</button>
    </details>
  );
};

export default CsvExport;
