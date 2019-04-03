package cache

import (
	"fmt"
	"reflect"
	"unsafe"
)

type HEAD struct {
	pre, next *HEAD
}

type T struct {
	a    int
	b    int
	head HEAD
}

func offset(typeOf reflect.Type, head *HEAD) {
	filed, ok := typeOf.FieldByName("head")
	if !ok {
		return
	}
	fmt.Println(filed.Offset)
	ptr := unsafe.Pointer(uintptr(unsafe.Pointer(head)) - filed.Offset)
	fmt.Println((*T)(ptr))
}

func main() {
	head := HEAD{}
	t := T{head: head, a: 1, b: 999}
	offset(reflect.TypeOf((*T)(nil)).Elem(), &t.head)
	fmt.Println(uintptr(unsafe.Pointer(&t)), uintptr(unsafe.Pointer(&t.head)), uintptr(unsafe.Pointer(&t.head))-uintptr(unsafe.Pointer(&t)))
}
