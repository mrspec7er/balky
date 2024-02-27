package master

import "github.com/mrspec7er/balky/app/model"

type MasterReportService struct {
	master model.MasterReport
}

func (s MasterReportService) Create(req *model.MasterReport) (int, error) {
	s.master = *req
	err := s.master.Create()
	if err != nil {
		return 500, err
	}

	return 201, nil
}

func (s MasterReportService) FindMany() ([]model.MasterReport, int, error) {
	masters, err := s.master.FindMany()
	if err != nil {
		return nil, 500, err
	}

	return masters, 201, nil
}

func (s MasterReportService) Delete(req *model.MasterReport) (int, error) {
	s.master = *req
	err := s.master.Delete()
	if err != nil {
		return 500, err
	}

	return 201, nil
}
