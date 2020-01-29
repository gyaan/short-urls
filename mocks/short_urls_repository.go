// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/gyaan/short-urls/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// ShortUrls is an autogenerated mock type for the ShortUrls type
type ShortUrls struct {
	mock.Mock
}

// CreateShortUrl provides a mock function with given fields: ctx, urlString
func (_m *ShortUrls) CreateShortUrl(ctx context.Context, urlString string) (*models.ShortUrl, error) {
	ret := _m.Called(ctx, urlString)

	var r0 *models.ShortUrl
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.ShortUrl); ok {
		r0 = rf(ctx, urlString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShortUrl)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, urlString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteShortUrl provides a mock function with given fields: ctx, shortUrlId
func (_m *ShortUrls) DeleteShortUrl(ctx context.Context, shortUrlId string) error {
	ret := _m.Called(ctx, shortUrlId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, shortUrlId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAShortUrl provides a mock function with given fields: ctx, shortUrlId
func (_m *ShortUrls) GetAShortUrl(ctx context.Context, shortUrlId string) (*models.ShortUrl, error) {
	ret := _m.Called(ctx, shortUrlId)

	var r0 *models.ShortUrl
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.ShortUrl); ok {
		r0 = rf(ctx, shortUrlId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShortUrl)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortUrlId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActualUrlOfAShortUrl provides a mock function with given fields: ctx, shortUrl
func (_m *ShortUrls) GetActualUrlOfAShortUrl(ctx context.Context, shortUrl string) (*models.ShortUrl, error) {
	ret := _m.Called(ctx, shortUrl)

	var r0 *models.ShortUrl
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.ShortUrl); ok {
		r0 = rf(ctx, shortUrl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShortUrl)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortUrl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllShortUrls provides a mock function with given fields: ctx
func (_m *ShortUrls) GetAllShortUrls(ctx context.Context) ([]models.ShortUrl, error) {
	ret := _m.Called(ctx)

	var r0 []models.ShortUrl
	if rf, ok := ret.Get(0).(func(context.Context) []models.ShortUrl); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ShortUrl)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementClickCountOfShortUrl provides a mock function with given fields: ctx, shortUrl
func (_m *ShortUrls) IncrementClickCountOfShortUrl(ctx context.Context, shortUrl string) error {
	ret := _m.Called(ctx, shortUrl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, shortUrl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateShortUrls provides a mock function with given fields: ctx, shortUrlId, url
func (_m *ShortUrls) UpdateShortUrls(ctx context.Context, shortUrlId string, url models.ShortUrl) error {
	ret := _m.Called(ctx, shortUrlId, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, models.ShortUrl) error); ok {
		r0 = rf(ctx, shortUrlId, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
