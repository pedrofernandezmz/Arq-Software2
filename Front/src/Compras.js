import React, { useState, List, Checkbox} from "react";
import "./Compras.css";
import "./css/materialize.css";
import logo from "./images/home.svg"
import cart from "./images/cart.svg"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
    return await fetch('http://localhost:8080/user/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}

async function getOrdersByUserId(id){
    return await fetch('http://localhost:8080/orders/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}

function goto(path){
  window.location = window.location.origin + path
}


function logout(){
  Cookie.set("user_id", -1, {path: "/"})
  document.location.reload()
}

function showDetails(id){
  let div = document.getElementsByClassName("o" + id)
  for(let detail of div){
    let show = true;
    detail.classList.value.split(" ").forEach(c=>{
      if (c == "hidden"){
        show = true;
      }
    })
    if(show){
      detail.classList.remove("hidden")
    }else{
      detail.classList.add("hidden")
    }
  }


}

function getDetails(details){
  return details.map((detail) =>
      <tbody>
        <tr className="detail">
          <td className="id">{detail.product_id}</td>
          <td className="name">{detail.name}</td>
          <td className="quantity">{detail.quantity}</td>
          <td className="price">{"$" + detail.price}</td>
        </tr>
      </tbody>
  )
}

function showUserOrders(orders){
  return orders.map((order) =>

  <div obj={order} key={order.order_id} className="order" onClick={()=>showDetails(order.order_id)}>
    <a className="ordern">Orden: {order.order_id}       </a>
    <a className="date">Fecha: {order.date.split("T")[0]}       </a>
    <a className="total">Total: <span>{"$" + order.total}  </span></a>

    <div className={"details" + " o" + order.order_id + " hidden"}>
      <table>
        <thead>
          <tr className="detail thead">
          <td className="name">ID</td>
            <td className="name">Producto</td>
            <td className="quantity">Cantidad</td>
            <td className="price">Precio</td>
          </tr>
        </thead>

      {getDetails(order.details)}
      </table>
    </div>

  </div>
 )
}



function User() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [cartItems, setCartItems] = useState("")
  const [orders, setOrders] = useState([])


  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
    if (!cartItems && Cookie.get("cartItems")){
      setCartItems(Cookie.get("cartItems"))
    }

    if(Cookie.get("orders") == undefined){
      getOrdersByUserId(Cookie.get("user_id")).then(response => {setOrders(response)})
    }


  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }

  const login = (

    <span>
    <a id="logout" onClick={logout}> <span>{user.first_name} </span> </a>
    </span>
  )

  const showUserInfo = (
    <div className="userInfo">
      <div> {user.first_name} {user.last_name} </div>
      <div> Username: {user.username} </div>
      <div>{orders.length > 0 ? showUserOrders(orders) : <a> NO HAY ORDENES</a>}</div>
    </div>
  )

  const pleaseLogin = (
    <div> INICIE SESION PARA VER SUS COMPRAS </div>
  )

  return (
    <div className="home">
<nav class=" yellow accent-2 ">
        <div class="nav-wrapper">
          <a href="/" class="brand-logo center blue-text text-darken-2"><img src={logo} width="50px" height="70px"/> </a>
          <ul id="nav-mobile" class="right hide-on-med-and-down">
          <li><a class="black-text" onClick={logout}>Cerrar Sesion</a></li>
          </ul>
        </div>
      </nav>
      <div id="main">
        {isLogged ? showUserInfo : pleaseLogin}
      </div>
    </div>
  );
}

export default User;