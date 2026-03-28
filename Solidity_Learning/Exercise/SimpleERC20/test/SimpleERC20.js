const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

describe("Test SimpleERC20", async function () {
    let simpleERC20;
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
        await deployments.fixture(["deploySimpleERC20"]);
        const deploySimpleERC20 = await deployments.get("deploySimpleERC20");
        console.log("合约部署地址:", deploySimpleERC20.address);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        simpleERC20 = await ethers.getContractAt(
            "SimpleERC20", 
            deploySimpleERC20.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await simpleERC20.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("转账", async function () {
        const tx = await simpleERC20.transfer(user1, 300);
        await tx.wait();

        const balance1 = await simpleERC20.balanceOf(deployer);
        expect(balance1).to.equal(9700);
        const balance2 = await simpleERC20.balanceOf(user1);
        expect(balance2).to.equal(300);
        console.log("✅ 转账300个:");
    });
});  