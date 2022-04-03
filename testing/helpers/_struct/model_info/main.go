package main

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/go-core/core/clients/db/helper"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"log"
	"reflect"
)

func main() {
	_id, _ := uuid.NewRandom()

	data := &Terminal{
		Name: "Super Terminal",
		ID:   _id,
	}

	modelStruct := _struct.GetPointerStructValue(data)

	t := &Test{
		//ModelStruct: Terminal{},
		ModelStruct: modelStruct,
		Model:       data,
	}

	log.Print("structure??", _struct.GetPointerStructValue(t.Model))

	log.Println("pointer struct type -> ", reflect.Indirect(reflect.ValueOf(&t.Model)).Elem().Type())
	log.Println("pointer struct kind -> ", reflect.Indirect(reflect.ValueOf(&t.Model)).Elem().Kind())
	log.Println("pointer struct kind -> ", reflect.Indirect(reflect.ValueOf(&t.Model)).Interface())
	log.Println("pointer struct kind -> ", _struct.GetPointerStructValue(t.Model))

	dHelper := _struct.New(data)

	log.Println("ModelStruct -> is pointer", _struct.IsPointer(t.ModelStruct))
	log.Println("NonPtrObj -> is pointer", _struct.IsPointer(t.Model))
	log.Println("NonPtrObj -> get field value Name -> ", dHelper.GetFieldValue("Name"))
	log.Println("NonPtrObj -> get field value ID -> ", dHelper.GetFieldValue("ID"))
	//primaryKeys := helper.GetModelPrimaryKeys(Terminal{})
	primaryKeys := helper.GetModelPrimaryKeys(t.ModelStruct)
	if len(primaryKeys) == 0 {
		panic("record model doesn't have any primary keys")
	}

	// create the vars

	//mHelper := _struct.New(Terminal{})
	mHelper := _struct.New(t.ModelStruct)

	for _, pkey := range primaryKeys {
		dbDefaultFieldVal := mHelper.GetFieldTagKeyValue(pkey, "gorm", "default")
		fieldType := mHelper.GetFieldTypeName(pkey)
		log.Println(
			"field -> "+pkey,
			"default val -> "+dbDefaultFieldVal,
			"field type -> "+fieldType,
		)
	}

}
