// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

//整数转罗马数字
contract Int2Roman {
    function intToRoman(uint256 num) public pure returns (string memory) {
        // 输入验证
        require(num > 0, "Input must be greater than zero");

        // 只验证输入是否在1到3999之间
        require(num < 4000, "Input must be less than 4000");

        // 从最大的罗马数字字符开始转换
        uint256[] memory values = new uint256[](13);
        values[0] = 1000;
        values[1] = 900;
        values[2] = 500;
        values[3] = 400;
        values[4] = 100;
        values[5] = 90;
        values[6] = 50;
        values[7] = 40;
        values[8] = 10;
        values[9] = 9;
        values[10] = 5;
        values[11] = 4;
        values[12] = 1;

        // 定义罗马数字字符数组
        string[] memory symbols = new string[](13);
        symbols[0] = "M";
        symbols[1] = "CM";
        symbols[2] = "D";
        symbols[3] = "CD";
        symbols[4] = "C";
        symbols[5] = "XC";
        symbols[6] = "L";
        symbols[7] = "XL";
        symbols[8] = "X";
        symbols[9] = "IX";
        symbols[10] = "V";
        symbols[11] = "IV";
        symbols[12] = "I";

        string memory result = "";
        // 从最大的罗马数字字符开始转换
        for (uint256 i = 0; i < values.length; i++) {
            while (num >= values[i]) {
                num -= values[i];
                result = string(abi.encodePacked(result, symbols[i]));
            }
        }

        return result;
    }
}
