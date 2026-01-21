const {ethers, deplo, deployments} = require("hardhat");
const {expect} = require("chai");

describe("Test upgrade", async function () {
    it("Upgrade", async function () {
        //1.部署业务合约
        await deployments.fixture(["deployNftAuction"]);
        console.log("1");
        const nftAuctionProxy = await deployments.get("NftAuctionProxy");
        console.log("2");
        expect(nftAuctionProxy).to.exist;

        //2.创建拍卖
        const nftAuction = await ethers.getContractAt("NftAuction", nftAuctionProxy.address);
        await nftAuction.createAuction(
            100 * 1000,
            ethers.parseEther("0.01"),
            ethers.zeroAddress,
            1
        );

        const auctionInfo = await nftAuction.getAuction(0);
        console.log("创建拍卖成功：", auctionInfo);
        
        //3.升级合约
        await deployments.fixture(["upgradeNftAuction"]);
        
        //4.读取拍卖信息
        const auctionInfo2 = await nftAuction.getAuction(0);
        console.log("升级后读取拍卖信息成功：", auctionInfo2);

        expect(auctionInfo2.startTime).to.equal(auctionInfo.startTime);
    });
});  