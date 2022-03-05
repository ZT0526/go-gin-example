package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/author_service"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
获取所有作者
*/
func GetAllAuthors(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authorServer := author_service.Author{
		PageSize: util.GetPage(c),
		PageNum:  setting.AppSetting.PageSize,
	}

	authors, err := authorServer.GetAll()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_GET_AUTHOR_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = authors
	data["total"] = len(authors)

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

/**
添加作者
*/
type AddAuthorForm struct {
	ID       int
	Name     string
	Age      int
	IsDelete int
}

func AddAuthor(c *gin.Context) {

	appG := app.Gin{C: c}
	valid := validation.Validation{}

	name := c.PostForm("name")
	age := com.StrTo(c.PostForm("age")).MustInt()
	valid.Required(name, "name").Message("name不能为空")
	valid.Range(age, 1, 100, "age").Message("年龄在1到100之间")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authorService := author_service.Author{
		Name:     name,
		Age:      age,
		IsDelete: 0,
	}

	err := authorService.AddAuthor()

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_POST_ADD_AUTHOR_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func EditAuthor(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	name := c.PostForm("name")
	age := com.StrTo(c.PostForm("age")).MustInt()
	valid.Min(id, 0, "id")
	valid.Required(name, "name").Message("name不能为空")
	valid.Min(name, 1, "name").Message("最小长度是1")
	valid.Range(age, 1, 100, "age").Message("年龄在1到100岁")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authorService := author_service.Author{
		ID:       id,
		Name:     name,
		Age:      age,
		IsDelete: 0,
	}
	data := make(map[string]interface{})
	data["name"] = name
	data["age"] = age
	if err := authorService.EditAuthor(id, data); err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_POST_EDIT_AUTHOR_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

/**
删除用户
*/
func DeleteAuthor(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.PostForm("id")).MustInt()
	valid.Min(id, 0, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authorService := author_service.Author{}

	err := authorService.DeleteAuthor(id)

	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_POST_DELETE_AUTHOR_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
