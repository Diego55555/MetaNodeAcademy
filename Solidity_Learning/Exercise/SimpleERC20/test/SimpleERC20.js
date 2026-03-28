const hre = require("hardhat");
const {expect} = require("chai");

describe("MyToken2", async() => { 
    const {ethers} = hre
    const initialSupply = 10000
    let MyTokenConstract
    let account1, account2

    beforeEach(async() => {
        [account1, account2] = await ethers.getSigners()
        console.log("测试账户1：", account1);
        console.log("测试账户2：", account2);

        const MyToken = await ethers.getContractFactory("MyToken");
        MyTokenConstract = await MyToken.deploy(initialSupply);
        await MyTokenConstract.waitForDeployment();

        const contractAddress = await MyTokenConstract.getAddress();
        expect(contractAddress).to.be.properAddress; 
        console.log("MyToken 部署地址：", contractAddress);
    })

    it("验证合约的name、symbol、decimals", async() => {
        const name = await MyTokenConstract.name();
        const symbol = await MyTokenConstract.symbol();
        const decimals = await MyTokenConstract.decimals();

        expect(name).to.equal("MyToken");
        expect(symbol).to.equal("MTK");
        expect(decimals).to.equal(18);

        console.log(decimals);
    })

    it("测试转账", async() => {
        const balanceOfAccount1 = await MyTokenConstract.balanceOf(account1)
        expect(balanceOfAccount1).to.equal(initialSupply)

        const resp = await MyTokenConstract.transfer(account2, 1000)
        console.log(resp);

        const balanceOfAccount2 = await MyTokenConstract.balanceOf(account2)
        expect(balanceOfAccount2).to.equal(1000)
    })
})