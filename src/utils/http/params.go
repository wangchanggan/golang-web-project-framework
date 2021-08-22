/*
    Different web frameworks(gin, iris, beego etc.) have different way to get params, which makes developers confused.
    So it provides unified method of getting params.
 */

package http

import (
	"github.com/gin-gonic/gin"
	"golang-web-project-framework/src/utils"
)

func GenerateOffset(pageNum, pageSize int64) int64 {
	return (pageNum - 1) * pageSize
}

// encapsulation of list query interface parameters
func GenerateListInfo(c *gin.Context) (pageNum, pageSize int64, order, orderBy, searchValue, searchBy string, err error) {
	pageNumStr, ok := c.GetQuery(utils.PageNum)
	if !ok {
		pageNum = utils.PageNumDefaultValue
	} else {
		pageNum, err = utils.StringToInt64(pageNumStr)
		if err != nil {
			return
		}
	}

	pageSizeStr, ok := c.GetQuery(utils.PageSize)
	if !ok {
		pageSize = utils.PageSizeDefaultValue
	} else {
		pageSize, err = utils.StringToInt64(pageSizeStr)
		if err != nil {
			return
		}
	}

	order, _ = c.GetQuery(utils.Order)
	orderBy, ok = c.GetQuery(utils.OrderBy)
	if !ok {
		orderBy = utils.OrderByDefaultValue
	}

	searchValue, _ = c.GetQuery(utils.SearchValue)
	searchBy, _ = c.GetQuery(utils.SearchBy)

	return
}

func CompleteListInfo(pageNum, pageSize int64, order, orderBy string) (int64, int64, string, string) {
	if pageNum < 1 {
		pageNum = utils.PageNumDefaultValue
	}
	if pageSize < 1 {
		pageSize = utils.PageSizeDefaultValue
	}
	if order != "" && orderBy != utils.OrderByDefaultValue && orderBy != utils.OrderByDesc {
		orderBy = utils.OrderByDefaultValue
	}

	return pageNum, pageSize, order, orderBy
}

func ValidateListInfo(searchValue, searchBy string) string {
	if searchValue != "" && searchBy == "" {
		return "searchBy"
	}

	return ""
}

func GetParamsFromBodyByJson(c *gin.Context, obj interface{}) (err error) {
	err = c.ShouldBindJSON(obj)
	if err != nil {
		return
	}

	return
}

func GetParamFromUrl(c *gin.Context, key string) string {
	return c.Param(key)
}

func GetQueryFromUrl(c *gin.Context, key string) (string, bool) {
	return c.GetQuery(key)
}

func GetDefaultQueryFromUrl(c *gin.Context, key, defaultValue string) string {
	return c.DefaultQuery(key, defaultValue)
}
