import React, { useState } from "react";
import "./Order.css";
import logo from "./images/home.svg"
import cart from "./images/cart.svg"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
  return fetch("http://127.0.0.1:8080/user/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getProductById(id){
  return fetch("http://127.0.0.1:8080/product/" + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getOrderById(id) {
  return fetch("http://127.0.0.1:8080/order/" + id, {
    method:"GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

function goto(path){
  window.location = window.location.origin + path
}

function logout(){
  Cookie.set("user_id", -1, {path:"/"})
  window.location.reload();
}

function showProducts(products){
  return products.map((product) =>
  <div class="row product">
    <hr></hr>
    <div obj={product} key={product.product_id} className="product">
      <div class="col s4">
        <img width="auto" height="300px" src={product.picture_url}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
      </div>
      <div class="col s4">
      <h3 className="name">{product.name}</h3>
      <h4 className="price">Precio: {"$" + product.base_price}</h4>
      <div className="right">
        <a className="category">{product.category.name}</a>
      </div>
      </div>
    </div>
    <div className="quantity" class="col s2 right">
      <h5 className="amount"> Cantidad: {product.quantity}</h5>
      <h5 className="subtotal"> Subtotal: ${product.quantity * product.base_price} </h5>
    </div>
</div>
 )

}

async function getOrderProducts(){
  let items = []
  let a = Cookie.get("order").split(";")

  for (let i = 0; i < a.length; i++){
    let item = a[i];
    if(item != ""){
      let array = item.split(",")
      let id = array[0]
      let quantity = array[1]
      let product = await getProductById(id)
      product.quantity = quantity;
      items.push(product)
    }
  }
  return items
}


async function setOrder(setOrder, setTotal){
  let total = 0;
  await getOrderProducts().then(response => {
    setOrder(response)
    response.forEach((item) => {
      total += item.base_price * item.quantity;
    });
    setTotal(total)
  })
}



function Order(){
  const [user, setUser] = useState({});
  const [isLogged, setIsLogged] = useState(false);
  const [orderProducts, setOrderProducts] = useState([])
  const [total, setTotal] = useState(0)

  const login = (

    <span>
    <a id="logout" onClick={logout}> <span>{user.first_name} </span> </a>
    </span>
  )

  if (Cookie.get("user_id") > -1 && !isLogged) {
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }


  if (orderProducts.length <= 0 && Cookie.get("user_id") > -1){
    setOrder(setOrderProducts, setTotal)
  }

  const complete = (
    <div>
      <div class="row">
        <h1 class="blue-text col s6">Pedido realizado</h1>
        <h1 class="blue-text right">Total: ${total}</h1>
      </div>
    {showProducts(orderProducts)}
    </div>
  )


  const error = (
    <div>
    <div> BOO ERROR :(((( </div>
    <div> Error {Cookie.get("orderError")} </div>
    <h1> X </h1>
    </div>
  )

  return (
    <div className="order">
      <nav class=" yellow accent-2 ">
        <div class="nav-wrapper">
          <a href="/" class="brand-logo center blue-text text-darken-2"><img src={logo} width="50px" height="70px" /></a>
          <ul id="nav-mobile" class="right hide-on-med-and-down">
          <li><a class="black-text" onClick={logout}>Cerrar Sesion</a></li>
          </ul>
        </div>
      </nav>



      <div id="main">
        {window.location.pathname.split("/")[2] == "complete" ? complete : error}
      </div>
    </div>
  );
}

export default Order;
