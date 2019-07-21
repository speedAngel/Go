package core


import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

type ProofOfWork struct {
	Block  *Block   //当前要验证的区块
	target *big.Int //大数存储，区块难度
}

//数据拼接，返回字节数组
func (pow *ProofOfWork) prePareData(nonce int) []byte {

	data := bytes.Join(
		[][]byte{
			pow.Block.PreBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.TimeStamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)
	return data
}

//256位Hash里面至少要有16个零0000 0000 0000 0000
const targetBit = 16

func (proofOfWork *ProofOfWork) Run(num int64) ([]byte, int64) {

	//3.判断Hash的有效性，如果满足条件循环体

	nonce := 0
	var hashInt big.Int //存储新生成的hash值
	var hash [32]byte

	for {
		//1. 将Block的属性拼接成字节数组
		databytes := proofOfWork.prePareData(nonce)

		//2.生成Hash
		hash = sha256.Sum256(databytes)
		//fmt.Printf("挖矿中..%x\n", hash)
		//3. 将hash存储至hashInt
		hashInt.SetBytes(hash[:])


		//4.判断hashInt是否小于Block里面的target
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//需要hashInt(y)小于设置的target(x)
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			//fmt.Println("挖矿成功", hashInt)
			fmt.Printf("第%d个区块，挖矿成功:%x\n",num,hash)
			fmt.Println(time.Now())
			time.Sleep(time.Second * 2)
			break

		}

		nonce ++

	}

	return hash[:], int64(nonce)

}

//创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	/*1.创建初始值为1的target
	  0000 0001
	  8 - 2
	*/

	target := big.NewInt(1)

	//2.左移256-targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}


//判断挖矿得到的区块是否有效
func (proofOfWork *ProofOfWork) IsValid() bool {
	//1.proofOfWork.Block.Hash
	//2.proofOfWork.Target
	var hashInt big.Int

	hashInt.SetBytes(proofOfWork.Block.Hash)

	if proofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}