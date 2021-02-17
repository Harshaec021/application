package basictypes

import (
	fm "fmt"

	sample "github.com/elliotforbes/test-package"
)

func main() {
	CompositeTypes()
}

// CompositeTypes func
func CompositeTypes() {

	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	num := [...]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(days); i++ {
		fm.Println(days[i])
	}
	for i := 0; i < len(num); i++ {
		fm.Println(num[i])
	}

	userAge := map[string]uint8{
		"Harry": 12,
		"James": 45,
		"Lilly": 13,
	}

	for k, v := range userAge {
		fm.Println(k, v)
	}

	type User struct {
		userName string
		userAge  uint8
	}

	myUser := User{userName: "Harry", userAge: 12}

	fm.Println(myUser.userAge)
	fm.Println(myUser.userName)
	myUser.userAge = 15
	fm.Println(myUser.userAge)

}

func basicTypes() {
	fm.Println("Hello world")
	sample.MySampleFunc()
	var flag bool
	for i := 0; i < 8; i++ {
		flag = i%2 == 0
		if flag {
			fm.Println("Even : ", i)
			subscribeToMycHANNEL()
		} else {
			fm.Println("odd : ", i)
			unSubscribeToMycHANNEL()
		}

		fm.Println(float64(i))

	}
}

func subscribeToMycHANNEL() {
	fm.Println("SUBSCRIBED TO MY CHANNEL")
}

func unSubscribeToMycHANNEL() {
	fm.Println("UnSUBSCRIBED TO MY CHANNEL")
}
