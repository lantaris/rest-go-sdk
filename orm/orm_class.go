package orm

import (
	"gorm.io/gorm"
	"reflect"
)

// ********************************************************
func GetStructAttr(obj interface{}, fieldName string) interface{} {

	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(fieldName)
	return f.Interface()
}

// ********************************************************
func (ORM *TORM) Select(RetData interface{}, SelectFields map[string]interface{}) (err error) {
	var (
		result *gorm.DB
	)

	result = ORM.ORM.Where(SelectFields).Find(RetData)
	err = result.Error

	return err
}

// ********************************************************
func (ORM *TORM) Update(Data interface{}, SelectCriteria map[string]interface{}, UpdateFieldList []string) error {
	var (
		UpdateData map[string]interface{}
		result     *gorm.DB
	)

	UpdateData = make(map[string]interface{})

	for _, UpdField := range UpdateFieldList {
		UpdateData[UpdField] = GetStructAttr(Data, UpdField)
	}

	result = ORM.ORM.Model(Data).Where(SelectCriteria).Updates(UpdateData)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ********************************************************
func (ORM *TORM) Insert(Data interface{}) error {
	var (
		result *gorm.DB
	)

	result = ORM.ORM.Create(Data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ********************************************************
func (ORM *TORM) Delete(StructName interface{}, SelectFields map[string]interface{}) (err error) {
	var (
		result *gorm.DB
	)

	result = ORM.ORM.Where(SelectFields).Delete(StructName)
	err = result.Error

	return err
}
