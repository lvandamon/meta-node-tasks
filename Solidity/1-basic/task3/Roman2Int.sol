// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 罗马数字转整数
contract Roman2Int {
    // 定义罗马数字字符和对应整数值的映射
    mapping(bytes1 => uint256) private ROMAN_VALUES;

    // 构造函数，初始化 ROMAN_VALUES
    constructor() {
        ROMAN_VALUES['I'] = 1;
        ROMAN_VALUES['V'] = 5;
        ROMAN_VALUES['X'] = 10;
        ROMAN_VALUES['L'] = 50;
        ROMAN_VALUES['C'] = 100;
        ROMAN_VALUES['D'] = 500;
        ROMAN_VALUES['M'] = 1000;
    }

    function romanToInt(string memory s) public view returns (uint256) {
        uint256 result = 0;
        bytes memory inputBytes = bytes(s);
        uint256 length = inputBytes.length;

        for (uint256 i = 0; i < length; i++) {
            bytes1 currentChar = inputBytes[i];
            // 验证输入字符是否有效
            require(ROMAN_VALUES[currentChar] != 0, "Invalid Roman numeral character");

            uint256 currentValue = ROMAN_VALUES[currentChar];
            if (i + 1 < length) {
                bytes1 nextChar = inputBytes[i + 1];
                uint256 nextValue = ROMAN_VALUES[nextChar];
                if (currentValue < nextValue) {
                    result += nextValue - currentValue;
                    i++; // 跳过下一个字符
                    continue;
                }
            }
            result += currentValue;
        }
        return result;
    }
}