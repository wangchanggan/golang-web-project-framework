package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-web-project-framework/src/utils"
	utilsHttp "golang-web-project-framework/src/utils/http"
)

func GenerateFilterAndFindOptions(pageNum, pageSize int64, order, orderBy, searchValue, searchBy string) (filter bson.M, findOpts *options.FindOptions) {
	filter = bson.M{}
	if searchValue != "" {
		filter[searchBy] = bson.M{"$regex": searchValue}
	}

	findOpts = options.Find()
	if order != "" {
		if orderBy == utils.OrderByDefaultValue {
			findOpts.SetSort(bson.M{order: 1})
		} else {
			findOpts.SetSort(bson.M{order: -1})
		}
	}

	findOpts.SetSkip(utilsHttp.GenerateOffset(pageNum, pageSize)).SetLimit(pageSize)
	return
}
