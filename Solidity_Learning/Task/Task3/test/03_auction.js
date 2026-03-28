const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");
const fs = require("fs");
const path = require("path");

describe("Test Auction", async function () {
    let nftAuction;
    let deployer;
    let user1;
    let user2;
    let nftPublisher;
    let nftAddress;
    let nftTokenId;
    let auctionId;

    before(async function () {
         // 获取命名账户
        const accounts = await getNamedAccounts();
        deployer = accounts.deployer;
        user1 = accounts.user1;
        user2 = accounts.user2;
        console.log("部署者地址:", deployer);
        console.log("用户1地址:", user1);
        console.log("用户2地址:", user2);

        //读取部署配置文件
        const inputStorePath = path.resolve(__dirname, "../deploy/.cache/deployNftAuction.json");
        const inputStoreData = JSON.parse(fs.readFileSync(inputStorePath, "utf-8"));
        const proxyAddress = inputStoreData.proxyAddress;

        console.log("拍卖合约代理地址:", proxyAddress);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        const nftAuctionFactory = await ethers.getContractFactory("NftAuction");
        nftAuction = nftAuctionFactory.attach(proxyAddress).connect(deployerSigner);
        // nftAuction = await ethers.getContractAt(
        //     "NftAuction",
        //     proxyAddress,
        //     deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        // );

        // 验证实现合约地址
        const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress);
        console.log("拍卖合约逻辑地址:", implAddress);

        //读取部署配置文件
        const inputStorePath2 = path.resolve(__dirname, "../deploy/.cache/deployNftPublisher.json");
        const inputStoreData2 = JSON.parse(fs.readFileSync(inputStorePath2, "utf-8"));
        const proxyAddress2 = inputStoreData2.address;

        console.log("NFT合约地址:", proxyAddress2);

        //获取合约实例
        nftPublisher = await ethers.getContractAt(
            "NftPublisher",
            proxyAddress2,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );

        nftAddress = proxyAddress2;
        nftTokenId = 1;
    });

    it("创建拍卖", async function () {
        //授权NFT给拍卖合约
        const auctionAddress = await nftAuction.getAddress();
        const approveTx = await nftPublisher.approve(auctionAddress, nftTokenId);
        console.log("授权交易哈希:", approveTx.hash);

        // 等待交易确认
        await approveTx.wait();
        console.log("✅ NFT授权成功");

        // 3. 验证授权
        const approvedAddress = await nftPublisher.getApproved(nftTokenId);
        console.log("验证授权 - 授权地址:", approvedAddress);
        console.log("验证授权 - 目标地址:", auctionAddress);

        //创建拍卖
        const txResponse = await nftAuction.createAuction(300, 10000, nftAddress, nftTokenId);
        const receipt = await txResponse.wait();
        
        // 从事件中获取拍卖ID
        const event = receipt.logs.find(log => {
            try {
                const parsedLog = nftAuction.interface.parseLog(log);
                return parsedLog?.name === "CreateAuctionResponse";
            } catch {
                return false;
            }
        });
        
        auctionId = event?.args.auctionId;
        console.log("✅ 拍卖创建成功,拍卖号：", auctionId);
    });

    it("拍卖流程", async function () {
        // 用户1出价
        const user1Signer = await ethers.getSigner(user1);
        const nftAuctionUser1 = nftAuction.connect(user1Signer);
        const bidAmount1 = 20000;
        const tx1 = await nftAuctionUser1.placeBid(auctionId, bidAmount1, ethers.ZeroAddress,
            {value: bidAmount1});
        // 等待交易确认（默认等 1 个区块确认）,需要修改状态的函数都需要交易确认
        await tx1.wait();
        console.log("✅ 用户1出价成功, 出价金额：20000Wei");

        //打印拍卖信息
        let auctionInfo = await nftAuction.auctions(auctionId);
        console.log("当前拍卖信息:", auctionInfo);

        // 用户2出价
        const user2Signer = await ethers.getSigner(user2);
        const nftAuctionUser2 = nftAuction.connect(user2Signer);
        const bidAmount2 = 30000;
        const tx2 = await nftAuctionUser2.placeBid(auctionId, bidAmount2, ethers.ZeroAddress,
            {value: bidAmount2});
        // 等待交易确认（默认等 1 个区块确认）,需要修改状态的函数都需要交易确认
        await tx2.wait();
        console.log("✅ 用户2出价成功, 出价金额：30000Wei");

        //打印拍卖信息
        auctionInfo = await nftAuction.auctions(auctionId);
        console.log("当前拍卖信息:", auctionInfo);
    });
});  