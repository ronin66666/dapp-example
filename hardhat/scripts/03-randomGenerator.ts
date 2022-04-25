
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import {  toUtf8Bytes } from "ethers/lib/utils";
import { deployments, ethers, getNamedAccounts } from "hardhat";
import { RandomGenerator } from "../typechain";

async function requestRandomNumber(tokenId: number, addr: string) {
  const { deployer, game } = await getNamedAccounts();

  const randomGenerator = await ethers.getContract<RandomGenerator>("RandomGenerator", deployer);
  console.log("randomGenerator address: ", randomGenerator.address)

  const result = await randomGenerator.requestRandomNumber(tokenId, addr).then(tx => tx.wait());
  console.log(result)
}

async function setConsumerRole(consumerRole: string) {

  // We get the contract to deploy
  const randomGenerator = await ethers.getContract<RandomGenerator>("RandomGenerator");
  console.log("randomGenerator address: ", randomGenerator.address)

  const role = ethers.utils.keccak256(toUtf8Bytes("CONSUMER_ROLE"));
  console.log("role: ", role);

  const result = await randomGenerator.grantRole(role, consumerRole).then(tx => tx.wait());
  console.log(result);

}

async function main() {

  const heroboxAddr =  (await deployments.get("DefinaHeroBox")).address;
   
  await setConsumerRole(heroboxAddr);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
