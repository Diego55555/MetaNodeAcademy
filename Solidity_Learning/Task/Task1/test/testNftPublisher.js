const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

describe("Test NftPublisher", async function () {
    let nftPublisher;
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
        await deployments.fixture(["deployNftPublisher"]);
        const deployNftPublisher = await deployments.get("deployNftPublisher");
        console.log("合约部署地址:", deployNftPublisher.address);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        nftPublisher = await ethers.getContractAt(
            "NftPublisher", 
            deployNftPublisher.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await nftPublisher.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("铸造一个NFT给deployer", async function () {
        const tokenURI = "https://ipfs.io/ipfs/QmYueiuRNmL4MiA2GwtVMm6ZagknXnSpQnB3z2gWbz36hP";
        
        console.log("正在为账户铸造NFT...");
        console.log("接收账户:", deployer);
        console.log("Token URI:", tokenURI);
        
        const tx = await nftPublisher.mintNFT(deployer, tokenURI);
        await tx.wait();

        console.log("✅ NFT铸造成功！交易哈希:", tx.hash);

        // 验证NFT所有权
        const tokenId = 1; // 第一个铸造的NFT tokenId 应该是1
        const owner = await nftPublisher.ownerOf(tokenId);
        expect(owner).to.equal(deployer);
        console.log("✅ NFT所有者验证通过:", owner);
        
        // 验证Token URI
        const retrievedURI = await nftPublisher.tokenURI(tokenId);
        expect(retrievedURI).to.equal(tokenURI);
        console.log("✅ Token URI验证通过:", retrievedURI);
    });
});  