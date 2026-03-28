const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

// 这一句的作用是定义一个测试套件，名称为 "Test BBCoin"，用于对BBCoin合约进行单元测试。
// describe 是 Mocha 测试框架用来组织和分组相关测试用例的函数。
// async function () { ... } 里面会写具体的测试内容（例如部署合约、验证功能等）。
describe("Test BBNft", async function () {
    let deployAddress;
    let BBNft;
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
        await deployments.fixture(["deployBBNft"]);
        const deployBBNft = await deployments.get("deployBBNft");
        deployAddress = deployBBNft.address;
        console.log("合约部署地址:", deployAddress);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        BBNft = await ethers.getContractAt(
            "BBNft", 
            deployBBNft.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await BBNft.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("铸造一个NFT给deployer", async function () {
        const tokenURI = "https://ipfs.io/ipfs/QmYueiuRNmL4MiA2GwtVMm6ZagknXnSpQnB3z2gWbz36hP";
        
        console.log("正在为账户铸造NFT...");
        console.log("接收账户:", deployer);
        console.log("Token URI:", tokenURI);
        
        const tx = await BBNft.mintNFT(deployer, tokenURI);
        await tx.wait();

        console.log("✅ NFT铸造成功！交易哈希:", tx.hash);

        // 验证NFT所有权
        const tokenId = 1; // 第一个铸造的NFT tokenId 应该是1
        const owner = await BBNft.ownerOf(tokenId);
        expect(owner).to.equal(deployer);
        console.log("✅ NFT所有者验证通过:", owner);
        
        // 验证Token URI
        const retrievedURI = await BBNft.tokenURI(tokenId);
        expect(retrievedURI).to.equal(tokenURI);
        console.log("✅ Token URI验证通过:", retrievedURI);
    });
});