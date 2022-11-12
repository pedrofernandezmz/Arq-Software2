// // import React, {useEffect, useState}  from "react";

// // import Cookies from 'universal-cookie';
// // import "./Results.css";
// // const cookies = new Cookies();
// // async function getSearch(query){
// //     // return (fetch('http://localhost:8983/solr/publicaciones/select?&defType=lucene&indent=true&q=description:"'+query+'"%0Atitle:"'+query+'"&q.op=OR', {method:"GET",
// //     return (fetch('http://localhost:8983/solr/publicaciones/select?&defType=lucene&indent=true&q=description:casa+%0Atitle:+casa+&q.op=OR', {method:"GET",
// //     mode: 'no-cors'}).then(response => response.json()));
// //      }

// // async function getStuff(){
// //     let items= await getSearch(cookies.get("busqueda_limpia"))
// //     return(
// //     <div>
// //     {items.map((item)=>
// //     <div>{item.title}</div>
// //     )}</div>)
// // }
   
// // function Results(){
// //     return(<div>{getStuff()}</div>)

// // }
// // export default Results;


// import React, {useEffect, useState } from "react";
// import "./Search.css";
// import Cookies from "universal-cookie";

// const Cookie = new Cookies();

// async function getProducts(){
//   return await fetch('https://shoppingapiacme.herokuapp.com/shopping', {
//     method: "GET",
//     headers: {
//       "Content-Type": "application/json"
//     }
//   }).then(response => response.json())
// }

// function goto(path){
//   window.location = window.location.origin + path
// }

// function showProducts(products){
//   return products.map((product) =>
//     <div class="col s2">
//       <div class="product large" key={product.id} className="product">
//         <div class="product-image">
//         <img width="128px" height="300px" src={product.image}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
//         </div>
//         <div class="product-content">
//           <span class="text-blue"><a className="name">{product.brand}</a></span>
//           <p>Cantidad disponible: {product.item}</p>
//         </div>
//       </div>
//     </div>
//  )
// }


// function Results() {
//   const [products, setProducts] = useState([])

//   if (products.length <= 0){
//     getProducts().then(response => {setProducts(response)})
//   }

//   return (
//     <div className="home">
//       <div class="row" id="main">
//         {showProducts(products)}
//       </div>
//     </div>
//   );
// }

// export default Results;
import React, {useEffect, useState}  from "react";

import Cookies from 'universal-cookie';
import "./Results.css";
const cookies = new Cookies();
async function getSearch(query){
    return (fetch('http://localhost:8983/solr/publicaciones/select?&defType=lucene&indent=true&q=description:"'+query+'"%0Atitle:"'+query+'"&q.op=OR', {method:"GET",
    
    
    mode: 'no-cors'}).then(response => response.json()));
     }

async function getStuff(){
    let items= await getSearch(cookies.get("busqueda_limpia"))
    return(
    <div>
    {items.map((item)=>
    <div>{item.title}</div>
    )}</div>)
}
   
function Results(){
    return(<div>{getStuff()}</div>)

}
export default Results;