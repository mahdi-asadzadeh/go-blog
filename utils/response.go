package utils

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statuscode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	StatusCode int         `json:"statuscode"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

func CreatePagedResponse(request *http.Request, resources []interface{}, resource_name string, page, page_size, count int) map[string]interface{} {
	response := CreatePageMeta(request, len(resources), page, page_size, count)
	response[resource_name] = resources
	return response
}

func CreatePageMeta(request *http.Request, loadedItemsCount, page, page_size, totalItemsCount int) map[string]interface{} {
	page_meta := map[string]interface{}{}
	page_meta["offset"] = (page - 1) * page_size
	page_meta["requested_page_size"] = page_size
	page_meta["current_page_number"] = page
	page_meta["current_items_count"] = loadedItemsCount

	page_meta["prev_page_number"] = 1
	total_pages_count := int(math.Ceil(float64(totalItemsCount) / float64(page_size)))
	page_meta["total_pages_count"] = total_pages_count

	if page < total_pages_count {
		page_meta["has_next_page"] = true
		page_meta["next_page_number"] = page + 1
	} else {
		page_meta["has_next_page"] = false
		page_meta["next_page_number"] = 1
	}
	if page > 1 {
		page_meta["prev_page_number"] = page - 1
	} else {
		page_meta["has_prev_page"] = false
		page_meta["prev_page_number"] = 1
	}

	page_meta["next_page_url"] = fmt.Sprintf("%v?page=%d&page_size=%d", request.URL.Path, page_meta["next_page_number"], page_meta["requested_page_size"])
	page_meta["prev_page_url"] = fmt.Sprintf("%s?page=%d&page_size=%d", request.URL.Path, page_meta["prev_page_number"], page_meta["requested_page_size"])

	response := gin.H{
		"success":   true,
		"page_meta": page_meta,
	}

	return response
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {
	jsonResponse := Response{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}
	ctx.JSON(StatusCode, jsonResponse)
}

func APIErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	jsonResponse := ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Error:      Error,
	}
	ctx.JSON(StatusCode, jsonResponse)
}

func CreateErrorWithMessage(message string) map[string]interface{} {
	return map[string]interface{}{
		"success":       false,
		"full_messages": []string{message},
	}
}

func CreateDetailedError(key string, err error) map[string]interface{} {
	return map[string]interface{}{
		"success":       false,
		"full_messages": []string{fmt.Sprintf("s -> %v", key, err.Error())},
		"errors":        err,
	}
}
