import { BigNumberish } from "ethers";
import { deployments, ethers } from "hardhat";
import { FinaMaster } from "../typechain";

//重置能提取的总的金额
async function resetAllowance(amount: BigNumberish) {
    const contract = await ethers.getContract<FinaMaster>("FinaMaster");
    const result = await contract.resetAllowance(amount).then(tx => tx.wait());
    console.log("result = ", result);
}

async function setFinaAddress() {
    const mecToken = (await deployments.get("FinaToken")).address;
    const contract = await ethers.getContract<FinaMaster>("FinaMaster");
    const result = await contract.setFinaAddress(mecToken).then(tx => tx.wait()); 
    console.log("result = ", result);
}

//存入金额
async function deposit(amount: BigNumberish) {
    const contract = await ethers.getContract<FinaMaster>("FinaMaster");
    const result = await contract.deposit(amount).then(tx => tx.wait());
    console.log("result = ", result);
}

//转账（用户提现）
async function multiSend(address: string[], amount: BigNumberish[]) {
    const contract = await ethers.getContract<FinaMaster>("FinaMaster");
    const result = await contract.multiSend(address, amount).then(tx => tx.wait());
    console.log("result = ", result);
}

async function main() {
   await resetAllowance(ethers.utils.parseEther("2000000"));
//    await deposit(ethers.utils.parseEther("2000000"));

}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
  