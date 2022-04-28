import { BigInt } from "@graphprotocol/graph-ts"
import {
  DefinaHeroBoxV2, MintMulti,
} from "../generated/DefinaHeroBoxV2/DefinaHeroBoxV2"
import { HeroBoxBuyRecord } from "../generated/schema";

export function handleMintMulti(event: MintMulti): void {

  //从合约地址中加载，如果代理模式，则传入代理地址
  let heroBox = HeroBoxBuyRecord.load("0x9d0Ad56df30ea438612EDed634D2BfEd021a5D85");
  if (heroBox === null) {
    heroBox = new HeroBoxBuyRecord("0x9d0Ad56df30ea438612EDed634D2BfEd021a5D85");
  }
  heroBox.owner = event.params._to;
  heroBox.amount = event.params._amount;
  heroBox.save()
}

