// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ArrayMerger {
    // 合并两个有序数组并返回新数组
    function mergeSortedArrays(uint[] memory arr1, uint[] memory arr2) public pure returns (uint[] memory) {

          // 输入验证，若任一数组为空，直接返回另一数组
        if (arr1.length == 0) {
            return arr2;
        }
        if (arr2.length == 0) {
            return arr1;
        }
        
        uint[] memory mergedArray = new uint[](arr1.length + arr2.length);
        uint i = 0; // arr1 的指针
        uint j = 0; // arr2 的指针
        uint k = 0; // mergedArray 的指针

        // 合并过程
        while (i < arr1.length && j < arr2.length) {
            if (arr1[i] < arr2[j]) {
                mergedArray[k] = arr1[i];
                i++;
            } else {
                mergedArray[k] = arr2[j];
                j++;
            }
            k++;
        }

        // 处理剩余元素
        while (i < arr1.length) {
            mergedArray[k] = arr1[i];
            i++;
            k++;
        }

        while (j < arr2.length) {
            mergedArray[k] = arr2[j];
            j++;
            k++;
        }

        return mergedArray;
    }
}