
# create project

create hardhat 

```bash
cd dapp-example/packages
npx hardhat
```

create react-app
```bash
cd dapp-example/packages
npx create-react-app react-app --template typescript
```

inistall dependencies

```bash
cd dapp-example/packages/hardhat
npm install

cd dapp-example/packages/react-app
npm install
```
# useage
install and start your 👷‍ Hardhat chain:

```bash
cd dapp-example
yarn install
yarn chain
```

in a second terminal window, start your 📱 frontend:
```bash
cd dapp-example
yarn start
```

in a third terminal window, 🛰 deploy your contract:

```bash
cd dapp-example
yarn deploy
```

run script

```bash
cd dapp-example
yarn hardhat-run scripts/01-greeter.ts
```

 Edit your smart contract YourContract.sol in packages/hardhat/contracts

 Edit your frontend App.jsx in packages/react-app/src

 Edit your deployment scripts in packages/hardhat/deploy

 Open http://localhost:3000 to see the app
 