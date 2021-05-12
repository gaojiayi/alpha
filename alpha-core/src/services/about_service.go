package services

import (
	"alpha-core/src/config"
	"alpha-core/src/repository"
	"alpha-core/src/repository/model"
	"alpha-core/src/views"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/**
 1 支持分页
 */
func AllMessages(pageIndex int64, pageSize int64) (views.MessageQueryResponse, error) {
	var resps views.MessageQueryResponse
	var rows []model.Message
	var child_rows []model.Message
	// 查询所有的消息
	mgo := repository.NewMgo(config.DBBASE_NAME, config.DB_COLL_MESSAGE)

	temp := mgo.CollectionRootComments(pageIndex, pageSize, -1)
	total := mgo.CollectionRootCommentsCount()
	//使用CURSOR.ALL  或者CURSOR.NEXT
	err := temp.All(context.Background(), &rows)
	if err != nil {
		log.Fatal(err)
	}
	var root_ids []string
	for _, row := range rows {
		root_ids = append(root_ids, row.Id)
	}
	filter := bson.D{{
		"parent_msg_id",
		bson.D{{
			"$in",
			root_ids}}}}
	child_cur := mgo.FindRows(filter, 1)
	err = child_cur.All(context.TODO(), &child_rows)
	if err != nil {
		log.Fatal(err)
	}
	//根据消息id 对所有的回复分组
	var wholeMessageArray []views.MessageResp
	for _, root := range rows {
		childMessageArray:= []views.MessageResp{}
		for _, child := range child_rows {
			if child.Parent_msg_id == root.Id {
				item := views.MessageResp{
					child.Id, child.HeadNum,child.Date,
					child.Context, child.Parent_msg_id,
					child.Parent_username,
					child.Username, []views.MessageResp{},
				}
				childMessageArray = append(childMessageArray, item)
			}
		}
		wholeMessage := views.MessageResp{
			root.Id, root.HeadNum,root.Date,
			root.Context, root.Parent_msg_id,
			root.Parent_username,
			root.Username, childMessageArray,
		}
		wholeMessageArray = append(wholeMessageArray, wholeMessage)
	}
	resps.MessageResp = wholeMessageArray
	resps.Total = total
	return resps, err
}


func AddMessage(request views.AddMessageRequest) string {
	message := make(map[string]interface{})
	message["date"] = time.Now().Format("2006-01-02 15:04:05")
	message["context"] = request.Context
	message["parent_msg_id"] = request.ParentMsgId
	message["username"] = request.Username
	message["parent_username"] = request.ParentUsername
	message["head_num"] = strconv.Itoa(rand.Intn(config.Head_photo_count)+1)
	mgo := repository.NewMgo(config.DBBASE_NAME, config.DB_COLL_MESSAGE)
	mgo.InsertOne(message)
	return "success"
}
