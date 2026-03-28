require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");
require('@openzeppelin/hardhat-upgrades');
require("dotenv").config();

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
  networks:{
    sepolia:{
      url:process.env.SEPOLIA_RPC_URL,
      accounts:[
        process.env.SEPOLIA_PRIVATE_KEY1,
        process.env.SEPOLIA_PRIVATE_KEY2,
        process.env.SEPOLIA_PRIVATE_KEY3
      ]
    }
  },
  namedAccounts: {
    deployer: {
      default: 0,           // 本地网络使用第一个账户
      sepolia: 0,           // sepolia 网络使用 accounts 数组的第一个账户（私钥1）
    },
    user1: {
      default: 1,           // 本地网络使用第二个账户
      sepolia: 1,           // sepolia 网络使用 accounts 数组的第二个账户（私钥2）
    },
    user2: {
      default: 2,           // 本地网络使用第三个账户
      sepolia: 2,           // sepolia 网络使用 accounts 数组的第三个账户（私钥3）
    },
  },
};
