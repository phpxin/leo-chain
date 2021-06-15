/*
 @Title
 @Description
 @Author  Leo
 @Update  2020/9/10 12:19 下午
*/

package blockchain

type Block struct {
	Id int64 `json:"block_id"`
	Hash string `json:"block_hash"`
	TxIds []string `json:"tx_ids"`
	TxRootHash string `json:"tx_root_hash"`
}

