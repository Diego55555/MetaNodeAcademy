const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

describe("Test NftAuction", async function () {
    let nftAuction;
    let nftAuctionV2;
    let deployer;

    before(async function () {
         // 获取命名账户
        const accounts = await getNamedAccounts();
        deployer = accounts.deployer;
        console.log("部署者地址:", deployer);

        //部署合约
        await deployments.fixture(["deployNftAuction"]);
        const deployNftAuction = await deployments.get("deployNftAuction");
        console.log("代理合约地址:", deployNftAuction.address);
        const implAddress = await upgrades.erc1967.getImplementationAddress(deployNftAuction.address);
        console.log("实现合约地址:", implAddress);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        nftAuction = await ethers.getContractAt(
            "NftAuction", 
            deployNftAuction.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await nftAuction.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("合约升级", async function () {
        //升级前检查版本
        expect(await nftAuction.getVersion()).to.equal("1.0.0");
        console.log("当前合约版本:", await nftAuction.getVersion());

        await deployments.fixture(["updateNftAuctionV2"]);
        const updateNftAuctionV2 = await deployments.get("updateNftAuctionV2");
        console.log("代理合约地址V2:", updateNftAuctionV2.address);
        const implAddressV2 = await upgrades.erc1967.getImplementationAddress(updateNftAuctionV2.address);
        console.log("实现合约地址V2:", implAddressV2);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        nftAuctionV2 = await ethers.getContractAt(
            "NftAuctionV2", 
            updateNftAuctionV2.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );

        expect(await nftAuctionV2.getVersion()).to.equal("2.0.0");
        console.log("✅ 合约升级成功，版本:", await nftAuctionV2.getVersion());
    });
});  