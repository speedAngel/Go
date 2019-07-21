package core

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
	"time"
)

const dbName = "blockchain.db"  //数据库名
const blockTableName = "blocks" //表名
type BlockChain struct {
	Tip []byte   //区块链里面最后一个区块的Hash
	DB  *bolt.DB //数据库
}

//迭代器
func (blockchain *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{blockchain.Tip, blockchain.DB}

}

//判断数据库是否存在
func DBExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}

func (blc *BlockChain) PrintChain() {

	blockchainIterator := blc.Iterator()

	for {
		block := blockchainIterator.Next()

		fmt.Printf("Height:%d\n", block.Height)
		fmt.Printf("PreBlockHash:%x\n", block.PreBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("TimeStamp:%s\n", time.Unix(block.TimeStamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Printf("Nonce:%d\n", block.Nonce)

		var hashInt big.Int
		hashInt.SetBytes(block.PreBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {

			break
		}

	}

}

//
func (blc *BlockChain) AddBlockChain(data string) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		//1.获取表
		b := tx.Bucket([]byte(blockTableName))
		//2.创建新区块
		if b != nil {
			//获取最新区块
			byteBytes := b.Get(blc.Tip)
			//反序列化
			block := DeserializeBlock(byteBytes)

			//3. 将区块序列化并且存储到数据库中
			newBlock := NewBlock(data, block.Height+1, block.Hash)
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			//4.更新数据库中"l"对应的Hash
			err = b.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			//5. 更新blockchain的Tip
			blc.Tip = newBlock.Hash
		}

		return nil

	})
	if err != nil {
		log.Panic(err)
	}

}

//1.创建带有创世区块的区块链
func CreateBlockChainWithGenesisBlock(data string) {

	//判断数据库是否存在
	if DBExists() {
		fmt.Println("创世区块已经存在")
		os.Exit(1)
	}

	//打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		//创建数据库表
		b, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}

		if b != nil {
			//创建创世区块
			genesisBlock := CreateGenesisBlock(data)
			//将创世区块存储至表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			//存储最新的区块链的hash
			err = b.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

}

//返回BlockChain对象
func BlockChainObject() *BlockChain {

	var tip []byte

	//打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			//读取最新区块的Hash
			tip = b.Get([]byte("l"))

		}
		return nil

	})

	return &BlockChain{tip, db}
}