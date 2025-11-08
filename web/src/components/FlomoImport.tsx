import { useState } from "react";
import uploadApi from "../api/uploadApi";

const FlomoImport = (props: { fetchMemoList: () => void }) => {
  const [uploadFileList, setUploadFileList] = useState<FileList | null>(null);

  const handleFileInputChange = (fileList: FileList | null) => {
    setUploadFileList(fileList);
  };

  const handleImportDataBtnClick = () => {
    if (uploadFileList === null) {
      return;
    }

    const formData = new FormData();
    for (let i = 0; i < uploadFileList.length; i++) {
      formData.append("uploadFileList[]", uploadFileList[i]);
    }

    uploadApi
      .upload(formData)
      .then(() => {
        props.fetchMemoList();
      })
      .finally(() => {
        setUploadFileList(null);
      });
  };

  return (
    <details>
      <summary>从 Flomo 导入</summary>
      <p>
        请选择从{" "}
        <a href="https://flomoapp.com/mine?source=account" target="_blank">
          Flomo
        </a>{" "}
        导出的 HTML 文件，可以一次性选择多个
      </p>
      <input
        type="file"
        name="uploadFileList"
        accept="text/html"
        multiple
        onChange={(e) => {
          handleFileInputChange(e.target.files);
        }}
      />
      <button onClick={handleImportDataBtnClick}>提交</button>
    </details>
  );
};

export default FlomoImport;
