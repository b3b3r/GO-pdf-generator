package cert

import (
	"fmt"
	"strings"
	"time"
)

var maxLenCourse = 20
var maxLenName = 30

// Cert is the struct of the certificate
type Cert struct {
	Course string
	Name   string
	Date   time.Time

	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

// Saver save the certificate
type Saver interface {
	Save(c Cert) error
}

func validateString(str string, maxLen int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 {
		return c, fmt.Errorf("Invalid string. Got='%s' len=%d", c, len(c))
	} else if len(c) >= maxLen {
		return c, fmt.Errorf("Invalid string. Got='%s' len=%d", c, len(c))
	}
	return c, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateString(course, maxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(course, "COURSE") {
		c = c + " COURSE"
	}
	return strings.ToTitle(c), nil
}

func validateName(name string) (string, error) {
	n, err := validateString(name, maxLenName)
	if err != nil {
		return "", err
	}
	return strings.ToTitle(n), nil
}

func parseDate(date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, err
	}
	return t, nil
}

// New generate a new certificate
func New(course, name, date string) (*Cert, error) {
	c, err := validateCourse(course)
	if err != nil {
		return nil, err
	}
	n, err := validateName(name)
	if err != nil {
		return nil, err
	}
	d, err := parseDate(date)
	if err != nil {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v certificate - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}
	return cert, nil
}
