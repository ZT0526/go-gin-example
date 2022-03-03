package author_service

import "github.com/EDDYCJY/go-gin-example/models"

type Author struct {
	ID       int
	Name     string
	Age      int
	IsDelete int

	PageNum  int
	PageSize int
}

/**
获取所有authors
*/
func (a *Author) GetAll() ([]*models.Author, error) {

	var authors []*models.Author

	authors, err := models.GecAuthors(a.PageNum, a.PageSize, a.getMaps())

	if err != nil {
		return nil, err
	}

	return authors, err
}

func (a *Author) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["is_delete"] = 0

	return maps
}

/**
添加作者
*/
func (a *Author) AddAuthor() error {

	author := map[string]interface{}{
		"name":      a.Name,
		"age":       a.Age,
		"is_delete": a.IsDelete,
	}

	err := models.AddAuthor(author)

	if err != nil {
		return err
	}

	return nil
}

func (a *Author) EditAuthor(id int, data map[string]interface{}) error {

	return nil
}

/**
删除作者
*/
func (a *Author) DeleteAuthor(id int) error {

	err := models.DeleteAuthor(id)
	if err != nil {
		return err
	}

	return nil
}
