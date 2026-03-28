// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;
// 导入OpenZeppelin的安全合约库
import "@openzeppelin/contracts/access/Ownable.sol";

contract Begging is Ownable {
    event Donation(address indexed from, uint256 value);

    mapping(address => uint256) private _donations; // 捐赠者地址与捐赠金额的映射
    mapping(address => bool) private _inserted; // 捐赠者地址是否已插入数组的映射
    address[] private _donators; // 捐赠者地址数组
    address[3] private _topDonators; // 前3名捐赠者地址数组

    constructor() Ownable(msg.sender) {}

    // 捐赠
    function donate() public payable {
        if(msg.value <= 0) {
            revert("Donation amount must be greater than 0");
        }

        // 判断当前区块时间是否在2026年2月25日到2026年2月28日之间
        uint256 startTime = 1771948800; // 2026-02-25 00:00:00 UTC
        uint256 endTime = 1772207999;   // 2026-02-28 23:59:59 UTC
        if (block.timestamp < startTime || block.timestamp > endTime) {
            revert("Donations are only accepted between 2026-02-25 and 2026-02-28");
        }

        // 更新捐赠者地址与捐赠金额的映射
        if(_inserted[msg.sender]){
            _donations[msg.sender] += msg.value;
        } else{
            _inserted[msg.sender] = true;
            _donators.push(msg.sender);
            _donations[msg.sender] = msg.value;
        }


        for(uint256 i = 0; i < 3; i++){
            if(_topDonators[i] == address(0)){// 如果当前位置为空，则直接插入
                _topDonators[i] = msg.sender;
                break;
            }

            if(_donations[_topDonators[i]] < _donations[msg.sender]){
                // 在i的位置插入msg.sender到_topHolders，将后面的整体向后移一位
                for(uint256 j = 2; j > i; j--){
                    _topDonators[j] = _topDonators[j - 1];
                }
                _topDonators[i] = msg.sender;
                // 删除最后一个元素：这里实际上等价于把排名靠后的“挤出”了
                break;
            }
        }
        
        emit Donation(msg.sender, msg.value);
    }

    // 提取
    function withdraw() public onlyOwner payable{
        payable(owner()).transfer(address(this).balance);
    }

    // 查询捐赠者地址的捐赠金额
    function getDonation(address holder) public view returns (uint256){
        return _donations[holder];
    }

    // 查询前3名捐赠者地址与捐赠金额
    function getTopDonators() public view returns (address[] memory donators, uint256[] memory values){
        uint256 count;
        for (uint256 i = 0; i < 3; i++) {
            if (_topDonators[i] == address(0)) {
                break;
            }
            count++;
        }

        donators = new address[](count);
        values = new uint256[](count);

        for (uint256 i = 0; i < count; i++) {
            address donator = _topDonators[i];
            donators[i] = donator;
            values[i] = _donations[donator];
        }
    }
}