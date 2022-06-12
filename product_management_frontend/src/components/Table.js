import React, { useEffect, useState } from "react";
import "../css/table.css";
import { Link, Navigate, useNavigate } from "react-router-dom";
import axios from "axios";
const url = "https://product-management-go.herokuapp.com/getProducts";
const Table = () => {
  const navigate = useNavigate();
  const token = localStorage.getItem("token");
  const [items, setItems] = useState([]);
  async function getProducts() {
    try {
      const response = await axios.get(url, {
        headers: {
          token: token,
        },
      });

      console.log(response);
      const result = response.data;
      console.log(result);
      setItems(result);
    } catch (error) {
      console.log(error);
    }
  }
  //
  //
  //
  async function handleDelete(deleteID) {
    const url2 = `https://product-management-go.herokuapp.com/deleteProducts/${deleteID}`;
    const response = await axios.delete(url2, {
      headers: {
        token: token,
      },
    });
    getProducts();
  }
  //
  //
  //
  async function handleUpdate(updateID) {
    navigate(`/updateProduct/${updateID}`);
  }
  //
  //
  //
  useEffect(() => {
    getProducts();
  }, []);
  return (
    <>
      <div>
        <Link to={"/addProduct"}>
          <button className="add-button">Add Product</button>
        </Link>
      </div>
      <div className="table-parent">
        <table>
          <thead>
            <tr>
              <th>Brand</th>
              <th>Processor</th>
              <th>color</th>
              <th>Os</th>
              <th>Hardisk</th>
              <th>Actions</th>
            </tr>
          </thead>
          <tbody>
            {items ? (
              items.map((item) => {
                return (
                  <tr key={item._id}>
                    <td>{item.brand}</td>
                    <td>{item.processor}</td>
                    <td>{item.color}</td>
                    <td>{item.os}</td>
                    <td>{item.hardisk}</td>
                    <td>
                      <button
                        className="actions"
                        onClick={() => handleUpdate(item._id)}
                      >
                        Update
                      </button>
                      <button
                        className="actions"
                        onClick={() => handleDelete(item._id)}
                      >
                        Delete
                      </button>
                    </td>
                  </tr>
                );
              })
            ) : (
              <tr>
                <td>N/A</td>
                <td>N/A</td>
                <td>N/A</td>
                <td>N/A</td>
                <td>N/A</td>
                <td>N/A</td>
              </tr>
            )}
          </tbody>
        </table>
      </div>
    </>
  );
};

export default Table;
