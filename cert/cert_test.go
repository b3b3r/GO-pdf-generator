package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2020-08-20")
	if err != nil {
		t.Errorf("Cert data should be valid. Error %v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. Expected=GOLANG COURSE got%s", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2020-08-20")
	if err == nil {
		t.Errorf("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "feragraevbhtazhtrzhhhhhhhhthtzhergraegraefgregzegefzafa"
	_, err := New(course, "Bob", "2020-08-20")
	if err == nil {
		t.Errorf("Error should be returned on a too long course (course =%s)", course)
	}
}

func TestEmptyNameValue(t *testing.T) {
	_, err := New("Golang", "", "2020-08-20")
	if err == nil {
		t.Errorf("Error should be returned on an empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "rezreaztrbtrnathr mryavnryeazvnmravtraeynyaytr"
	_, err := New("Golang", name, "2020-08-20")
	if err == nil {
		t.Errorf("Error should be returned on a too long name (name=%s)", name)
	}
}

func TestValidDateFormat(t *testing.T) {
	c, _ := New("Golang", "John", "2020-08-20")
	if c.LabelDate != "Date: 20/08/2020" {
		t.Errorf("LabelDate should be 20/08/2020 and we got value='%s'", c.LabelDate)
	}
}
