package core

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -data DATA --交易数据")
	fmt.Println("\taddBlock -data DATA -- 交易数据")
	fmt.Println("\tprintChain --输出区块信息")
}

func isValid() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {

	if !DBExists() {
		fmt.Println("数据库不存在")
		os.Exit(1)

	}

	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.AddBlockChain(data)

}

func (cli *CLI) printChain() {
	if !DBExists() {
		fmt.Println("数据库不存在")
		os.Exit(1)

	}

	blockchain := BlockChainObject()
	defer blockchain.DB.Close()
	blockchain.PrintChain()
}

func (cli *CLI) createGenenisBlockChain(data string) {

	CreateBlockChainWithGenesisBlock(data)

}

func (cli *CLI) Run() {

	isValid()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	createBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "", "交易数据")
	flagCreateBlockChainWithData := createBlockCmd.String("data", "", "交易数据")
	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)

	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}

		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if createBlockCmd.Parsed() {
		if *flagCreateBlockChainWithData == "" {
			fmt.Println("交易数据不能为空")
			printUsage()
			os.Exit(1)
		}

		cli.createGenenisBlockChain(*flagCreateBlockChainWithData)

	}

}