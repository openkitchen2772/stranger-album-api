package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"stranger-album-api/model/comment"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentController struct {
	client   *mongo.Client
	database string
}

func NewCommentController(client *mongo.Client, database string) CommentController {
	return CommentController{client, database}
}

func (cc CommentController) NewComment(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var c comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		fmt.Println(err)
		http.Error(rw, "JSON convertion error", http.StatusInternalServerError)
		return
	}

	if err := comment.InsertOne(cc.client, cc.database, c); err != nil {
		fmt.Println(err)
		http.Error(rw, "Insertion error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Comment insertion done for photo <%v>.", c.PhotoId)

	resJson := struct {
		PhotoId string
	}{
		c.PhotoId,
	}

	resByte, err := json.Marshal(resJson)
	if err != nil {
		fmt.Println("Comment Json convertion fails", err)
	}
	rw.Write(resByte)
}

func (cc CommentController) GetCommentByPhotoId(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
	photoId := params.ByName("photoId")
	var (
		c   = make([]comment.Comment, 0)
		err error
	)
	c, err = comment.FindAll(cc.client, cc.database, photoId)

	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Get photo comment error", http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "Json convertion error", http.StatusInternalServerError)
		return
	}

	rw.Write(data)
}
