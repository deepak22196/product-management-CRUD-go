import React from "react";
import Header from "./Header";
import Table from "./Table";

const AllProducts = ({ setToken }) => {
  return (
    <>
      <Header setToken={setToken}></Header>
      <Table></Table>
    </>
  );
};

export default AllProducts;
