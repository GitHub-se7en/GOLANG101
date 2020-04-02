package main

type student struct {
	Name string
	Age  int
}

func pase_student() {
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range stus {
		stu.Age = stu.Age + 10
		stus[i] = stu
	}
	for k, v := range stus {
		println(k, "=>", v.Age)
	}
}
func pase_student2() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = &stu
		println(&stu)
	}
	for k, v := range m {
		println(k, "=>", v.Age)
	}
}
func pase_student3() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// 正确
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for k, v := range m {
		println(k, "=>", v.Name)
	}
}

func main() {
	pase_student2()
}
