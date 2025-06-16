// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting {
    // 使用mapping存储候选人得票数
    mapping(string => uint256) private votesReceived;
    
    // 存储所有候选人名单
    string[] private candidates;
    
    // 合约所有者
    address private owner;
    
    // 构造函数，初始化候选人名单
    constructor(string[] memory candidateNames) {
        candidates = candidateNames;
        owner = msg.sender;
    }
    
    // 投票函数
    function vote(string memory candidate) public {
        require(isValidCandidate(candidate), "Invalid candidate");
        votesReceived[candidate] += 1;
    }
    
    // 获取候选人得票数
    function getVotes(string memory candidate) public view returns (uint256) {
        require(isValidCandidate(candidate), "Invalid candidate");
        return votesReceived[candidate];
    }
    
    // 重置所有候选人得票数（仅限合约所有者）
    function resetVotes() public {
        require(msg.sender == owner, "Only owner can reset votes");
        for (uint i = 0; i < candidates.length; i++) {
            votesReceived[candidates[i]] = 0;
        }
    }
    
    // 检查候选人是否有效
    function isValidCandidate(string memory candidate) private view returns (bool) {
        for (uint i = 0; i < candidates.length; i++) {
            if (keccak256(bytes(candidates[i])) == keccak256(bytes(candidate))) {
                return true;
            }
        }
        return false;
    }
    
    // 获取所有候选人名单
    function getAllCandidates() public view returns (string[] memory) {
        return candidates;
    }
}