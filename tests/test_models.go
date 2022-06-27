package test_models 

import (
	"models"
	"testing"
	"assert"
)

var customers []models.Customer 
var products []models.Product
var currencies []models.Currency 

var (
	AllModels = map[string]interface{}{
		"Customer": customers, "Product": products, "Currency": currencies 
	}
)	

func TestModelCreate(t *testing.T, model struct{}, model_data map[string]interface{}){
	object := models.database.Model(&model).create(model_data)
	models.database.Save(object)
	assert.Greater(t, len(customers), 0)
}
func TestModelUpdate(model struct{}, obj_id int, updated_data map[string]interface{}){
	obj := models.database.Model(&model).Where("id = ?", obj_id).First()
	new_obj := obj.Update(updated_data)
	assert.NotEqual(t, obj, new_obj)
}

func TestModelDelete(t *testing.T, model struct{}, obj_id int){
	models.database.Model(&model).Where("id = ?", obj_id).Delete()
	assert.Equal(t, len(AllModels[model]), 0)
}