import React, { Component, useContext, useEffect, useState} from "react";
import { PropertyItems } from "../Main/PropertyItems";
import swal from "sweetalert2";
import './main.css';



export const Main = () => {
  const [properties,setProperties] = useState([]);
  const [SearchCountries,setSearchContry] = useState([]);
  const [SearchCities,setSearchCity] = useState([]);
  const [SearchServices,setSearchService] = useState([]);
  const [valueCity, setValueCity] = useState("");
  const [valueCountry, setValueCountry] = useState("");
  const [valueService, setValueService] = useState("");
  const fetchApi1 = async()=>{
    const response = await fetch('http://localhost:8090/properties/country')
    .then((response) => response.json());
    setSearchContry(response);
    };
    useEffect(()=>{
    fetchApi1();
    },[])

    const fetchApi2 = async()=>{
      const response = await fetch('http://localhost:8090/properties/city')
      .then((response) => response.json());
      setSearchCity(response);
      };
      useEffect(()=>{
      fetchApi2();
      },[])

      const fetchApi3 = async()=>{
        const response = await fetch('http://localhost:8090/properties/service')
        .then((response) => response.json());
        setSearchService(response);
        };
        useEffect(()=>{
        fetchApi3();
        },[])

      const fetchApiProperty = async() => {
          const search = await fetch("http://localhost:8983/solr/Properties/select?defType=lucene&fq=city%3A%22"+valueCity+"%22&fq=country%3A%22"+valueCountry+"%22&fq=service%3A%22"+valueService+"%22&indent=true&q.op=OR&q=%3A")
          .then((res) => res.json())
          setProperties(search.response.docs)
          console.log(search.response.docs);
          };

        

      const handleChange=e=>{
      setValueCity(e.target.value);
       setValueCountry(e.target.value);
       setValueService(e.target.value);
        };
  
      const handleSubmit= (event)=>{
          event.preventDefault();
          fetchApiProperty();
      };
  return (
    <main>
        <div class="search-padre">
        <div class="search-hijo">
          <div class="select">
            <select list="pais-PT" name="pais-PT" id="pais-PT" valueCountry={ valueCountry } onChange={ (event) => setValueCountry(event.target.value)}>
            <option value="" selected disabled>Pais</option>
            {
                SearchCountries.map(country =>(
                  <option value={country} onChange={handleChange}>{country}
         </option> 
                ))
            }
            </select>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT" valueCity={ valueCity } onChange={ (event) => setValueCity(event.target.value)}>
              <option value="" selected disabled>Ciudad</option>
            {
                SearchCities.map(city =>(
                  <option value={city} onChange={handleChange}>{city}
                  </option> 
                ))
            }
              </select>
            </div>
          </div>
          <div class="lista">
            <div class="select">
              <select list="ciudad-PT" name="ciudad-PT" id="ciudad-PT" valueService={ valueService } onChange={ (event) => setValueService(event.target.value)}>
              <option value="" selected disabled>Servicio</option>
            {
                SearchServices.map(service =>(
                  <option value={service} onChange={handleChange}>{service}
                  </option> 
                ))
            }
              </select>
            </div>
          </div>
          <div class="div-search">
            <button class="btn-search" type="button" onClick = {handleSubmit}>
              Buscar
            </button>
          </div>
        </div>
      </div>
      <div className="Property">
      {
                properties.map((property) =>(
                  <PropertyItems key={property.id}
                  id={property.id}
                  tittle ={property.tittle}
                  size={property.size}
                  bathrooms={property.bathrooms}
                  service={property.service}
                  city={property.city}
                  state={property.state}
                  country={property.country}
                  street={property.street}
                  price={property.price}
                  rooms={property.rooms}
                  image={property.image}
                  description={property.description}
                  /> 
                ))
            }
        </div>
    </main>
  )
}