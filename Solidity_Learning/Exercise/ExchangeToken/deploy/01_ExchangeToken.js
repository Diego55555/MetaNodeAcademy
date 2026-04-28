const {deployments, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();

  //直接部署合约
  const deployerSigner = await ethers.getSigner(deployer);
  const exchangeTokenFactory = await ethers.getContractFactory(
    "ExchangeToken", 
    deployerSigner        //可选，不提供则使用hardhat.config.js中配置的第一个账户
  );
  const totalSupply = BigInt(10000 * 1e18);
  const exchangeTokenContract = await exchangeTokenFactory.deploy(totalSupply);
  await exchangeTokenContract.waitForDeployment();
  const contractAddress = await exchangeTokenContract.getAddress();

  // 保存合约地址和ABI到本地文件
  const storePath = path.resolve(__dirname, "./.cache/deployExchangeToken.json");
  await fs.writeFileSync(storePath, JSON.stringify({
    address: contractAddress,
    abi: exchangeTokenFactory.interface.format("json"),  
  }));

  await deployments.save("deployExchangeToken", {
    address: contractAddress,
    abi: exchangeTokenFactory.interface.format("json"),
  });
}
module.exports.tags = ["deployExchangeToken"];