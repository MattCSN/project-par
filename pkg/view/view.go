package view

import (
	"github.com/MattCSN/project-par/pkg/course"
	"github.com/MattCSN/project-par/pkg/golf"
	"github.com/MattCSN/project-par/pkg/hole"
	"github.com/MattCSN/project-par/pkg/tee"
)

type GolfView struct {
	golf.Model
	Courses []CourseView `json:"courses"`
}

type CourseView struct {
	course.Model
	Holes []HoleView `json:"holes"`
}

type HoleView struct {
	hole.Model
	Tees []tee.Model `json:"tees"`
}
