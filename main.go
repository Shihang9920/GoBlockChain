package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

// 定义区块的结构体
type Block struct {
	Timestamp int64  //时间戳
	Hash      []byte //本身的哈希
	PrevHash  []byte //指向上一个区块的哈希
	Data      []byte //区块中的数据
}

// 定义区块链
type BlockChain struct {
	Blocks []*Block
}

// 对区块包含的信息做哈希，相当于区块的id
func (b *Block) SetHash() {
	//拼接字节串
	information := bytes.Join([][]byte{ToHexInt(b.Timestamp), b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(information)
	b.Hash = hash[:]
}

// 对时间戳做哈希
func ToHexInt(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// 创建区块的函数
func CreateBlock(prevhash, data []byte) *Block {
	block := Block{time.Now().Unix(), []byte{}, prevhash, data}
	block.SetHash()
	return &block
}

// 创世区块
func GenesisBlock() *Block {
	genesisWords := "创世区块"
	return CreateBlock([]byte{}, []byte(genesisWords))
}

// 添加区块
func (bc *BlockChain) AddBlock(data string) {
	newBlock := CreateBlock(bc.Blocks[len(bc.Blocks)-1].Hash, []byte(data))
	bc.Blocks = append(bc.Blocks, newBlock)
}

// 初始化区块链
func CreateBlockChain() *BlockChain {
	blockchain := BlockChain{}
	blockchain.Blocks = append(blockchain.Blocks, GenesisBlock())
	return &blockchain
}
func main() {
	//初始化区块链
	blockchain := CreateBlockChain()
	time.Sleep(time.Second)
	blockchain.AddBlock("第二个区块...")
	time.Sleep(time.Second)
	blockchain.AddBlock("第三个区块...")
	time.Sleep(time.Second)
	blockchain.AddBlock("第四个区块...")
	time.Sleep(time.Second)

	for _, block := range blockchain.Blocks {
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("hash: %x\n", block.Hash)
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("data: %s\n", block.Data)

	}
}
