package main

import (
	"flag"
	"fmt"
	"log"
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
			log.Panic(err)
		}
	case "print":
		err := printCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
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
	bc, err := blockchain.NewBlockchain()
	if err != nil {
		log.Panic(err)
	}
	defer bc.Close()

	err = bc.AddBlock(data)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Success!")
}

// printChain prints all the blocks in the blockchain
func (cli *CLI) printChain() {
	bc, err := blockchain.NewBlockchain()
	if err != nil {
		log.Panic(err)
	}
	defer bc.Close()

	blocks, err := bc.GetBlocks()
	if err != nil {
		log.Panic(err)
	}

	for _, block := range blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
	}
}

func main() {
	cli := CLI{}
	cli.Run()
}
