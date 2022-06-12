import React, { useState } from "react";
import "../css/signin.css";
import axios from "axios";
import { useNavigate, Link } from "react-router-dom";

const url = "https://product-management-go.herokuapp.com/login";

const Signin = ({ setToken }) => {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  async function handleSubmit() {
    const data = { email, password };
    try {
      const response = await axios.post(url, data);
      console.log(response);
      const token = response.data.token;
      console.log(token);
      setToken(token);
      localStorage.setItem("token", token);
      navigate("/allProducts");
    } catch (error) {
      alert(error.response.data.message);
    }
  }

  return (
    <>
      <>
        <div className="signin-parent">
          <div className="signin-form-container">
            <div className="signin-heading">
              <h4>Login here</h4>
            </div>

            <input
              type="email"
              className="signin-email"
              placeholder="Email ID"
              name="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <input
              type="password"
              className="signin-password"
              placeholder="Password"
              name="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />

            <button
              type="button"
              className="signin-submit"
              onClick={() => {
                handleSubmit();
              }}
            >
              Sign in
            </button>
          </div>
          <div className="signup-option">
            Not a member?<Link to={"/signup"}>Register</Link>
          </div>
        </div>
      </>
    </>
  );
};

export default Signin;
