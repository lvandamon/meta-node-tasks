// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ReverseString {
    function reverse(string memory input) public pure returns (string memory) {
        //输入字符串
        bytes memory inputBytes = bytes(input);
        //输出字符串
        bytes memory outputBytes = new bytes(inputBytes.length);

        // 反转字符串
        for (uint256 i = 0; i < inputBytes.length; i++) {
            outputBytes[i] = inputBytes[inputBytes.length - 1 - i];
        }

        // 转换为字符串
        return string(outputBytes);
    }
}