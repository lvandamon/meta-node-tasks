// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 在一个有序数组中查找目标值。
contract BinarySearch {
    // 二分查找函数
    function binarySearch(uint[] memory arr, uint target) public pure returns (bool) {
        uint left = 0;
        uint right = arr.length - 1;

        while (left <= right) {
            uint mid = left + (right - left) / 2;

            if (arr[mid] == target) {
                return true;
            } else if (arr[mid] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        return false;
    }
}
