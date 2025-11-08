import { useState } from "react";
import csvApi from "../api/csvApi";

const CsvImport = (props: { fetchMemoList: () => void }) => {
  const [csvFile, setCsvFile] = useState<File | null>(null);

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

    const formData = new FormData();
    formData.append("csvFile", csvFile);

    csvApi
      .csvImport(formData)
      .then((response) => {
        const success = response.data.success;
        if (success) {
          props.fetchMemoList();
        } else {
          alert(response.data.message);
        }
      })
      .catch((e) => {
        console.log(e);
      })
      .finally(() => {
        setCsvFile(null);
      });
  };

  return (
    <details>
      <summary>CSV 导入</summary>
      <input
        type="file"
        name="file"
        accept="text/csv"
        onChange={(e) => {
          const files = e.target.files;
          if (files != null && files.length > 0) {
            handleCsvFileInputChange(files);
          }
        }}
      />
      <button onClick={handleCsvImportBtnClick}>提交</button>
    </details>
  );
};

export default CsvImport;
