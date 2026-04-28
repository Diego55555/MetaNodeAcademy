// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "hardhat/console.sol";

contract ExchangeToken is ERC20 {
    uint256 public exchangeRate = 1000;

    constructor(uint256 initialSupply) ERC20("ExchangeToken", "ETK") {
        console.log("initialSupply:", initialSupply);
        _mint(address(this), initialSupply);
    } 

    function exchangeETHToETK() public payable {
        require(msg.value > 0, "Value must be greater than 0");
        uint256 tokenRemain = balanceOf(address(this));
        uint256 tokenAmount = msg.value * exchangeRate;
        console.log("tokenRemain:", tokenRemain);
        console.log("tokenAmount:", tokenAmount);
        require(tokenRemain >= tokenAmount, "Insufficient ETK");
        _transfer(address(this), msg.sender, tokenAmount);
    }

    function exchangeETKToETH(uint256 tokenAmount) public {
        require(tokenAmount >= 1000, "Amount must be greater than or equal to 1000");
        require(balanceOf(msg.sender) >= tokenAmount, "Insufficient ETK");
        _transfer(msg.sender, address(this), tokenAmount);
        payable(msg.sender).transfer(tokenAmount / exchangeRate);
    }
}