// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract ERC20 {
    string public constant name = ""; // 正确使用了字符串常量
    string public constant symbol = "";
    uint8 public constant decimals = 0;

    function totalSupply() public view returns (uint) { // 更正为 view 和 returns
        // 函数体应包含返回值或抛出异常
        return 0; // 示例返回值
    }

    function balanceOf(address tokenOwner) public view returns (uint balance) {
        // 函数体应包含返回值或抛出异常
        return 0; // 示例返回值
    }

    function allowance(address tokenOwner, address spender) public view returns (uint remaining) {
        // 函数体应包含返回值或抛出异常
        return 0; // 示例返回值
    }

    function transfer(address to, uint tokens) public returns (bool success) {
        // 函数体应包含返回值或抛出异常
        return false; // 示例返回值
    }

    function approve(address spender, uint tokens) public returns (bool success) {
        // 函数体应包含返回值或抛出异常
        return false; // 示例返回值
    }

    function transferFrom(address from, address to, uint tokens) public returns (bool success) {
        // 函数体应包含返回值或抛出异常
        return false; // 示例返回值
    }

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}