const {ethers, deplo, deployments} = require("hardhat");
const {expect} = require("chai");

describe("Test upgrade", async function () {
    it("Upgrade", async function () {
        //1.部署业务合约
        await deployments.fixture(["deployNftAuction"]);
        const nftAuctionProxy = await deployments.get("NftAuctionProxy");
        console.log("nftAuctionProxy", nftAuctionProxy);

        //2.创建拍卖
        const nftAuction = await ethers.getContractAt("NftAuctionV3", nftAuctionProxy.address);
        console.log("1");
        await nftAuction.createAuction(
            100 * 1000,
            ethers.parseEther("0.01"),
            ethers.zeroAddress,
            1
        );
        console.log("2");

        const auctionInfo = await nftAuction.getAuction(0);
        console.log("创建拍卖成功：", auctionInfo);

        const implAddress = await upgrades.erc1967.getImplementationAddress(nftAuctionProxy.address);
        console.log("升级前实现合约地址：", implAddress);
        
        //3.升级合约
        await deployments.fixture(["upgradeNftAuction"]);
        const implAddress2 = await upgrades.erc1967.getImplementationAddress(nftAuctionProxy.address);
        console.log("升级后实现合约地址：", implAddress2);
        
        //4.读取拍卖信息
        const auctionInfo2 = await nftAuction.getAuction(0);
        console.log("升级后读取拍卖信息成功：", auctionInfo2);

        const nftAuction2 = await ethers.getContractAt("NftAuctionV4", nftAuctionProxy.address);
        const hello = await nftAuction2.hell
        o();
        console.log("调用新合约方法hello():", hello);

        expect(auctionInfo2.startTime).to.equal(auctionInfo.startTime);
    });
});  