import { deployments, ethers, getNamedAccounts } from "hardhat";
import { DefinaNFTMaster } from "../typechain";

async function deployDefinaNFTMaster() {

    const { deployer } = await getNamedAccounts();
    const deploy = deployments.deploy;

    const cardAddr = (await deployments.get("DefinaCard721")).address;

    const result = await deploy("DefinaNFTMaster", {
        from: deployer,
        log: true,
        proxy: {//如果已部署过TransparentProxy， 下次就会走升级
            owner: deployer,
            proxyContract: "OpenZeppelinTransparentProxy",
            execute: {
                init: {
                    methodName: "initialize",
                    args: [cardAddr]
                }
            }
        }
    });

    console.log("deploy address = ", result.address);
}

export default deployDefinaNFTMaster;
deployDefinaNFTMaster.tags = ["definaNFTMaster"]