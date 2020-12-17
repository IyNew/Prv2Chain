package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

var testNum int = 10
var testTreeNum int = 10
var straightTreeNum int = 2

type Record struct {
	ID       string `json:"ID"`
	Previous string `json:"previous"`
	Future   string `json:"future"`
	Data     string `json:"data"`
}
type RecordList struct {
	Records []Record
}

type FutureRecord struct {
	ID       string `json:"ID"`
	Previous string `json:"previous"`
	Future   string `json:"future"`
	Data     string `json:"data"`
}

func RecordToFutureRecord(record Record) FutureRecord {
	// var futureRecord FutureRecord
	futureRecord := FutureRecord{
		ID:       record.ID,
		Previous: record.Previous,
		Future:   "",
		Data:     record.Data,
	}
	return futureRecord
}

func PrintRecordList(list []*Record) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("{\nID: %s,\nPrevious:%s,\nData:%s\n} \n", list[i].ID, list[i].Previous, list[i].Data)
	}
}
func PrintFutureRecordList(list []*FutureRecord) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("{\nID: %s,\nPrevious:%s,\nFuture:%s,\nData:%s\n} \n", list[i].ID, list[i].Previous, list[i].Future, list[i].Data)
		// fmt.Printf("{\nID: %s,\nPrevious:%s,\nData:%s\n} \n", list[i].ID, list[i].Previous, list[i].Future, list[i].Data)
	}
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func GenerateRecord(seed *rand.Rand, seq int) *Record {

	sha := sha1.New()
	sha.Write(IntToBytes(seq))
	// fmt.Println(hex.EncodeToString(sha.Sum(nil)))
	record := Record{
		ID:       hex.EncodeToString(sha.Sum(nil)),
		Previous: "",
		Future:   "",
		Data:     fmt.Sprint(seq),
	}
	return &record
}

func RollDice(seed *rand.Rand, max int) int {
	return seed.Intn(max)
}
func RollDiceWithoutSeed(max int) int {
	return rand.Intn(max)
}

func GetRandomTree(recordNum int) []*Record {

	var list []*Record

	rr := rand.New(rand.NewSource(time.Now().Unix()))
	// rr2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < recordNum; i++ {
		newRecord := GenerateRecord(rr, i)
		if i != 0 {
			// previousNo := RollDice(rr2, i)
			previousNo := RollDiceWithoutSeed(i)
			println("Dice = ", previousNo)
			newRecord.Previous = list[previousNo].ID

			// parse the future part of previous record
			list[previousNo].Future += "|" + newRecord.ID
		}

		list = append(list, newRecord)
	}

	// naiveJson, err := json.Marshal(list)
	// if err != nil {
	// }
	// fmt.Println(string(naiveJson))
	return list
}

func GetRandomTreeWithJump(recordNum int, jump int) []*Record {

	var list []*Record

	var previousNo int

	rr := rand.New(rand.NewSource(time.Now().Unix()))
	// rr2 := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < recordNum; i++ {
		recordInicator := recordNum*jump + i

		newRecord := GenerateRecord(rr, recordInicator)
		if i > 0 {
			// previousNo := RollDice(rr2, i)
			if jump < straightTreeNum {
				previousNo = i - 1
			} else {
				previousNo = RollDiceWithoutSeed(i)
			}

			newRecord.Previous = list[previousNo].ID

			// parse the future part of previous record
			list[previousNo].Future += "|" + newRecord.ID
		}

		list = append(list, newRecord)
	}

	return list
}

func GetMultipleRandomTrees(recordNum int, treeNum int) []*Record {

	var list []*Record
	var nodeListForEachTree [][]*Record
	var treeNodeIndicator []int

	for i := 0; i < treeNum; i++ {
		tree := GetRandomTreeWithJump(recordNum, i)
		nodeListForEachTree = append(nodeListForEachTree, tree)
		treeNodeIndicator = append(treeNodeIndicator, 0)
	}
	rr := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < treeNum*recordNum; i++ {
		dice := rr.Intn(treeNum)
		for treeNodeIndicator[dice] >= recordNum {
			dice = (dice + 1) % treeNum
		}
		list = append(list, nodeListForEachTree[dice][treeNodeIndicator[dice]])
		treeNodeIndicator[dice]++
	}
	// naiveJson, err := json.Marshal(list)
	// if err != nil {
	// }
	// fmt.Println(string(naiveJson))
	return list
}

func GetForwardTestSequence() []int {
	var testSeq []int
	for i := straightTreeNum; i < testTreeNum; i++ {
		testSeq = append(testSeq, i*testNum)
	}
	return testSeq
}

func GetBackwardTestSequence() []int {
	var testSeq []int
	for i := 0; i < straightTreeNum; i++ {
		testSeq = append(testSeq, (i+1)*testNum-1)
	}
	return testSeq
}

func main() {
	// treeNum := 10
	// recordNumInTree := 10
	// recordList := GetRandomListForAllTrees(treeNum, recordNumInTree)
	// q, err := json.Marshal(recordList)
	// if err != nil {
	// }
	list := GetMultipleRandomTrees(testNum, testTreeNum)
	naiveJson, err := json.Marshal(list)
	if err != nil {
	}
	fmt.Println(string(naiveJson))
	fmt.Println(GetForwardTestSequence())
	fmt.Println(GetBackwardTestSequence())

}
