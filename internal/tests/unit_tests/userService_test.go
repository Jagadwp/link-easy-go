package tests

import (
	"testing"
)

// var IsUserAllowedToEditUrl = services.IsUserAllowedToEditUrl
// var IsUrlValid = services.IsUrlValid
// var NewUrlService = services.NewUrlService
// var Now = services.Now
// var GenerateLink = helper.GenerateLink

func Login(t *testing.T) {
	// oldUrl := &models.User{
	// 	ID:        1,
	// 	Username:  "Jagadul",
	// 	Fullname:  "Jagad Doel dul dul",
	// 	Email:     "jagad@mail.com",
	// 	Password:  "admin",
	// 	Admin:     1,
	// 	CreatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	// 	UpdatedAt: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
	// }

	// cases := []struct {
	// 	name                   string
	// 	IsUserAllowedToEditUrl func(userID int, userIDInUrl int) bool
	// 	IsUrlValid             func(toTest string) bool
	// 	args1                  int
	// 	args2                  int
	// 	out1                   *models.Url
	// 	out2                   error
	// }{
	// 	{
	// 		name: "success no error",
	// 		IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
	// 			return true
	// 		},
	// 		IsUrlValid: func(toTest string) bool {
	// 			return true
	// 		},
	// 		args1: oldUrl.ID,
	// 		args2: oldUrl.UserID,
	// 		out1:  oldUrl,
	// 		out2:  nil,
	// 	},
	// 	{
	// 		name: "should error when url not found",
	// 		IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
	// 			return true
	// 		},
	// 		IsUrlValid: func(toTest string) bool {
	// 			return true
	// 		},
	// 		args1: oldUrl.ID,
	// 		args2: oldUrl.UserID,
	// 		out1:  &models.Url{},
	// 		out2:  shared.ErrUrlNotFound,
	// 	},
	// 	{
	// 		name: "should error when user not allowed to edit url",
	// 		IsUserAllowedToEditUrl: func(userID int, userIDInUrl int) bool {
	// 			return false
	// 		},
	// 		IsUrlValid: func(toTest string) bool {
	// 			return true
	// 		},
	// 		args1: oldUrl.ID,
	// 		args2: oldUrl.UserID,
	// 		out1:  &models.Url{},
	// 		out2:  shared.ErrForbiddenToAccess,
	// 	},
	// }

	// for _, tc := range cases {
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		mockRepo := new(mocks.IUrlRepository)

	// 		// mock helper function
	// 		services.IsUserAllowedToEditUrl = tc.IsUserAllowedToEditUrl
	// 		services.IsUrlValid = tc.IsUrlValid

	// 		if tc.name == "should error when url not found" {
	// 			mockRepo.On("GetUrlById", oldUrl.ID).Return(&models.Url{ID: 0}, nil)
	// 		} else {
	// 			mockRepo.On("GetUrlById", oldUrl.ID).Return(oldUrl, nil)
	// 		}

	// 		urlService := services.NewUrlService(mockRepo)
	// 		got, gotErr := urlService.GetUrlUserById(tc.args1, tc.args2)
	// 		assert.Equal(t, tc.out1, got)
	// 		assert.Equal(t, tc.out2, gotErr)
	// 	})
	// }
}

func InsertUser(t *testing.T) {

}

func GetAllUsers(t *testing.T) {

}

func GetUserById(t *testing.T) {

}

func UpdateUser(t *testing.T) {

}

func DeleteUser(t *testing.T) {

}
