const {deployments, ethers} = require("hardhat");
const fs = require("fs");
const path = require("path");

// 这一行的意思是将一个异步函数导出为模块的接口，用于部署脚本。
// 该函数接收一个包含 getNamedAccounts 和 deployments 的对象作为参数，
// 这样 hardhat 运行部署脚本时会自动传入相关工具和账户信息。
module.exports = async ({getNamedAccounts, deployments}) => {
    //获取账户
    const {deployer} = await getNamedAccounts();
    const deployerSigner = await ethers.getSigner(deployer);

    const BBCoinFactory = await ethers.getContractFactory("BBCoin");
    const BBCoin = await BBCoinFactory.connect(deployerSigner).deploy("BBCoin", "BBC");
    await BBCoin.waitForDeployment();
    const contractAddress = await BBCoin.getAddress();

    // 保存合约地址和ABI到本地文件
    const storePath = path.resolve(__dirname, "./.cache/deployBBCoin.json");
    await fs.writeFileSync(storePath, JSON.stringify({
        address: contractAddress,
        abi: BBCoinFactory.interface.format("json"),  
    }));

    await deployments.save("deployBBCoin", {
        address: contractAddress,
        abi: BBCoinFactory.interface.format("json"),
    });
}
module.exports.tags = ["deployBBCoin"];