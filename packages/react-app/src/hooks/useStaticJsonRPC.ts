import { StaticJsonRpcProvider } from "@ethersproject/providers";
import { TEthersProvider } from "eth-hooks/models";
import { ethers } from "ethers"
import { useCallback, useEffect, useState } from "react";
import { Web3Provider } from "walletlink/dist/provider/Web3Provider";

const createProvider = async (url?: string | ethers.utils.ConnectionInfo | undefined, network?: ethers.providers.Networkish | undefined) => {
    const  p = new ethers.providers.StaticJsonRpcProvider(url);
    await p.ready;
    return p;
}

export default function useStaticJsonRPC(urlArray: string[]) {

    const [provider, setProvider] = useState<TEthersProvider>();
    
    // useCallback: 只有urlArray值改变时才重新渲染
    const handleProviders = useCallback(async () => {
        try {
            const p = await Promise.race(urlArray.map(createProvider));
            setProvider(p);
        } catch (error) {
            console.log(error);
        }
    }, [urlArray]);

    //初始化
    useEffect(() => {
        handleProviders();
    }, [JSON.stringify(urlArray)])

    return provider as TEthersProvider;
}