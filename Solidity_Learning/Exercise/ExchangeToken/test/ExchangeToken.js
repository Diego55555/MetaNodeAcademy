const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

describe("Test ExchangeToken", async function () {
    let exchangeToken;
    let deployer;
    let user1;
    let user2;

    before(async function () {
         // 获取命名账户
        const accounts = await getNamedAccounts();
        deployer = accounts.deployer;
        user1 = accounts.user1;
        user2 = accounts.user2;
        console.log("部署者地址:", deployer);
        console.log("用户1地址:", user1);
        console.log("用户2地址:", user2);

        //部署合约
        await deployments.fixture(["deployExchangeToken"]);
        const deployExchangeToken = await deployments.get("deployExchangeToken");
        console.log("合约部署地址:", deployExchangeToken.address);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        exchangeToken = await ethers.getContractAt(
            "ExchangeToken", 
            deployExchangeToken.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await exchangeToken.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("转换ETH为ETK", async function () {
        const tx = await exchangeToken.exchangeETHToETK({value: ethers.parseEther("0.001")});
        await tx.wait();

        const balance = await exchangeToken.balanceOf(deployer);
        expect(balance).to.equal(BigInt(1e18));
        console.log("✅ 得到1个ETK");
    });
});  