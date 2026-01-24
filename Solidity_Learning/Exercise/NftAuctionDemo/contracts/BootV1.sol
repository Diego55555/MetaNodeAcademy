// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "hardhat/console.sol";

contract BootV1 is Initializable, UUPSUpgradeable {
    // 序号
    uint256 public index;
    // 管理员地址
    address public admin;

    function initialize() public initializer {
        admin = msg.sender;
    }

    function _authorizeUpgrade(address) internal view override {
        // 只有管理员可以升级合约
        require(msg.sender == admin, "Only admin can upgrade");
    }

    // 打招呼
    function sayHello() public virtual {
        index++;
        console.log("Hello, BootV1! Index:", index);
    }
}