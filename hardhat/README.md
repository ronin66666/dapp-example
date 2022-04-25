
## 依赖

部署
```bash
npm install --save-dev hardhat-deploy
npm install --save-dev  @nomiclabs/hardhat-ethers@npm:hardhat-deploy-ethers ethers
```
`openzeppelin`:

```bash
npm install --save-dev @openzeppelin/contracts @openzeppelin/contracts-upgradeable 

```
代理插件
```bash
npm install --save-dev @openzeppelin/hardhat-upgrades
```

随机数 `chainlikn`
该合约使用的是v1版本
文档
v2: https://docs.chain.link/docs/chainlink-vrf/
v1： https://docs.chain.link/docs/chainlink-vrf/v1/

link代币地址
https://docs.chain.link/docs/link-token-contracts/

bsctest 水龙头：[https://faucets.chain.link/chapel](https://faucets.chain.link/chapel?_ga=2.32896190.1454616603.1650874080-118985434.1649648357)

部署参数配置
v1: https://docs.chain.link/docs/vrf-contracts/v1/
v2: https://docs.chain.link/docs/vrf-contracts/

```bash
npm install --save-dev @chainlink/contracts
```

## game 

### game介绍
https://docs-cn.defina.finance/

### game 合约地址

#### bsc链地址

代币 FINA
Defina Token(FINA)： 0x426c72701833fdDBdFc06c944737C6031645c708

质押代币 `FinaMaster`（用于玩游戏）

代理地址 `proxy`： 0x01AA61FfC3058929CE9bDA8303a9D9B0Af74934c

实现地址 ：0xC459AC3fc36373A3fC665fC06E3A1Abf2871af7A

NFT `DefinaCard721`：
0xa318d9a2D6900A652FD0C9fea8c57a29b2a63709

NFT质押
DefinaNFTMaster 0xFfD7dA7E1CE01e6649aEC0524b255be903A0B2aA

随机数`Random Generator`
Random Generator： 0x58fb9D315658Ef7D784DB241e1e7595f4318F022

盲盒v2 (使用chainlink随机数) DefinaHeroBoxV2：
代理(prox)地址: 0x9d0Ad56df30ea438612EDed634D2BfEd021a5D85

实现地址：0xDaDA6200eB3aEe1f7F9D3972C267929b0c51D38C

账户注册，绑定邮箱`GameAccountRegister`

代理(proxy)地址: 0xeE6e5ccBE2309e3002fEC7972E6111fE1B9cb827

实现地址：0x58e41b89b5ECc41A5De323cF718D74C213974507

----------------------------------------------------------------
其他：

盲盒合约v1(没使用chainlink随机数)
DefinaHeroBox/Mystery Box (V1) 0x3F50dA5128D712b7C5c7B329a03901bcCA3dDAaE


市场、盲盒
NFTMarket for BlindBox/Mystery box 0x590bb8e6a9F2E0423ae93b14d63F4a5786C4046D


盲盒v1合约 blindBox 合约地址(简单版，这里没用到)
Mystery Box (V1) 0xde52486aE60e820127F5c17cCFF06352B87144dD
Mystery Box (OKEX Version) 0x37097D5329898df5e24979139fcF61E7581c9a94 


Defina Genesis NFT Card (Sold through TheForce Launchpad)[https://launchpad.theforce.trade/]
0x766709e2FA0C43352B6298c0605cAE3c84b77a49

DefinaHeroAdapter 0xD439F0A9881DccA056F7492f519B97dc6a4EA6Bf
Fina Voucher (FINAVOU) 代金券
0x627A6b8604bC075eeDD5eCB5a8A7B986BC347eda
Fina Voucher Counter 0x0d6BB1aF0a39AC347eE6Ed93d51fd577B6491329

TimeLock Contract 流动性锁定
Contract Address 0x9190BE2aF07c2e59FE9e3eCf7836c4d370aeA84D

## 流程
1.  注册
绑定邮箱 -> 服务器验证邮箱验证码(进入待验证状态) -> 客户端调用合约绑定邮箱 -> 通知服务端调用合约验证

2. 购买盲盒

3. 开盲盒

4. 质押卡片和代币

5. 玩游戏

### 部署

#### 部署`FinaToken`代币

脚本发布代币数量

```
 function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
        require(totalSupply() < maxSupply, "ERC20: total supply overflowing");
        _mint(to, amount);
    }
```

#### 部署NFT合约：`DefinaCard721`

新建card（种类，数量，品质等）

```solidity
function newCard(string memory name_, uint cardId_, uint camp_, uint rarity_,
        uint maxAmount_, string memory cardURI_) external onlyAdmin
```

#### 部署随机函数`Random Generator`

部署时构造函数参数见： https://docs.chain.link/docs/vrf-contracts/v1/

部署完成后向该合约质押 `Link` 代币

#### 部署盲盒合约：DefinaHeroV2

使用代理部署，部署时初始化盲盒相关配置，部署完成后完成以下设置

##### DefinaCard721相关设置

1. NFT（DefinaCard721）合约相关设置
   初始化盲盒合约地址， baseTokenURI_

   ```solidity
   function initialize(address blindBox_, string memory baseTokenURI_) onlyAdmin external initializer {}
   ```

3. 设置setSuperMinter（这里设置为盲盒合约地址，用于开盲盒铸造NFT），或者setMinter（根据需要设置）

   ```solidity
   //newSuperMinter_: DefinaHeroV2
   function setSuperMinter(address newSuperMinter_) external onlyAdmin
   ```
   
3. `Random Generator`随机合约设置consumerRole角色

   将盲盒`DefinaHeroV2`设置为`CONSUMER_ROLE`角色，才有权限调用随机函数

   ```solidity
   
   function grantRole(bytes32 role, address account) public virtual override onlyRole(getRoleAdmin(role)) {
           _grantRole(role, account);
       }
   ```

#### 部署`DefinaNFTMaster`NFT质押合约 

使用代理部署

#### 部署`FinaMaster`质押合约

使用代理部署





