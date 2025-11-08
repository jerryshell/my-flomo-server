import { useRecoilState } from "recoil";
import { atoms } from "../atoms/atoms";

const MemoKeywordFilter = () => {
  const [memoKeyword, setMemoKeyword] = useRecoilState(atoms.memoKeyword);

  return (
    <details open>
      <summary>关键字模糊筛选</summary>
      <div style={{ display: "flex", flexWrap: "wrap" }}>
        <input
          placeholder="多个关键字使用空格分割"
          type="text"
          value={memoKeyword}
          onChange={(e) => setMemoKeyword(e.target.value)}
        />
        <button onClick={() => setMemoKeyword("")}>重置</button>
      </div>
    </details>
  );
};

export default MemoKeywordFilter;
