import { Button } from "antd";
import { TEthersProvider } from "eth-hooks/models";
import { ethers } from "ethers";
import Web3Modal from "web3modal";
import Address from "./Address";
import Balance from "./Balance";

export default function Account(
    useBurner: boolean,
    address: string,
    userSigner: ethers.Signer,
    localProvider: TEthersProvider,
    mainnetProvider: TEthersProvider,
    price: number,
    web3Modal: Web3Modal,
    loadWeb3Modal: React.MouseEventHandler<HTMLElement> | undefined,
    logoutOfWeb3Modal: React.MouseEventHandler<HTMLElement> | undefined,
    blockExplorer: string,
    isContract: boolean,
    minimized? :string //最小化？
    
    ) {

    const modalButtons = [];
    if (web3Modal) {
        if (web3Modal.cachedProvider) { //缓存provider
            modalButtons.push(
                <Button
                  key="logoutbutton"
                  style={{ verticalAlign: "top", marginLeft: 8, marginTop: 4 }}
                  shape="round"
                  size="large"
                  onClick={logoutOfWeb3Modal}
                >
                  logout
                </Button>,
              );
        }else {
            modalButtons.push(
                <Button
                  key="loginbutton"
                  style={{ verticalAlign: "top", marginLeft: 8, marginTop: 4 }}
                  shape="round"
                  size="large"
                  /* type={minimized ? "default" : "primary"}     too many people just defaulting to MM and having a bad time */
                  onClick={loadWeb3Modal}
                >
                  connect
                </Button>,
              );
        }
    }

    const display = minimized ? (
        ""
      ) : (
        <span>
          {web3Modal && web3Modal.cachedProvider ? (
            <>
              {address && <Address address={address} ensProvider={mainnetProvider} blockExplorer={blockExplorer} />}
              <Balance address={address} provider={localProvider} price={price} />
              {/* <Wallet
                address={address}
                provider={localProvider}
                signer={userSigner}
                ensProvider={mainnetProvider}
                price={price}
                color={currentTheme === "light" ? "#1890ff" : "#2caad9"}
              /> */}
            </>
          ) : useBurner ? (
            ""
          ) : isContract ? (
            <>
              {address && <Address address={address} ensProvider={mainnetProvider} blockExplorer={blockExplorer} />}
              <Balance address={address} provider={localProvider} price={price} />
            </>
          ) : (
            ""
          )}
          {useBurner && web3Modal && !web3Modal.cachedProvider ? (
            <>
              <Address address={address} ensProvider={mainnetProvider} blockExplorer={blockExplorer} />
              <Balance address={address} provider={localProvider} price={price} />
              {/* <Wallet
                address={address}
                provider={localProvider}
                signer={userSigner}
                ensProvider={mainnetProvider}
                price={price}
                color={currentTheme === "light" ? "#1890ff" : "#2caad9"}
              /> */}
            </>
          ) : (
            <></>
          )}
        </span>
      );
    
      return (
        <div>
          {display}
          {modalButtons}
        </div>
      );
    
}