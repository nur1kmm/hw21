package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/nur1kmm/hw21/blockchain"
)

// CLI is the command-line interface for the blockchain
type CLI struct{}

// printUsage prints the usage instructions
func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add -data <data> - add a new block to the blockchain")
	fmt.Println("  print - print all the blocks in the blockchain")
}

// validateArgs validates the command-line arguments
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run runs the command-line interface
func (cli *CLI) Run() {
	cli.validateArgs()

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printCmd := flag.NewFlagSet("print", flag.ExitOnError)

	addBlockData := addCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error parsing arguments:", err)
			os.Exit(1)
		}
	case "print":
		err := printCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("Error parsing arguments:", err)
			os.Exit(1)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addCmd.Parsed() {
		if *addBlockData == "" {
			addCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printCmd.Parsed() {
		cli.printChain()
	}
}

// addBlock adds a new block to the blockchain
func (cli *CLI) addBlock(data string) {
	bc := blockchain.NewBlockchain()
	defer bc.Close()
	bc.AddBlock(data)
	fmt.Println("Success!")
}

// printChain prints all the blocks in the blockchain
func (cli *CLI) printChain() {
	bc := blockchain.NewBlockchain()
	defer bc.Close()
	bci := bc.Iterator()

	for {
		block := bci.Next()
		if block == nil {
			break
		}

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

// main is the command-line entry point for the blockchain
func main() {
	cli := CLI{}
	cli.Run()
}
