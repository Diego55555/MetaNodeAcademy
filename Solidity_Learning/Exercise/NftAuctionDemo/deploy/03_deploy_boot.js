const {deployments, upgrades, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();
  console.log("部署用户地址：", deployer);

  //通过代理部署合约
  const bootFactory = await ethers.getContractFactory("BootV1");
  const bootProxy = await upgrades.deployProxy(bootFactory, [], {
    initializer: "initialize",
  });
  await bootProxy.waitForDeployment();

  const proxyAddress = await bootProxy.getAddress()
  console.log("代理合约地址:", proxyAddress);
  const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
  console.log("实现合约地址:", implAddress);

  const storePath = path.resolve(__dirname, "./.cache/proxyBoot.json");

  await fs.writeFileSync(storePath, JSON.stringify({
    proxyAddress,
    implAddress,
    abi: bootFactory.interface.format("json"),
  }));

  await deployments.save("BootProxy", {
    address: proxyAddress,
    abi: bootFactory.interface.format("json"),
  });
}
module.exports.tags = ["deployBoot"];