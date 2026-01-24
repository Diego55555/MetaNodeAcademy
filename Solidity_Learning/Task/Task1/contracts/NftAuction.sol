// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "hardhat/console.sol";


contract NftAuction is Initializable, UUPSUpgradeable {
    // 结构体
    struct Auction {
        // 卖家
        address seller;
        // 拍卖持续时间
        uint256 duration;
        // 起始价格
        uint256 startPrice;
        // 开始时间
        uint256 startTime;
        // 是否结束
        bool ended;
        // 最高出价者
        address highestBidder;
        // 最高价格
        uint256 highestBid;
        // NFT合约地址
        address nftContract;
        // NFT ID
        uint256 tokenId;
        // 参与竞价的资产类型 0x 地址表示eth，其他地址表示erc20
        address tokenAddress;
    }

    // 状态变量
    mapping(uint256 => Auction) public auctions;
    // 下一个拍卖ID
    uint256 public nextAuctionId;
    // 管理员地址
    address public admin;
    // 价格预言机映射
    mapping(address => AggregatorV3Interface) public priceFeeds;

    function initialize() public initializer {
        admin = msg.sender;
        setPriceFeed(address(0), 0x694AA1769357215DE4FAC081bf1f309aDC325306); // ETH/USD
        setPriceFeed(0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238, 0xA2F78ab2355fe2f984D808B5CeE7FD0A93D5270E); // USDC/USD
    }

    function _authorizeUpgrade(address) internal view override {
        // 只有管理员可以升级合约
        require(msg.sender == admin, "Only admin can upgrade");
    }

    function setPriceFeed(address tokenAddress, address _priceFeed) public {
        priceFeeds[tokenAddress] = AggregatorV3Interface(_priceFeed);
    }

    function getChainlinkDataFeedLatestAnswer(address tokenAddress) public view returns (int) {
        AggregatorV3Interface priceFeed = priceFeeds[tokenAddress];
        (   ,
            int256 answer,
            ,
            ,
        ) = priceFeed.latestRoundData();
        return answer;
    }
}