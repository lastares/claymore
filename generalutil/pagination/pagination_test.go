package pagination

import (
	"encoding/json"
	"testing"

	"github.com/lastares/claymore/protobuf/pagination"
)

func TestNewPaginator(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	users := []*User{
		{"aaa", 23},
		{"bbb", 24},
		{"ccc", 25},
		{"ddd", 26},
		{"eee", 27},
		{"fff", 28},
		{"ggg", 29},
		{"hhh", 30},
		{"iii", 31},
	}
	var want *Paginator[[]*User]
	a := `{"List":[{"Name":"aaa","Age":23},{"Name":"bbb","Age":24},{"Name":"ccc","Age":25},{"Name":"ddd","Age":26},{"Name":"eee","Age":27},{"Name":"fff","Age":28},{"Name":"ggg","Age":29},{"Name":"hhh","Age":30},{"Name":"iii","Age":31}],"Pagination":{"Page":1,"PageSize":5,"Total":10,"TotalPages":2,"Next":2,"Prev":0,"HasMore":true}}`
	json.Unmarshal([]byte(a), &want)
	tests := []struct {
		name  string
		total int
		list  []*User
		p     *pagination.Pagination
		want  *Paginator[[]*User]
	}{
		{
			name:  "First page with more items",
			total: 10,
			list:  users,
			p:     &pagination.Pagination{Page: 1, PageSize: 5},
			want:  want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPaginator[[]*User](tt.total, tt.list, tt.p)
			if got.Pagination.Prev != want.Pagination.Prev {
				t.Errorf("NewPaginator() got Prev = %v, want %v", got.Pagination.Prev, want.Pagination.Prev)
			}
			if got.Pagination.Next != want.Pagination.Next {
				t.Errorf("NewPaginator() got Next = %v, want %v", got.Pagination.Next, want.Pagination.Next)
			}
			if got.Pagination.HasMore != want.Pagination.HasMore {
				t.Errorf("NewPaginator() got HasMore = %v, want %v", got.Pagination.HasMore, want.Pagination.HasMore)
			}
			if got.Pagination.TotalPages != want.Pagination.TotalPages {
				t.Errorf("NewPaginator() got Total = %v, want %v", got.Pagination.TotalPages, want.Pagination.TotalPages)
			}
		})
	}

}

// TestNewPaginationBuilder 测试 NewPaginationBuilder 函数
func TestNewPaginationBuilder(t *testing.T) {
	tests := []struct {
		name        string
		total       int
		page        int
		pageSize    int
		wantPage    int
		wantSize    int
		wantTotal   int
		wantTotalPg int
		wantNext    int
		wantPrev    int
		wantMore    bool
	}{
		{
			name:        "First page with more items",
			total:       10,
			page:        1,
			pageSize:    5,
			wantPage:    1,
			wantSize:    5,
			wantTotal:   10,
			wantTotalPg: 2,
			wantNext:    2,
			wantPrev:    0,
			wantMore:    true,
		},
		{
			name:        "Middle page with more items",
			total:       10,
			page:        2,
			pageSize:    5,
			wantPage:    2,
			wantSize:    5,
			wantTotal:   10,
			wantTotalPg: 2,
			wantNext:    0,
			wantPrev:    1,
			wantMore:    false,
		},
		{
			name:        "Single page with no more items",
			total:       5,
			page:        1,
			pageSize:    5,
			wantPage:    1,
			wantSize:    5,
			wantTotal:   5,
			wantTotalPg: 1,
			wantNext:    0,
			wantPrev:    0,
			wantMore:    false,
		},
		{
			name:        "Last page with no more items",
			total:       10,
			page:        2,
			pageSize:    3,
			wantPage:    2,
			wantSize:    3,
			wantTotal:   10,
			wantTotalPg: 4,
			wantNext:    3,
			wantPrev:    1,
			wantMore:    true,
		},
		{
			name:        "Last page with no more items",
			total:       10,
			page:        3,
			pageSize:    3,
			wantPage:    3,
			wantSize:    3,
			wantTotal:   10,
			wantTotalPg: 4,
			wantNext:    4,
			wantPrev:    2,
			wantMore:    true,
		},
		{
			name:        "Last page with no more items",
			total:       10,
			page:        4,
			pageSize:    3,
			wantPage:    4,
			wantSize:    3,
			wantTotal:   10,
			wantTotalPg: 4,
			wantNext:    0,
			wantPrev:    3,
			wantMore:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewPaginationBuilder(tt.total).WithPagination(&pagination.Pagination{Page: int32(tt.page), PageSize: int32(tt.pageSize)}).Build()
			if builder.Page != tt.wantPage {
				t.Errorf("NewPaginationBuilder() got Page = %v, want %v", builder.Page, tt.wantPage)
			}

			if builder.PageSize != tt.wantSize {
				t.Errorf("NewPaginationBuilder() got PageSize = %v, want %v", builder.PageSize, tt.wantSize)
			}

			if builder.Total != tt.wantTotal {
				t.Errorf("NewPaginationBuilder() got Total = %v, want %v", builder.Total, tt.wantTotal)
			}

			if builder.TotalPages != tt.wantTotalPg {
				t.Errorf("NewPaginationBuilder() got TotalPages = %v, want %v", builder.TotalPages, tt.wantTotalPg)
			}

			if builder.Next != tt.wantNext {
				t.Errorf("NewPaginationBuilder() got Next = %v, want %v", builder.Next, tt.wantNext)
			}

			if builder.Prev != tt.wantPrev {
				t.Errorf("NewPaginationBuilder() got Prev = %v, want %v", builder.Prev, tt.wantPrev)
			}

			if builder.HasMore != tt.wantMore {
				t.Errorf("NewPaginationBuilder() got HasMore = %v, want %v", builder.HasMore, tt.wantMore)
			}
		})
	}
}
