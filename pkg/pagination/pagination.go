package pagination

import "errors"

type Response struct {
	TotalPage    int64       `json:"total_page"`
	CurrentPage  int64       `json:"current_page"`
	NextPage     int64       `json:"next_page"`
	PreviousPage int64       `json:"previous_page"`
	TotalItem    int64       `json:"total_item"`
	Data         interface{} `json:"data"`
}

type page struct {
	currentPage int64
	totalItem   int64
	data        interface{}
	perPageItem int
}

type Pagination interface {
	GetPagination() (Response, error)
	GetLastPage() (int64, error)
	GetNextPage() int64
	GetPreviousPage() int64
}

func New(totalItem int64, currentPage int64, data interface{}, perPageItem int) Pagination {
	return &page{
		currentPage: currentPage,
		totalItem:   totalItem,
		data:        data,
		perPageItem: perPageItem,
	}
}

func (p *page) GetPagination() (Response, error) {
	response := Response{}
	response.CurrentPage = p.currentPage
	response.PreviousPage = p.GetPreviousPage()
	response.Data = p.data
	response.NextPage = p.GetNextPage()
	response.TotalItem = p.totalItem
	lastPage, err := p.GetLastPage()
	if err != nil {
		return response, err
	}
	response.TotalPage = lastPage
	return response, nil
}

func (p *page) GetLastPage() (int64, error) {
	if p.perPageItem <= 0 {
		return 0, errors.New("per page item can't be zero or less then zero")
	}
	return (p.totalItem / int64(p.perPageItem))+1, nil
}

func (p *page) GetNextPage() int64 {
	lastPage, _ := p.GetLastPage()
	if p.currentPage >= lastPage {
		return p.currentPage
	}
	return p.currentPage + 1
}

func (p *page) GetPreviousPage() int64 {
	if p.currentPage <= 1 {
		return p.currentPage
	}

	return p.currentPage - 1
}
