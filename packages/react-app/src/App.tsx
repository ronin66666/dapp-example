import React, { useCallback, useEffect, useState } from 'react';
import './App.css';
import useStaticJsonRPC from './hooks/useStaticJsonRPC';
import { ALCHEMY_KEY, NETWORKS } from './constants';
import web3ModalSetup from './heplers/Web3ModalSetup';
import Web3Modal from "web3modal";
import { ethers } from 'ethers';
import { useGasPrice, useUserProviderAndSigner } from 'eth-hooks';

import { useExchangeEthPrice } from "eth-hooks/dapps/dex";
import NetworkDisplay from './components/NetworkDisplay';

const DEBUG = true;
const NETWORKCHECK = true;
const USE_BURNER_WALLET = true; // toggle burner wallet feature
const USE_NETWORK_SELECTOR = false;

//providers: 固定providers，也可以使用类似小狐狸提供的provider
const providers = [
  "https://eth-mainnet.gateway.pokt.network/v1/lb/611156b4a585a20035148406",
  `https://eth-mainnet.alchemyapi.io/v2/${ALCHEMY_KEY}`,
  "https://rpc.scaffoldeth.io:48544",
  "https://data-seed-prebsc-1-s1.binance.org:8545/"
]
const localAppProvider = "http://localhost:8545";
const web3Modal: Web3Modal = web3ModalSetup();

function App() {

  //ethers.providers.ExternalProvider | ethers.providers.JsonRpcFetchFunc
  const [injectedProvider, setInjectedProvider] = useState<any>();


  const [selectedNetwrok, setSlectedNetwrok] = useState<string>();

  const targetNetwork = NETWORKS[selectedNetwrok as string];

  //load all your providers
  const localProvider = useStaticJsonRPC([targetNetwork.rpcUrl]);
  const mainnetProvider = useStaticJsonRPC(providers);


  //signer
  const userProviderAndSigner = useUserProviderAndSigner(injectedProvider, localProvider, USE_BURNER_WALLET)
  const userSigner = userProviderAndSigner?.signer;

  const localChainId = localProvider && localProvider._network && localProvider._network.chainId;
  // const selectedChainId =
  //   userSigner && userSigner.provider && userSigner.provider._network && userSigner.provider._network.chainId;
  const selectedChainId = userProviderAndSigner?.providerNetwork?.chainId

  const [address, setAddress] = useState<string>();

  /* 💵 This hook will get the price of ETH from 🦄 Uniswap: */
  const price = useExchangeEthPrice(targetNetwork, mainnetProvider);
  /* 🔥 This hook will get the price of Gas from ⛽️ EtherGasStation */
  const gasPrice = useGasPrice(targetNetwork, "fast");

  const blockExplorer = targetNetwork.blockExplorer;


  //断开web3Modal连接
  const logoutOfWeb3Modal = async () => {
    web3Modal.clearCachedProvider();
    if (injectedProvider && injectedProvider.provider && typeof injectedProvider.provider.disconnect == "function") {
      await injectedProvider.provider.disconnect();
    }
    setTimeout(() => {
      window.localtion.reload();
    }, 1);
  }

  //加載web3Modal
  const loadWeb3Modal = useCallback(async () => {

    const provider = await web3Modal.connect();
    setInjectedProvider(new ethers.providers.Web3Provider(provider));

    //chainId 监听
    provider.on("chainChanged", (chainId: number) => {
      console.log(`chain changed to ${chainId}! updating providers`);
      setInjectedProvider(new ethers.providers.Web3Provider(provider));
    });
    //账户改变
    provider.on("accountsChanged", () => {
      console.log(`account changed!`);
      setInjectedProvider(new ethers.providers.Web3Provider(provider));
    });

    // Subscribe to session disconnection
    provider.on("disconnect", (code: number | string, reason: any) => {
      console.log(code, reason);
      logoutOfWeb3Modal();
    });

  }, [setInjectedProvider])

  //加载的时候访问地址
  useEffect(() => {
    async function getAddress() {
      if (userSigner) {
        const newAddress = await userSigner.getAddress();
        setAddress(newAddress);
      }
    }
  }, [userSigner])

  return (
    <div className="App">
      <NetworkDisplay
        NETWORKCHECK={NETWORKCHECK}
        localChainId={localChainId}
        selectedChainId={selectedChainId as number}
        targetNetwork={targetNetwork}
        USE_NETWORK_SELECTOR={USE_NETWORK_SELECTOR}
      />
    </div>
  );
}

export default App;
