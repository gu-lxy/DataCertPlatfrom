package blockchain

import "github.com/bolt-master/bolt-master"

const BLOCKCHAIN = "blockchain.db"
const BUCKET_NAME = "blocks"
const LAST_HASH = "lasthash"

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


	return bc
}

/**
 * 保存数据到区块链中: 先生成一个新区块,然后将新区块添加到区块链中
 */
func (bc BlockChain) SaveData(data []byte) {
	//1、从文件中读取到最新的区块
	db := bc.BoltDb
	var lastBlock *Block
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BUCKET_NAME))
		if bucket == nil {
			panic("读取区块链数据失败")
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
		//把新创建的区块存储到boltdb数据中
		bucket.Put(newBlock.Hash,newBlock.Serialize())
		//更新LASTHASH对应的值，更新为 最新存储的区块的hash值
		bucket.Put([]byte(LAST_HASH),newBlock.Hash)
		bc.LastHash = lastBlock.Hash
		return nil
	})
}
