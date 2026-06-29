package parser

// Resume represents a structured candidate resume with all major sections.
type Resume struct {
	Name       string           `json:"name"`
	Summary    string           `json:"summary"`
	Experience []WorkExperience `json:"experience"`
	Education  []EducationInfo  `json:"education"`
	Projects   []ProjectInfo    `json:"projects"`
	Skills     []string         `json:"skills"`
}

// WorkExperience represents a single professional role held by the candidate.
type WorkExperience struct {
	Company      string   `json:"company"`
	Title        string   `json:"title"`
	StartDate    string   `json:"start_date"`
	EndDate      string   `json:"end_date"`
	Achievements []string `json:"achievements"`
}

// EducationInfo represents an educational qualification of the candidate.
type EducationInfo struct {
	Institution string `json:"institution"`
	Degree      string `json:"degree"`
	Major       string `json:"major"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}

// ProjectInfo represents a notable project the candidate has worked on.
type ProjectInfo struct {
	Name        string   `json:"name"`
	Role        string   `json:"role"`
	TechStack   []string `json:"tech_stack"`
	Description string   `json:"description"`
}
