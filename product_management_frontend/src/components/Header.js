import React from "react";
import { useNavigate } from "react-router-dom";
import "../css/header.css";

const Header = ({ setToken }) => {
  const navigate = useNavigate();
  function handleLogout() {
    setToken(localStorage.clear());
    navigate("/");
  }
  return (
    <>
      <div className="header">
        <button
          className="logout"
          onClick={() => {
            handleLogout();
          }}
        >
          Log Out
        </button>
      </div>
    </>
  );
};

export default Header;
