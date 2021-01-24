package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the runningMedian function below.
 */
func runningMedian(a []int32) []float64 {

	result := []float64{}
	mh := NewHeap()
	for _, v := range a {

		result = append(result, mh.Update(int(v)))
	}

	return result

}

type minHeap []int

func (mh minHeap) Len() int {
	return len(mh)
}

func (mh minHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}

func (mh minHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *minHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *minHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]

	return x
}

func (mh minHeap) Peek() int {
	return mh[0]
}

type maxHeap []int

func (mh maxHeap) Len() int {
	return len(mh)
}

func (mh maxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh maxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *maxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *maxHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]

	return x
}

func (mh maxHeap) Peek() int {
	return mh[0]
}

type mediansHeap struct {
	max *maxHeap
	min *minHeap
}

func NewHeap() *mediansHeap {
	h := &mediansHeap{
		max: &maxHeap{},
		min: &minHeap{},
	}

	heap.Init(h.max)
	heap.Init(h.min)

	return h
}

func (h mediansHeap) Add(n interface{}) {
	if h.max.Len() == 0 {
		heap.Push(h.max, n)
		return
	}

	if n.(int) > h.max.Peek() {
		heap.Push(h.min, n)
	} else {
		heap.Push(h.max, n)
	}

	if h.max.Len()-h.min.Len() == 2 {
		heap.Push(h.min, heap.Pop(h.max))
	} else if h.min.Len()-h.max.Len() == 1 {
		heap.Push(h.max, heap.Pop(h.min))
	}
}

func (h mediansHeap) Median() float64 {
	fmt.Println(h.max.Len(), h.min.Len())
	if h.max.Len() == 0 {
		return 0.0
	} else if h.max.Len() == h.min.Len() {
		fmt.Println(h.max.Peek(), h.min.Peek())
		return (float64(h.max.Peek()) + float64(h.min.Peek())) / 2
	}

	return float64(h.max.Peek())
}

func (h mediansHeap) Update(n interface{}) float64 {
	h.Add(n)
	return h.Median()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	aCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var a []int32

	for aItr := 0; aItr < int(aCount); aItr++ {
		aItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := runningMedian(a)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%f", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
