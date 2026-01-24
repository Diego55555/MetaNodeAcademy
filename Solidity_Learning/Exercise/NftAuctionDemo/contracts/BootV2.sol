// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "./BootV1.sol";

contract BootV2 is BootV1 {
    function sayHello() public override {
        index++;
        console.log("Hello, BootV2! Index:", index);
    }
}