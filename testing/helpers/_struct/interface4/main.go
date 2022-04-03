package main

import (
	"github.com/KyaXTeam/go-core/v2/core/helpers/_struct"
	"github.com/jinzhu/copier"
	"log"
	"reflect"
)

type InputData struct {
	Name string
	Age  int
}

type Animal struct {
	ID       string
	Name     string
	Age      int
	LastName string
}

type Record struct {
	// Data -> can be a map[string]interface{} or a real any Struct from where we will copy the data!
	Data interface{}
	// this case is used when the data is structure! but it should be saved as plain structure, because we don't
	// want to modify the input data that came in! So it should be created a copy!
	dataStr interface{}
	// this var is the modelStruct with copied dataStr (all the data from dataStr is copied to dataStrCopied)
	dataStrCopied interface{}

	// this is the structure of the set NonPtrObj, it's a pointer!
	ModelStruct interface{}
	// this is not a pointer! it's a plain one!
	modelStruct interface{}

	dbData interface{}
}

func CloneInterfaceItem(obj interface{}) interface{} {
	var i interface{}
	if _struct.IsPointer(obj) {
		i = reflect.Indirect(reflect.ValueOf(obj)).Interface()
	} else {
		i = obj
	}

	p := reflect.New(reflect.TypeOf(i))
	p.Elem().Set(reflect.ValueOf(i))
	return p.Interface()
}

func main() {
	inputData := &InputData{
		Name: "Octavian",
		Age:  28,
	}
	inputData2 := InputData{
		Name: "Octavian",
		Age:  28,
	}
	var i interface{}
	i = inputData2
	// It will create another one as pointer... to the structure?!
	p := reflect.New(reflect.TypeOf(i))
	p.Elem().Set(reflect.ValueOf(i))

	log.Println("p", p.Interface())

	log.Println("&i", &i)
	//log.Println("i addr", reflect.ValueOf(&i).Addr())
	log.Println("&inputData2", &inputData2)

	// Define the record!
	r := &Record{
		Data:        inputData,
		ModelStruct: &Animal{},
	}

	// We need the value because we want afterwards to copy the model!
	r.modelStruct = _struct.GetPointerStructValue(r.ModelStruct)
	// Get the value of the address that is pointed to structure
	//r.dataStr = _struct.GetPointerStructValue(r.Data)
	r.dataStr = CloneInterfaceItem(r.Data)
	//r.dataStr = reflect.ValueOf(r.Data).Elem().Interface()
	//r.dataStr = r.Data
	log.Println("r.Data", r.Data) // this is a pointer!
	log.Println("r.modelStruct", r.modelStruct)
	log.Println("r.dataStr ", r.dataStr)
	//log.Println("&r.dataStr ", &r.dataStr)
	log.Println("&r.dataStr ", CloneInterfaceItem(r.dataStr))

	// TODO: how to get a pointer of dataStr?!!

	// Create the model structure
	r.dataStrCopied = r.modelStruct
	r.dataStrCopied = CloneInterfaceItem(r.dataStrCopied)

	log.Println("r.dataStrCopied", r.dataStrCopied)
	log.Println("&r.dataStrCopied", &r.dataStrCopied)
	log.Println("r.dataStrCopied reflection", reflect.ValueOf(&r.dataStrCopied).Elem().CanAddr())

	_err := copier.Copy(r.dataStrCopied, CloneInterfaceItem(r.dataStr))
	if _err != nil {
		panic("failed to copy data from dataStr to dataStrCopied -> " + _err.Error())
	}

	log.Println("r.dataStrCopied", r.dataStrCopied)
}
