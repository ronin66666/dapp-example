// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { BigNumberish } from "ethers";
import { deployments, ethers, getNamedAccounts } from "hardhat";
import { debugPort } from "process";
import { DefinaHeroBoxV2, FinaToken } from "../typechain";

async function mint(to: string, number: BigNumberish) {
  // We get the contract to deploy
  const token = await ethers.getContract<FinaToken>("FinaToken");
  const result = await token.mint(to, number).then(tx => tx.wait())

  console.log(result)
}

async function approve(who:string, spender: string, amount: BigNumberish) {

  // We get the contract to deploy
  const token = await ethers.getContract<FinaToken>("FinaToken", who);

  const result = await token.approve(spender, amount).then(tx => tx.wait());

  console.log(result)
}

async function allowance(owner: string, spender: string) {

  // We get the contract to deploy
  const token = await ethers.getContract<FinaToken>("FinaToken", owner);

  let approveAmount = await token.allowance(owner, spender);
  let amount = ethers.utils.formatEther(approveAmount)
  console.log("amount = ", amount);
}

async function transferFrom(from: string, to: string, amount: BigNumberish) {
  const token = await ethers.getContract<FinaToken>("FinaToken", from);
  const reuslt = await token.transfer(to, amount).then(tx => tx.wait());
  console.log(reuslt);
}

async function main() {
  
  const {deployer, game } = await getNamedAccounts();

  const heroBox =  (await deployments.get("DefinaHeroBox")).address;

  // await mint(deployer, ethers.utils.parseEther("100000000"))
  await approve(game, heroBox, ethers.utils.parseEther("1000000"));
  await allowance(game, heroBox);

  // await transferFrom(deployer, game, ethers.utils.parseEther("2000000"));

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
