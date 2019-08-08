开发环境介绍：

软件安装

- go1.12.5.windows-amd64.msi
- goland-2019.1.3.exe
- JetBrains.GoLand.2019_jb51.rar


		C:\Users\Administrator>go version
		go version go1.12.5 windows/amd64


一个GO工程中主要包含以下三个目录：

- src：源代码文件
- pkg：包文件
- bin：相关bin文件

## 项目构建 ##
	1. 建立工程文件夹 MagicCoin
	2. 在工程文件夹中建立src,pkg,bin文件夹
	3. 在File->Setting->Go->GOPATH>projiect中添加MagicCoin路径
	4. 程序的代码都是放在src目录下，可以对自己的代码进行分包管理
	5. go语言必须有一个文件的包名为package main的go文件作为程序的入口

## 编译打包： ##

-  go build coin
-  go test coin
-  go install coin


编译后可在%GOPATH%\bin目录下找到coin.exe

此程序是一个命令行程序，运行如下：
1. 在此目录下打开命令行

C:\Go\bin>coin
	
	====返回结果======	
	Usage:	
	        createblockchain -data DATA --交易数据	
	        addBlock -data DATA -- 交易数据	
	        printChain --输出区块信息

C:\Go\bin>coin addBlock -data "haha"
	
	====返回结果======	
	数据库不存在

C:\Go\bin>coin createblockchain -data  "创世区块"

	====返回结果======	
	第1个区块，挖矿成功:00000041ebecfb21cb7e167d28af92f2ae9da669c823f0b78a941a62ede6d18d	
	2019-07-21 20:14:15.602062 +0800 CST m=+0.133093001

C:\Go\bin>coin addBlock -data "第一个区块"

	====返回结果======	
	第2个区块，挖矿成功:00003a367f5ead94b5eae8a6e35c738cb6a650afeaf2aee8b7a512ff474f4b3a	
	2019-07-21 20:14:32.6231128 +0800 CST m=+0.124090101	
	C:\Go\bin>coin printChain

====返回结果======

	Height:3
	PreBlockHash:00003a367f5ead94b5eae8a6e35c738cb6a650afeaf2aee8b7a512ff474f4b3a
	Data:第二个区块
	TimeStamp:2019-07-21 08:25:14 PM
	Hash:00001039613178060f580d2923b3f7aee8be7cae9268c1bc8784e5a19b016fd0
	Nonce:239791
	
	Height:2
	PreBlockHash:00000041ebecfb21cb7e167d28af92f2ae9da669c823f0b78a941a62ede6d18d	
	Data:第一个区块	
	TimeStamp:2019-07-21 08:14:32 PM	
	Hash:00003a367f5ead94b5eae8a6e35c738cb6a650afeaf2aee8b7a512ff474f4b3a	
	Nonce:96610
	
	Height:1	
	PreBlockHash:0000000000000000000000000000000000000000000000000000000000000000	
	Data:创世区块	
	TimeStamp:2019-07-21 08:14:15 PM	
	Hash:00000041ebecfb21cb7e167d28af92f2ae9da669c823f0b78a941a62ede6d18d	
	Nonce:101456

## 说明 ##

	
	- Nonce是共识算法，也就挖矿产生的证明工作量的 数值，此数值一个有效区块的唯一值	
	
	- Height是区块在链中的索引