import { deployments, getChainId, getNamedAccounts, network } from "hardhat";

async function deployGreeter() {

    const { deployer } = await getNamedAccounts();
    const chainId = await getChainId();
    const networkName = network.name;

    console.log(`deployer = ${deployer} network = ${networkName} chainId = ${chainId} `);
    
    const deploy = deployments.deploy;
    const deployment = await deploy("Greeter", {
        from: deployer,
        log: true,
        args: ["greeter"]
    })

    console.log("address = ", deployment.address);
    
}

export default deployGreeter;
deployGreeter.tags = ["greeter"];