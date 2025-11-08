import { useState } from "react";
import memoApi from "../api/memoApi";
import Memo from "../interfaces/Memo";
import { useRecoilState } from "recoil";
import { atoms } from "../atoms/atoms";
import dayjs from "dayjs";

const MemoListItem = (props: { memo: Memo }) => {
  const [editModeFlag, setEditModeFlag] = useState(false);
  const [memo, setMemo] = useState({ ...props.memo });
  const [memoList, setMemoList] = useRecoilState(atoms.memoList);

  const handleTextareaChange = (e: { target: { value: string } }) => {
    const content = e.target.value;
    const newMemo = { ...memo, content };
    setMemo(newMemo);
  };

  const handleUpdateBtnClick = () => {
    setEditModeFlag(false);
    memoApi.update(memo).then((response) => {
      const success = response.data.success;
      if (success) {
        const memo = response.data.data;
        setMemoList(
          memoList.map((item) => (item.id === memo.id ? memo : item))
        );
      }
    });
  };

  const handleMemoDeleteBtnClick = (id: string) => {
    setMemoList(memoList.filter((item) => item.id !== id));
    memoApi.deleteById(id).then((response) => {
      console.log("delete memo response", response);
    });
  };

  const handleCancelBtnClick = () => {
    setEditModeFlag(false);
    setMemo({ ...props.memo });
  };

  return (
    <details open key={memo.id}>
      <summary>{dayjs(memo.createdAt).format("YYYY-MM-DD HH:mm:ss")}</summary>
      <p style={{ whiteSpace: "pre-wrap" }}>
        {editModeFlag ? (
          <textarea value={memo.content} onChange={handleTextareaChange} />
        ) : (
          memo.content
        )}
      </p>
      <p>
        {editModeFlag ? (
          <>
            <button style={{ float: "right" }} onClick={handleUpdateBtnClick}>
              更新
            </button>
            <button style={{ float: "right" }} onClick={handleCancelBtnClick}>
              取消
            </button>
          </>
        ) : (
          <button
            style={{ float: "right" }}
            onClick={() => setEditModeFlag(true)}
          >
            编辑
          </button>
        )}

        <button
          style={{ color: "#9E3B37", float: "right" }}
          onClick={() => handleMemoDeleteBtnClick(memo.id)}
        >
          删除
        </button>
      </p>
    </details>
  );
};

export default MemoListItem;
