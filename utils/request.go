package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

/*
# VerifyPageQuery

This function verifies if the page query parameter is valid. If the page query parameter is not present, it defaults to 1. If the page query parameter is invalid, it returns an error.

# Parameters
  - r: *http.Request
  - w: http.ResponseWriter

# Returns
  - int: page number (default 1) or 0 if the page number is invalid
  - error: error if the page number is invalid
*/
func VerifyPageQuery(r *http.Request, w http.ResponseWriter) (int, error) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid page number"))
		return 0, err
	}

	return pageInt, nil
}

/*
# VerifyItemsPerPageQuery

This function verifies if the itemsPerPage query parameter is valid. If the itemsPerPage query parameter is not present, it defaults to 30. If the itemsPerPage query parameter is invalid, it returns an error.

# Parameters
  - r: *http.Request
  - w: http.ResponseWriter

# Returns
  - int: itemsPerPage number (default 30) or 0 if the itemsPerPage number is invalid
  - error: error if the itemsPerPage number is invalid
*/
func VerifyItemsPerPageQuery(r *http.Request, w http.ResponseWriter) (int, error) {
	itemsPerPage := r.URL.Query().Get("itemsPerPage")
	if itemsPerPage == "" {
		itemsPerPage = "30"
	}

	itemsPerPageInt, err := strconv.Atoi(itemsPerPage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid itemsPerPage number"))
		return 0, err
	}

	return itemsPerPageInt, nil
}

/*
# CalculateTotalPages

This function calculates the total number of pages based on the total number of products and the number of items per page.

# Parameters
  - productsCount: int
  - itemsPerPage: int

# Returns
  - int: total number of pages
*/
func CalculateTotalPages(productsCount int, itemsPerPage int) int {
	totalPages := productsCount / itemsPerPage
	if productsCount%itemsPerPage != 0 {
		totalPages++
	}
	return totalPages
}

/*
# WritePrevAndNextPages

This function writes the previous and next pages URLs based on the current page number, the total number of products, and the number of items per page.

# Parameters
  - page: int
  - productsCount: int
  - itemsPerPage: int

# Returns
  - string: previous page URL
  - string: next page URL
*/
func WritePrevAndNextPages(page int, productsCount int, itemsPerPage int) (string, string) {
	var prevPage string
	if page > 1 {
		prevPage = fmt.Sprintf("/products?page=%d", page-1)
		if itemsPerPage != DefaultItemsPerPage {
			prevPage = fmt.Sprintf("%s&itemsPerPage=%d", prevPage, itemsPerPage)
		}
	}

	var nextPage string
	if productsCount > page*itemsPerPage {
		nextPage = fmt.Sprintf("/products?page=%d", page+1)
		if itemsPerPage != DefaultItemsPerPage {
			nextPage = fmt.Sprintf("%s&itemsPerPage=%d", nextPage, itemsPerPage)
		}
	}

	return prevPage, nextPage
}
