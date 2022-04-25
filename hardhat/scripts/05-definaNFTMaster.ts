import { deployments, ethers, getNamedAccounts } from "hardhat";
import {DefinaCard721, DefinaNFTMaster } from "../typechain";


async function getNftTokenAddr() {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster");
    const nftToken = await nftMaster.nftToken();
    console.log("nftToken = ", nftToken);
}

async function changeNFTTokenAddress() {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster");
    const nftAddr = (await deployments.get("DefinaCard721")).address
    const result = await nftMaster.changeLockTime(nftAddr).then(tx => tx.wait());
    console.log("result = ", result);
}

async function stake(who: string, tokenId: number) {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster", who);
    const result = await nftMaster.stake(tokenId).then(tx => tx.wait());
    console.log("result = ", result);
}

async function stakeMulti(who: string, tokenIds: number[]) {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster", who);
    const result = await nftMaster.stakeMulti(tokenIds).then(tx => tx.wait());
    console.log("result = ", result);
}

async function getTokensStakedByAddress(who: string) {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster");
    const tokenIds =  (await nftMaster.getTokensStakedByAddress(who)).map(tokenId => tokenId.toNumber());
    console.log("staked tokenIds = ", tokenIds);
}


async function withdraw(who:string, tokenId: number) {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster", who);
    const result = await nftMaster.withdraw(tokenId).then(tx => tx.wait());
    console.log("result = ", result);
}

async function withdrawMulti(who:string, tokenId: number[]) {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster", who);
    const result = await nftMaster.withdrawMulti(tokenId).then(tx => tx.wait());
    console.log("result = ", result);
}

async function isStaked(tokenId: number) {
    const nftMaster = await ethers.getContract<DefinaNFTMaster>("DefinaNFTMaster");
    const isStaked =  await nftMaster.isStaked(tokenId);
    console.log("staked tokenIds = ", isStaked);
}


async function main() {
    const {deployer, game } = await getNamedAccounts();

    // await getNftTokenAddr();
    
    // await stake(game, 1);
    // await getTokensStakedByAddress(game);

    // await withdraw(game, 1);

    await isStaked(1);
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
  });
  

