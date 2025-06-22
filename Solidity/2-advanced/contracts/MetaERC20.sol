// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

// 题目 1
contract MetaERC20 is IERC20 {
	// State variables
	address private _owner;
	mapping(address => uint256) private _balances;
	mapping(address => mapping(address => uint256)) private _allowances;

	uint256 private _totalSupply;
	string private _name;
	string private _symbol;

	constructor(uint256 totalSupply_, string memory name_, string memory symbol_) {
		_owner = msg.sender;
		_totalSupply = totalSupply_;
		_name = name_;
		_symbol = symbol_;

		_balances[msg.sender] = totalSupply_;
	}

	// 增发代币
	function mint(address to, uint256 amount) external {
		require(msg.sender == _owner, "MetaERC20: mint from non-owner");
		require(to != address(0), "MetaERC20: mint to the zero address");
		require(amount > 0, "MetaERC20: mint amount must be greater than zero");

		_totalSupply += amount;
		_balances[to] += amount;

		emit Transfer(address(0), to, amount);
	}

	// 总供应量
	function totalSupply() external view override returns (uint256) {
		return _totalSupply;
	}

	// 余额
	function balanceOf(address account) external view override returns (uint256) {
		return _balances[account];
	}

	// 转账
	function transfer(address to, uint256 value) external override returns (bool) {
		require(_balances[msg.sender] >= value, "MetaERC20: transfer amount exceeds balance");
		require(to != address(0), "MetaERC20: transfer to the zero address");
		require(value > 0, "MetaERC20: transfer amount must be greater than zero");
		// Update balances
		_balances[msg.sender] -= value;
		_balances[to] += value;

		emit Transfer(msg.sender, to, value);
		return true;
	}

	// 授权额度
	function allowance(address owner, address spender) external view override returns (uint256) {
		return _allowances[owner][spender];
	}

	// 授权
	function approve(address spender, uint256 value) external override returns (bool) {
		require(msg.sender != spender, "MetaERC20: approve to self");
		require(spender != address(0), "MetaERC20: approve to the zero address");
		require(value > 0, "MetaERC20: approve amount must be greater than zero");
		// Update allowance
		_allowances[msg.sender][spender] = value;

		emit Approval(msg.sender, spender, _allowances[msg.sender][spender]);
		return true;
	}

	// 授权转账
	function transferFrom(address from, address to, uint256 value) external override returns (bool) {
		require(from != _owner, "MetaERC20: transfer from the owner");
		require(from != address(0), "MetaERC20: transfer from the zero address");
		require(_balances[from] >= value, "MetaERC20: transfer amount exceeds balance");
		if (from != msg.sender && _allowances[from][msg.sender] != type(uint256).max) {
			require(_allowances[from][msg.sender] >= value, "MetaERC20: transfer amount exceeds allowance");
		}
		require(to != address(0), "MetaERC20: transfer to the zero address");
		require(value > 0, "MetaERC20: transfer amount must be greater than zero");
		// Update balances
		_balances[from] -= value;
		_balances[to] += value;
		// Update allowance
		_allowances[from][msg.sender] -= value;

		emit Transfer(from, to, value);
		return true;
	}
}
