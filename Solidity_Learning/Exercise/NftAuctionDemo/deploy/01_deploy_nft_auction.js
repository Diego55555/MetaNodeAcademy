const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {save} = deployments;
  const {deployer} = await getNamedAccounts();
  console.log("部署用户地址：", deployer);

  //通过代理部署合约
  const nftAuctionFactory = await ethers.getContractFactory("NftAuction");
  const nftAuctionProxy = await upgrades.deployProxy(nftAuctionFactory, [], {
    initializer: "initialize",
  });
  await nftAuctionProxy.waitForDeployment();

  const proxyAddress = await nftAuctionProxy.getAddress()
  console.log("代理合约地址:", proxyAddress);
  const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log("实现合约地址:", implAddress);

  const storePath = path.resolve(__dirname, "./.cache/proxyNftAuction.json");

  fs.writeFileSync(storePath, JSON.stringify({
    proxyAddress,
    implAddress,
    abi: nftAuctionFactory.interface.format("json"),  
  }));

  await save("NftAuctionProxy", {
    address: proxyAddress,
    abi: nftAuctionFactory.interface.format("json"),
  });
}
module.exports.tags = ["deployNftAuction"];