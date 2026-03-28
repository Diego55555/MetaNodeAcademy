const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();
  console.log("部署用户地址：", deployer);

  //通过代理部署合约
  const simpleERC20Factory = await ethers.getContractFactory("SimpleERC20");
  const simpleERC20Contract = await simpleERC20Factory.deploy(1000000);
  await simpleERC20Contract.waitForDeployment();
  const contractAddress = await simpleERC20Contract.getAddress()

  console.log("合约部署地址:", contractAddress);


  const storePath = path.resolve(__dirname, "./.cache/proxysimpleERC20.json");
  await fs.writeFileSync(storePath, JSON.stringify({
    contractAddress,
    abi: simpleERC20Factory.interface.format("json"),  
  }));

  await deployments.save("simpleERC20", {
    address: contractAddress,
    abi: simpleERC20Factory.interface.format("json"),
  });
}
module.exports.tags = ["deploySimpleERC20"];