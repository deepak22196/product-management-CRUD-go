import React from "react";
import Header from "./Header";
import "../css/form.css";
import { Link, useNavigate, useParams } from "react-router-dom";
import { useState } from "react";
import axios from "axios";

const Form = ({ setToken }) => {
  const token = localStorage.getItem("token");
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    brand: "",
    color: "",
    os: "",
    hardisk: "",
    processor: "",
  });
  const { editID } = useParams();
  const url1 = `https://product-management-go.herokuapp.com/addProduct`;
  async function handleSubmit(e) {
    e.preventDefault();
    if (editID) {
      console.log("updating one record");
      const url2 = `https://product-management-go.herokuapp.com/updateProduct/${editID}`;
      try {
        const response = await axios.put(url2, formData, {
          headers: {
            token: token,
          },
        });
        navigate("/allProducts");
      } catch (error) {
        console.log(error);
      }
    } else {
      try {
        const response = await axios.post(url1, formData, {
          headers: {
            token: token,
          },
        });
        navigate("/allProducts");
      } catch (error) {
        console.log(error);
      }
    }
  }

  return (
    <>
      <Header setToken={setToken}></Header>
      <div className="form-parent">
        <div className="form-heading">
          {" "}
          <h3>Enter Product details here</h3>
        </div>
        <form
          onSubmit={(e) => {
            handleSubmit(e);
          }}
        >
          <input
            type="text"
            id="brand"
            name="brand"
            placeholder="Brand"
            required
            value={formData.brand}
            onChange={(e) => {
              setFormData({ ...formData, brand: e.target.value });
            }}
          />

          <input
            type="text"
            id="processor"
            name="processor"
            placeholder="Processor"
            required
            value={formData.processor}
            onChange={(e) => {
              setFormData({ ...formData, processor: e.target.value });
            }}
          />

          <input
            type="text"
            id="color"
            name="color"
            placeholder="Color"
            required
            value={formData.color}
            onChange={(e) => {
              setFormData({ ...formData, color: e.target.value });
            }}
          />

          <input
            type="text"
            id="os"
            name="os"
            placeholder="OS"
            required
            value={formData.os}
            onChange={(e) => {
              setFormData({ ...formData, os: e.target.value });
            }}
          />

          <input
            type="text"
            id="hardisk"
            name="hardisk"
            placeholder="Hardisk"
            required
            value={formData.hardisk}
            onChange={(e) => {
              setFormData({ ...formData, hardisk: e.target.value });
            }}
          />
          <button type="submit" className="form-submit">
            Submit
          </button>
        </form>
      </div>
    </>
  );
};

export default Form;
