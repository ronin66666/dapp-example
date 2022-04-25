import { ethers } from "hardhat";
import { GameAccountRegister } from "../typechain";

async function getEmailHash() {
    const gameAccountRegister = await ethers.getContract<GameAccountRegister>("GameAccountRegister");
    const emailHash = await gameAccountRegister.calcEmailHash("promaxnice@gmail.com");
    console.log("emailHash = ", emailHash);
}

async function main() {
    await getEmailHash();
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
  
