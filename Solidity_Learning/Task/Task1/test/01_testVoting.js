const {ethers, deployments} = require("hardhat");
const {expect} = require("chai");

// 这一句的作用是定义一个测试套件，名称为 "Test Voting"，用于对Voting合约进行单元测试。
// describe 是 Mocha 测试框架用来组织和分组相关测试用例的函数。
// async function () { ... } 里面会写具体的测试内容（例如部署合约、验证功能等）。
describe("Test Voting", async function () {
    let voting;
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
        // 1. await deployments.fixture(["deployVoting"]);
        //    这句会根据 deployment 脚本标签名（这里是 "deployVoting"），自动部署并初始化需要的合约。它确保在测试前目标合约已经正确部署，并且可以重复多次以保证测试的隔离性。
        // 2. const deployVoting = await deployments.get("deployVoting");
        //    这句通过部署管理工具获取名为 "deployVoting" 的已部署合约的信息（例如合约地址、ABI 等），方便后续与合约进行交互。
        await deployments.fixture(["deployVoting"]);
        const deployVoting = await deployments.get("deployVoting");
        console.log("合约部署地址:", deployVoting.address);

        //获取合约实例
        const deployerSigner = await ethers.getSigner(deployer);
        voting = await ethers.getContractAt(
            "Voting", 
            deployVoting.address,
            deployerSigner      //签名者（可选）,不提供则只能读取，不能发送交易
        );
    });

    it("应该正确部署合约", async function () {
        const contractAddress = await voting.getAddress();
        expect(contractAddress).to.be.properAddress;
        console.log("✅ 合约部署成功，地址:", contractAddress);
    });

    it("投票", async function(){
        const candidate1 = user1;
        const candidate2 = user2;
        await voting.vote(candidate1);
        await voting.vote(candidate1);
        const votes1 = await voting.getVotes(candidate1);
        const votes2 = await voting.getVotes(candidate2);
        expect(votes1).to.equal(2);
        expect(votes2).to.equal(0);

        await voting.resetVotes();
        const votes1New = await voting.getVotes(candidate1);
        const votes2New = await voting.getVotes(candidate2);
        expect(votes1New).to.equal(0);
        expect(votes2New).to.equal(0);
    });

    it("反转字符串", async function(){
        const result = await voting.reverseString("hello");
        expect(result).to.equal("olleh");
    });

    it("整数转罗马数字", async function(){
        let result = await voting.uintToRoman(3);
        expect(result).to.equal("III");
        result = await voting.uintToRoman(4);
        expect(result).to.equal("IV");
        result = await voting.uintToRoman(9);
        expect(result).to.equal("IX");
        result = await voting.uintToRoman(58);
        expect(result).to.equal("LVIII");
        result = await voting.uintToRoman(1994);
        expect(result).to.equal("MCMXCIV");
    });

    it("罗马数字转整数", async function(){
        let result = await voting.romanToUint("III");
        expect(result).to.equal(3);
        result = await voting.romanToUint("IV");
        expect(result).to.equal(4);
        result = await voting.romanToUint("IX");
        expect(result).to.equal(9);
        result = await voting.romanToUint("LVIII");
        expect(result).to.equal(58);
        result = await voting.romanToUint("MCMXCIV");
        expect(result).to.equal(1994);
    });

    
    it("合并两个有序数组", async function(){
        let result = await voting.mergeSortedArray([1,2,3], [4,5,6]);
        expect(result.map(x => Number(x))).to.deep.equal([1,2,3,4,5,6]);
        result = await voting.mergeSortedArray([1,3,5], [2,4,6]);
        expect(result.map(x => Number(x))).to.deep.equal([1,2,3,4,5,6]);
    });

    it("二分查找 ", async function(){
        let result = await voting.binarySearch([1,2,3,4,5,6,7,8,9,10], 3);
        expect(result).to.deep.equal([true, 2]);
        result = await voting.binarySearch([1,2,3,4,5,6,7,8,9,10], 11);
        expect(result).to.deep.equal([false, 0] );
    });
});