package domain

import "context"

// TechBook is just []byte at this time.
type TechBook []byte

// TechBookRepository defines what should be implemented as repository.
type TechBookRepository interface {
	SetTechBookURL(ctx context.Context, techBookURL string) error
	GetTechBookURL(ctx context.Context) (string, error)
	GetTechBook(ctx context.Context) (TechBook, error)
}
