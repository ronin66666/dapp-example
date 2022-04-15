import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';
import useStaticJsonRPC from './hooks/useStaticJsonRPC';


const DEBUG = true;
const NETWORKCHECK = true;
const USE_BURNER_WALLET = true; // toggle burner wallet feature
const USE_NETWORK_SELECTOR = false;

//providers: 固定providers，也可以使用类似小狐狸提供的provider
const providers = ["https://data-seed-prebsc-1-s1.binance.org:8545/"]
const localAppProvider = "https://data-seed-prebsc-1-s1.binance.org:8545/";

function App() {

  //load providers
  const localProvider = useStaticJsonRPC([localAppProvider])
  
  const [address, setAddress] = useState();


  return (
    <div className="App">
      
    </div>
  );
}

export default App;
