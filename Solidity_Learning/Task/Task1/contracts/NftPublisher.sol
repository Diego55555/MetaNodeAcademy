// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;
// 导入OpenZeppelin的安全合约库
import "@openzeppelin/contracts/access/Ownable.sol";
// 导入OpenZeppelin的ERC-721基础实现，包含元数据扩展
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract NftPublisher is ERC721URIStorage, Ownable {
    // 计数器，用于生成唯一的Token ID
    uint256 private _tokenIds;

    constructor() ERC721("Diego's NFT", "DIG") Ownable(msg.sender) {}

    // 铸造函数：仅合约所有者可调用，为用户铸造一个带有元数据链接的新NFT
    // 测试tokenURI：https://ipfs.io/ipfs/QmYueiuRNmL4MiA2GwtVMm6ZagknXnSpQnB3z2gWbz36hP
    function mintNFT(address recipient, string memory tokenURI) 
        public onlyOwner returns (uint256) {
        _tokenIds++; // 递增ID
        uint256 newItemId = _tokenIds;

        // 核心铸造操作
        _safeMint(recipient, newItemId);
        // 将Token ID与元数据URI（如图片链接）绑定
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }
}