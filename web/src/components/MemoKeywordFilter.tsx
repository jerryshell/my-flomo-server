import { useRecoilState } from "recoil";
import { atoms } from "../atoms/atoms";

const MemoKeywordFilter = () => {
  const [memoKeyword, setMemoKeyword] = useRecoilState(atoms.memoKeyword);

  return (
    <div className="card bg-base-100 shadow-md mb-4">
      <div className="card-body p-4">
        <h3 className="card-title text-lg mb-2">关键字模糊筛选</h3>
        <div className="form-control">
          <div className="input-group">
            <input
              type="text"
              placeholder="多个关键字使用空格分割"
              className="input input-bordered flex-1"
              value={memoKeyword}
              onChange={(e) => setMemoKeyword(e.target.value)}
            />
            <button
              className="btn btn-outline"
              onClick={() => setMemoKeyword("")}
              disabled={!memoKeyword}
            >
              重置
            </button>
          </div>
        </div>
        {memoKeyword && (
          <div className="text-sm text-gray-500 mt-2">
            当前筛选关键字: {memoKeyword}
          </div>
        )}
      </div>
    </div>
  );
};

export default MemoKeywordFilter;
