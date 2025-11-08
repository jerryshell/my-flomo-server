import { useState } from "react";
import memoApi from "../api/memoApi";
import Memo from "../interfaces/Memo";
import { useRecoilState } from "recoil";
import { atoms } from "../atoms/atoms";
import dayjs from "dayjs";
import { useModal } from "../hooks/useModal";
import Modal from "./Modal";

const MemoListItem = (props: { memo: Memo }) => {
  const [editModeFlag, setEditModeFlag] = useState(false);
  const [memo, setMemo] = useState({ ...props.memo });
  const [memoList, setMemoList] = useRecoilState(atoms.memoList);
  const [isLoading, setIsLoading] = useState(false);
  const { modalOpen, modalConfig, showModal, hideModal } = useModal();

  const handleTextareaChange = (e: { target: { value: string } }) => {
    const content = e.target.value;
    const newMemo = { ...memo, content };
    setMemo(newMemo);
  };

  const handleUpdateBtnClick = () => {
    setIsLoading(true);
    setEditModeFlag(false);
    memoApi
      .update(memo)
      .then((response) => {
        const success = response.data.success;
        if (success) {
          const memo = response.data.data;
          setMemoList(
            memoList.map((item) => (item.id === memo.id ? memo : item))
          );
        }
      })
      .finally(() => {
        setIsLoading(false);
      });
  };

  const handleMemoDeleteBtnClick = (id: string) => {
    showModal({
      title: "确认删除",
      message: "确定要删除这条备忘录吗？",
      type: "warning",
      confirmText: "删除",
      cancelText: "取消",
      onConfirm: () => {
        setIsLoading(true);
        setMemoList(memoList.filter((item) => item.id !== id));
        memoApi
          .deleteById(id)
          .then((response) => {
            console.log("delete memo response", response);
          })
          .finally(() => {
            setIsLoading(false);
          });
      },
    });
  };

  const handleCancelBtnClick = () => {
    setEditModeFlag(false);
    setMemo({ ...props.memo });
  };

  return (
    <>
      <div className="card bg-base-100 shadow-md mb-4">
        <div className="card-body p-4">
          <div className="flex justify-between items-center mb-2">
            <h3 className="card-title text-sm font-normal text-base-content/60">
              {dayjs(memo.createdAt).format("YYYY-MM-DD HH:mm:ss")}
            </h3>

            <div className="flex space-x-2">
              {editModeFlag ? (
                <>
                  <button
                    className="btn btn-sm btn-success"
                    onClick={handleUpdateBtnClick}
                    disabled={isLoading}
                  >
                    {isLoading ? (
                      <span className="loading loading-spinner"></span>
                    ) : (
                      "更新"
                    )}
                  </button>
                  <button
                    className="btn btn-sm btn-outline"
                    onClick={handleCancelBtnClick}
                    disabled={isLoading}
                  >
                    取消
                  </button>
                </>
              ) : (
                <button
                  className="btn btn-sm btn-outline"
                  onClick={() => setEditModeFlag(true)}
                >
                  编辑
                </button>
              )}

              <button
                className="btn btn-sm btn-error"
                onClick={() => handleMemoDeleteBtnClick(memo.id)}
                disabled={isLoading}
              >
                {isLoading ? (
                  <span className="loading loading-spinner"></span>
                ) : (
                  "删除"
                )}
              </button>
            </div>
          </div>

          <div className="mt-2">
            {editModeFlag ? (
              <textarea
                className="textarea textarea-bordered w-full h-32"
                value={memo.content}
                onChange={handleTextareaChange}
                placeholder="编辑备忘录内容..."
              />
            ) : (
              <p className="whitespace-pre-wrap text-base">{memo.content}</p>
            )}
          </div>
        </div>
      </div>

      <Modal
        isOpen={modalOpen}
        onClose={hideModal}
        title={modalConfig.title}
        message={modalConfig.message}
        type={modalConfig.type}
        confirmText={modalConfig.confirmText}
        cancelText={modalConfig.cancelText}
        onConfirm={modalConfig.onConfirm}
        showCancel={modalConfig.showCancel}
      />
    </>
  );
};

export default MemoListItem;
