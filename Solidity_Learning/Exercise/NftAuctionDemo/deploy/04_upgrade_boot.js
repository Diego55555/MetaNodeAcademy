const {ethers, upgrades} = require("hardhat");
const fs = require("fs");
const path = require("path");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {deployer} = await getNamedAccounts();
  console.log("部署用户地址：", deployer);
  
  //读取部署配置文件
  const storePath = path.resolve(__dirname, "./.cache/proxyBoot.json");
  const proxyData = JSON.parse(fs.readFileSync(storePath, "utf-8"));
  const {proxyAddress, implAddress, abi} = proxyData;

  //升级合约
  console.log("升级前的代理合约地址:", proxyAddress);
  const bootFactoryV2 = await ethers.getContractFactory("BootV2");
  const bootProxyV2 = await upgrades.upgradeProxy(proxyAddress, bootFactoryV2);
  await bootProxyV2.waitForDeployment();
  const proxyAddressV2 = await bootProxyV2.getAddress();

  console.log("升级后的代理合约地址:", proxyAddressV2); 

  await deployments.save("BootProxyV2", {
    address: proxyAddressV2,
    abi: bootFactoryV2.interface.format("json"),
  });
}
module.exports.tags = ["upgradeBoot"];