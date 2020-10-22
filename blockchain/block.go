package blockchain

import (
	"time"
	"encoding/gob"
	"bytes"
)

/**
 * 定义区块结构体, 用于表示区块
 */
type Block struct {
	Height    int64  //区块的高度,第几个区块
	TimeStamp int64  // 区块产生的时间戳
	PrevHash  []byte //前一个区块hash
	Data      []byte //数据字段
	Hash      []byte //当前区块的hash值
	Version   string //版本号
	Nonce     int64  // 区块对应的nonce值
}

/**
 * 创建一个新区块
 */
func NewBlock(height int64, prevHash []byte, data []byte) Block {
	block := Block{
		Height:    height,
		TimeStamp: time.Now().Unix(),
		PrevHash:  prevHash,
		Data:      data,
		Version:   "0x01",
	}

	//找nonce值,通过工作量证明算法计算寻找
	//挖矿竞争，获得记账权
	xiaolong := NewPoW(block)
	hash, nonce := xiaolong.Run()
	block.Nonce = nonce
	block.Hash = hash

	//1、将block结构体数据转换为[]byte类型
	//heightBytes, _ := Int64ToByte(block.Height)
	//timeStampBytes, _ := Int64ToByte(block.TimeStamp)
	//versionBytes := StringToBytes(block.Version)
	//nonceBytes, _ := utils.Int64ToByte(block.Nonce)
	//var blockBytes []byte
	////bytes.Join拼接
	//bytes.Join([][]byte{
	//	heightBytes,
	//	timeStampBytes,
	//	block.PrevHash,
	//	block.Data,
	//	versionBytes,
	//	nonceBytes,
	//}, []byte{})
	//
	////调用Hash计算,对区块进行sha256哈希值计算
	//block.Hash = utils.SHA256HashBlock(blockBytes)

	return block
}

/**
 * 创建创世区块
 */
func CreateGenesisBlock() Block {
	genesisBlock := NewBlock(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, nil)
	return genesisBlock
}

/**
 * 对区块进行序列化操作
 */
func (b Block) Serialize() ([]byte) {
	buff := new(bytes.Buffer) //缓冲区
	encoder := gob.NewEncoder(buff)
	encoder.Encode(b) //将区块b放入到序列化编码器中
	return buff.Bytes()
}

/**
 * 区块反序列化操作
 */
func DeSerialize(data []byte) (*Block, error) {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}
