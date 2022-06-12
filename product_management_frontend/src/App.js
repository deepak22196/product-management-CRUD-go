import { BrowserRouter, Routes, Route } from "react-router-dom";
import Signin from "./components/Signin";
import Signup from "./components/Signup";
import AllProducts from "./components/AllProducts";
import Form from "./components/Form";
import { useState } from "react";

function App() {
  const [token, setToken] = useState(localStorage.getItem("token"));
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Signin setToken={setToken} />}></Route>
          <Route path="/signup" element={<Signup />}></Route>
          {token ? (
            <Route
              path="/allProducts"
              element={<AllProducts setToken={setToken} />}
            ></Route>
          ) : (
            <Route
              path="/allProducts"
              element={<Signin setToken={setToken} />}
            ></Route>
          )}
          {token ? (
            <Route
              path="/addProduct"
              element={<Form setToken={setToken} />}
            ></Route>
          ) : (
            <Route
              path="/addProduct"
              element={<Signin setToken={setToken} />}
            ></Route>
          )}
          {token ? (
            <Route
              path="/updateProduct/:editID"
              element={<Form setToken={setToken} />}
            ></Route>
          ) : (
            <Route
              path="/updateProduct/:editID"
              element={<Signin setToken={setToken} />}
            ></Route>
          )}
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
