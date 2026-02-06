// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "./NftAuction.sol";

contract NftAuctionV2 is NftAuction {
    function getVersion() public pure override returns (string memory) {
        return "2.0.0";
    }
}