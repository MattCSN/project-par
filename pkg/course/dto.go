package course

import (
	"github.com/MattCSN/project-par/pkg/golf"
	"github.com/MattCSN/project-par/pkg/hole"
	"github.com/MattCSN/project-par/pkg/tee"
)

type CourseDetailsDTO struct {
	Model
	Golf  GolfDetailsDTO   `json:"golf"`
	Holes []HoleDetailsDTO `json:"holes"`
}

type GolfDetailsDTO struct {
	golf.Model
	Courses []CourseSummaryDTO `json:"courses"`
}

type CourseSummaryDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type HoleDetailsDTO struct {
	hole.Model
	Tees []tee.Model `json:"tees"`
}

func newCoursesSummaryDTOs(courses []Model) []CourseSummaryDTO {
	var courseDTOs []CourseSummaryDTO
	for _, course := range courses {
		courseDTOs = append(courseDTOs, CourseSummaryDTO{
			ID:   course.ID,
			Name: course.Name,
		})
	}
	return courseDTOs
}

func newGolfDetailsDTO(golf golf.Model, courses []Model) GolfDetailsDTO {
	return GolfDetailsDTO{
		Model:   golf,
		Courses: newCoursesSummaryDTOs(courses),
	}
}

func newHoleDetailsDTO(hole hole.Model, tees []tee.Model) HoleDetailsDTO {
	var holeDetail HoleDetailsDTO
	holeDetail.Model = hole
	holeDetail.Tees = tees

	return holeDetail
}

func newCourseDetailsDTO(course Model, otherCourses []Model, golf golf.Model, holes []HoleDetailsDTO) CourseDetailsDTO {
	return CourseDetailsDTO{
		Model: course,
		Golf:  newGolfDetailsDTO(golf, otherCourses),
		Holes: holes,
	}
}
