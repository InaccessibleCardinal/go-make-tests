package svc

type Person struct {
	Name string
	Age  int
}

func DoTheNeedful(person Person) string {
	if person.Age > 40 {
		return "old"
	}
	return "young"
}
