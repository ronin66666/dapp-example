
官网：https://thegraph.com/

## quickstart

https://thegraph.com/docs/en/developer/quick-start/

以太主网（Ethereum mainnet）使用 Subgraph Studio
其他网络使用：Hosted Service

支持网络：
https://thegraph.com/docs/en/developer/create-subgraph-hosted/#supported-networks

## 创建项目

先在`thegraph`平台创建项目
主网： https://thegraph.com/studio/
其他: https://thegraph.com/hosted-service/dashboard

## 初始化项目：

1. 如果没有安装脚手架的先安装

```bash
# NPM
$ npm install -g @graphprotocol/graph-cli

# Yarn
$ yarn global add @graphprotocol/graph-cli

```

2. 初始化项目

以太主网：`graph init --product subgraph-studio`
其他：`graph init--product hosted-service`

可选参数：
`--from-contract`: 指定合约地址
`--network`: 指定网络，
`--abi`: 指定abi文件
`<SUBGRAPH_SLUG>`: subgraph ID, 可以在你的subgraph详情页查看（一般填写： 用户名/项目名）

```bash

# 以太主网

graph init \
  --product subgraph-studio
  --from-contract <CONTRACT_ADDRESS> \
  [--network <ETHEREUM_NETWORK>] \
  [--abi <FILE>] \
  <SUBGRAPH_SLUG> [<DIRECTORY>]

# 其他
graph init \
  --product hosted-service
  --from-contract <CONTRACT_ADDRESS> \
  [--network <ETHEREUM_NETWORK>] \
  [--abi <FILE>] \
  <SUBGRAPH_SLUG> [<DIRECTORY>]

```

初始化项目

```bash
PS F:\blockchain\dapp-example> graph init --product hosted-service Game
√ Protocol · ethereum
× Subgraph name · Game
PS F:\blockchain\dapp-example> graph init --product hosted-service Game
√ Protocol · ethereum
√ Subgraph name · Ronin/Game
√ Directory to create the subgraph in · theGaph
√ Ethereum network · bsc
√ Contract address · 0x00000000000000
√ Fetching ABI from Etherscan
√ Contract Name · Game
———
  Generate subgraph
  Write subgraph to directory
√ Create subgraph scaffold
√ Initialize networks config
√ Initialize subgraph repository
√ Install dependencies with yarn
√ Generate ABI and schema types with yarn codegen

Subgraph Ronin/Game created in theGaph
Next steps:

  1. Run `graph auth` to authenticate with your deploy key.

  2. Type `cd theGaph` to enter the subgraph.

  3. Run `yarn deploy` to deploy the subgraph.

Make sure to visit the documentation on https://thegraph.com/docs/ for further information.
```

编写好相关配置后部署

```bash
graph auth --product hosted-service <ACCESS_TOKEN>

graph deploy --product hosted-service <GITHUB_USER>/<SUBGRAPH NAME>

```

## graph-cli
https://github.com/graphprotocol/graph-cli

graph项目脚手架相关命令

- `graph init` — Creates a new subgraph project from an example or an existing contract.
- `graph create` — Registers a subgraph name with a Graph Node.
- `graph remove` — Unregisters a subgraph name with a Graph Node.
- `graph codegen` — Generates AssemblyScript types for smart contract ABIs and the subgraph schema.
- `graph build` — Compiles a subgraph to WebAssembly.
- `graph deploy` — Deploys a subgraph to a [Graph Node](https://github.com/graphprotocol/graph-node).
- `graph auth` — Stores a [Graph Node](https://github.com/graphprotocol/graph-node) access token in the system's keychain.
- `graph local` — Runs tests against a [Graph Node](https://github.com/graphprotocol/graph-node) test environment (using Ganache by default).
- `graph test` — Downloads and runs the [Matchstick](https://github.com/LimeChain/matchstick) rust binary in order to test a subgraph.


## graph-node

Graph Node 是一个开源的 Rust 实现，可以通过 GraphQL 端点查询 来自 Ethereum 区块链 事件 的数据存储。

https://github.com/graphprotocol/graph-node

可以搭建本地节点测试，相对比较麻烦 

### 环境搭建

#### rust 安装
https://www.rust-lang.org/tools/install

其他安装方法
https://forge.rust-lang.org/infra/other-installation-methods.html

mac推荐使用`homebrew`安装
```bash
  brew install rust
```
安装完成验证
``` bash
$ rustc --version
  rustc 1.59.0
```

#### PostgreSQL 安装
https://www.postgresql.org/download/
mac推荐使用`homebrew`
```bash
$ brew install postgresql
```


#### IPFS 安装
https://docs.ipfs.io/install/

测试可以直接下载桌面版本

### 本地运行Graph 节点

https://github.com/graphprotocol/graph-node#running-a-local-graph-node





