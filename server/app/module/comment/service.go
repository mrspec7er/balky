package comment

import (
	"context"

	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
)

type CommentService struct {
	comment model.Comment
	publish utility.Publisher
}

func (s CommentService) Create(req *model.Comment) (int, error) {
	s.comment = *req
	err := s.comment.Create()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s CommentService) FindMany(appNumber string) ([]*model.Comment, int, error) {
	comments, err := s.comment.FindMany(appNumber)
	if err != nil {
		return nil, 500, err
	}

	return comments, 201, nil
}

func (s CommentService) Delete(req *model.Comment) (int, error) {
	s.comment = *req
	err := s.comment.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s CommentService) Publish(data []byte, queueName string, userEmail string) (int, error) {
	ctx := context.Background()

	payload := utility.Payload{
		Body:      data,
		UserEmail: userEmail,
	}
	err := s.publish.SendMessage(ctx, queueName, &payload)
	if err != nil {
		return 500, err
	}

	return 201, nil
}
