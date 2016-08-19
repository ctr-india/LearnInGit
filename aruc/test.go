package main

// test 
import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"time"
)

// hawky hawky

func findSumTillNum(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}

func findMulTillNum(num int) (mul int) {
	mul, n := 1, 2
	for n <= num {
		mul *= n
		n++
	}
	return
}

func checkOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Os is darwin")
	case "linux":
		fmt.Println("Os is linux")
	default:
		fmt.Println("Os is ?!! : ", os)
	}
}

func gudGreeting() {
	switch {
	case time.Hour < 12:
		fmt.Println("Good morning !!")
	case time.Hour < 17:
		fmt.Println("Good after noon !!")
	default:
		fmt.Println("Good evening")
	}
}

func copyFiles(srcFile, dstFile string) (writter int64, err error) {
	fmt.Println("test")
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}

func main() {
	fmt.Println("Control statements")
	fmt.Println("Sum till 10 : -- ", findSumTillNum(10))
	fmt.Println("Mul till 5  : -- ", findMulTillNum(5))
	if findSumTillNum(3) < findMulTillNum(3) {
		fmt.Println("Hoo.. multiple is biggy it seems")
	} else {
		fmt.Println("Hoo Addition is biggy it seems")
	}

	if v := math.Pow(4, 2); v > 15 {
		fmt.Println("of course 16 is bigger than 15 man :)")
	}

	checkOS()

	gudGreeting()

	// defer test
	fmt.Println("1 st")
	fmt.Println("2 nd")
	defer fmt.Println("3rd but I am deffered")
	fmt.Println("4th but became 3rd")
	fmt.Println("5th but became 4th")

	fmt.Println("\n\n++++++ Defer test ++++++++")
	copied, err := copyFiles("/Users/arunbabu/GoStudy/src/Learn/testSrc", "/Users/arunbabu/GoStudy/src/Learn/testDst")
	if err == nil {
		fmt.Println("Copied successfully : ", copied)
	} else {
		fmt.Println("error : ", err)
	}
}
