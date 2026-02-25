const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

// 这一句的作用是定义一个测试套件，名称为 "Test BBCoin"，用于对BBCoin合约进行单元测试。
// describe 是 Mocha 测试框架用来组织和分组相关测试用例的函数。
// async function () { ... } 里面会写具体的测试内容（例如部署合约、验证功能等）。
describe("Test begging", async function () {
    let deployAddress;
    let begging;
    let deployer;
    let user1;
    let user2;

    before(async function () {
         // 获取命名账户
        // getNamedAccounts 是 Hardhat Runtime Environment (HRE) 提供的内置异步函数，用于获取 hardhat.config.js 中命名的账户地址
        const accounts = await getNamedAccounts();
        deployer = accounts.deployer;
        user1 = accounts.user1;
        user2 = accounts.user2;
        console.log("部署者地址:", deployer);
        console.log("用户1地址:", user1);
        console.log("用户2地址:", user2);

        //部署合约
        // 这两行代码的作用如下：
        // 1. await deployments.fixture(["deployBBCoin"]);
        //    这句会根据 deployment 脚本标签名（这里是 "deployBBCoin"），自动部署并初始化需要的合约。它确保在测试前目标合约已经正确部署，并且可以重复多次以保证测试的隔离性。
        // 2. const deployBBCoin = await deployments.get("deployBBCoin");
        //    这句通过部署管理工具获取名为 "deployBBCoin" 的已部署合约的信息（例如合约地址、ABI 等），方便后续与合约进行交互。
        await deployments.fixture(["deployBegging"]);
        const deployBegging = await deployments.get("deployBegging");
        deployAddress = deployBegging.address;
        console.log("合约部署地址:", deployAddress);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        begging = await ethers.getContractAt(
            "Begging", 
            deployBegging.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await begging.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("捐赠并查询捐赠额", async function () {
        // 使用 user1 捐赠
        const user1Amount = ethers.parseEther("0.01"); // 捐赠0.01ETH
        const user1Signer = await ethers.getSigner(user1);
        const tx1 = await begging.connect(user1Signer).donate({ value: user1Amount });
        await tx1.wait();

        // 使用 user2 捐赠
        const user2Amount = ethers.parseEther("0.02"); // 捐赠0.02ETH
        const user2Signer = await ethers.getSigner(user2);
        const tx2 = await begging.connect(user2Signer).donate({ value: user2Amount });
        await tx2.wait();
        
        const donation1 = await begging.getDonation(user1);
        expect(donation1).to.equal(user1Amount);
        const donation2 = await begging.getDonation(user2);
        expect(donation2).to.equal(user2Amount);

        console.log("✅ 捐赠成功，捐赠额:", donation1, donation2);
    });

    it("获取捐赠前3名", async function () {
        const [donators, values] = await begging.getTopDonators();
        expect(donators.length).to.equal(2);
        expect(donators[0]).to.equal(user2);
        expect(donators[1]).to.equal(user1);
        console.log("✅ 获取捐赠前3名成功，地址:", donators, "，金额:", values);
    });

    it("提取", async function () {
        const amount = ethers.parseEther("0.03");
        const oldBalance = await ethers.provider.getBalance(deployAddress);
        expect(oldBalance).to.equal(amount);

        const tx = await begging.withdraw();    
        await tx.wait();
        const newBalance = await ethers.provider.getBalance(deployAddress);
        expect(newBalance).to.equal(0);
        console.log("✅ 提取成功，提取金额:", amount);
    });
});