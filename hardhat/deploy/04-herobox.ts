// herobox 盲盒合约部署脚本

import { deployments, ethers, getNamedAccounts, network, upgrades } from "hardhat";
import { DeployFunction } from "hardhat-deploy/types";

const contractName = "HeroBoxV1";

export default async function deployHero () {

    const { deployer } = await getNamedAccounts();

    console.log(`deployer = ${deployer} \n  network = ${network.name} chainId = ${network.config.chainId}`);

    const deploy = deployments.deploy;
    

    const result1 = await deploy("DefinaHeroBox", {
        from: deployer,
        contract: "DefinaHeroBoxV2",
        log: true,
        proxy: {//如果已部署过TransparentProxy， 下次就会走升级
            owner: deployer,
            proxyContract: "OpenZeppelinTransparentProxy",
            execute: {
                init: {
                    methodName: "initialize",
                    args: await initializeArgs()
                }
            }
        }
    })
    console.log("heroBox =  ", result1.address);

    //升级
    // const upgradesResult = await deploy("DefinaHeroBox", {
    //     contract:"DefinaHeroBoxV2",
    //     from: deployer,
    //     log: true,
    //     proxy: {//如果已部署过TransparentProxy， 下次就会走升级
    //         owner: deployer,
    //         proxyContract: "OpenZeppelinTransparentProxy",
    //     }
    // })
    // console.log("heroBoxV2 =  ", upgradesResult.address);

}

//获取初始化数据
async function initializeArgs(): Promise<any[]> {

    const nftToken =  (await deployments.get("DefinaCard721")).address;
    const baseTokenURI = "ipfs://";
    const currToken = (await deployments.get("FinaToken")).address;
    const ntfPrice = ethers.utils.parseEther("1");

    const { deployer } = await getNamedAccounts();
    const tokenReceiveAddress = deployer;

    const cardids = [
        1001, 1002, 1003, 1004,  
        2001, 2002, 2003, 2004, 
        3001, 3002, 3003, 3004, 
        4001, 4002, 4003, 4004,
    ]; 

    const cardnums = [
        111, 111, 111, 111, 
          55, 55, 55, 56, 
          27, 27, 28, 28, 
          5, 5, 6, 6, 
        ];

    const randomGenerator = (await deployments.get("RandomGenerator")).address;

    return [nftToken, baseTokenURI, currToken, ntfPrice, tokenReceiveAddress, cardids, cardnums, randomGenerator];
}


deployHero.tags = [
    "herobox"
];
