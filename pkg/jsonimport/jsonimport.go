package jsonimport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/MattCSN/project-par/pkg/course"
	"github.com/MattCSN/project-par/pkg/golf"
	"github.com/MattCSN/project-par/pkg/hole"
	"github.com/MattCSN/project-par/pkg/tee"
)

type JSONGolfData struct {
	Count int        `json:"count"`
	Golfs []JSONGolf `json:"golfs"`
}

type JSONGolf struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	Latitude  string       `json:"latitude"`
	Longitude string       `json:"longitude"`
	Courses   []JSONCourse `json:"courses"`
}

type JSONCourse struct {
	Holes        string `json:"holes"`
	Compact      string `json:"compact"`
	PitchAndPutt string `json:"pitchAndPutt,omitempty"`
}

func ReadJSONFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Error closing file: %v", err)
		}
	}()

	return io.ReadAll(file)
}

func UnmarshalJSONData(data []byte) (*JSONGolfData, error) {
	var jsonGolfData JSONGolfData
	err := json.Unmarshal(data, &jsonGolfData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	return &jsonGolfData, nil
}

func MapAndSaveData(jsonGolfData *JSONGolfData, golfService *golf.Service, courseService *course.Service, holeService *hole.Service, teeService *tee.Service) error {
	for _, jsonGolf := range jsonGolfData.Golfs {
		if err := mapAndSaveGolf(jsonGolf, golfService, courseService, holeService, teeService); err != nil {
			log.Printf("Error processing golf: %v", err)
			continue
		}
	}
	return nil
}

func mapAndSaveGolf(jsonGolf JSONGolf, golfService *golf.Service, courseService *course.Service, holeService *hole.Service, teeService *tee.Service) error {
	golfModel := &golf.Model{
		Name:           jsonGolf.Name,
		Country:        "France",
		GoogleMapLinks: "https://maps.google.com/?q=" + jsonGolf.Name,
		Latitude:       parseFloat(jsonGolf.Latitude),
		Longitude:      parseFloat(jsonGolf.Longitude),
	}

	if err := golfService.CreateGolf(golfModel); err != nil {
		return fmt.Errorf("error saving golf: %w", err)
	}

	courseNumber := 1
	for _, jsonCourse := range jsonGolf.Courses {
		if err := mapAndSaveCourse(jsonCourse, golfModel.ID, courseNumber, courseService, holeService, teeService); err != nil {
			log.Printf("Error processing course: %v", err)
			continue
		}
		courseNumber++
	}
	return nil
}

func mapAndSaveCourse(jsonCourse JSONCourse, golfID string, courseNumber int, courseService *course.Service, holeService *hole.Service, teeService *tee.Service) error {
	courseModel := &course.Model{
		GolfID:       golfID,
		Name:         fmt.Sprintf("unnamed Course n°%d", courseNumber),
		NumHoles:     parseInt(jsonCourse.Holes),
		Compact:      parseBool(jsonCourse.Compact),
		PitchAndPutt: parseBool(jsonCourse.PitchAndPutt),
	}

	if err := courseService.CreateCourse(courseModel); err != nil {
		return fmt.Errorf("error saving course: %w", err)
	}

	for i := 1; i <= parseInt(jsonCourse.Holes); i++ {
		if err := mapAndSaveHole(i, courseModel.ID, holeService, teeService); err != nil {
			log.Printf("Error processing hole: %v", err)
			continue
		}
	}
	return nil
}

func mapAndSaveHole(holeNumber int, courseID string, holeService *hole.Service, teeService *tee.Service) error {
	holeModel := &hole.Model{
		CourseID:   courseID,
		HoleNumber: holeNumber,
		Par:        0,
	}

	if err := holeService.CreateHole(holeModel); err != nil {
		return fmt.Errorf("error saving hole: %w", err)
	}

	if err := mapAndSaveTee(holeModel.ID, teeService); err != nil {
		return fmt.Errorf("error processing tee: %w", err)
	}
	return nil
}

func mapAndSaveTee(holeID string, teeService *tee.Service) error {
	teeModel := &tee.Model{
		HoleID:   holeID,
		Color:    "colorless",
		Distance: 0,
	}

	if err := teeService.CreateTee(teeModel); err != nil {
		return fmt.Errorf("error saving tee: %w", err)
	}
	return nil
}

func parseFloat(value string) float64 {
	parsedValue, _ := strconv.ParseFloat(value, 64)
	return parsedValue
}

func parseInt(value string) int {
	parsedValue, _ := strconv.Atoi(value)
	return parsedValue
}

func parseBool(value string) bool {
	parsedValue, _ := strconv.ParseBool(value)
	return parsedValue
}

func ImportGolfData(filePath string) {
	golfService := golf.NewGolfService(golf.NewRepository())
	courseService := course.NewCourseService(course.NewRepository(), golfService)
	holeService := hole.NewHoleService(hole.NewRepository())
	teeService := tee.NewTeeService(tee.NewRepository())

	if golfService == nil || courseService == nil || holeService == nil || teeService == nil {
		log.Println("Error: One or more services are nil")
		return
	}

	data, err := ReadJSONFile(filePath)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	jsonGolfData, err := UnmarshalJSONData(data)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	if err := MapAndSaveData(jsonGolfData, golfService, courseService, holeService, teeService); err != nil {
		log.Println("Error:", err)
	}
}
