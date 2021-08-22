package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang-web-project-framework/src/modules/user/dao"
	"golang-web-project-framework/src/utils"
	"golang-web-project-framework/src/utils/db/mongo"
)

type UserDetailInfo struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	CreatedTime int64  `bson:"created_time"`
}

type User struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	CreatedTime int64  `bson:"created_time"`
}

type UserUpdate struct {
	Name string `bson:"name"`
}

func AddUser(userDetailInfo UserDetailInfo) (string, *utils.Error) {
	_, err := dao.UserDaoMgo.Mgo.InsertOne(userDetailInfo)
	if err != nil {
		utils.Warnf("add user failed: error=%s", err.Error())
		return "", utils.NewErrorWithInnerError(err.Error())
	}

	return userDetailInfo.Id, nil
}

func DeleteUserById(userId string) (utilsErr *utils.Error) {
	// TODO do not delete directly but update status.
	_, err := dao.UserDaoMgo.Mgo.DeleteOne(bson.M{utils.MongoDbId: userId})
	if err != nil {
		utils.Warnf("delete user failed: error=%s", err.Error())
		return utils.NewErrorWithInnerError(err.Error())
	}

	return
}

func UpdateUserById(userId string, userUpdate UserUpdate) (utilsErr *utils.Error) {
	_, err := dao.UserDaoMgo.Mgo.UpdateOne(bson.M{utils.MongoDbId: userId}, bson.M{utils.MongoDbSet: userUpdate})
	if err != nil {
		utils.Warnf("update user failed: error=%s", err.Error())
		return utils.NewErrorWithInnerError(err.Error())
	}

	return
}

func GetUsers(pageNum, pageSize int64, order, orderBy, searchValue, searchBy string) (users []User, totalSize int64, utilsErr *utils.Error) {
	filter, opts := mongo.GenerateFilterAndFindOptions(pageNum, pageSize, order, orderBy, searchValue, searchBy)
	cursor, err := dao.UserDaoMgo.Mgo.Find(filter, opts)
	if err != nil {
		utils.Warnf("get users failed: error=%s", err.Error())
		return nil, 0, utils.NewErrorWithInnerError(err.Error())
	}

	err = cursor.All(context.Background(), &users)
	if err != nil {
		utils.Warnf("get users decode failed: error=%s", err.Error())
		return nil, 0, utils.NewErrorWithInnerError(err.Error())
	}
	defer cursor.Close(context.Background())

	totalSize, err = dao.UserDaoMgo.Mgo.CountDocuments(filter)
	if err != nil {
		utils.Warnf("get users count documents failed: error=%s", err.Error())
		return nil, 0, utils.NewErrorWithInnerError(err.Error())
	}

	return
}

func GetUserById(userId string) (userDetailInfo *UserDetailInfo, utilsErr *utils.Error) {
	singleResult := dao.UserDaoMgo.Mgo.FindOne(bson.M{utils.MongoDbId: userId})
	if singleResult == nil {
		utils.Warnf("get user by id failed: error=result is empty")
		return nil, utils.NewErrorWithNotFound()
	}

	userDetailInfo = new(UserDetailInfo)
	err := singleResult.Decode(&userDetailInfo)
	if err != nil {
		utils.Warnf("get user by id decode failed: error=%s", err.Error())
		return nil, utils.NewErrorWithInnerError(err.Error())
	}

	return
}

