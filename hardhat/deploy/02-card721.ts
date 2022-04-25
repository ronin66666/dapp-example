
import { deployments, getNamedAccounts, network } from "hardhat";

const contractName = "DefinaCard721";

export default async function card721 () {
    const { deployer } = await getNamedAccounts();

    console.log(`deployer = ${deployer} \n  network = ${network.name} chainId = ${network.config.chainId}`);

    const deploy = deployments.deploy;
    const deployResult = await deploy(contractName, {
        from: deployer,
        log: true
    });
    console.log("deploy " + contractName + " address = ", deployResult.address);
}

card721.tags = [
    "card721"
];

