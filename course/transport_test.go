package course_test

import (
	"encoding/json"
	"errors"
	"fmt"
	c "github.com/ncostamagna/go_http_client/client"
	courseSdk "github.com/og11423074s/go_course_sdk/course"
	"github.com/og11423074s/gocourse_domain/domain"
	"net/http"
	"os"
	"strings"
	"testing"
)

var header http.Header
var sdk courseSdk.Transport

func TestMain(m *testing.M) {
	header = http.Header{}
	header.Set("Content-Type", "application/json")
	sdk = courseSdk.NewHttpClient("base-url/", "")
	os.Exit(m.Run())
}

func TestGet_Response404Error(t *testing.T) {
	expectedErr := courseSdk.ErrNotFound{Message: "course '1' not found"}

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 404,
		RespBody: fmt.Sprintf(`{
		"status": 404,
		"message": "%s"
        }`, expectedErr.Error()),
	})

	if err != nil {
		t.Errorf("expeted nil, got %v", err)
	}

	course, err := sdk.Get("1")

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected nil, got %v", err)
	}

	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}

}

func TestGet_Response500Error(t *testing.T) {
	expectedErr := errors.New("internal server error")

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 500,
		RespBody: fmt.Sprintf(`{
		"status": 500,
		"message": "%s"
        }`, expectedErr.Error()),
	})

	if err != nil {
		t.Errorf("expeted nil, got %v", err)
	}

	course, err := sdk.Get("1")
	if err == nil || !strings.Contains(err.Error(), expectedErr.Error()) {
		t.Errorf("expected nil, got %v", err)
	}
	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}

func TestGet_ResponseMarshalError(t *testing.T) {
	expectedErr := errors.New("unexpected end of JSON input")

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 200,
		RespBody:     `{`,
	})

	if err != nil {
		t.Errorf("expeted nil, got %v", err)
	}

	course, err := sdk.Get("1")
	if err == nil || !strings.Contains(err.Error(), expectedErr.Error()) {
		t.Errorf("expected nil, got %v", err)
	}

	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}

func TestGet_ClientError(t *testing.T) {
	expectedErr := errors.New("client error")

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 400,
		Err:          expectedErr,
	})

	if err != nil {
		t.Errorf("expeted nil, got %v", err)
	}

	course, err := sdk.Get("1")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected nil, got %v", err)
	}

	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}

func TestGet_ResponseSuccess(t *testing.T) {
	expectedCourse := &domain.Course{
		ID:   "1",
		Name: "course 1",
	}

	expectedCourseJson, err := json.Marshal(expectedCourse)
	if err != nil {
		t.Errorf("expeted nil, got %v", err)
	}

	err = c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 200,
		RespBody: fmt.Sprintf(`{
		"status": 200,
		"message": "success",
		"data": %s
		}`, expectedCourseJson),
	})

	if err != nil {
		t.Errorf("expeted nil, got %v", err)
	}

	course, err := sdk.Get("1")
	if err != nil {
		t.Errorf("expeted nil, got %v", err)
	}

	if course == nil {
		t.Errorf("expected %v, got nil", expectedCourse)
	}

	if course.ID != expectedCourse.ID {
		t.Errorf("expected %s, got %s", expectedCourse.ID, course.ID)
	}

	if course.Name != expectedCourse.Name {
		t.Errorf("expected %s, got %s", expectedCourse.Name, course.Name)
	}
}
