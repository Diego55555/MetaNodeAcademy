// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract BBCoin {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply;
    string private _name;
    string private _symbol;
    address private _owner;

    constructor(string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
        _owner = msg.sender;
    }
    
    //查询余额
    function balanceOf(address owner) public view returns (uint256) {
        return _balances[owner];
    }

    //查询授权
    function allowance(address owner, address spender) public view returns (uint256) {
        return _allowances[owner][spender];
    }

    function _approve(address owner, address spender, uint256 value, bool emitEvent) internal {
        if (owner == address(0)) {
            revert("Invalid owner address");
        }
        if (spender == address(0)) {
            revert("Invalid spender address");
        }

        _allowances[owner][spender] = value;
        if(emitEvent){
            emit Approval(owner, spender, value);
        }
    }

    //授权额度
    function approve(address spender, uint256 value) public {
        _approve(msg.sender, spender, value, true);
    }

    //更新余额
    function _update(address from, address to, uint256 value) internal virtual{
        if (from == address(0)) {
            //增发
            _totalSupply += value;
        } else {
            uint256 fromBalance = _balances[from];
            if (fromBalance < value) {
                revert("Not sufficient funds");
            }
            unchecked {
                _balances[from] = fromBalance - value;
            }
        }

        if (to == address(0)) {
            unchecked {
                //销毁
                _totalSupply -= value;
            }
        } else {
            unchecked {
                _balances[to] += value;
            }
        }

        emit Transfer(from, to, value);
    }

    function _transfer(address from, address to, uint256 value) internal {
        if (from == address(0)) {
            revert("Invalid from address");
        }
        if (to == address(0)) {
            revert("Invalid to address");
        }

        _update(from, to, value);
    }

    //转账
    function transfer(address to, uint256 value) public {
        _transfer(msg.sender, to, value);
    }

    //授权转账
    function transferFrom(address from, address to, uint256 value) public {
        _transfer(from, to, value);
    }

    function mint(uint256 value) public {
        if(msg.sender != _owner){
            revert("Only owner can mint");
        }
        _update(address(0), _owner, value);
    }
}