const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();

  //直接部署合约
  const deployerSigner = await ethers.getSigner(deployer);
  const nftAuctionFactory = await ethers.getContractFactory(
    "NftAuction", 
    deployerSigner        //可选，不提供则使用hardhat.config.js中配置的第一个账户
  );
  const nftAuctionProxy = await upgrades.deployProxy(nftAuctionFactory, [], {
    initializer: "initialize",
  });
  await nftAuctionProxy.waitForDeployment();
  const proxyAddress = await nftAuctionProxy.getAddress();
  const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);

  // 保存合约地址和ABI到本地文件
  const storePath = path.resolve(__dirname, "./.cache/deployNftAuction.json");
  await fs.writeFileSync(storePath, JSON.stringify({
    proxyAddress:proxyAddress,
    implAddress:implAddress,
    abi: nftAuctionFactory.interface.format("json"),
  }));

  await deployments.save("deployNftAuction", {
    address: proxyAddress,  // 合约地址（必需字段）
    abi: nftAuctionFactory.interface.format("json"),  // 合约ABI
    _custom: {
      proxyAddress:proxyAddress,
      implAddress:implAddress,
    },               // 自定义字段（可选）
  });
}
module.exports.tags = ["deployNftAuction"];