const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

// 这一句的作用是定义一个测试套件，名称为 "Test BBCoin"，用于对BBCoin合约进行单元测试。
// describe 是 Mocha 测试框架用来组织和分组相关测试用例的函数。
// async function () { ... } 里面会写具体的测试内容（例如部署合约、验证功能等）。
describe("Test BBCoin", async function () {
    let deployAddress;
    let BBCoin;
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
        await deployments.fixture(["deployBBCoin"]);
        const deployBBCoin = await deployments.get("deployBBCoin");
        deployAddress = deployBBCoin.address;
        console.log("合约部署地址:", deployAddress);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        BBCoin = await ethers.getContractAt(
            "BBCoin", 
            deployBBCoin.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await BBCoin.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("增发代币", async function () {
        const tx = await BBCoin.mint(10000);
         // 实时交易使用的是Pending块，等待交易打包并确认，再查余额
        await tx.wait();

        // 只读操作使用latest块，所以需要等待交易打包完成，生成新块，才能查到增发代币后的余额
        const balance = await BBCoin.balanceOf(deployer);
        expect(balance).to.equal(10000);
        console.log("✅ 增发代币10000个:");
    });

    it("转账", async function () {
        const tx = await BBCoin.transfer(user1, 300);
        await tx.wait();

        const balance1 = await BBCoin.balanceOf(deployer);
        expect(balance1).to.equal(9700);
        const balance2 = await BBCoin.balanceOf(user1);
        expect(balance2).to.equal(300);
        console.log("✅ 转账300个:");
    });

    it("授权和代扣转账", async function () {
        //获取合约实例
        const BBCoinFactory = await ethers.getContractFactory("BBCoin");
        const user1Signer = await ethers.getSigner(user1);
        const user1Contract = await BBCoinFactory.attach(deployAddress).connect(user1Signer);
        await user1Contract.approve(user2, 100);

        const user2Signer = await ethers.getSigner(user2);
        const user2Contract = await BBCoinFactory.attach(deployAddress).connect(user2Signer);
        const tx = await user2Contract.transferFrom(user1, user2, 100);
        await tx.wait();

        const balance1 = await BBCoin.balanceOf(deployer);
        expect(balance1).to.equal(9700);
        const balance2 = await BBCoin.balanceOf(user1);
        expect(balance2).to.equal(200);
        const balance3 = await BBCoin.balanceOf(user2);
        expect(balance3).to.equal(100);
        console.log("✅ 授权和代扣转账");
    });
});