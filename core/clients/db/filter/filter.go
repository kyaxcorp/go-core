package filter

import (
	"github.com/kyaxcorp/go-core/core/clients/db/helper"
	"github.com/kyaxcorp/go-core/core/helpers/_struct"
	"github.com/kyaxcorp/go-core/core/helpers/err/define"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	//"log"
	"strings"
	"sync"
)

type InputModel struct {
	Name  string
	Model interface{}
}

func (f *Input) Apply() *Input {
	f.check()
	f.ApplyConditions()
	f.applyDefaultScope()
	f.ApplyPagination()
	f.ApplyOrdering()
	return f
}

func (f *Input) check() {
	// Check and set default values!
	if f.PageNr == nil {
		f.PageNr = new(int)
		*f.PageNr = 1
	}
	if f.NrOfItems == nil {
		f.NrOfItems = new(int)
		*f.NrOfItems = 10
	}

	if *f.PageNr <= 0 {
		*f.PageNr = DefaultPageNr
	}

	if *f.NrOfItems <= 0 {
		*f.NrOfItems = DefaultNrOfItems
	}

	f.checkContext()
}

/*type InputModel interface {
	// Each GORM NonPtrObj has this function or should have at least...
	TableName() string
}*/

func (f *Input) SetModels(models ...InputModel) *Input {
	if f.models == nil {
		f.models = make(map[string]cachedModel)
	}
	if f.cachedDBFields == nil {
		f.cachedDBFields = make(map[string]string)
	}

	//f.db
	//f.db.NewScope(model).TableName()

	/*
		TODO:
			We should represent:
			- What's the model name in GraphQL (as the user sees it)
			- What's the model structure to which is bind!
			-
	*/

	// Parse the model table name

	for _, inputModel := range models {
		// Get the model native Type/Name

		model := inputModel.Model
		modelName := _struct.GetName(model)

		// Get the model DB Table Name
		tableName, _err := helper.GetModelDBTableName(f.db, model)
		if _err != nil {
			panic("failed to parse model table name")
		}
		// Save the model as it is

		inputModelName := strings.ToLower(inputModel.Name)

		if f.primaryModelName == "" {
			f.primaryModelName = inputModelName
		}

		f.models[inputModelName] = cachedModel{
			modelName:   modelName,
			model:       model,
			dbColumns:   make(map[string]string),
			dbTableName: tableName,
		}

		// Let's also parse model Columns!
		//modelFields := _struct.GetExportableFieldsByJSONTag(model)
		//lowerCaseModelFields := make(map)
		//for _,fieldName := range modelFields{
		//
		//}

		s, err := schema.Parse(model, &sync.Map{}, schema.NamingStrategy{})
		if err != nil {
			panic("failed to create schema")
		}

		//m := make(map[string]string)
		found := make(map[string]bool)
		for _, field := range s.Fields {
			dbFieldName := field.DBName
			modelFieldName := field.Name
			//m[modelFieldName] = dbFieldName

			filteredDbFieldName := strings.ToLower(modelFieldName)

			// We will lowercase the field name from the structure/model!
			f.models[inputModelName].dbColumns[filteredDbFieldName] = dbFieldName

			if f.enableDBFieldsCaching {
				// Save as it is
				f.cachedDBFields[filteredDbFieldName] = dbFieldName
				// Save with model name and field name
				f.cachedDBFields[inputModelName+"."+filteredDbFieldName] = dbFieldName
			}

			found[modelFieldName] = true

			//log.Println(modelFieldName + " -> " + dbFieldName)
		}

		// Try also getting from tags...

		/*mHelper := _struct.New(model)
		fields := mHelper.Map()
		for fieldName, _ := range fields {
			if _, ok := found[fieldName]; !ok {
				dbFieldName := mHelper.GetFieldTagKeyValue(fieldName, "gorm", "column")
				if dbFieldName != "" {
					// it exists, add manually

					modelFieldName := fieldName
					filteredDbFieldName := strings.ToLower(modelFieldName)

					// We will lowercase the field name from the structure/model!
					f.models[inputModelName].dbColumns[filteredDbFieldName] = dbFieldName

					if f.enableDBFieldsCaching {
						// Save as it is
						f.cachedDBFields[filteredDbFieldName] = dbFieldName
						// Save with model name and field name
						f.cachedDBFields[inputModelName+"."+filteredDbFieldName] = dbFieldName
					}

					found[modelFieldName] = true
				}
			}

		}*/

		//

		//
		/*result, _ := f.db.Migrator().ColumnTypes(model)
		for _, v := range result {
			// TODO: what to do with the column tag which can change entirely the name
			// 		for which we can't reverse the name...
			//		we need the function column types to return the original field name for which the conversion
			//		has been made!
			// 		or we should loop by ourselves and get the names

			dbFieldName := v.Name()
			filteredDbFieldName := strings.ReplaceAll(dbFieldName, "_", "")
			//log.Println(v.Name(), v.DatabaseTypeName(), filteredDbFieldName)

			f.models[inputModelName].dbColumns[filteredDbFieldName] = dbFieldName
			// Let's get it back to original name and after that map!
			//log.Println()
		}*/
	}

	return f
}

func (f *Input) getDBFieldName(fieldName string) (string, error) {
	// transform
	/*
		1. Check if there is a DOT (.) in the string
		if there is one, check the by this model name!
		Loop there to find the field name! if there is no
		field, then throw an error that field doesn't exist!

		2. If there is no DOT (.) search consecutively through
		all existing models for this field, the first found, first delivered...
		3. if still not found then throw an error...

	*/

	lowerFieldName := strings.ToLower(fieldName)

	//  Check cache here...
	if dbField, ok := f.cachedDBFields[lowerFieldName]; ok {
		return dbField, nil
	}

	if strings.Contains(fieldName, ".") {
		// Case 1
		splitted := strings.Split(lowerFieldName, ".")
		modelName := splitted[0]
		fName := splitted[1]
		if _, ok := f.models[modelName]; !ok {
			return "", define.Err(1, "this model doesn't exist -> "+modelName)
		}
		if _, ok := f.models[modelName].dbColumns[fName]; !ok {
			return "", define.Err(2, "this field doesn't exist -> "+fName)
		}

		return f.models[modelName].dbColumns[fName], nil
	} else {
		// Case 2

		found := false
		dbFieldName := ""
		for _, model := range f.models {
			//log.Println("model", model.dbColumns)
			if dbField, ok := model.dbColumns[lowerFieldName]; ok {
				found = true
				dbFieldName = dbField
				break
			}
		}
		if !found {
			return "", define.Err(3, "this field doesn't exist -> "+lowerFieldName)
		}
		return dbFieldName, nil
	}
}

func (f *Input) getDBFieldNameOrPanic(fieldName string) string {
	fName, _err := f.getDBFieldName(fieldName)
	if _err != nil {
		panic(_err.Error())
	}
	return fName
}

func (f *Input) DB() *gorm.DB {
	return f.db
}
