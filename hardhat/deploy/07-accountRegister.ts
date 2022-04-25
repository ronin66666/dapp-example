import { deployments, ethers, getNamedAccounts } from "hardhat";
import { DefinaNFTMaster } from "../typechain";

async function deployAccountRegist() {

    const { deployer } = await getNamedAccounts();
    const deploy = deployments.deploy;


    const result = await deploy("GameAccountRegister", {
        from: deployer,
        log: true,
        proxy: {//如果已部署过TransparentProxy， 下次就会走升级
            owner: deployer,
            proxyContract: "OpenZeppelinTransparentProxy",
            execute: {
                init: {
                    methodName: "initialize",
                    args: []
                }
            }
        }
    });

    console.log("deploy address = ", result.address);
}

export default deployAccountRegist;

deployAccountRegist.tags = ["accountRegister"]