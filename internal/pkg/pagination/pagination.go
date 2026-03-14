package pagination

type Request struct {
	Page  int
	Limit int
}

func Normalize(page, limit int) Request {
	if page < 1 {
		page = 1
	}

	switch limit {
	case 10, 20, 50:
	default:
		limit = 10
	}

	return Request{
		Page:  page,
		Limit: limit,
	}
}

func Offset(page, limit int) int {
	return (page - 1) * limit
}

func TotalPages(total int64, limit int) int64 {
	if limit <= 0 {
		return 1
	}
	pages := total / int64(limit)
	if total%int64(limit) != 0 {
		pages++
	}
	if pages == 0 {
		pages = 1
	}
	return pages
}
