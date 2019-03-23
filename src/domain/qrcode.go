package domain

import (
	"bytes"
	"context"
	"image/png"
	"os"

	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
)

// QRCode is just Generated ID to validate the permission to download techbook.
type QRCode struct {
	ID    string
	Index int
}

func NewQRCode(index int) *QRCode {
	return &QRCode{
		ID:    uuid.Must(uuid.NewRandom()).String(),
		Index: index,
	}
}

// GetURL returns techbook download URL.
func (q *QRCode) GetURL() string {
	return os.Getenv("QRCODE_BASE_URL") + q.ID
}

// GetQRCode generates QRCode of techbook download URL.
func (q *QRCode) GetQRCode() []byte {
	qrCode, _ := qr.Encode(q.GetURL(), qr.H, qr.Auto)
	var buf bytes.Buffer
	png.Encode(&buf, qrCode)
	return buf.Bytes()
}

// QRCodeRepository defines what should be implemented as repository.
type QRCodeRepository interface {
	GetQRCodes(ctx context.Context) ([]*QRCode, error)
	IsExistQRCode(ctx context.Context, code *QRCode) bool
	// GenerateQRCodes(ctx context.Context, len int) error
	SaveQRCodes(ctx context.Context, qrs []*QRCode) error
}
