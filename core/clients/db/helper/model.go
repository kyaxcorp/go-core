package helper

import "github.com/KyaXTeam/go-core/v2/core/helpers/_struct"

func GetModelPrimaryKeys(model interface{}) []string {
	return _struct.New(model).GetFieldNamesByTagKeyExistence("gorm", "primaryKey")
}
