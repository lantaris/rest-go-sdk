package orm

import (
	"gorm.io/gorm"
	"reflect"
)

// ********************************************************
func (SELF *TORM) GetStructAttr(obj interface{}, fieldName string) interface{} {

	r := reflect.ValueOf(obj)
	f := reflect.Indirect(r).FieldByName(fieldName)
	return f.Interface()
}

// ********************************************************
func (SELF *TORM) Select(RetData interface{}, SelectFields map[string]interface{}) (err error) {
	var (
		result *gorm.DB
	)

	result = SELF.ORM.Where(SelectFields).Find(RetData)
	err = result.Error

	return err
}

// ********************************************************
func (SELF *TORM) Update(Data interface{}, SelectCriteria map[string]interface{}, UpdateFieldList []string) (err error) {
	var (
		UpdateData map[string]interface{}
		result     *gorm.DB
	)

	if UpdateFieldList == nil {
		result = SELF.ORM.Model(Data).Where(SelectCriteria).Updates(Data)
	} else {
		UpdateData = make(map[string]interface{})
		for _, UpdField := range UpdateFieldList {
			UpdateData[UpdField] = SELF.GetStructAttr(Data, UpdField)
		}
		result = SELF.ORM.Model(Data).Where(SelectCriteria).Updates(UpdateData)
	}
	err = result.Error

	return err
}

// ********************************************************
func (SELF *TORM) Insert(Data interface{}) (err error) {
	var (
		result *gorm.DB
	)

	result = SELF.ORM.Create(Data)
	err = result.Error

	return err
}

// ********************************************************
func (SELF *TORM) Delete(StructName interface{}, SelectFields map[string]interface{}) (err error) {
	var (
		result *gorm.DB
	)

	result = SELF.ORM.Where(SelectFields).Delete(StructName)
	err = result.Error

	return err
}
