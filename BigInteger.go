package BigInteger
//package main

import (
	"fmt"
)

//大整数
type BigInteger string

//乘以-1，改变符号，0不变
func (num BigInteger) Flip() BigInteger {
	if num == "0" {
		return num
	}
	if num[0] == '-' {
		return num[1:]
	}
	return "-" + num
}

//绝对值
func (num BigInteger) Abs() BigInteger {
	if num[0] == '-' {
		return num[1:]
	}
	return num
}

//a,b的最大公因数
func Gcd(a, b BigInteger) BigInteger {
	if _, k := a.Divide(b); k == "0" {
		return b
	} else {
		return Gcd(b, k)
	}
}

//num1是否大于num2,是返回true
func (num1 BigInteger) GreaterThan(num2 BigInteger) bool {
	if num1[0] == '-' && num2[0] == '-' {
		return num2.Abs().GreaterThan(num1.Abs())
	}
	if num1[0] == '-' {
		return false
	}
	if num2[0] == '-' {
		return true
	}
	if len(num1) > len(num2) {
		return true
	}
	if len(num1) < len(num2) {
		return false
	}
	for i := 0; i < len(num1); i++ {
		if num1[i] > num2[i] {
			return true
		}
		if num1[i] < num2[i] {
			return false
		}
	}
	return false //两个数相等的情况
}

//乘法
func (num1 BigInteger) Multiply(num2 BigInteger) BigInteger {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	negative := false //0规定为true
	if num1[0] == '-' && num2[0] == '-' {
		num1 = num1[1:]
		num2 = num2[1:]
	} else if num1[0] == '-' {
		negative = true
		num1 = num1[1:]
	} else if num2[0] == '-' {
		negative = true
		num2 = num2[1:]
	}

	//n位数乘以m位数,结果最多有n+m位数
	length := len(num1) + len(num2)
	ans := make([]int, length)
	ret := make([]byte, 0, length+1) //负数多一个符号位
	//模拟乘法运算
	for i := 0; i < len(num2); i++ {
		for j := 0; j < len(num1); j++ {
			ans[length-1-j-i] += int(num2[len(num2)-1-i]-'0') * int(num1[len(num1)-1-j]-'0')
		}
	}
	flag := 0 //排除前导0
	for ans[flag] == 0 {
		flag++
	}
	carry := 0 //进位
	for i := length - 1; ; i-- {
		ans[i] += carry
		carry = ans[i] / 10
		ans[i] = ans[i] % 10
		if i < flag && carry == 0 {
			break
		}
	}
	//排除前导0，将int切片转换为byte切片
	flag = 0
	for ans[flag] == 0 {
		flag++
	}
	if negative {
		ret = append(ret, '-')
	}
	for i := 0; flag < length; i++ {
		ret = append(ret, byte(ans[flag]+'0'))
		flag++
	}
	return BigInteger(string(ret))
}

//加法
func (num1 BigInteger) Add(num2 BigInteger) BigInteger {

	if num1[0] == '-' && num2[0] == '-' {
		return "-" + num1.Abs().Add(num2.Abs())
	}
	if num1[0] == '-' {
		return num2.Subtract(num1.Abs())
	}
	if num2[0] == '-' {
		return num1.Subtract(num2.Abs())
	}
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}
	length := max(len(num1), len(num2))
	ans := make([]int, length+1)
	i := 0
	for (len(num1)-1-i) >= 0 && (len(num2)-1-i) >= 0 {
		ans[len(ans)-1-i] = int(num1[len(num1)-1-i]-'0') + int(num2[len(num2)-1-i]-'0')
		i++
	}
	if len(num1) > len(num2) {
		for len(num1)-1-i >= 0 {
			ans[len(ans)-1-i] = int(num1[len(num1)-1-i] - '0')
			i++
		}
	} else if len(num1) < len(num2) {
		for len(num2)-1-i >= 0 {
			ans[len(ans)-1-i] = int(num2[len(num2)-1-i] - '0')
			i++
		}
	}
	flag := 0 //排除前导0
	for ans[flag] == 0 {
		flag++
	}
	carry := 0
	for i = len(ans) - 1; ; i-- {
		ans[i] += carry
		carry = ans[i] / 10
		ans[i] = ans[i] % 10
		if i < flag && carry == 0 {
			break
		}
	}
	ret := make([]byte, 0, len(ans)+1) //负数多一个符号位
	//排除前导0，将int切片转换为byte切片
	flag = 0
	for ans[flag] == 0 {
		flag++
	}
	for i := 0; flag < len(ans); i++ {
		ret = append(ret, byte(ans[flag]+'0'))
		flag++
	}
	return BigInteger(string(ret))
}

