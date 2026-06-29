package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/johncegom/resume-adaptation/internal/parser"
)

func TestResumeModelJSONRoundTrip(t *testing.T) {
	original := parser.Resume{
		Name:    "Jane Doe",
		Summary: "Senior software engineer with 8 years of experience.",
		Experience: []parser.WorkExperience{
			{
				Company:      "Acme Corp",
				Title:        "Senior Engineer",
				StartDate:    "2020-01",
				EndDate:      "2024-06",
				Achievements: []string{"Led migration to microservices"},
			},
		},
		Education: []parser.EducationInfo{
			{
				Institution: "MIT",
				Degree:      "B.S.",
				Major:       "Computer Science",
				StartDate:   "2012-09",
				EndDate:     "2016-06",
			},
		},
		Projects: []parser.ProjectInfo{
			{
				Name:         "Open Source CLI",
				Role:         "Creator",
				TechStack:    []string{"Go", "Cobra"},
				Description:  "A CLI tool for resume adaptation.",
			},
		},
		Skills: []string{"Go", "Python", "Kubernetes"},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("failed to marshal Resume: %v", err)
	}

	var decoded parser.Resume
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("failed to unmarshal Resume: %v", err)
	}

	if decoded.Name != original.Name {
		t.Errorf("Name: got %q, want %q", decoded.Name, original.Name)
	}
	if decoded.Summary != original.Summary {
		t.Errorf("Summary: got %q, want %q", decoded.Summary, original.Summary)
	}
	if len(decoded.Experience) != len(original.Experience) {
		t.Fatalf("Experience length: got %d, want %d", len(decoded.Experience), len(original.Experience))
	}
	if decoded.Experience[0].Company != original.Experience[0].Company {
		t.Errorf("Experience[0].Company: got %q, want %q", decoded.Experience[0].Company, original.Experience[0].Company)
	}
	if len(decoded.Education) != len(original.Education) {
		t.Fatalf("Education length: got %d, want %d", len(decoded.Education), len(original.Education))
	}
	if decoded.Education[0].Institution != original.Education[0].Institution {
		t.Errorf("Education[0].Institution: got %q, want %q", decoded.Education[0].Institution, original.Education[0].Institution)
	}
	if len(decoded.Projects) != len(original.Projects) {
		t.Fatalf("Projects length: got %d, want %d", len(decoded.Projects), len(original.Projects))
	}
	if decoded.Projects[0].Name != original.Projects[0].Name {
		t.Errorf("Projects[0].Name: got %q, want %q", decoded.Projects[0].Name, original.Projects[0].Name)
	}
	if len(decoded.Skills) != len(original.Skills) {
		t.Fatalf("Skills length: got %d, want %d", len(decoded.Skills), len(original.Skills))
	}
}


func TestResumeModelRequiredFields(t *testing.T) {
	r := parser.Resume{
		Name:       "Test",
		Summary:    "Summary",
		Experience: []parser.WorkExperience{},
		Education:  []parser.EducationInfo{},
		Projects:   []parser.ProjectInfo{},
		Skills:     []string{},
	}
	if r.Name == "" || r.Summary == "" {
		t.Fatal("Resume must support Name and Summary fields")
	}
	if r.Experience == nil || r.Education == nil || r.Projects == nil || r.Skills == nil {
		t.Fatal("Resume must support Experience, Education, Projects, Skills slices")
	}
}

func TestWorkExperienceModelFields(t *testing.T) {
	w := parser.WorkExperience{
		Company:      "Acme",
		Title:        "Engineer",
		StartDate:    "2020-01",
		EndDate:      "2024-01",
		Achievements: []string{"Built systems"},
	}
	data, _ := json.Marshal(w)
	m := make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	for _, key := range []string{"company", "title", "start_date", "end_date", "achievements"} {
		if _, ok := m[key]; !ok {
			t.Errorf("missing JSON key %q", key)
		}
	}
}

func TestEducationInfoModelFields(t *testing.T) {
	e := parser.EducationInfo{
		Institution: "MIT",
		Degree:      "B.S.",
		Major:       "CS",
		StartDate:   "2012-09",
		EndDate:     "2016-06",
	}
	data, _ := json.Marshal(e)
	m := make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	for _, key := range []string{"institution", "degree", "major", "start_date", "end_date"} {
		if _, ok := m[key]; !ok {
			t.Errorf("missing JSON key %q", key)
		}
	}
}

func TestProjectInfoModelFields(t *testing.T) {
	p := parser.ProjectInfo{
		Name:        "CLI Tool",
		Role:        "Lead",
		TechStack:   []string{"Go"},
		Description: "A tool.",
	}
	data, _ := json.Marshal(p)
	m := make(map[string]interface{})
	_ = json.Unmarshal(data, &m)
	for _, key := range []string{"name", "role", "tech_stack", "description"} {
		if _, ok := m[key]; !ok {
			t.Errorf("missing JSON key %q", key)
		}
	}
}

