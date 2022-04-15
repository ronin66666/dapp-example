import { ethers, getNamedAccounts } from "hardhat";
import { Greeter } from "../typechain";

async function main() {
    const { deployer } = await getNamedAccounts();
    const greeter = await ethers.getContract<Greeter>("Greeter", deployer);
    const result = await greeter.setGreeting("new greeting").then(tx => tx.wait());
    console.log(result);
    
    const newValue = await greeter.greet();
    console.log(newValue);
    
}

main().catch((error) => {
    console.log(error);
    process.exit(1);
})