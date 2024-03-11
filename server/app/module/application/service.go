package application

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/lib/pq"
	"github.com/mrspec7er/balky/app/model"
	"github.com/mrspec7er/balky/app/utility"
	"gorm.io/gorm"
)

type ApplicationService struct {
	app      model.Application
	publish  utility.Publisher
	content  model.Content
	reaction model.Reaction
}

func (s *ApplicationService) Create(req *model.Application) (int, error) {
	s.app = *req
	err := s.app.Create()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s *ApplicationService) FindMany() ([]*model.Application, int, error) {
	apps, err := s.app.FindMany()
	if err != nil {
		return nil, 500, err
	}

	return apps, 201, nil
}

func (s *ApplicationService) FindOne(number string) (*model.Application, int, error) {
	s.app.Number = number
	app, err := s.app.FindOne()
	if err != nil {
		return nil, 500, err
	}

	return app, 201, nil
}

func (s *ApplicationService) Delete(req *model.Application) (int, error) {
	s.app = *req
	err := s.app.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s *ApplicationService) Publish(data []byte, queueName string, userEmail string) (int, error) {
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

func (s *ApplicationService) CreateContent(req []*model.Content) (int, error) {

	err := s.content.Create(req)
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s *ApplicationService) FindManyContent(appId string) ([]*model.Content, int, error) {
	contents, err := s.content.FindMany(appId)
	if err != nil {
		return nil, 500, err
	}

	return contents, 201, nil
}

func (s *ApplicationService) DeleteContent(req *model.Content) (int, error) {
	s.content = *req
	err := s.content.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s *ApplicationService) CreateReaction(req *InsertReaction) (int, error) {
	s.reaction = model.Reaction{
		ApplicationNumber: req.ApplicationNumber,
	}
	r, err := s.reaction.FindOne()

	if err == nil {
		if !slices.Contains(r.LikesBy, req.UserEmail) {
			s.reaction.LikesBy = append(s.reaction.LikesBy, req.UserEmail)
			err := s.reaction.Create(&s.reaction)
			if err != nil {
				return 500, err
			}
		} else {
			return 400, fmt.Errorf("user already like this application")
		}
	}

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		reactionPayload := &model.Reaction{
			ApplicationNumber: req.ApplicationNumber,
			LikesBy:           pq.StringArray{req.UserEmail},
		}
		err := s.reaction.Create(reactionPayload)
		if err != nil {
			return 500, err
		}
	} else {
		return 500, err
	}

	return 201, nil
}
