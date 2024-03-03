// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
 * @title Storage
 * @dev 存储和检索一个变量值
 * @dev Store and retrieve a variable value
 */
contract Storage {

  uint256 number;

  /**
   * @dev 存储一个变量
   * @dev store a variable
   * @param num store
   */
  function store(uint256 num) public {
    number += num;
  }

  /**
   * @dev
     * @return number
     */
  function retrieve() public view returns (uint256){
    return number;
  }
}
