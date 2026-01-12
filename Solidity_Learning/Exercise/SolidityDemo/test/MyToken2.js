const hre = require("hardhat");
const {expect} = require("chai");

describe("MyToken2", async() => { 
    const {ethers} = hre
    const initialSupply = 10000
    let MyTokenConstract

    beforeEach(async() => {
        const MyToken = await ethers.getContractFactory("MyToken");
        MyTokenConstract = await MyToken.deploy(initialSupply);
        await MyTokenConstract.waitForDeployment();

        const contractAddress = await MyTokenConstract.getAddress();
        expect(contractAddress).to.be.properAddress; 
        console.log("MyToken 部署地址：", contractAddress);
    })

    it("test1", async() => {
        console.log("我是 test1");
    })

    it("test2", async() => {
        console.log("我是 test2");
    })
})