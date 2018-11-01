package main

import (
	"fmt"
	"time"
)

var sli = []int{3, 45, 456, 1, 3489, 123, 13, 4894, 21, 3, 45, 12, 3213}

//冒泡
func bubbleSort(sli []int) []int {
	len := len(sli)

	for i := 1; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
	return sli
}

//选择排序
func selectSort(sli []int) []int {
	len := len(sli)
	for i := 0; i < len-1; i++ {
		k := i
		for j := i + 1; j < len; j++ {
			if sli[k] > sli[j] {
				sli[k], sli[j] = sli[j], sli[k]
			}
		}
	}
	return sli
}

//快速排序
func quickSort(sli []int) []int {
	len := len(sli)
	if len <= 1 {
		return sli
	}
	//选择第一个元素作为基准
	baseNum := sli[0]
	//遍历除了标尺外的所有元素，按照大小关系放入左右两个切片内
	//初始化两个切片
	leftSlice := []int{}
	rightSlice := []int{}
	for i := 1; i < len; i++ {
		if baseNum > sli[i] {
			leftSlice = append(leftSlice, sli[i])
		} else {
			rightSlice = append(rightSlice, sli[i])
		}
	}
	//分别对应左右边的切片进行相同排序 递归进行
	leftSlice = quickSort(leftSlice)
	rightSlice = quickSort(rightSlice)
	//把两个分组对应合并起来
	leftSlice = append(leftSlice, baseNum)
	return append(leftSlice, rightSlice...)
}

//插入排序
func insertSort(sli []int) []int {
	len := len(sli)
	for i := 0; i < len; i++ {
		tmp := sli[i]
		for j := i - 1; j >= 0; j-- {
			if tmp < sli[j] {
				//发现插入的元素要小，交换位置，将后边的元素与前面的元素互换
				sli[j+1], sli[j] = sli[j], tmp
			} else {
				break
			}
		}
	}
	return sli
}

//睡眠排序
//通过为待排序的元素启动独立的任务，每个任务按照待排元素的key执行相对应的睡眠时间，然后及时的将序列中的元素收集到一起，达到排序的目的
func sleepSort(sli []int) {
	ch := make(chan int)
	for _, value := range sli {
		go func(val int) {
			time.Sleep(time.Duration(val) * 1000000)
			fmt.Println(val)
			ch <- val
		}(value)
	}

	for _ = range sli {
		<-ch
	}
}

func main() {
	// res := bubbleSort(sli)
	// res := selectSort(sli)
	// res := quickSort(sli)
	// res := insertSort(sli)
	// fmt.Println(res)
	sleepSort(sli)
}
