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
	defer file.Close()

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
		golfModel := &golf.Model{
			Name:      jsonGolf.Name,
			Latitude:  parseFloat(jsonGolf.Latitude),
			Longitude: parseFloat(jsonGolf.Longitude),
		}

		if err := golfService.CreateGolf(golfModel); err != nil {
			log.Printf("Error saving golf: %v", err)
			continue
		}

		for _, jsonCourse := range jsonGolf.Courses {
			courseModel := &course.Model{
				GolfID:       golfModel.ID,
				NumHoles:     parseInt(jsonCourse.Holes),
				Compact:      parseBool(jsonCourse.Compact),
				PitchAndPutt: parseBool(jsonCourse.PitchAndPutt),
			}

			if err := courseService.CreateCourse(courseModel); err != nil {
				log.Printf("Error saving course: %v", err)
				continue
			}

			for i := 1; i <= parseInt(jsonCourse.Holes); i++ {
				holeModel := &hole.Model{
					CourseID:   courseModel.ID,
					HoleNumber: i,
					Par:        0,
				}

				if err := holeService.CreateHole(holeModel); err != nil {
					log.Printf("Error saving hole: %v", err)
					continue
				}

				teeModel := &tee.Model{
					HoleID:   holeModel.ID,
					Color:    "colorless",
					Distance: 0,
				}

				if err := teeService.CreateTee(teeModel); err != nil {
					log.Printf("Error saving tee: %v", err)
					continue
				}
			}
		}
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

func ImportGolfData(filePath string, golfService *golf.Service, courseService *course.Service, holeService *hole.Service, teeService *tee.Service) {
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
