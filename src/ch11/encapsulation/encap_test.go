package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

//func (e Employee) StringFormat() string {
//	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID: %d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
//}

func (e *Employee) StringFormat() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %d-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{1, "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = 2
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)

	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

func TestStructOperations(t *testing.T) {
	e := Employee{5, "Bob", 20}
	fmt.Printf("Address is %x\n", &e.Name)
	t.Log(e.StringFormat())
}
