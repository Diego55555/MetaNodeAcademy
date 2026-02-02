// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

event DebugInfo(string message);

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

    // 创建拍卖
    function createNewAuction(
        uint256 _duration,
        uint256 _startPrice,
        address _nftAddress,
        uint256 _tokenId
    ) public {
        emit DebugInfo("createAuction: called");

        // 只有管理员可以创建拍卖
        require(msg.sender == admin, "Only admin can create auctions");
        // 检查参数
        require(_duration >= 10, "Duration must be greater than 10s");
        require(_startPrice > 0, "Start price must be greater than 0");

        emit DebugInfo("createAuction: checked");

        // 转移NFT到合约
        IERC721(_nftAddress).safeTransferFrom(msg.sender, address(this), _tokenId);

        emit DebugInfo("createAuction: safeTransferFrom");

        auctions[nextAuctionId] = Auction({
            seller: msg.sender,
            duration: _duration,
            startPrice: _startPrice,
            ended: false,
            highestBidder: msg.sender,
            highestBid: 0,
            startTime: block.timestamp,
            nftContract: _nftAddress,
            tokenId: _tokenId,
            tokenAddress: address(0)
        });

        nextAuctionId++;
    }

    // 买家参与拍卖
    function placeBid(
        uint256 _auctionID,
        uint256 amount,
        address _tokenAddress
    ) external payable {
        // 获取拍卖信息
        Auction storage auction = auctions[_auctionID];
        
        // 判断当前拍卖是否结束
        require(
            !auction.ended &&
                auction.startTime + auction.duration > block.timestamp,
            "Auction has ended"
        );

        // 判断出价是否大于当前最高出价，统一的价值尺度到USD
        if (_tokenAddress == address(0)) {
            // 处理 ETH
            amount = msg.value;
        }
        uint payValue = amount * uint(getChainlinkDataFeedLatestAnswer(_tokenAddress));
        
        //获取起拍价格（USD）
        uint startPriceValue = auction.startPrice *
            uint(getChainlinkDataFeedLatestAnswer(auction.tokenAddress));

        //获取当前最高价（USD）
        uint highestBidValue = auction.highestBid *
            uint(getChainlinkDataFeedLatestAnswer(auction.tokenAddress));

        require(
            payValue >= startPriceValue && payValue > highestBidValue,
            "Bid must be higher than the current highest bid"
        );

        // 转移 ERC20 到合约
        if (_tokenAddress != address(0)) {
            IERC20(_tokenAddress).transferFrom(msg.sender, address(this), amount);
        }

        // 退还前最高价
        if (auction.highestBid > 0) {
            if (auction.tokenAddress == address(0)) {
                //退还ETH
                payable(auction.highestBidder).transfer(auction.highestBid);
            } else {
                // 退还ERC20
                IERC20(auction.tokenAddress).transfer(
                    auction.highestBidder,
                    auction.highestBid
                );
            }
        }
        
        auction.tokenAddress = _tokenAddress;
        auction.highestBid = amount;
        auction.highestBidder = msg.sender;
    }

    // 结束拍卖
    function endAuction(uint256 _auctionID) external {
        // 获取拍卖信息
        Auction storage auction = auctions[_auctionID];

        // 判断当前拍卖是否结束
        require(
            !auction.ended &&
                (auction.startTime + auction.duration) <= block.timestamp,
            "Auction has not ended"
        );

        // 转移NFT到最高出价者
        IERC721(auction.nftContract).safeTransferFrom(
            address(this),
            auction.highestBidder,
            auction.tokenId
        );

        // 转移剩余的资金到卖家
        auction.ended = true;

        if (auction.highestBid > 0) {
            if (auction.tokenAddress == address(0)) {
                // 转移ETH
                payable(auction.seller).transfer(auction.highestBid);
            } else {
                // 转移ERC20
                IERC20(auction.tokenAddress).transfer(
                    auction.seller,
                    auction.highestBid
                );
            }
        }
    }

    function onERC721Received(
        address,    //调用 safeTransferFrom 的地址，发起者
        address,    //NFT 拥有者
        uint256,    //NFT 的 Token ID
        bytes calldata  //额外的数据
    ) external pure returns (bytes4) {
        // 必须返回这个固定的 magic value，否则转移会被拒绝
        return this.onERC721Received.selector;
    }

    function getVersion() public pure virtual returns (string memory) {
        return "1.0.0";
    }

    // 测试NFT信息
    function nftTest(
        address _nftAddress,
        uint256 _tokenId
    ) public {
        address owner = IERC721(_nftAddress).ownerOf(_tokenId);
        address approved = IERC721(_nftAddress).getApproved(_tokenId);
    
        emit DebugInfo(string.concat("NFT Owner: ", Strings.toHexString(uint256(uint160(owner)))));
        emit DebugInfo(string.concat("Your address: ", Strings.toHexString(uint256(uint160(msg.sender)))));
        emit DebugInfo(string.concat("Approved address: ", Strings.toHexString(uint256(uint160(approved)))));
        emit DebugInfo(string.concat("Admin address: ", Strings.toHexString(uint256(uint160(admin)))));
    } 

    function fullDiagnosis(address _nftAddress, uint256 _tokenId) public {
        emit DebugInfo(unicode"=== 完整NFT诊断 ===");
        
        // 1. 基础信息
        emit DebugInfo(string.concat(unicode"诊断调用者:", Strings.toHexString(uint256(uint160(msg.sender)))));
        emit DebugInfo(string.concat(unicode"NFT合约地址:", Strings.toHexString(uint256(uint160(_nftAddress)))));
        emit DebugInfo(string.concat(unicode"Token ID:", Strings.toString(_tokenId)));
        emit DebugInfo(string.concat(unicode"本合约地址:", Strings.toHexString(uint256(uint160(address(this))))));
        
        // 2. NFT所有权检查
        try IERC721(_nftAddress).ownerOf(_tokenId) returns (address owner) {
            emit DebugInfo(string.concat(unicode"NFT当前所有者:", Strings.toHexString(uint256(uint160(owner)))));
            emit DebugInfo(string.concat(unicode"调用者是否是所有者?", owner == msg.sender ? "true":"false"));
            
            if (owner != msg.sender) {
                emit DebugInfo(unicode"❌ 问题：你不是NFT所有者！");
                emit DebugInfo(string.concat(unicode"你需要用地址", Strings.toHexString(uint256(uint160(owner))), unicode"来调用此函数"));
                return;
            }
            
            emit DebugInfo(unicode"✅ NFT所有权验证通过");
            
            // 3. 详细授权检查
            emit DebugInfo(unicode"=== 授权状态检查 ===");
            
            // 单一授权
            address approved;
            try IERC721(_nftAddress).getApproved(_tokenId) returns (address _approved) {
                approved = _approved;
                emit DebugInfo(string.concat(unicode"单一授权地址:", Strings.toHexString(uint256(uint160(approved)))));
                emit DebugInfo(string.concat(unicode"本合约是否被单一授权?", approved == address(this) ? "true":"false"));
            } catch {
                emit DebugInfo(unicode"⚠️ getApproved调用失败");
            }
            
            // 全局授权
            bool isApprovedForAll = false;
            try IERC721(_nftAddress).isApprovedForAll(msg.sender, address(this)) returns (bool _isApproved) {
                isApprovedForAll = _isApproved;
                emit DebugInfo(string.concat(unicode"是否有全局授权?", isApprovedForAll ? "true":"false"));
            } catch {
                emit DebugInfo(unicode"⚠️ isApprovedForAll调用失败");
            }
            
            emit DebugInfo(string.concat(unicode"最终授权状态（任一即可）:", approved == address(this) || isApprovedForAll ? "true":"false"));
            
            // 4. 尝试模拟转移
            emit DebugInfo(unicode"=== 模拟转移测试 ===");
            
            if (approved == address(this) || isApprovedForAll) {
                emit DebugInfo(unicode"✅ 授权检查通过，理论上可以转移");
            } else {
                emit DebugInfo(unicode"❌ 授权失败：合约未被授权");
                emit DebugInfo(unicode"解决方案：");
                emit DebugInfo(string.concat(unicode"1. 调用 approve(", Strings.toHexString(uint256(uint160(address(this)))), ", ", Strings.toString(_tokenId), ")"));
                emit DebugInfo(string.concat(unicode"2. 或调用 setApprovalForAll(", Strings.toHexString(uint256(uint160(address(this)))), ", true)"));
            }
            
        } catch Error(string memory reason) {
            emit DebugInfo(string.concat(unicode"❌ ownerOf错误:", reason));
        } catch {
            emit DebugInfo(string.concat(unicode"❌ 未知错误，可能NFT不存在"));
        }
    }
}
