const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();

  //直接部署合约
  const deployerSigner = await ethers.getSigner(deployer);
  const simpleERC20Factory = await ethers.getContractFactory(
    "SimpleERC20", 
    deployerSigner        //可选，不提供则使用hardhat.config.js中配置的第一个账户
  );
  const simpleERC20Contract = await simpleERC20Factory.deploy(10000);
  await simpleERC20Contract.waitForDeployment();
  const contractAddress = await simpleERC20Contract.getAddress();

  // 保存合约地址和ABI到本地文件
  const storePath = path.resolve(__dirname, "./.cache/deploySimpleERC20.json");
  await fs.writeFileSync(storePath, JSON.stringify({
    address: contractAddress,
    abi: simpleERC20Factory.interface.format("json"),  
  }));

  await deployments.save("deploySimpleERC20", {
    address: contractAddress,
    abi: simpleERC20Factory.interface.format("json"),
  });
}
module.exports.tags = ["deploySimpleERC20"];