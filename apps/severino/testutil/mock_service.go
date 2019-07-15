package testutil

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"

	"backend/apps/severino/mocks"
	app "backend/apps/severino/src"
	"backend/apps/severino/src/services"
	"backend/apps/severino/src/structs"
	"backend/libs/json"
)

type TestUtil struct {
	ServiceContainer services.Container
}

func WithAlbumServiceMocked(method string, mockedResult interface{}, mockedErr error) TestUtil {
	service := new(mocks.AlbumService)
	service.On(method, mock.Anything).Return(mockedResult, mockedErr)
	serviceContainer := services.Container{
		Album: service,
	}
	return TestUtil{serviceContainer}
}

func WithTrackServiceMocked(method string, mockedResult interface{}, mockedErr error) TestUtil {
	service := new(mocks.Service)
	service.On(method, mock.Anything).Return(mockedResult, mockedErr)
	serviceContainer := services.Container{
		Track: service,
	}
	return TestUtil{serviceContainer}
}

func WithShelfServiceMocked(method string, mockedResult interface{}, mockedErr error) TestUtil {
	service := new(mocks.Service)
	service.On(method, mock.Anything).Return(mockedResult, mockedErr)
	serviceContainer := services.Container{
		Shelf: service,
	}
	return TestUtil{serviceContainer}
}

func WithUserServiceMocked(method string, mockedResult interface{}, mockedErr error) TestUtil {
	service := new(mocks.Service)
	service.On(method, mock.Anything).Return(mockedResult, mockedErr)
	serviceContainer := services.Container{
		User: service,
	}
	return TestUtil{serviceContainer}
}

func WithAccountServiceMocked(method string, mockedResult interface{}, mockedErr error) TestUtil {
	service := new(mocks.Service)
	service.On(method, mock.Anything).Return(mockedResult, mockedErr)
	serviceContainer := services.Container{
		Account: service,
	}
	return TestUtil{serviceContainer}
}

func WithPreferenceServiceMocked(method string, mockedResult interface{}, mockedErr error) TestUtil {
	service := new(mocks.Service)
	service.On(method, mock.Anything).Return(mockedResult, mockedErr)
	serviceContainer := services.Container{
		Preference: service,
	}
	return TestUtil{serviceContainer}
}

func WithSubjectServiceMocked(method string, mockedResult interface{}, mockedErr error) TestUtil {
	service := new(mocks.Service)
	service.On(method, mock.Anything).Return(mockedResult, mockedErr)
	serviceContainer := services.Container{
		Subject: service,
	}
	return TestUtil{serviceContainer}
}

func (this TestUtil) ServeHTTP(req *http.Request) (*http.Response, structs.Response) {
	application := app.New(this.ServiceContainer)
	server := application.GetServer()
	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)
	result := resp.Result()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		panic("Failed to read body data in testutil.")
	}
	var response structs.Response
	json.Unmarshal(body, &response)
	return result, response
}
