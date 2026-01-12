const hre = require("hardhat");

describe("MyToken Test", async() => {
    beforeEach(async() => {
        console.log("等待 2 s");
        
        await new Promise((resolve) => {
            setTimeout(()=>{
                resolve(1)
            }, 2000)
        });

        console.log("开始运行测试用例");
    })

    it("test1", async() => {
        console.log("我是 test1");
    })

    it("test2", async() => {
        console.log("我是 test2");
    })
})