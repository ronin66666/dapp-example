// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// When running the script with `npx hardhat run <script>` you'll find the Hardhat
// Runtime Environment's members available in the global scope.
import { deployments, ethers, getNamedAccounts } from "hardhat";
import { DefinaCard721 } from "../typechain";

// 名字，卡牌类型编号， 阵营， 品质， 数量， cardURI
async function newCard(name: string, cardId: number, camp: number, rarity: number,
  maxAmount: number, cardURI: string) {

  const { deployer } = await getNamedAccounts();
  console.log(deployer);

  // We get the contract to deploy
  const card = await ethers.getContract<DefinaCard721>("DefinaCard721", deployer);
  
  const result = await card.newCard(name, cardId, camp, rarity, maxAmount, cardURI).then(tx => tx.wait());
  console.log(result)
}

async function initialize(baseTokenURI: string) {
  const { deployer } = await getNamedAccounts();
  console.log(deployer);

  const blindBoxAddress = (await deployments.get("DefinaHeroBox")).address
  console.log("blindBoxAddress = ", blindBoxAddress);
  
  // We get the contract to deploy
  const card = await ethers.getContract<DefinaCard721>("DefinaCard721", deployer);
  
  const result = await card.initialize(blindBoxAddress, baseTokenURI).then(tx => tx.wait());
  console.log(result)
}

async function echoNewCard() {
  await newCard("George Patton",      1001, 100, 11, 2000, "R")
  await newCard("Akechi Mitsuhide",   1002, 100, 11, 2000, "R")
  await newCard("Zorro",              1003, 100, 11, 2000, "R")
  await newCard("Erwin Schrödinger",  1004, 100, 11, 2000, "R")

  await newCard("Beowulf",            2001, 100, 12, 1000, "SR ")
  await newCard("Homer",              2002, 100, 12, 1000, "SR ")
  await newCard("Hattori Hanzo",      2003, 100, 12, 1000, "SR ")
  await newCard("Hasan",              2004, 100, 12, 1000, "SR ")

  await newCard("Alfred Nobel",       3001, 100, 13, 500, "SSR")
  await newCard("Cervantes",          3002, 100, 13, 500, "SSR")
  await newCard("Helen of Troy",      3003, 100, 13, 500, "SSR")
  await newCard("Paris",              3004, 100, 13, 500, "SSR")

  await newCard("Morgan le Fay",      4001, 100, 14, 100, "UR")
  await newCard("Solomon",            4002, 100, 14, 100, "UR")
  await newCard("Guanyu",             4003, 100, 14, 100, "UR")
  await newCard("Dante Alighieri",    4004, 100, 14, 100, "UR")

}

async function setSuperMinter(address: string) {
  const { deployer } = await getNamedAccounts();
  console.log(deployer);

  // We get the contract to deploy
  const card = await ethers.getContract<DefinaCard721>("DefinaCard721", deployer);

  const result = await card.setSuperMinter(address).then(tx => tx.wait());
  console.log(result);
}

async function getCardInfo(tokenId: number) {
  const { deployer } = await getNamedAccounts();
  console.log(deployer);

  // We get the contract to deploy
  const card = await ethers.getContract<DefinaCard721>("DefinaCard721", deployer);

  const cardId = await card.cardIdMap(tokenId);

  const info = await card.cardInfoes(cardId);
  console.log(info);
}

async function mint(cardId: number) {
  const { deployer } = await getNamedAccounts();
  console.log(deployer);

  // We get the contract to deploy
  const card = await ethers.getContract<DefinaCard721>("DefinaCard721", deployer);

  const result = await card.mint(cardId).then(tx => tx.wait);
  console.log(result);
}

async function balance(who: string) {
  const card = await ethers.getContract<DefinaCard721>("DefinaCard721");
  //获取tokenIds
  const tokenIds = await card.getTokenIDsByAddress(who);
  console.log("tokenIds = ", tokenIds.toString());
  
  for (let index = 0; index < tokenIds.length; index++) {
    const tokenId = tokenIds[index];
    const cardId = await card.cardIdMap(tokenId);
    const cardInfo = await card.cardInfoes(cardId)
    console.log(cardInfo.toString());
  }
}

async function setApprovalForAll(who: string, addr: string, approved: boolean) {
  const {deployer} = await getNamedAccounts();
  console.log(deployer);

  // We get the contract to deploy
  const heroBoxv1 = await ethers.getContract<DefinaCard721>("DefinaCard721", who);
  console.log(heroBoxv1.address)

  let result = await heroBoxv1.setApprovalForAll(addr, approved).then(tx => tx.wait());
  console.log(result);
}


async function main() {
  //  await echoNewCard();
  // getCardInfo(2);
  // await initialize("ipfs://");

  // const superMinter = (await deployments.get("DefinaHeroBox")).address;
  // await setSuperMinter(superMinter);

  // mint(1001)
  const { game } = await getNamedAccounts();

  await balance(game);

  //授权
  // const nftMasterAddr = (await deployments.get("DefinaNFTMaster")).address;
  // await setApprovalForAll(game, nftMasterAddr, true );

}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
