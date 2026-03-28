const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();

  //直接部署合约
  const deployerSigner = await ethers.getSigner(deployer);
  const nftPublisherFactory = await ethers.getContractFactory(
    "NftPublisher", 
    deployerSigner        //可选，不提供则使用hardhat.config.js中配置的第一个账户
  );
  const nftPublisherContract = await nftPublisherFactory.deploy();
  await nftPublisherContract.waitForDeployment();
  const contractAddress = await nftPublisherContract.getAddress();

  // 保存合约地址和ABI到本地文件
  const storePath = path.resolve(__dirname, "./.cache/deployNftPublisher.json");
  await fs.writeFileSync(storePath, JSON.stringify({
    address: contractAddress,
    abi: nftPublisherFactory.interface.format("json"),  
  }));

  await deployments.save("deployNftPublisher", {
    address: contractAddress,
    abi: nftPublisherFactory.interface.format("json"),
  });
}
module.exports.tags = ["deployNftPublisher"];