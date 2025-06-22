// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts/access/Ownable.sol";

// 题目 3
contract BeggingContract is Ownable {
	// 定义捐赠事件
	event Donation(address indexed donator, uint256 amount);

	// 定义捐赠映射
	mapping(address donator => uint256 amount) public donations;
	// 定义捐赠总额
	uint256 public totalDonations;
	// 定义开始时间
	uint256 public startTime;
	// 定义结束时间
	uint256 public endTime;

	// 定义捐赠者结构体
	struct Donor {
		address donator;
		uint256 amount;
	}

	// 定义Top3捐赠者数组
	Donor[3] private top3Donors;

	constructor(uint256 duration) Ownable(msg.sender) {
		startTime = block.timestamp;
		endTime = block.timestamp + duration;
	}

	// 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠
	modifier onlyDuringDonationPeriod() {
		require(block.timestamp >= startTime && block.timestamp <= endTime, "Donations are not allowed during this period");
		_;
	}

	// 捐赠函数：添加一个捐赠函数，允许用户捐赠以太币
	function donate() public payable onlyDuringDonationPeriod {
		require(msg.value > 0, "Donation amount must be greater than 0");
		donations[msg.sender] += msg.value;
		totalDonations += msg.value;
		_updateTop3Donors(msg.sender);
		emit Donation(msg.sender, msg.value);
	}

	// 提现函数：添加一个提现函数，允许合约所有者提取所有捐赠的以太币
	function withdraw() public onlyOwner {
		uint256 balance = address(this).balance;
		require(balance > 0, "No funds available");
		// 使用 send 或 call 代替 transfer 以避免 gas 不足问题
		(bool success, ) = payable(owner()).call{value: balance}("");
		require(success, "Transfer failed");
		totalDonations = 0;
	}

	// 查询捐赠函数：添加一个查询捐赠函数，允许用户查询自己的捐赠金额
	function getDonation(address donator) public view returns (uint256) {
		return donations[donator];
	}

	// 更新Top3捐赠者函数：添加一个内部函数，用于更新Top3捐赠者数组
	function _updateTop3Donors(address newDonor) private {
		uint256 newDonation = donations[newDonor];
		bool isExisting = false;
		uint256 existingIndex;

		// 检查捐献者是否已经存在于 top3Donors 数组中
		for (uint256 i = 0; i < 3; i++) {
			if (top3Donors[i].donator == newDonor) {
				isExisting = true;
				existingIndex = i;
				break;
			}
		}

		if (isExisting) {
			// 如果存在，更新其捐赠金额
			top3Donors[existingIndex].amount = newDonation;
		} else {
			// 如果不存在，按原逻辑插入
			for (uint256 i = 0; i < 3; i++) {
				if (newDonation > top3Donors[i].amount) {
					// 移动后续元素
					for (uint256 j = 2; j > i; j--) {
						top3Donors[j] = top3Donors[j - 1];
					}
					top3Donors[i] = Donor(newDonor, newDonation);
					break;
				}
			}
		}

		// 重新排序 top3Donors 数组
		for (uint256 i = 0; i < 2; i++) {
			for (uint256 j = i + 1; j < 3; j++) {
				if (top3Donors[i].amount < top3Donors[j].amount) {
					Donor memory temp = top3Donors[i];
					top3Donors[i] = top3Donors[j];
					top3Donors[j] = temp;
				}
			}
		}
	}

	// 查询Top3捐赠者函数：添加一个查询Top3捐赠者函数，返回当前排名前三的捐赠者
	function getTop3Donors() public view returns (Donor[3] memory) {
		return top3Donors;
	}

	// 合约应包含一个 fallback 函数，用于接收捐赠以太币。
	fallback() external payable {
		donate();
	}

	// 合约应包含一个 receive 函数，用于接收捐赠以太币。
	receive() external payable {
		donate();
	}
}