//减法(主体逻辑为num1减num2,num1,num2均为正数)
func (num1 BigInteger) Subtract(num2 BigInteger) BigInteger {
	if num1 == num2 {
		return "0"
	}
	if num1[0] == '-' && num2[0] == '-' {
		return num2.Abs().Subtract(num1.Abs())
	}
	if num1[0] == '-' {
		return "-" + num1.Abs().Add(num2.Abs())
	}
	if num2[0] == '-' {
		return num1.Add(num2.Abs())
	}
	//以下操作需确保num1,num2不为负数
	if num2.GreaterThan(num1) {
		return "-" + num2.Subtract(num1)
	}
	length := max(len(num1), len(num2))
	ans := make([]int, length+1)
	//此时num1一定大于等于num2
	i := 0
	for (len(num2) - 1 - i) >= 0 {
		ans[len(ans)-1-i] = int(num1[len(num1)-1-i]-'0') - int(num2[len(num2)-1-i]-'0')
		i++
	}
	for len(num1)-1-i >= 0 {
		ans[len(ans)-1-i] = int(num1[len(num1)-1-i] - '0')
		i++
	}
	flag := 0 //排除前导0
	for ans[flag] == 0 {
		flag++
	}
	carry := 0
	for i = len(ans) - 1; ; i-- {
		ans[i] += carry
		if ans[i] < 0 {
			carry = -1
			ans[i] += 10
		} else {
			carry = 0
		}
		if i < flag && carry == 0 {
			break
		}
	}
	ret := make([]byte, 0, len(ans))
	//排除前导0，将int切片转换为byte切片
	flag = 0
	for ans[flag] == 0 {
		flag++
	}
	for i := 0; flag < len(ans); i++ {
		ret = append(ret, byte(ans[flag]+'0'))
		flag++
	}
	return BigInteger(string(ret))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//除法 num1/num2,返回商和余数
//负数的除法仍然满足 A=kB+c ,k为商，c为余数
func (num1 BigInteger) Divide(num2 BigInteger) (BigInteger, BigInteger) {
	if num2 == "0" {
		panic("除数不能为0")
	}
	if num1 == num2 {
		return "1", "0"
	}
	if num1[0] == '-' && num2[0] == '-' {
		return num1.Abs().Divide(num2.Abs())
	}
	if num1[0] == '-' {
		Q, R := num1.Abs().Divide(num2)
		return Q.Flip(), R.Flip() //被除数为负数，商和余数结果全部为负数(0为本身)
	}
	if num2[0] == '-' {
		Q, R := num1.Divide(num2.Abs())
		return Q.Flip(), R //除数为负数，商的结果为负数，余数为正数(0为本身)
	}
	if num2.GreaterThan(num1) {
		return "0", num1
	}
	//此时num1大于num2且都为正数,下面步骤模拟除法
	ret := make([]byte, 0, len(num1))
	var Quotients int               //商
	var Remainders BigInteger = "0" //余数
	var tem BigInteger = num1[0:1]
	for i := 0; i < len(num1); i++ {
		if Remainders != "0" {
			tem = Remainders + num1[i:i+1] //字符串拼接
		} else {
			tem = num1[i : i+1]
		}
		//此时Quotients商一定是个位数
		Quotients, Remainders = tem.simpleDivide(num2)
		ret = append(ret, byte(Quotients+'0'))
	}
	//排除前导0
	flag := 0
	for ret[flag] == '0' {
		flag++
	}
	return BigInteger(ret[flag:]), Remainders
}

//暴力简单，减法代替除法 ,返回商和余数，速度非常慢，只适用于商比较小的情况，不能处理负数的除法
func (num1 BigInteger) simpleDivide(num2 BigInteger) (int, BigInteger) {
	if num2 == "0" {
		panic("除数不能为0")
	}
	var i int = 0
	for !num2.GreaterThan(num1) {
		num1 = num1.Subtract(num2)
		i++
	}
	return i, num1
}

//High Precision 高精度除法,num1/num2，结果不会四舍五入
//n表示精度，计算到小数点后n位,若n太大可以改成直接输出，不反回string
func HP_Division(num1, num2 BigInteger, n int) string {
	if num1[0] == '-' && num2[0] == '-' {
		return HP_Division(num1.Abs(), num2.Abs(), n)
	}
	if num1[0] == '-' {
		return "-" + HP_Division(num1.Abs(), num2, n)
	}
	if num2[0] == '-' {
		return "-" + HP_Division(num1, num2.Abs(), n)
	}
	Q, R := num1.Divide(num2)
	if R == "0" {
		return string(Q) + ".0" //可以整除，无小数部分
	}
	ret := make([]byte, 0, len(Q)+n+1)
	ret = append(ret, Q...) //整数部分
	ret = append(ret, '.')  //小数点
	R = R.Pow10(n)
	Q, R = R.Divide(num2) //如果Q的值为"0",代表给定的精度不够高，返回的都是0.0000...0
	//小数部分
	//加上小数开头的0，解决0.00的问题(如2/1000)
	if len(Q) != n {
		for i := 0; i < n-len(Q); i++ {
			ret = append(ret, '0')
		}
	}
	if R == "0" { //表示除的尽，是有限小数
		fmt.Println("能除尽，是有限小数")
		//排除末尾的0
		flag := len(Q) - 1
		for Q[flag] == '0' {
			flag--
		}
		ret = append(ret, Q[:flag+1]...)
	} else {
		ret = append(ret, Q...)
	}
	return string(ret)
}

//乘以10的n次方
func (num1 BigInteger) Pow10(n int) BigInteger {
	zero := make([]byte, n)
	for i := 0; i < n; i++ {
		zero[i] = '0'
	}
	return num1 + BigInteger(zero)
}

//小数相加，只处理正数
func DecimalAdd(a, b string) string {
	flagA := 0
	flagB := 0
	for flagA = 0; a[flagA] != '.'; flagA++ {
	}
	for flagB = 0; a[flagB] != '.'; flagB++ {
	}
	integerA := a[0:flagA]
	decimalA := a[flagA+1:]
	integerB := b[0:flagB]
	decimalB := b[flagB+1:]
	//a为0.0000这种0值，直接返回b
	if integerA == "0" {
		i := len(decimalA) - 1
		for ; i > -1; i-- {
			if decimalA[i] != '0' {
				break
			}
		}
		if i == -1 {
			return b
		}
	}
	//b为0.0000这种0值，直接返回a
	if integerB == "0" {
		i := len(decimalB) - 1
		for ; i > -1; i-- {
			if decimalB[i] != '0' {
				break
			}
		}
		if i == -1 {
			return a
		}
	}
	if len(decimalA) > len(decimalB) {
		zeroB := make([]byte, len(decimalA)-len(decimalB))
		for i := 0; i < len(zeroB); i++ {
			zeroB[i] = '0'
		}
		decimalB = decimalB + string(zeroB)
	} else if len(decimalA) < len(decimalB) {
		zeroA := make([]byte, len(decimalB)-len(decimalA))
		for i := 0; i < len(zeroA); i++ {
			zeroA[i] = '0'
		}
		decimalA = decimalA + string(zeroA)
	}

	length := len(decimalA)
	ans := make([]int, length+1)
	i := 0
	for (len(decimalA)-1-i) >= 0 && (len(decimalB)-1-i) >= 0 {
		ans[len(ans)-1-i] = int(decimalA[len(decimalA)-1-i]-'0') + int(decimalB[len(decimalB)-1-i]-'0')
		i++
	}
	flag := 0 //排除前导0
	for ans[flag] == 0 {
		flag++
	}
	carry := 0
	for i = len(ans) - 1; ; i-- {
		ans[i] += carry
		carry = ans[i] / 10
		ans[i] = ans[i] % 10
		if i < flag && carry == 0 {
			break
		}
	}
	ret := make([]byte, 0, len(ans))
	for i := 0; i < len(ans); i++ {
		ret = append(ret, byte(ans[i]+'0'))
	}
	if ret[0] == '0' {
		return string(BigInteger(integerA).Add(BigInteger(integerB))) + "." + string(ret[1:])
	} else {
		return string(BigInteger(integerA).Add(BigInteger(integerB)).Add("1")) + "." + string(ret[1:])
	}
}

// func main() {
// 	var (
// 		num1 BigInteger
// 		num2 BigInteger
// 	)
// 	for i := 0; i < 100; i++ {
// 		fmt.Scan(&num1, &num2)
// 		fmt.Println("multiply:", num1.Multiply(num2))
// 		fmt.Println("add:", num1.Add(num2))
// 		fmt.Println("subtract:", num1.Subtract(num2))
// 		Quotients, Remainders := num1.Divide(num2)
// 		fmt.Println("Divide:", Quotients, Remainders)
// 		s := HP_Division(num1, num2, 5000)
// 		fmt.Println("HP_Division:", s)
// 	}
// }

// func main() {
// 	var (
// 		s1 BigInteger
// 		s2 BigInteger
// 	)
// 	for i := 0; i < 100; i++ {
// 		fmt.Scan(&s1, &s2)
// 		s := HP_Division(s1, s2, 5000)
// 		fmt.Println("HP_Division:", s)
// 	}
// }
