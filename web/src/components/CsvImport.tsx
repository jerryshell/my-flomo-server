import { useState } from "react";
import csvApi from "../api/csvApi";

const CsvImport = (props: { fetchMemoList: () => void }) => {
  const [csvFile, setCsvFile] = useState<File | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const handleCsvFileInputChange = (fileList: FileList | null) => {
    if (fileList && fileList.length > 0) {
      setCsvFile(fileList[0]);
    } else {
      setCsvFile(null);
    }
  };

  const handleCsvImportBtnClick = () => {
    if (csvFile === null) {
      return;
    }

    setIsLoading(true);
    const formData = new FormData();
    formData.append("csvFile", csvFile);

    csvApi
      .csvImport(formData)
      .then((response) => {
        const success = response.data.success;
        if (success) {
          props.fetchMemoList();
          alert("CSV导入成功！");
        } else {
          alert(response.data.message);
        }
      })
      .catch((e) => {
        console.log(e);
        alert("CSV导入失败，请检查文件格式");
      })
      .finally(() => {
        setCsvFile(null);
        setIsLoading(false);
      });
  };

  return (
    <div className="space-y-2">
      <h4 className="font-semibold">CSV 导入</h4>
      <p className="text-sm text-gray-600">从 CSV 文件导入备忘录数据</p>

      <div className="form-control">
        <input
          type="file"
          name="file"
          accept="text/csv"
          className="file-input file-input-bordered w-full"
          onChange={(e) => {
            const files = e.target.files;
            if (files != null && files.length > 0) {
              handleCsvFileInputChange(files);
            }
          }}
        />
      </div>

      {csvFile && (
        <div className="text-sm text-success">已选择文件: {csvFile.name}</div>
      )}

      <button
        className="btn btn-primary btn-sm"
        onClick={handleCsvImportBtnClick}
        disabled={!csvFile || isLoading}
      >
        {isLoading ? (
          <span className="loading loading-spinner"></span>
        ) : (
          "提交导入"
        )}
      </button>
    </div>
  );
};

export default CsvImport;
