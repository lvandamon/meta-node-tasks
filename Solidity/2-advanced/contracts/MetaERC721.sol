// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

// 题目 2
contract MetaERC721 is ERC721URIStorage, Ownable {
	uint256 private _nextTokenId;

	constructor(string memory name, string memory symbol) ERC721(name, symbol) Ownable(msg.sender) {}

	function mintNFT(address recipient, string memory tokenURI_) public onlyOwner returns (uint256) {
		require(recipient != address(0), "Invalid address");
		require(bytes(tokenURI_).length > 0, "Token URI must be provided");
		// 铸造 NFT
		uint256 tokenId = _nextTokenId;
		_nextTokenId++;
		_mint(recipient, tokenId);
		// 设置 NFT 的 URI
		_setTokenURI(tokenId, tokenURI_);
		return tokenId;
	}

	function tokenURI(uint256 tokenId) public view override returns (string memory) {
		return super.tokenURI(tokenId);
	}
}
