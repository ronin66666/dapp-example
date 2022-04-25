// herobox 盲盒合约部署脚本

import { deployments, ethers, getNamedAccounts, network, upgrades } from "hardhat";

export default async function deployMecMaster () {

    const { deployer } = await getNamedAccounts();

    console.log(`deployer = ${deployer} \n  network = ${network.name} chainId = ${network.config.chainId}`);

    const deploy = deployments.deploy;

    const mecToken = (await deployments.get("FinaToken")).address;

    const result1 = await deploy("FinaMaster", {
        from: deployer,
        contract: "FinaMaster",
        log: true,
        proxy: {//如果已部署过TransparentProxy， 下次就会走升级
            owner: deployer,
            proxyContract: "OpenZeppelinTransparentProxy",
            execute: {
                init: {
                    methodName: "initialize",
                    args: [mecToken]
                }
            }
        }
    })
    console.log("finaMaster =  ", result1.address);

    //升级, 如果合约名一样，只是该了代码则不需要调用这种方法
    // const upgradesResult = await deploy("EPKMaster", {
    //     contract:"EPKMasterV2",
    //     from: deployer,
    //     log: true,
    //     proxy: {//如果已部署过TransparentProxy， 下次就会走升级
    //         owner: deployer,
    //         proxyContract: "OpenZeppelinTransparentProxy",
    //     }
    // })
    // console.log("heroBoxV2 =  ", upgradesResult.address);

}

deployMecMaster.tags = [
    "finaMaster"
];
