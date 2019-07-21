package core


import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

//定义区块
type Block struct {
	//1.区块高度，也就是区块的编号，第几个区块
	Height int64
	//2.上一个区块的Hash值
	PreBlockHash []byte
	//3.交易数据（最终都属于transaction 事务）
	Data []byte
	//4.创建时间的时间戳
	TimeStamp int64
	//5.当前区块的Hash值
	Hash []byte
	//6.Nonce 随机数，用于验证工作量证明
	Nonce int64
}

//1. 创建新的区块
func NewBlock(data string, height int64, PreBlockHash []byte) *Block {
	//创建区块
	block := &Block{
		height,
		PreBlockHash,
		[]byte(data),
		time.Now().Unix(),
		nil,
		0,
	}
	//调用工作量证明的方法，并且返回有效的Hash和Nonce值
	//创建pow对象
	pow := NewProofOfWork(block)
	//挖矿验证
	hash, nonce := pow.Run(height)

	block.Hash = hash[:]
	block.Nonce = nonce
	return block

}

//2.生成创世区块
func CreateGenesisBlock(data string) *Block {

	return NewBlock(data, 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

}


// 定义Block的方法Serialize(),将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	//1.定义result的字节buffer，用于存储序列化后的区块
	var result bytes.Buffer
	//2.初始化序列化对象encoder
	encoder := gob.NewEncoder(&result)
	//3.通过Encode()方法对区块进行序列化
	err:=encoder.Encode(block)
	if err !=nil {
		log.Panic(err)
	}
	//4.返回result的字节数组
	return result.Bytes()
}

//定义函数DeserializeBlock()，传入参数为字节数组，返回值为Block
func DeserializeBlock(blockBytes []byte) *Block {
	//1.定义一个Block指针对象
	var block Block
	//2.初始化反序列化对象decoder
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	//3.通过Decode()进行反序列化
	err := decoder.Decode(&block)

	if err !=nil {
		log.Panic(err)
	}
	//4.返回block对象
	return &block
}