package models

import "github.com/jinzhu/gorm"

type Author struct {
	Model
	//ID int `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	IsDelete int    `json:"is_delete"`
}

func GecAuthors(pageNum int, pageSize int, maps interface{}) ([]*Author, error) {

	var authors []*Author

	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&authors).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return authors, nil

}

/**
添加作者
*/
func AddAuthor(data map[string]interface{}) error {

	author := Author{
		Name:     data["name"].(string),
		Age:      data["age"].(int),
		IsDelete: 0,
	}

	err := db.Create(&author).Error

	if err != nil {
		return err
	}

	return nil
}

func EditAuhor(id int, data map[string]interface{}) error {

	var author = Author{}
	err := db.Model(&author).Where("id = ?", id).Updates(data).Error

	if err != nil {
		return err
	}

	return nil
}

/**
删除作者
*/
func DeleteAuthor(id int) error {
	err := db.Where("id = ?", id).Delete(Author{}).Error

	if err != nil {
		return err
	}

	return nil
}
