package home

import (
	"bi-activity/dao/home"
	"bi-activity/response/errors"
	"context"
	"github.com/sirupsen/logrus"
)

type HelpService struct {
	hr  home.HelpRepo
	log *logrus.Logger
}

func NewHelpService(hr home.HelpRepo, log *logrus.Logger) *HelpService {
	return &HelpService{
		hr:  hr,
		log: log,
	}
}

func (hs *HelpService) HelpList(ctx context.Context) (list []*Help, err error) {
	resp, err := hs.hr.GetHelpList(ctx)
	if err != nil {
		return nil, errors.GetHelpError
	}

	for _, v := range resp {
		list = append(list, &Help{
			Problem: v.Name,
			Answer:  v.Answer,
		})
	}

	return list, nil
}

func (hs *HelpService) SearchHelp(ctx context.Context, params string) (list []*Help, err error) {
	if params == "" {
		return hs.HelpList(ctx)
	}

	resp, err := hs.hr.SearchHelp(ctx, params)
	if err != nil {
		return nil, errors.GetHelpError
	}

	for _, v := range resp {
		list = append(list, &Help{
			Problem: v.Name,
			Answer:  v.Answer,
		})
	}

	return list, nil
}
