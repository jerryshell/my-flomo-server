import logo from "../favicon.svg";
import ServerStatusCheck from "./ServerStatusCheck";

const Header = () => {
  return (
    <>
      <h1>
        <img src={logo} alt="logo" />
        <span> </span>
        <span>My Flomo</span>
      </h1>
      <ServerStatusCheck />
    </>
  );
};

export default Header;
