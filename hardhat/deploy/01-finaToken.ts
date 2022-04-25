
import { deployments, getNamedAccounts, network } from "hardhat";

const contractName = "FinaToken";

export default async function finaToken () {
    const { deployer } = await getNamedAccounts();

    console.log(`deployer = ${deployer} \n  network = ${network.name} chainId = ${network.config.chainId}`);

    const deploy = deployments.deploy;
    const deployResult = await deploy(contractName, {
        from: deployer,
        log: true
    });
    console.log("deploy " + contractName + " address = ", deployResult.address);
}


finaToken.tags = [
    "finaToken"
];
