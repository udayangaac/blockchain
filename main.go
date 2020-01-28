package main

import (
	"flag"
	"fmt"
	"github.com/udayangaac/blockchain/blockchain"
	"os"
	"runtime"
	"strconv"
)

type CommandLine struct {
	blockChain *blockchain.BlockChain
}

func (cli *CommandLine) printUsage() {
	fmt.Println(" Usage:")
	fmt.Println(" add -block BLOCK_DATA - add a block to chain")
	fmt.Println(" print - Print the blocks in the chain")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) addBlock(data string) {
	cli.blockChain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (cli *CommandLine) printChain() {
	iter := cli.blockChain.Iterator()
	for {
		block := iter.Next()
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		pow := blockchain.NewProof(block)
		fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevHash) == 0 {

		}
	}
}

func (cli *CommandLine) run() {
	cli.validateArgs()
	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")
	val := os.Args[1]
	fmt.Printf("Console Command : %v\n", val)
	switch val {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.ErrorHandle(err)
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.ErrorHandle(err)
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	defer func() {
		err := chain.Database.Close()
		blockchain.ErrorHandle(err)
	}()

	cli := CommandLine{blockChain: chain}
	cli.run()
}
