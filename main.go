package main

import (
	"fmt"
	"time"
)

func main() {
	// e là khai báo của 1 thang máy, với presentFloor là vị trí hiện tại của thang máy trong tòa nhà
	// ElevatorRunTime là hàm mô phỏng sự di chuyển của thang máy từ vị trí hiện tại (presentFloor) đến giá trị tầng cần đến

	Test()

}
func Test() {
	e := Elevators{presentFloor: 0}
	ElevatorRunTime(&e, 2)
	go ElevatorRunTime(&e, 7)

}

type Elevators struct {
	presentFloor int
	run          bool
}

func GetPresentFloor(e *Elevators) int {
	return e.presentFloor
}

func ElevatorRunTime(e *Elevators, floor int) {
	presentFloor := GetPresentFloor(e)

	fmt.Println("thang máy đang chạy")
	if presentFloor < floor {
		e.run = true
		for i := presentFloor; i <= floor; i++ {
			fmt.Println("--------------->", i)
			time.Sleep(2 * time.Second)
		}
		e.presentFloor = floor
		e.run = false
		fmt.Println("Bạn đã đến")
	} else if presentFloor > floor {
		e.run = true
		for i := presentFloor; i >= floor; i-- {
			fmt.Println("--------------->", i)
			time.Sleep(2 * time.Second)
		}
		e.presentFloor = floor
		e.run = false
		fmt.Println("Bạn đã đến")
	} else if presentFloor == floor {
		fmt.Println("Bạn đã đến")
	}
}
func CallElevator(e *Elevators, floor int) {
	// gọi thang máy  - nếu thang máy đang di chuyển và sẽ ghé ngang tầng này - thang máy sẽ dừng lại để nhận khách
	// bài toán sẽ có thể mở rộng nếu thang chứa đủ hành khách hay các yếu tố khác
}
