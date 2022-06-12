import axios from "axios";
import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import "../css/signup.css";

const url = "https://product-management-go.herokuapp.com/register";

const Signup = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  async function handleSubmit(e) {
    e.preventDefault();
    if (password !== confirmPassword) {
      alert("passwords do not match!!");
      return;
    }
    const data = { email, name, password };

    try {
      const response = await axios.post(url, data);
      navigate("/");
    } catch (error) {
      alert(error.response.data.message);
    }
  }
  return (
    <>
      <div className="signup-parent">
        <div className="signup-form-container">
          <div className="signup-heading">
            <h4>Register here</h4>
          </div>
          <form
            onSubmit={(e) => {
              handleSubmit(e);
            }}
          >
            <input
              type="email"
              className="signup-email"
              placeholder="Email ID"
              name="email"
              required
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
            />

            <input
              type="text"
              className="name"
              placeholder="Name"
              name="name"
              required
              value={name}
              onChange={(e) => {
                setName(e.target.value);
              }}
            />

            <input
              type="password"
              className="signup-password"
              placeholder="Password"
              name="password"
              required
              value={password}
              onChange={(e) => {
                setPassword(e.target.value);
              }}
            />
            <input
              type="password"
              className="signup-confirm-password"
              placeholder="Confirm Password"
              name="confirmPassword"
              required
              value={confirmPassword}
              onChange={(e) => {
                setConfirmPassword(e.target.value);
              }}
            />
            <button type="submit" className="signup-submit">
              Sign Up
            </button>
          </form>
        </div>
        <div className="signin-option">
          Already a member?<Link to={"/"}>Sign In</Link>
        </div>
      </div>
    </>
  );
};

export default Signup;
