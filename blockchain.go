
package blockchain

import(
	"fmt"
	"crypto/sha256"
)

type Block struct{
	data string
	hash string
	prev *Block
}

func CalculateHash(data string) string{
	h:= sha256.New()
	h.Write([] byte(data))
	valueInHex := fmt.Sprintf("%x", h.Sum(nil))
	return valueInHex
}

func InsertBlock(transaction string, chainHead *Block) *Block {

	var newBLock Block

	newBLock.data = transaction
	newBLock.hash = CalculateHash(transaction)

	if chainHead == nil{
		newBLock.prev = nil
	}else{
		newBLock.prev = chainHead
	}

	chainHead = &newBLock;
	return chainHead
}

func ListBlocks(chainHead *Block) {

	fmt.Println(chainHead)
	for chainHead.prev!=nil {
		chainHead = chainHead.prev
		fmt.Println(chainHead)
	}
}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {

	for chainHead!=nil {

		if chainHead.data == oldTrans {
			chainHead.data = newTrans
			break
		}else{
			chainHead = chainHead.prev
		}	
	}
}

func VerifyChain(chainHead *Block) {
	var verified = true

	for chainHead!=nil {

		if chainHead.hash != CalculateHash(chainHead.data) {
			verified = false
			break
		}
		chainHead = chainHead.prev	
	}

	if verified == true {
		fmt.Println("Chain is correct and verified")
	}else{
		fmt.Println("Chain is tampered")
	}
}
