package core

import (
	"github.com/boltdb/bolt" //导入BoltDB库
	"log"
)

type BlockChainIterator struct {
	CurrentHash []byte   // 保存当前的区块Hash值
	DB          *bolt.DB //DB对象
}

func (blockchainIterator *BlockChainIterator) Next() *Block {
	//1.定义Block对象block
	var block *Block
	//2.操作DB对象blockchainIterator.DB
	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {
		//3.打开表对象blockTableName
		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			//Get()方法通过Key:当前区块的Hash值获取当前区块的序列化信息
			currentBlockBytes := b.Get(blockchainIterator.CurrentHash)
			//反序列化出当前的区块
			block = DeserializeBlock(currentBlockBytes)
			//更新迭代器里面的CurrentHash
			blockchainIterator.CurrentHash = block.PreBlockHash
		}
		return nil

	})

	if err != nil {
		log.Panic(err)
	}
	return block
}