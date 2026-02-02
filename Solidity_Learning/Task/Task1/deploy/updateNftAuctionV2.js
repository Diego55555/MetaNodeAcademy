const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();

  //读取部署配置文件
  const inputStorePath = path.resolve(__dirname, "./.cache/deployNftAuction.json");
  const inputStoreData = JSON.parse(fs.readFileSync(inputStorePath, "utf-8"));
  const {proxyAddress, implAddress} = inputStoreData;

  //升级合约
  const deployerSigner = await ethers.getSigner(deployer);
  const nftAuctionV2Factory = await ethers.getContractFactory("NftAuctionV2", deployerSigner);
  const nftAuctionV2Proxy = await upgrades.upgradeProxy(proxyAddress, nftAuctionV2Factory);
  await nftAuctionV2Proxy.waitForDeployment();
  console.log("**版本号:", await nftAuctionV2Proxy.getVersion());

  //获取升级后的合约实例
  const proxyAddressV2 = await nftAuctionV2Proxy.getAddress();
  const implAddressV2 = await upgrades.erc1967.getImplementationAddress(proxyAddressV2);

  // 保存合约地址和ABI到本地文件
  const storePath = path.resolve(__dirname, "./.cache/updateNftAuctionV2.json");
  await fs.writeFileSync(storePath, JSON.stringify({
    proxyAddress: proxyAddressV2,
    implAddress: implAddressV2,
    abi: nftAuctionV2Factory.interface.format("json"),
  }));

  await deployments.save("updateNftAuctionV2", {
    address: proxyAddressV2,
    abi: nftAuctionV2Factory.interface.format("json"),
  });
}
module.exports.tags = ["updateNftAuctionV2"];