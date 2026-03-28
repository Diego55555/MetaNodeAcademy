// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract Voting{
    //得票数
    mapping(address => uint16) public votesReceived;
    
    //候选人地址数组（用于遍历mapping）
    address[] public candidates;
    
    //标记地址是否已经是候选人
    mapping(address => bool) public isCandidate;
    
    //合约所有者
    address public owner;
    
    constructor() {
        owner = msg.sender;
    }
    
    // 这是Solidity中的修饰符（modifier）语法，用于复用检查条件的逻辑。
    // modifier onlyOwner() 表示定义了一个名为 onlyOwner 的修饰符，用于限制某些函数只有合约拥有者(owner)可以调用。
    // require(...) 语句判断调用者是否为owner，如果不是则会revert并显示错误信息。
    // 下划线 _; 表示被修饰的函数的主体会插入在此位置。
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    //投票
    function vote(address candidate) public {
        // 如果是新候选人，添加到数组
        if (!isCandidate[candidate]) {
            candidates.push(candidate);
            isCandidate[candidate] = true;
        }
        votesReceived[candidate] += 1;
    }

    //获取得票数
    // 如果当前candidate不存在（即尚未获得任何投票），mapping会自动返回默认值0，不会报错
    function getVotes(address candidate) public view returns (uint16) {
        return votesReceived[candidate]; // 不存在的候选人返回0
    }
    
    //获取所有候选人
    function getAllCandidates() public view returns (address[] memory) {
        return candidates;
    }
    
    //获取所有候选人及其得票数
    function getAllVotes() public view returns (address[] memory, uint16[] memory) {
        uint16[] memory votes = new uint16[](candidates.length);
        for (uint i = 0; i < candidates.length; i++) {
            votes[i] = votesReceived[candidates[i]];
        }
        return (candidates, votes);
    }

    //重置所有候选人得票数
    function resetVotes() public onlyOwner {
        for (uint i = 0; i < candidates.length; i++) {
            votesReceived[candidates[i]] = 0;
        }
    }

    //反转字符串
    function reverseString(string memory str) pure public returns (string memory) {
        bytes memory strBytes = bytes(str);
        uint len = strBytes.length;
        for (uint i = 0; i < len / 2; i++) {
            bytes1 temp = strBytes[i];
            strBytes[i] = strBytes[len - 1 - i];
            strBytes[len - 1 - i] = temp;
        }
        
        return string(strBytes);
    }

    //整数转罗马数字
    function uintToRoman(uint16 num) pure public returns (string memory) {
        require(num < 4000, "Number must less than 4000");
        string[13] memory romans = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        uint16[13] memory nums = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];

        string memory roman = "";
        for (uint i = 0; i < nums.length; i++) {
            while (num >= nums[i]) {
                num -= nums[i];
                roman = string.concat(roman, romans[i]);
            }
        }
        return roman;
    }

    //罗马数字转整数
    function romanToUint(string memory roman) pure public returns (uint16) {
        string[13] memory romans = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"];
        uint16[13] memory nums = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];

        uint16 num = 0;
        bytes memory romanBytes = bytes(roman);
        uint offset = 0;
        while (offset < romanBytes.length) {
            bool matched = false;
            for (uint i = 0; i < romans.length; i++) {
                bytes memory sym = bytes(romans[i]);
                if (romanBytes.length - offset >= sym.length) {
                    bool equal = true;
                    for (uint j = 0; j < sym.length; j++) {
                        if (romanBytes[offset + j] != sym[j]) {
                            equal = false;
                            break;
                        }
                    }
                    if (equal) {
                        num += nums[i];
                        offset += sym.length;
                        matched = true;
                        break;
                    }
                }
            }
            // 如果都没匹配上, 说明出现无效字符，不是有效罗马数字
            if (!matched) {
                revert("Invalid Roman numeral input");
            }
        }
        return num;
    }

    //合并两个有序数组
    function mergeSortedArray(uint16[] memory nums1, uint16[] memory nums2) pure public returns(uint16[] memory) {
        uint256 length1 = nums1.length;
        uint256 length2 = nums2.length;
        uint256 offset1 = 0;
        uint256 offset2 = 0;
        uint256 count = length1 + length2;
        uint16[] memory numsTotal = new uint16[](count);

        for(uint256 i = 0; i < count; i++){
            if(offset1 >= length1){
                numsTotal[i] = nums2[offset2];
                offset2++;
                continue;
            }

            if(offset2 >= length2){
                numsTotal[i] = nums1[offset1];
                offset1++;
                continue;
            }

            if(nums1[offset1] <= nums2[offset2]){
                numsTotal[i] = nums1[offset1];
                offset1++;
            }
            else{
                numsTotal[i] = nums2[offset2];
                offset2++;
            }
        }

        return numsTotal;
    }

    //二分查找 
    function binarySearch(uint16[] memory nums, uint16 num) public pure returns(bool, uint256){
        uint256 positionStart = 0;
        uint256 positionEnd = nums.length - 1;
        while(positionStart < positionEnd){
            uint256 middle = (positionStart + positionEnd) / 2;
            if(num == nums[middle]){
                return (true, middle);
            }

            if(num < nums[middle]){
                positionEnd = middle - 1;
            }

            if(num > nums[middle]){
                positionStart = middle + 1;
            }
        }

        return (false, 0);
    }
}