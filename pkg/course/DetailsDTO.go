package course

import "github.com/MattCSN/project-par/pkg/common"

type DetailsDTO struct {
	common.Base
	GolfName     string `json:"golfName"`
	CourseName   string `json:"name"`
	NumHoles     int    `json:"numHoles"`
	PitchAndPutt bool   `json:"pitchAndPutt"`
	Compact      bool   `json:"compact"`
	PostalCode   string `json:"postalCode"`
	City         string `json:"city"`
	Country      string `json:"country"`
} // @Name CourseDetails
