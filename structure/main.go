package main

import (
	"fmt"
)

// // User의 크기는 sting + int 의 바이트 크기
// type User struct {
// 	Name string
// 	Age  int
// }

// func nameAndAge(uid int) User {
// 	switch uid {
// 	case 0:
// 		return User{"User0", 00}
// 	case 1:
// 		return User{"User1", 11}
// 	default:
// 		return User{"Unknown", -1}
// 	}

// }

// // User 구조체를 복사할 뿐. 실제로는 값 변경X
// func incrementAge(user User) {
// 	user.Age++
// 	fmt.Println(user.Age)
// }

// // 포인터를 사용하는 것이 아니라 참조하는 것임. 이것이 Go....
// func incrementAge2(user *User) {
// 	user.Age++
// 	fmt.Println(user.Age)
// }

// // 함수 키워드와 함수 이름사이에 과롷로 써인 형태의 새로운 인자 이름과 타입이 존재
// // Go에게 user 구조체 위에 인스턴스 메서드를 만들고 해당 인스턴스를 user라는 이름의 함수에 전달해~ 라는 의미
// func (user User) prettyString() string {
// 	return fmt.Sprintf("Name: %s, Age: %d", user.Name, user.Age)
// }

// func (user *User) incrementAge3() {
// 	user.Age++
// 	fmt.Println(user.Age)
// }

// func main() {
// 	//user := nameAndAge(123)
// 	jun := User{"Jun", 20}
// 	//incrementAge(jun)
// 	incrementAge2(&jun)
// 	fmt.Println(jun.Age)
// 	fmt.Println(jun.prettyString())
// 	jun.incrementAge3()
// 	fmt.Println(jun.Age)
// }

// =========================================================
// 인터페이스
// =========================================================
type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name  string
	Owner *Person
	Age   int
}

func (person *Person) incrementAge() {
	person.Age++
}

func (person *Person) getAge() int {
	return person.Age
}

func (dog *Dog) incrementAge() {
	dog.Age++
}

func (dog *Dog) getAge() int {
	return dog.Age
}

// 정의된 기능만 있고 구현은 없는 것
// 모두 포인터 리시버이거나 모두 값 리시버임. 섞을 수 없음.
type Living interface {
	incrementAge()
	getAge() int
}

func incrementAndPrintAge(being Living) {
	being.incrementAge()
	fmt.Println(being.getAge())
}

func main() {
	a := Person{"Jun", 20}
	b := Dog{"Max", &a, 2}
	incrementAndPrintAge(&a)
	incrementAndPrintAge(&b)
}
