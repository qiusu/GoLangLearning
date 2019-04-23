//helloworld 程序

package main
import (
	"fmt"
	_ "math"
)
/*
var ja = 11
var jt = 11.11
var hy = "Syty"
*/
var (
	ja = 11 
	jt = 11.11
	hy = "Syty"
)
func main(){
	// fmt.Println("hello World")
	// fmt.Println("==================")
	// fmt.Println("姓名\t年龄\t籍贯\t住址\njohn\t12\t河北\t北京")
	// fmt.Println("==================")
	// var i int
	// fmt.Println("i=", i)
	// var t = 22.22
	// fmt.Println("t=", t)
	// name := "tom"
	// fmt.Println("name==", name)
	// fmt.Println("==================")
	// val1, val2, val3 := 11.11, "JOJO", 58
	// fmt.Println("val1=",val1, ",va2=",val2, ",val3=",val3)
	// fmt.Println("==================")
	// ja = 22
	// fmt.Println("ja=",ja, ",jt=", jt, ",hy=", hy)
	// fmt.Println("==================")
	// var j int = 1
	// var x int = 3
	// var r int = j + x
	// fmt.Println("r=", r)
	// fmt.Printf("r=%T", r)
	// var r2 = hy + val2
	// fmt.Println("r2=", r2)
	// fmt.Println("==================")
	// var n2 int64 = 10
	// fmt.Println("n2的数据类型是 %T , 占用字节数是 %d", n2,unsafe.Sizeof(n2))
	// var x_i float32 = 99.0 / 10
	// fmt.Println("x_i==", x_i)

	// var allDay int = 97
	// var weekNum int = allDay / 7
	// var dayNum int = allDay % 7
	// fmt.Printf("还有%d周%d天放假", weekNum, dayNum)

	// var age int = 10
	// if age > 3{
	// 	fmt.Println("age==", age)
	// }
	// if !(age>99){
	// 	fmt.Println("age=======", age)
	// }

	// var (
	// 	name string
	// 	age int
	// 	sal float32
	// 	isPass bool
	// ) 
	// fmt.Println("请输入姓名:")
	// fmt.Scanln(&name)
	// fmt.Println("请输入年龄:")
	// fmt.Scanln(&age)
	// fmt.Println("请输入薪资:")
	// fmt.Scanln(&sal)
	// fmt.Println("请输入是否通过考试:")
	// fmt.Scanln(&isPass)
	// fmt.Printf("姓名:%v \n年龄:%v \n薪资:%v \n是否通过考试:%v \n", name, age, sal, isPass)

	// fmt.Println("姓名： 年龄: 薪资: 是否通过考试: ")
	// fmt.Scanf("%s %d %f %t",&name, &age, &sal, &isPass)
	// fmt.Printf("姓名:%v \n年龄:%v \n薪资:%v \n是否通过考试:%v \n", name, age, sal, isPass)

	// var age int
	// fmt.Println("请输入年龄")
	// fmt.Scanf("%d",&age)

	// if age > 18 {
	// 	fmt.Println("你已经成年了，你要养活自己")
	// } else {	
	// 	fmt.Println("你还是个孩子")
	// }

	// var yearNum int
	// fmt.Println("请输入年份：")
	// fmt.Scanln(&yearNum)

	// if (yearNum % 4 ==0 && yearNum % 100 !=0) || (yearNum % 400 ==0) {
	// 	fmt.Printf("%d是闰年",yearNum)
	// } else {
	// 	fmt.Printf("%d不是闰年",yearNum)
	// }

	// var score float32
	// fmt.Println("请输入成绩")
	// fmt.Scanln(&score)

	// if score == 100 {
	// 	fmt.Println("宝马拿去吧")
	// } else if score>80 && score <= 99 {
	// 	fmt.Println("iphone7用用好了")
	// } else if score >= 60 && score <=80 {
	// 	fmt.Println("ipad凑活着用用吧")
	// } else {
	// 	fmt.Println("滚去学习")
	// }

	// var Skey byte 
	// fmt.Println("请输入：a,b,c,d 中的一个")
	// fmt.Scanf("%c",&Skey)

	// switch Skey {
	// 	case 'a':
	// 		fmt.Printf("周一===>%c", Skey)
	// 	case 'b':
	// 		fmt.Printf("周二===》%c", Skey)
	// 	case 'c':
	// 		fmt.Printf("周三===》%c", Skey)
	// 	case 'd':
	// 		fmt.Printf("周四===》%c", Skey)
	// 	default:
	// 		fmt.Printf("你输入的%c不在范围内", Skey)
	// }
	
	// var num int = 10
	// switch num {
	// 	case 10 :
	// 		fmt.Printf("周一===>%d", num)
	// 		fallthrough
	// 	case 20 :
	// 		fmt.Printf("周二===》%d", num)
	// 	case 30 :
	// 		fmt.Printf("周三===》%d", num)
	// 	default:
	// 		fmt.Println("你输入错误")
	// }

	// for i := 1;i <= 10; i++ {
	// 	fmt.Println("你好，每光")
	// }

	// var str string = "hello world上海"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c \n",str[i])
	// }

	// for index, val := range str {
	// 	fmt.Printf("index=%d,val=%c \n",index, val)
	// }

	// str2 := []rune(str)
	// for i := 0;i < len(str2);i++ {
	// 	fmt.Printf("%c \n",str2[i])
	// }

	// var (
	// 	sum uint32
	// 	count uint32
	// 	i uint32 = 1
	// )

	// for ; i <= 100; i++ {
	// 	if i % 9 == 0 {
	// 		count++
	// 		sum += i
	// 		fmt.Println("========",i)
	// 	}
	// }
	// fmt.Printf("总共%d个数，总和为 %d",count, sum)

	// result := cal(11.11, 22.22, '+')
	// fmt.Printf("结果是：%f", result)

	// n1, n2 := 10, 22
	// test(n1)
	// result := getSum(n1, n2)
	// fmt.Println("result==",result)

	// _, sub := getSumAndSub(n1,n2)
	// fmt.Printf("sub==%d",sub)

	var n int
	fmt.Println("请输入数字")
	//fmt.Scanln(&n)
	fmt.Scanf("%d",&n)

	fmt.Println("斐波那契数===",getFeiboNum(n))
	

}

func cal(n1 float32, n2 float32, operate byte) float32 {
	switch operate {
		case '+' : 
			return n1 + n2
		
		case '-' : 
			return n1 - n2
		
		case '*' : 
			return n1 * n2
		
		case '/' : 
			return n1 / n2
		
		default : 
			return 0
		
	}
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum方法中 sum==",sum)
	return sum
}

func test(n1 int) {
	n1++
	fmt.Println("test方法中 n1==",n1)
}

func getSumAndSub(n1 int ,n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func getFeiboNum(n int) int {
	if (n == 1 || n==2) {
		return 1
	} else {
		return getFeiboNum(n-1) + getFeiboNum(n-2)
	}
}