// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// UserStore is an autogenerated mock type for the UserStore type
type UserStore struct {
	mock.Mock
}

// AnalyticsActiveCount provides a mock function with given fields: time
func (_m *UserStore) AnalyticsActiveCount(time int64) (int64, *model.AppError) {
	ret := _m.Called(time)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64) int64); ok {
		r0 = rf(time)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int64) *model.AppError); ok {
		r1 = rf(time)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// AnalyticsGetInactiveUsersCount provides a mock function with given fields:
func (_m *UserStore) AnalyticsGetInactiveUsersCount() (int64, *model.AppError) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// AnalyticsGetSystemAdminCount provides a mock function with given fields:
func (_m *UserStore) AnalyticsGetSystemAdminCount() (int64, *model.AppError) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// ClearAllCustomRoleAssignments provides a mock function with given fields:
func (_m *UserStore) ClearAllCustomRoleAssignments() *model.AppError {
	ret := _m.Called()

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func() *model.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// ClearCaches provides a mock function with given fields:
func (_m *UserStore) ClearCaches() {
	_m.Called()
}

// Count provides a mock function with given fields: options
func (_m *UserStore) Count(options model.UserCountOptions) (int64, *model.AppError) {
	ret := _m.Called(options)

	var r0 int64
	if rf, ok := ret.Get(0).(func(model.UserCountOptions) int64); ok {
		r0 = rf(options)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(model.UserCountOptions) *model.AppError); ok {
		r1 = rf(options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *UserStore) Get(id string) (*model.User, *model.AppError) {
	ret := _m.Called(id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *UserStore) GetAll() ([]*model.User, *model.AppError) {
	ret := _m.Called()

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func() []*model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAllAfter provides a mock function with given fields: limit, afterId
func (_m *UserStore) GetAllAfter(limit int, afterId string) ([]*model.User, *model.AppError) {
	ret := _m.Called(limit, afterId)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(int, string) []*model.User); ok {
		r0 = rf(limit, afterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int, string) *model.AppError); ok {
		r1 = rf(limit, afterId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAllProfiles provides a mock function with given fields: options
func (_m *UserStore) GetAllProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) *model.AppError); ok {
		r1 = rf(options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAllProfilesInChannel provides a mock function with given fields: channelId, allowFromCache
func (_m *UserStore) GetAllProfilesInChannel(channelId string, allowFromCache bool) (map[string]*model.User, *model.AppError) {
	ret := _m.Called(channelId, allowFromCache)

	var r0 map[string]*model.User
	if rf, ok := ret.Get(0).(func(string, bool) map[string]*model.User); ok {
		r0 = rf(channelId, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, bool) *model.AppError); ok {
		r1 = rf(channelId, allowFromCache)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAllUsingAuthService provides a mock function with given fields: authService
func (_m *UserStore) GetAllUsingAuthService(authService string) ([]*model.User, *model.AppError) {
	ret := _m.Called(authService)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(authService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(authService)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetAnyUnreadPostCountForChannel provides a mock function with given fields: userId, channelId
func (_m *UserStore) GetAnyUnreadPostCountForChannel(userId string, channelId string) (int64, *model.AppError) {
	ret := _m.Called(userId, channelId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(userId, channelId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(userId, channelId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByAuth provides a mock function with given fields: authData, authService
func (_m *UserStore) GetByAuth(authData *string, authService string) (*model.User, *model.AppError) {
	ret := _m.Called(authData, authService)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*string, string) *model.User); ok {
		r0 = rf(authData, authService)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*string, string) *model.AppError); ok {
		r1 = rf(authData, authService)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: email
func (_m *UserStore) GetByEmail(email string) (*model.User, *model.AppError) {
	ret := _m.Called(email)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *UserStore) GetByUsername(username string) store.StoreChannel {
	ret := _m.Called(username)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetChannelGroupUsers provides a mock function with given fields: channelID
func (_m *UserStore) GetChannelGroupUsers(channelID string) ([]*model.User, *model.AppError) {
	ret := _m.Called(channelID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(channelID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetEtagForAllProfiles provides a mock function with given fields:
func (_m *UserStore) GetEtagForAllProfiles() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetEtagForProfiles provides a mock function with given fields: teamId
func (_m *UserStore) GetEtagForProfiles(teamId string) string {
	ret := _m.Called(teamId)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(teamId)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetEtagForProfilesNotInTeam provides a mock function with given fields: teamId
func (_m *UserStore) GetEtagForProfilesNotInTeam(teamId string) string {
	ret := _m.Called(teamId)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(teamId)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetForLogin provides a mock function with given fields: loginId, allowSignInWithUsername, allowSignInWithEmail
func (_m *UserStore) GetForLogin(loginId string, allowSignInWithUsername bool, allowSignInWithEmail bool) (*model.User, *model.AppError) {
	ret := _m.Called(loginId, allowSignInWithUsername, allowSignInWithEmail)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(string, bool, bool) *model.User); ok {
		r0 = rf(loginId, allowSignInWithUsername, allowSignInWithEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, bool, bool) *model.AppError); ok {
		r1 = rf(loginId, allowSignInWithUsername, allowSignInWithEmail)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetNewUsersForTeam provides a mock function with given fields: teamId, offset, limit, viewRestrictions
func (_m *UserStore) GetNewUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	ret := _m.Called(teamId, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamId, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int, *model.ViewUsersRestrictions) *model.AppError); ok {
		r1 = rf(teamId, offset, limit, viewRestrictions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfileByGroupChannelIdsForUser provides a mock function with given fields: userId, channelIds
func (_m *UserStore) GetProfileByGroupChannelIdsForUser(userId string, channelIds []string) (map[string][]*model.User, *model.AppError) {
	ret := _m.Called(userId, channelIds)

	var r0 map[string][]*model.User
	if rf, ok := ret.Get(0).(func(string, []string) map[string][]*model.User); ok {
		r0 = rf(userId, channelIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, []string) *model.AppError); ok {
		r1 = rf(userId, channelIds)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfileByIds provides a mock function with given fields: userIds, options, allowFromCache
func (_m *UserStore) GetProfileByIds(userIds []string, options *store.UserGetByIdsOpts, allowFromCache bool) ([]*model.User, *model.AppError) {
	ret := _m.Called(userIds, options, allowFromCache)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func([]string, *store.UserGetByIdsOpts, bool) []*model.User); ok {
		r0 = rf(userIds, options, allowFromCache)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func([]string, *store.UserGetByIdsOpts, bool) *model.AppError); ok {
		r1 = rf(userIds, options, allowFromCache)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfiles provides a mock function with given fields: options
func (_m *UserStore) GetProfiles(options *model.UserGetOptions) ([]*model.User, *model.AppError) {
	ret := _m.Called(options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(*model.UserGetOptions) []*model.User); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.UserGetOptions) *model.AppError); ok {
		r1 = rf(options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfilesByUsernames provides a mock function with given fields: usernames, viewRestrictions
func (_m *UserStore) GetProfilesByUsernames(usernames []string, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	ret := _m.Called(usernames, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func([]string, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(usernames, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func([]string, *model.ViewUsersRestrictions) *model.AppError); ok {
		r1 = rf(usernames, viewRestrictions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfilesInChannel provides a mock function with given fields: channelId, offset, limit
func (_m *UserStore) GetProfilesInChannel(channelId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(channelId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(channelId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetProfilesInChannelByStatus provides a mock function with given fields: channelId, offset, limit
func (_m *UserStore) GetProfilesInChannelByStatus(channelId string, offset int, limit int) store.StoreChannel {
	ret := _m.Called(channelId, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, int, int) store.StoreChannel); ok {
		r0 = rf(channelId, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetProfilesNotInChannel provides a mock function with given fields: teamId, channelId, groupConstrained, offset, limit, viewRestrictions
func (_m *UserStore) GetProfilesNotInChannel(teamId string, channelId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	ret := _m.Called(teamId, channelId, groupConstrained, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, bool, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamId, channelId, groupConstrained, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, bool, int, int, *model.ViewUsersRestrictions) *model.AppError); ok {
		r1 = rf(teamId, channelId, groupConstrained, offset, limit, viewRestrictions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfilesNotInTeam provides a mock function with given fields: teamId, groupConstrained, offset, limit, viewRestrictions
func (_m *UserStore) GetProfilesNotInTeam(teamId string, groupConstrained bool, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	ret := _m.Called(teamId, groupConstrained, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, bool, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamId, groupConstrained, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, bool, int, int, *model.ViewUsersRestrictions) *model.AppError); ok {
		r1 = rf(teamId, groupConstrained, offset, limit, viewRestrictions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetProfilesWithoutTeam provides a mock function with given fields: offset, limit, viewRestrictions
func (_m *UserStore) GetProfilesWithoutTeam(offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) store.StoreChannel {
	ret := _m.Called(offset, limit, viewRestrictions)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int, *model.ViewUsersRestrictions) store.StoreChannel); ok {
		r0 = rf(offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetRecentlyActiveUsersForTeam provides a mock function with given fields: teamId, offset, limit, viewRestrictions
func (_m *UserStore) GetRecentlyActiveUsersForTeam(teamId string, offset int, limit int, viewRestrictions *model.ViewUsersRestrictions) ([]*model.User, *model.AppError) {
	ret := _m.Called(teamId, offset, limit, viewRestrictions)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, int, int, *model.ViewUsersRestrictions) []*model.User); ok {
		r0 = rf(teamId, offset, limit, viewRestrictions)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, int, int, *model.ViewUsersRestrictions) *model.AppError); ok {
		r1 = rf(teamId, offset, limit, viewRestrictions)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetSystemAdminProfiles provides a mock function with given fields:
func (_m *UserStore) GetSystemAdminProfiles() (map[string]*model.User, *model.AppError) {
	ret := _m.Called()

	var r0 map[string]*model.User
	if rf, ok := ret.Get(0).(func() map[string]*model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetTeamGroupUsers provides a mock function with given fields: teamID
func (_m *UserStore) GetTeamGroupUsers(teamID string) ([]*model.User, *model.AppError) {
	ret := _m.Called(teamID)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string) []*model.User); ok {
		r0 = rf(teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(teamID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// GetUnreadCount provides a mock function with given fields: userId
func (_m *UserStore) GetUnreadCount(userId string) (int64, error) {
	ret := _m.Called(userId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUnreadCountForChannel provides a mock function with given fields: userId, channelId
func (_m *UserStore) GetUnreadCountForChannel(userId string, channelId string) store.StoreChannel {
	ret := _m.Called(userId, channelId)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(userId, channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetUsersBatchForIndexing provides a mock function with given fields: startTime, endTime, limit
func (_m *UserStore) GetUsersBatchForIndexing(startTime int64, endTime int64, limit int) ([]*model.UserForIndexing, *model.AppError) {
	ret := _m.Called(startTime, endTime, limit)

	var r0 []*model.UserForIndexing
	if rf, ok := ret.Get(0).(func(int64, int64, int) []*model.UserForIndexing); ok {
		r0 = rf(startTime, endTime, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserForIndexing)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(int64, int64, int) *model.AppError); ok {
		r1 = rf(startTime, endTime, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// InferSystemInstallDate provides a mock function with given fields:
func (_m *UserStore) InferSystemInstallDate() (int64, *model.AppError) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func() *model.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// InvalidatProfileCacheForUser provides a mock function with given fields: userId
func (_m *UserStore) InvalidatProfileCacheForUser(userId string) {
	_m.Called(userId)
}

// InvalidateProfilesInChannelCache provides a mock function with given fields: channelId
func (_m *UserStore) InvalidateProfilesInChannelCache(channelId string) {
	_m.Called(channelId)
}

// InvalidateProfilesInChannelCacheByUser provides a mock function with given fields: userId
func (_m *UserStore) InvalidateProfilesInChannelCacheByUser(userId string) {
	_m.Called(userId)
}

// PermanentDelete provides a mock function with given fields: userId
func (_m *UserStore) PermanentDelete(userId string) *model.AppError {
	ret := _m.Called(userId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// ResetLastPictureUpdate provides a mock function with given fields: userId
func (_m *UserStore) ResetLastPictureUpdate(userId string) *model.AppError {
	ret := _m.Called(userId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// Save provides a mock function with given fields: user
func (_m *UserStore) Save(user *model.User) (*model.User, *model.AppError) {
	ret := _m.Called(user)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.User) *model.AppError); ok {
		r1 = rf(user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Search provides a mock function with given fields: teamId, term, options
func (_m *UserStore) Search(teamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	ret := _m.Called(teamId, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(teamId, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(teamId, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchInChannel provides a mock function with given fields: channelId, term, options
func (_m *UserStore) SearchInChannel(channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	ret := _m.Called(channelId, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(channelId, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(channelId, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchNotInChannel provides a mock function with given fields: teamId, channelId, term, options
func (_m *UserStore) SearchNotInChannel(teamId string, channelId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	ret := _m.Called(teamId, channelId, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(teamId, channelId, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(teamId, channelId, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchNotInTeam provides a mock function with given fields: notInTeamId, term, options
func (_m *UserStore) SearchNotInTeam(notInTeamId string, term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	ret := _m.Called(notInTeamId, term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(notInTeamId, term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(notInTeamId, term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// SearchWithoutTeam provides a mock function with given fields: term, options
func (_m *UserStore) SearchWithoutTeam(term string, options *model.UserSearchOptions) ([]*model.User, *model.AppError) {
	ret := _m.Called(term, options)

	var r0 []*model.User
	if rf, ok := ret.Get(0).(func(string, *model.UserSearchOptions) []*model.User); ok {
		r0 = rf(term, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.User)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, *model.UserSearchOptions) *model.AppError); ok {
		r1 = rf(term, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: user, allowRoleUpdate
func (_m *UserStore) Update(user *model.User, allowRoleUpdate bool) (*model.UserUpdate, *model.AppError) {
	ret := _m.Called(user, allowRoleUpdate)

	var r0 *model.UserUpdate
	if rf, ok := ret.Get(0).(func(*model.User, bool) *model.UserUpdate); ok {
		r0 = rf(user, allowRoleUpdate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserUpdate)
		}
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(*model.User, bool) *model.AppError); ok {
		r1 = rf(user, allowRoleUpdate)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateAuthData provides a mock function with given fields: userId, service, authData, email, resetMfa
func (_m *UserStore) UpdateAuthData(userId string, service string, authData *string, email string, resetMfa bool) (string, *model.AppError) {
	ret := _m.Called(userId, service, authData, email, resetMfa)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, *string, string, bool) string); ok {
		r0 = rf(userId, service, authData, email, resetMfa)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string, *string, string, bool) *model.AppError); ok {
		r1 = rf(userId, service, authData, email, resetMfa)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// UpdateFailedPasswordAttempts provides a mock function with given fields: userId, attempts
func (_m *UserStore) UpdateFailedPasswordAttempts(userId string, attempts int) *model.AppError {
	ret := _m.Called(userId, attempts)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, int) *model.AppError); ok {
		r0 = rf(userId, attempts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdateLastPictureUpdate provides a mock function with given fields: userId
func (_m *UserStore) UpdateLastPictureUpdate(userId string) *model.AppError {
	ret := _m.Called(userId)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string) *model.AppError); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdateMfaActive provides a mock function with given fields: userId, active
func (_m *UserStore) UpdateMfaActive(userId string, active bool) store.StoreChannel {
	ret := _m.Called(userId, active)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, bool) store.StoreChannel); ok {
		r0 = rf(userId, active)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateMfaSecret provides a mock function with given fields: userId, secret
func (_m *UserStore) UpdateMfaSecret(userId string, secret string) *model.AppError {
	ret := _m.Called(userId, secret)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) *model.AppError); ok {
		r0 = rf(userId, secret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: userId, newPassword
func (_m *UserStore) UpdatePassword(userId string, newPassword string) *model.AppError {
	ret := _m.Called(userId, newPassword)

	var r0 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) *model.AppError); ok {
		r0 = rf(userId, newPassword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AppError)
		}
	}

	return r0
}

// UpdateUpdateAt provides a mock function with given fields: userId
func (_m *UserStore) UpdateUpdateAt(userId string) (int64, *model.AppError) {
	ret := _m.Called(userId)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(userId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// VerifyEmail provides a mock function with given fields: userId, email
func (_m *UserStore) VerifyEmail(userId string, email string) (string, *model.AppError) {
	ret := _m.Called(userId, email)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(userId, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 *model.AppError
	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(userId, email)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}
