import deleteMyAccountApi from "../api/deleteMyAccountApi";

const DangerousArea = (props: { logout: () => void }) => {
  const deleteMyAccount = () => {
    deleteMyAccountApi.deleteMyAccount().then((response) => {
      const success = response.data.success;
      if (success) {
        alert(response.data.message);
        props.logout();
      }
    });
  };

  return (
    <details>
      <summary>️⚠️ 危险区 ⚠️</summary>
      <button onClick={deleteMyAccount}>
        ⚠️ 账号注销，永久抹除数据，无法恢复，点击立刻生效 ⚠️
      </button>
    </details>
  );
};

export default DangerousArea;
