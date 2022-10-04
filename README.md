# Tendermint Practice

## 编译
```
git clone git@github.com:ssuchichen/tendermint.git
cd tendermint
make build # build目录下会生成可执行文件tendermint
```

## 什么是Tendermint？
Tendermint is software for securely and consistently replicating an application on many machines.  
Tendermint是用于在许多机器上安全、一致地复制应用程序的软件。
* **安全性**  
多达1/3的机器以任意方式出现故障，Tendermint也能正常工作。
* **一致性**  
每台无故障的机器都会看到相同的事务日志并计算相同的状态。

## Tendermint组成
Tendermint consists of two chief technical components: a blockchain consensus engine and a generic application interface.
* **Tendermint Core**  
确保在每台机器上以相同的顺序记录相同的事务。
* **Application BlockChain Interface (ABCI)**  
允许用任何编程语言处理事务。

## Tendermint的作用
构建需要状态复制的应用，如区块链、存储等。


