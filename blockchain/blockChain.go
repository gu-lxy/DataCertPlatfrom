package blockchain

import (
	"errors"
	"fmt"
	"github.com/bolt-master/bolt-master"
	"math/big"
)

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME = "blocks"
const LAST_HASH = "lasthash"

var CHAIN *BlockChain
/**
 * 区块链结构体的定义，代表的是一条区块链
 */
type BlockChain struct {
	LastHash []byte   // 表示区块链中最新区块的哈希，用于查找最新的区块内容
	BoltDb   *bolt.DB //区块链中操作区块数据文件的数据库操作对象
}

/**
 * 创建一条区块链
 */
func NewBlockChain() BlockChain {
	var bc BlockChain
	//先打开文件
	db, err := bolt.Open(BLOCKCHAIN, 0600, nil)
	//查看chain.db文件
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte(BUCKET_NAME))
  		    if err != nil {
  		    	panic(err.Error())
			}
		}
		lastHash := bucket.Get([]byte(LAST_HASH))
		if len(lastHash) == 0 {//桶中没有lastHash记录，需要新建创世区块，并保存
			//	//创世区块
			genesis := CreateGenesisBlock()
			//区块序列化以后的数据
			gensisBytes := genesis.Serialize()
			//创世区块保存到boltdb中
			bucket.Put(genesis.Hash,gensisBytes)
			//更新最新区块的哈希值记录
			bucket.Put([]byte(LAST_HASH),gensisBytes)
			bc = BlockChain{
				LastHash: lastHash,
				BoltDb:   db,
			}
		}else {
			lastHash1 := bucket.Get([]byte(LAST_HASH))
			bc = BlockChain{
				LastHash: lastHash1,
				BoltDb:   db,
			}
		}
		return nil
	})
	CHAIN = bc
	return bc
}
//该方法用于遍历区块链chain.db文件，并将所有的区块查出，并返回
func (bc BlockChain) QueryAllAlocks() ([]Block,error)  {
	blocks := make([]Block,0)//blocks是一个切片容器，用于盛放查询拿到的区块

	db := bc.BoltDb
	var err error
	//从chain.db文件查询到有的区块
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("查询区块数据失败" )
			return nil
		}
		eachHash := bc.LastHash
		eachBig := new(big.Int)
		zeroBig := big.NewInt(0)
		for {
			//根据区块的hash值获取对应的区块
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化操作
		    eachBlock, _ := DeSerialize(eachBlockBytes)
			//将遍历到每一个区块放入到切片容器当中
		    blocks = append(blocks, )

		    eachBig.SetBytes(eachBlock.PrevHash)
		    if eachBig.Cmp(zeroBig) == 0 {
		    	break//跳出循环
			}
			//不满足条件，没有找到创世区块
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return blocks,err
}

//该方法用户完成根据用户输入的区块高度查询对应的区块信息
func (bc BlockChain) QueryBlockByHeight(height int64) (*Block, error) {
	db := bc.BoltDb

	var errs error
	var eachBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			errs = errors.New("都区区块信息失败")
			return errs
		}
		eachHash := bc.LastHash
		for {
			//获取的搜最后一个区块
			eachBlockBytes := bucket.Get(eachHash)
			//反序列化操作
			eachBlock, errs := DeSerialize(eachBlockBytes)
			if errs != nil {
				//fmt.Println("遍历遇到错误")
				return errs
			}
			if eachBlock.Height < height {
				break
			}
			if eachBlock.Height == height {
				break
			}
			eachHash = eachBlock.PrevHash
		}
		return nil
	})
	return eachBlock,errs
}

/**
 * 保存数据到区块链中: 先生成一个新区块,然后将新区块添加到区块链中
 */
func (bc BlockChain) SaveData(data []byte) (Block,error) {
	//1、从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
	//error的自定义
	var err error
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			err = errors.New("读取区块链数据失败")
			//panic("读取区块链数据失败")
			return err
		}
		//lastHash := bucket.Get([]byte(LAST_HASH))
		lastBlockBytes := bucket.Get(bc.LastHash)
		//反序列化
		lastBlock, _ = DeSerialize(lastBlockBytes)
		return nil
	})

	//新建一个区块
	newBlock := NewBlock(lastBlock.Height+1, lastBlock.Hash, data)
	//把新区块存到文件中
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		//序列化后的区块的数据
		blockBytes := newBlock.Serialize()
		fmt.Printf("序列化后的数据",blockBytes)
		//把新创建的区块存储到boltdb数据中
		bucket.Put(newBlock.Hash,newBlock.Serialize())
		//更新LASTHASH对应的值，更新为 最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH),newBlock.Hash)
		bc.LastHash = lastBlock.Hash
		return nil
	})
	//返回值语句：newclock，err可能含有错误信息
	return newBlock,err
}
