package pagination

import (
	"math"

	"github.com/lastares/claymore/protobuf/pagination"
)

type Paginator[T any] struct {
	List       T
	Pagination *Builder
}

type Builder struct {
	Page       int
	PageSize   int
	Total      int
	TotalPages int
	Next       int
	Prev       int
	HasMore    bool
}

func NewPaginator[T any](total int, list T, p *pagination.Pagination) *Paginator[T] {
	return &Paginator[T]{
		List: list,
		Pagination: NewPaginationBuilder(total).WithPagination(&pagination.Pagination{
			Page:     p.Page,
			PageSize: p.PageSize,
		}).Build(),
	}
}

// NewPaginationBuilder 创建一个新的空的 PaginationBuilder 实例
func NewPaginationBuilder(total int) *Builder {
	return &Builder{
		Total: total,
	}
}

func (pb *Builder) setPage(page int) {
	if page == 0 {
		page = 1
	}
	pb.Page = page
}

func (pb *Builder) setPageSize(pageSize int) {
	if pageSize == 0 {
		pageSize = 10
	}
	pb.PageSize = pageSize
}

func (pb *Builder) WithPagination(p *pagination.Pagination) *Builder {
	pb.setPage(int(p.Page))
	pb.setPageSize(int(p.PageSize))
	return pb
}

func (pb *Builder) setTotalPages() {
	pb.TotalPages = int(math.Ceil(float64(pb.Total) / float64(pb.PageSize)))
}

func (pb *Builder) setPrev() {
	if pb.Page > 1 {
		pb.Prev = max(1, pb.Page-1)
	}
}

func (pb *Builder) setNext() {
	if pb.TotalPages > 1 && pb.Page < pb.TotalPages {
		pb.Next = min(pb.TotalPages, pb.Page+1)
	}
}

// HasMore
func (pb *Builder) setHasMore() {
	pb.HasMore = pb.Page < pb.TotalPages
}

// Build 构建最终的 PaginationBuilder 实例
func (pb *Builder) Build() *Builder {
	pb.setTotalPages()
	pb.setPrev()
	pb.setNext()
	pb.setHasMore()
	return pb
}
