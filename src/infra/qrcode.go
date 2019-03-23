package infra

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
	"github.com/juntaki/techbook-qrcode/src/domain"
)

type QRCodeRepositoryDatastoreImpl struct {
	db *datastore.Client
}

func NewQRCodeRepositoryDatastoreImpl(db *datastore.Client) domain.QRCodeRepository {
	return &QRCodeRepositoryDatastoreImpl{
		db: db,
	}
}

func (i *QRCodeRepositoryDatastoreImpl) GetQRCodes(ctx context.Context) ([]*domain.QRCode, error) {
	ret := []*domain.QRCode{}
	q := datastore.NewQuery("QRCode").Order("Index")
	_, err := i.db.GetAll(ctx, q, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (i *QRCodeRepositoryDatastoreImpl) IsExistQRCode(ctx context.Context, code *domain.QRCode) bool {
	q := datastore.NewQuery("QRCode").Filter("ID = ", code.ID)
	c, err := i.db.Count(ctx, q.KeysOnly())
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (i *QRCodeRepositoryDatastoreImpl) GenerateQRCodes(ctx context.Context, len int) error {
	keys := make([]*datastore.Key, len)
	qrs := make([]*domain.QRCode, len)
	for i := range qrs {
		id := uuid.Must(uuid.NewRandom()).String()
		qrs[i] = &domain.QRCode{ID: id}
		keys[i] = datastore.NameKey("QRCode", id, nil)
	}

	_, err := i.db.PutMulti(ctx, keys, qrs)
	if err != nil {
		return err
	}
	return nil
}

func (i *QRCodeRepositoryDatastoreImpl) SaveQRCodes(ctx context.Context, qrs []*domain.QRCode) error {
	keys := make([]*datastore.Key, len(qrs))
	for i, qr := range qrs {
		keys[i] = datastore.NameKey("QRCode", qr.ID, nil)
	}
	_, err := i.db.PutMulti(ctx, keys, qrs)
	if err != nil {
		return err
	}
	return nil
}
