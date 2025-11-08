import MemoListItem from "./MemoListItem";
import { useRecoilValue } from "recoil";
import { atoms } from "../atoms/atoms";
import MemoKeywordFilter from "./MemoKeywordFilter";
import { useMemo } from "react";

const MemoList = () => {
  const memoList = useRecoilValue(atoms.memoList);
  const memoKeyword = useRecoilValue(atoms.memoKeyword);
  const memoShowList = useMemo(() => {
    const memoKeywordSplit = memoKeyword
      .split(" ")
      .filter((item) => item.length > 0);
    if (memoKeywordSplit.length <= 0) {
      return memoList;
    }
    return memoList.filter((item) => {
      return Object.values(item).some((value) => {
        return memoKeywordSplit.some(
          (keyword) => value && value.includes(keyword)
        );
      });
    });
  }, [memoList, memoKeyword]);

  return (
    <>
      <MemoKeywordFilter />
      {memoShowList.map((memo) => (
        <MemoListItem memo={memo} key={memo.id} />
      ))}
    </>
  );
};

export default MemoList;
