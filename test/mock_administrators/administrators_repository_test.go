package mock_administrators

import (
	"app/domain"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestFindByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var admin domain.Administrator
	var err error

	mockSample := NewMockAdministratorsRepository(ctrl)

	mockSample.EXPECT().FindByEmail("test@example.com").Return(admin, err)
	mockSample.FindByEmail("test@example.com")
}
