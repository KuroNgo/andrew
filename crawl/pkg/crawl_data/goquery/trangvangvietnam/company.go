package trangvangvietnam

import (
	"crawl/domain"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type CompanyCrawl struct {
	companies domain.Companies
}

type ICompanyCrawl interface {
	GetByURL(url string) error
	GetByTotalPages(url string) error
	GetAll(currentUrl string) error
}

func NewCompaniesCrawl(companies domain.Companies) ICompanyCrawl {
	return &CompanyCrawl{companies: companies}
}

func (c *CompanyCrawl) GetByURL(url string) error {
	//TODO implement me
	panic("implement me")
}

func (c *CompanyCrawl) GetByTotalPages(url string) error {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	lastPageLink, _ := doc.Find("paging").Attr("href") // Đọc dữ liệu từ thẻ a của ul.pagination
	if lastPageLink == "javascript:void();" {          // Trường hợp chỉ có 1 page thì sẽ không có url
		c.companies.TotalPages = 1
		return nil
	}
	split := strings.Split(lastPageLink, "?page=")
	totalPages, _ := strconv.Atoi(split[1])
	c.companies.TotalPages = totalPages
	return nil
}

func (c *CompanyCrawl) GetAll(currentUrl string) error {
	//TODO implement me
	panic("implement me")
}
