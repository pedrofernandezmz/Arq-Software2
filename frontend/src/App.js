import React from 'react';
import './App.css';
import Search from './components/Search';
import Results from './components/Results';


import {BrowserRouter as Router, Routes, Route} from 'react-router-dom'
function App() {
  return (
    <>
    <Router>
    
   
    <Routes>
    <Route exact path="/" element={<Search/>}/>
    <Route exact path="/results" element={<Results/>}/>
    </Routes>
 
    
    
    </Router>
      
  
    </>
  );
}

export default App;