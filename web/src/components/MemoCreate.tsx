import { useState } from "react";
import memoApi from "../api/memoApi";

const MemoCreate = (props: { fetchMemoList(): void }) => {
  const [newMemo, setNewMemo] = useState("");

  const handleSaveBtnClick = () => {
    memoApi
      .create({
        content: newMemo,
      })
      .then((response) => {
        const success = response.data.success;
        if (success) {
          props.fetchMemoList();
        }
      })
      .finally(() => {
        setNewMemo("");
      });
  };

  return (
    <>
      <textarea
        placeholder="开始记录你的想法..."
        value={newMemo}
        onChange={(e) => setNewMemo(e.target.value)}
      />
      <button onClick={handleSaveBtnClick}>保存</button>
    </>
  );
};

export default MemoCreate;
