const {ethers, upgrades} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {save} = deployments;
  const {deployer} = await getNamedAccounts();
  console.log("部署用户地址：", deployer);
  
  //读取部署配置文件
  const storePath = path.resolve(__dirname, "./.cache/proxyNftAuction.json");
  const proxyData = JSON.parse(fs.readFileSync(storePath, "utf-8"));
  const {proxyAddress, implAddress, abi} = proxyData;

  //升级合约
  const nftAuctionFactoryV2 = await ethers.getContractFactory("NftAuctionV4");
  const nftAuctionProxyV2 = await upgrades.upgradeProxy(proxyAddress, nftAuctionFactoryV2);
  await nftAuctionProxyV2.waitForDeployment();
  const proxyAddressV2 = await nftAuctionProxyV2.getAddress();

  console.log("升级后的代理合约地址:", proxyAddressV2); 

  await save("NftAuctionProxyV2", {
    address: proxyAddressV2,
    abi: nftAuctionFactoryV2.interface.format("json"),
  });
}
module.exports.tags = ["upgradeNftAuction"];