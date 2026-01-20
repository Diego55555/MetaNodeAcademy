const {deployments, upgrades, ethers} = require("hardhat");

module.exports = async ({getNamedAccounts, deployments}) => {
  const {save} = deployments;
  const {deployer} = await getNamedAccounts();

  console.log("部署用户地址：", deployer);
  const nftAuctionFactory = await ethers.getContractFactory("NftAuction");

  //通过代理部署合约
  const nftAuctionProxy = await upgrades.deployProxy(nftAuctionFactory, [], {
    initializer: "initialize",
  });
  await nftAuctionProxy.waitForDeployment();
  console.log("代理合约地址:", await nftAuctionProxy.getAddress());

//   const nftAuction = await deploy("NftAuction", {
//     from: deployer,
//     args: ["Hello"],
//     log: true,
//   });
};
module.exports.tags = ["deployNftAuction"];