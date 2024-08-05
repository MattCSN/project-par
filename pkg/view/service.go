package view

import (
	"errors"
	"github.com/MattCSN/project-par/pkg/course"
	"github.com/MattCSN/project-par/pkg/golf"
	"github.com/MattCSN/project-par/pkg/hole"
	"github.com/MattCSN/project-par/pkg/tee"
)

var (
	viewService   *Service
	golfService   *golf.Service
	courseService *course.Service
	holeService   *hole.Service
	teeService    *tee.Service
)

func InitViewService() {
	viewService = &Service{}
	golfService = golf.NewGolfService(golf.NewRepository())
	courseService = course.NewCourseService(course.NewRepository(), golfService)
	holeService = hole.NewHoleService(hole.NewRepository())
	teeService = tee.NewTeeService(tee.NewRepository())
}

type Service struct{}

func (s *Service) GetGolfWithDetails(golfID string) (*GolfView, error) {
	if golfService == nil || courseService == nil || holeService == nil || teeService == nil {
		return nil, errors.New("services are not properly initialized")
	}

	golfModel, err := golfService.GetGolfByID(golfID)
	if err != nil {
		return nil, err
	}

	courses, err := courseService.GetCoursesByGolfID(golfID, 1, 100)
	if err != nil {
		return nil, err
	}

	var courseViews []CourseView
	for _, courseModel := range courses {
		holes, err := holeService.GetHolesByCourseID(courseModel.ID, 1, 100)
		if err != nil {
			return nil, err
		}

		var holeViews []HoleView
		for _, holeModel := range holes {
			tees, err := teeService.GetTeesByHoleID(holeModel.ID, 1, 100)
			if err != nil {
				return nil, err
			}
			holeViews = append(holeViews, HoleView{Model: holeModel, Tees: tees})
		}
		courseViews = append(courseViews, CourseView{Model: courseModel, Holes: holeViews})
	}

	return &GolfView{Model: *golfModel, Courses: courseViews}, nil
}
