package main

import (
	"fmt"
	"go_homework_project/taskOne"
	"go_homework_project/taskThree"
	"go_homework_project/taskThree/config"
	"go_homework_project/taskThree/repository"
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

func taskThreeTest() {
	//初始化数据库
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	//gorm.DB
	if err := repository.InitDB(&cfg); err != nil {
		panic(err)
	}

	//Sqlx.DB
	if err := repository.InitDB_Sqlx(&cfg); err != nil {
		panic(err)
	}

	//taskThree.CreateStudent()
	//taskThree.GetAgeOver18()
	//taskThree.UpdateGrand("张三", "四年级")
	//taskThree.DeleteAgeSmall15()

	//转账
	/*
		if err := taskThree.TxTransfer("B", "A", 100); err != nil {
			fmt.Println("转账失败：", err)
		} else {
			fmt.Println("转账成功")
		}
	*/

	//使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息
	/*employees := taskThree.GetEmpByDept("技术部")
	for i := range employees {
		fmt.Println("技术部员工信息：", employees[i].ID, employees[i].Name, employees[i].Department, employees[i].Salary)
	}*/

	//使用Sqlx查询 employees 表中工资最高的员工信息
	/*employee := taskThree.GetHighestSalary()
	fmt.Println("工资最高的员工信息：", employee.ID, employee.Name, employee.Department, employee.Salary)
	*/

	//查询价格大于 50 元的书籍
	/*books := taskThree.GetBookByPriceOver(50)
	for i := range books {
		fmt.Println("价格大于50的书籍信息是：", books[i].ID, books[i].Title, books[i].Author, books[i].Price)
	}*/

	//查询某个用户发布的所有文章及其对应的评论信息
	//taskThree.GetAllInfoByUserName("作者3")

	//查询评论数量最多的文章信息
	/*hottestPosts := taskThree.GetHottestPostInfo()
	for i := range hottestPosts {
		fmt.Printf("评论最多的文章id：%d，文章名称：%s，作者：%s，评论数量：%d \n",
			hottestPosts[i].ID, hottestPosts[i].Title, hottestPosts[i].Author, hottestPosts[i].CommentNum)
	}*/

	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	/*posts := []taskThree.Post{
		{Title: "作者1的第四篇文章", Content: "思思上", UserID: 27},
		{Title: "作者1的第五篇文章", Content: "舞舞舞", UserID: 27},
	}
	taskThree.AddPost(posts)*/

	/*newUsers := []taskThree.User{
		{Name: "作者4",
			Posts: []taskThree.Post{
				{Title: "作者4的第一篇文章", Content: "思乡"},
				{Title: "作者4的第二篇文章", Content: "晓"},
			}},
		{Name: "作者5",
			Posts: []taskThree.Post{
				{Title: "作者5的第一篇文章", Content: "咏楼"},
			}},
	}
	taskThree.AddUserAndPost(newUsers)*/

	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	taskThree.DeleteCommentsById([]uint{25, 27})

}

func main() {

	//taskOneTest()

	//taskTwoTest()

	taskThreeTest()
	//taskThree.InitTable()
	//taskThree.InitData()
	//taskThree.InitDataBySqlx()
	//taskThree.InitDataBySqlx1()
	//taskThree.InitDataBlog()

}
