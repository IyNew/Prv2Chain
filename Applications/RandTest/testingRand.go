package main

var testNum int = 10000
var testStepSize int = testNum / 20

func GetTestSequence() []int {
	var testSeq []int
	for i := 1; i < testNum; i += testStepSize {
		testSeq = append(testSeq, i)
	}
	return testSeq
}
func main() {
	a := GetTestSequence()
	for i := 0; i < len(a); i++ {
		println(a[i])
	}
}
