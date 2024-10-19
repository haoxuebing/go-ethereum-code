

## windows

###  1.Solidity编译器

```
npm install -g solc
```

> 命令行可执行文件名为solcjs

```shell
solcjs --version
```

### 2.安装 abigen 工具

```shell
go get -u github.com/ethereum/go-ethereum
```

进入：github.com\ethereum\go-ethereum\cmd\abigen 目录

然后 go build

等待下载库和编译的时间，会生成一个exe的文件，放到 go bin目录

### 3.编写一个简单的智能合约

store.go

```sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) { // 修改参数类型和可见性
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
```

### 4.利用solc生成 .abi 文件

```shell
solcjs --abi Store.sol  
// or
solc --abi Store.sol
```

### 5.编译为EVM字节码 .bin 文件

```shell
solcjs --bin Store.sol
// or
solc --bin Store.sol
```

### 6.利用abigen生成 .go 文件

```shell
abigen --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```

### 7.生成包括deploy方法的 .go 文件

```shell
abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=Store.go
```
