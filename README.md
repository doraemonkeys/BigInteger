# BigInteger
- golang实现的低性能大整数运算




```go
	var (
		num1 BigInteger
		num2 BigInteger
	)
```

## 加

```go
//func (BigInteger).Add(num2 BigInteger) BigInteger
num1.Add(num2)
```

## 减

```go
//func (BigInteger).Subtract(num2 BigInteger) BigInteger
num1.Subtract(num2)
```

## 乘

```go
//func (BigInteger).Multiply(num2 BigInteger) BigInteger
num1.Multiply(num2)
```

## 除

```go
//func (BigInteger).Divide(num2 BigInteger) (BigInteger, BigInteger)
num1.Divide(num2)
```

## 高精度除

```go
//func HP_Division(num1 BigInteger, num2 BigInteger, n int) string
HP_Division(num1, num2, 5000)
```

## 小数加法

```go
func DecimalAdd(a, b string) string 
```