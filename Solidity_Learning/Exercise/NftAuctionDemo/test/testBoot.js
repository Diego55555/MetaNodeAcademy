const {ethers, deplo, deployments} = require("hardhat");
const {expect} = require("chai");

describe("Test boot upgrade", async function () {
    it("Boot  Upgrade", async function () {
        //1.部署业务合约
        await deployments.fixture(["deployBoot"]);
        const bootProxy = await deployments.get("BootProxy");
        console.log("bootProxy", bootProxy);

        //2.调用合约hello方法
        const bootV1 = await ethers.getContractAt("BootV1", bootProxy.address);
        await bootV1.sayHello();

        //3.升级合约
        const implAddress = await upgrades.erc1967.getImplementationAddress(bootProxy.address);
        console.log("升级前实现合约地址：", implAddress);
        await deployments.fixture(["upgradeBoot"]);
        const implAddress2 = await upgrades.erc1967.getImplementationAddress(bootProxy.address);
        console.log("升级后实现合约地址：", implAddress2);
        
        //4.调用合约hello方法
        const bootV2 = await ethers.getContractAt("BootV2", bootProxy.address);
        await bootV2.sayHello();
    });
});  