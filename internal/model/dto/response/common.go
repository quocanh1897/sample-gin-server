package dto

import (
	apperror "github.com/quocanh1897/sample-gin-server/internal/error"
)

type Response[T interface{}] struct {
	Data     T                `json:"data,omitempty"`
	Meta     any              `json:"meta,omitempty"`
	Links    any              `json:"links,omitempty"`
	Included any              `json:"included,omitempty"`
	Errors   []apperror.Error `json:"errors,omitempty"`
}

func SuccessResponse(data interface{}) Response[interface{}] {
	return Response[interface{}]{
		Data: data,
	}
}

func FailedResponse(err apperror.Error) Response[interface{}] {
	return Response[interface{}]{
		Errors: []apperror.Error{err},
		Data:   nil,
	}
}

type PaginationResponse struct {
	Data     any                    `json:"data"`
	Links    PaginationLinkResponse `json:"links"`
	Errors   []apperror.Error       `json:"errors,omitempty"`
	Meta     PaginationMetaResponse `json:"meta"`
	Included []interface{}          `json:"included,omitempty"`
}

type PaginationMetaResponse struct {
	Sort   string `json:"sort,omitempty"`
	Offset int32  `json:"offset"`
	Count  int32  `json:"count"`
	Limit  int32  `json:"limit"`
}

type PaginationLinkResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	First    string `json:"first"`
	Last     string `json:"last"`
}
