require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  namedAccounts: {
    deployer: {
      default: 0,           // 本地网络使用第一个账户
    },
    user1: {
      default: 1,           // 本地网络使用第二个账户
    },
    user2: {
      default: 2,           // 本地网络使用第三个账户
    },
  },
};
