const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();
  console.log("部署用户地址：", deployer);

  //直接部署合约
  const nftPublisherFactory = await ethers.getContractFactory("NftPublisher");
  const nftPublisherContract = await nftPublisherFactory.deploy();
  await nftPublisherContract.waitForDeployment();
  const contractAddress = await nftPublisherContract.getAddress()

  console.log("合约部署地址:", contractAddress);
  
  // 保存合约地址和ABI到本地文件
  const storePath = path.resolve(__dirname, "./.cache/NftPublisher.json");
  await fs.writeFileSync(storePath, JSON.stringify({
    contractAddress,
    abi: nftPublisherFactory.interface.format("json"),  
  }));

  await deployments.save("deployNftPublisher", {
    address: contractAddress,
    abi: nftPublisherFactory.interface.format("json"),
  });
}
module.exports.tags = ["deployNftPublisher"];