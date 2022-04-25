// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>`
import { deployments, ethers, getNamedAccounts } from "hardhat";
import { DefinaCard721, FinaToken, RandomGenerator, DefinaHeroBoxV2 } from "../typechain";

async function setCurrToken(address: string) {
  
    // We get the contract to deploy
    const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox");

    let result = await heroBox.pause().then(tx => tx.wait());
    console.log(result);

    result = await heroBox.setCurrToken(address).then(tx => tx.wait());
    console.log(result);

    result = await heroBox.unpause().then(tx => tx.wait());
    console.log(result);

    console.log("set currToken address: ", heroBox.currToken)
}

async function setPublicBuyableQuota(amount: number) {
  // We get the contract to deploy
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox");

  let result = await heroBox.pause().then(tx => tx.wait());
  console.log(result);
  result = await heroBox.setPublicBuyableQuota(amount).then(tx => tx.wait());
  console.log(result);
  result = await heroBox.unpause().then(tx => tx.wait());
  console.log(result);
}

async function mintMulti(who: string, amount: number) {
  // We get the contract to deploy
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox", who);

  const result = await heroBox.mintMulti(amount).then(tx => tx.wait());

  console.log(result);
}

async function setApprovalForAll(who:string, operator: string, approved: boolean) {

  // We get the contract to deploy
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox");
  console.log(heroBox.address)

  let result = await heroBox.setApprovalForAll(operator, approved).then(tx => tx.wait());
  console.log(result);
}

async function setRandomGenerator() {
  
  const {deployer} = await getNamedAccounts();
  console.log(deployer);

  const randomGenerator = await ethers.getContract<RandomGenerator>("RandomGenerator", deployer);
  console.log("randomGenerator address: ", randomGenerator.address)

  // We get the contract to deploy
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox", deployer);
  console.log(heroBox.address)
  let result = await heroBox.pause().then(tx => tx.wait());
  console.log(result);
  let result1 = await heroBox.setRandomGenerator(randomGenerator.address).then(tx => tx.wait());
  console.log(result1);
  result = await heroBox.unpause().then(tx => tx.wait());
  console.log(result);
}

async function open(tokenId: number) {
  const { game } = await getNamedAccounts();
  console.log("game address: ", game);

  // We get the contract to deploy
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox", game);
  console.log("heroBox address: ", heroBox.address)

  let result = await heroBox.open(tokenId).then(tx => tx.wait());
  console.log(result);
}


async function totalBlindBoxQuota() {
  const { deployer } = await getNamedAccounts();
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox", deployer);
 
  const totalBlindBoxQuota = await heroBox.totalBlindBoxQuota();
  const totalBlindBoxSold = await heroBox.totalBlindBoxSold();
  const totalBlindBoxRemaining = await heroBox.totalBlindBoxRemaining();

  console.log(totalBlindBoxQuota.toNumber(), totalBlindBoxSold.toNumber(), totalBlindBoxRemaining.toNumber() );
  
}


async function balance(who: string) {
  const { game } = await getNamedAccounts();
  console.log("game address: ", game);

  // We get the contract to deploy
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox", game);
  console.log("heroBox address: ", heroBox.address)

  let count = (await heroBox.balanceOf(who)).toNumber();

  console.log("count = ", count);
  
  for (let index = 0; index < count; index++) {
    const tokenId = await heroBox.tokenOfOwnerByIndex(game, index);
    console.log(tokenId.toString());
  }
}

async function nftToken() {
  const heroBox = await ethers.getContract<DefinaHeroBoxV2>("DefinaHeroBox");
  const nft = await heroBox.nftToken()
  console.log("nft = ", nft);
}

async function main() {

  const {deployer, game } = await getNamedAccounts();
  const heroBoxAddr = (await deployments.get("DefinaHeroBox")).address;
  // await initialize();
  // await setCurrToken('0xCf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9');
  // await setPublicBuyableQuota(1000);

  // await setApprovalForAll(game, heroBoxAddr , true)
  // await mintMulti(game, 3);
  await balance(game);
  // setRandomGenerator()
  // await open(2);
  // info();

  // await nftToken()
  // await totalBlindBoxQuota();

}
// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
