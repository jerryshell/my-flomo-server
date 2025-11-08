import api from "../api/api";

const CsvExport = () => {
  const handleCsvExportBtnClick = () => {
    const token = localStorage.getItem("token");
    window.open(`${api.defaults.baseURL}/csvExport/token/${token}`);
  };

  return (
    <div className="space-y-2">
      <h4 className="font-semibold">CSV 导出</h4>
      <p className="text-sm text-gray-600">导出所有备忘录数据为 CSV 格式文件</p>
      <button
        className="btn btn-outline btn-sm"
        onClick={handleCsvExportBtnClick}
      >
        CSV 导出
      </button>
    </div>
  );
};

export default CsvExport;
