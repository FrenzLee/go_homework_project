package main

import (
	"fmt"
	"go_homework_project/taskOne"
	"go_homework_project/taskTwo"
)

func taskOneTest() {
	/*
		//136. 只出现一次的数字
		res1 := taskOne.SingleNumber([]int{1, 2, 3, 4, 4, 2, 1})
		fmt.Println(res1)

		//9. 回文数
		res2 := taskOne.IsPalindrome(1331)
		fmt.Println(res2)

		//20. 有效的括号
		var s string = "{[}]"
		res3 := taskOne.IsValid(s)
		fmt.Println(res3)


		//14. 最长公共前缀
		var strs []string = []string{"flower", "fl", "flo"}
		//res41 := taskOne.LongestCommonPrefix1(strs)
		//fmt.Println(res41)
		res42 := taskOne.LongestCommonPrefix2(strs)
		fmt.Println(res42)

		//66. 加一
		digits := []int{1, 2, 9}
		res5 := taskOne.PlusOne1(digits)
		fmt.Println(res5)

		//26. 删除排序数组中的重复项
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		res6 := taskOne.RemoveDuplicates(nums)
		fmt.Println(res6)

		//56. 合并区间
		intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
		res7 := taskOne.Merge(intervals)
		fmt.Println(res7)
	*/

	//1. 两数之和
	res8 := taskOne.TwoSum([]int{3, 2, 4}, 6)
	fmt.Println(res8)

}

func taskTwoTest() {
	/*

		//一、指针
		num := 8
		res1 := taskTwo.PointerPlus(&num)
		fmt.Println("res1===", res1)

		intSlice := []int{1, 2, 3}
		res2 := taskTwo.PointerSlice(&intSlice)
		fmt.Println("res2===", res2)

		//二、Goroutine
		var wg sync.WaitGroup

		wg.Add(2)
		var mu sync.Mutex
		taskTwo.PrintNum(&wg, &mu)
		wg.Wait()


		task1 := func() { time.Sleep(2 * time.Second); fmt.Println("任务 1 完成") }
		task2 := func() { time.Sleep(3 * time.Second); fmt.Println("任务 2 完成") }
		task3 := func() { time.Sleep(4 * time.Second); fmt.Println("任务 3 完成") }

		scheduler := taskTwo.NewScheduler(task1, task2, task3)

		task4 := func() { time.Sleep(5 * time.Second); fmt.Println("任务 4 完成") }
		scheduler.Add(task4)

		scheduler.Run(&wg)

		wg.Wait()
		fmt.Println("所有任务已完成")

		//三、面向对象
		rectangle := taskTwo.Rectangle{}
		rectangle.Perimeter()
		rectangle.Area()

		circle := taskTwo.Circle{}
		circle.Perimeter()
		circle.Area()

		person := taskTwo.Person{Name: "王某", Age: 23}
		employee := taskTwo.Employee{EmployeeID: 1, Person: person}
		employee.PrintInfo()

		//四、Channel
		ch := make(chan int)
		go taskTwo.SendChannel(ch)
		taskTwo.RecieveChannel(ch)

		ch1 := make(chan int, 10)
		go taskTwo.ProduceChannel(ch1)
		go taskTwo.ConsumerChannel(ch1)

		time.Sleep(1 * time.Second)

	*/

	//五、锁机制
	//taskTwo.Calculate()

	taskTwo.Increment()
}

func main() {

	//taskOneTest()

	taskTwoTest()

}
