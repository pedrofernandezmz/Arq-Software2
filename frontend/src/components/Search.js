import React from "react";
import "./Search.css";
import Cookies from 'universal-cookie';

import {useState} from 'react';
const cookies = new Cookies();

function gopath(path){
    window.location = window.location.origin + path
  }

  
async function SearchByQuery(){

    let current= cookies.get('busqueda')
    let chain=''
    let a= current.split("?")
    let item = a[1];
    let b=item.split("=")
    item=b[1]
    let c= item.split("+")

    for (let i = 0; i < c.length; i++){

        chain = `${chain} `+`${c[i]} `;
        cookies.set("busqueda_limpia", chain)

    }
    gopath("/results")
}
function Search(){

   
  
   //var {query} = document.forms[0];
const renderForm = (
    <div id="cover">
  <form method="get" action="">
    <div class="tb">
      <div class="td">
        <input type="text" id="search_query" placeholder="Buscar" name="search" required /></div>
      <div class="td" id="s-cover">
      <button  id="search_button" onClick={SearchByQuery(cookies.set("busqueda", window.location.search))} type="input">
          <div id="s-circle"></div>
          <span></span>
        </button>
      </div>
    </div>
  </form>
</div>
);

      return (
      <div className="app">
      <div className="search-form">

      {renderForm}

      </div>
      </div>
      );
    
}export default Search;