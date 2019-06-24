package infra

import (
	"context"
	"github.com/juntaki/techbook-qrcode/src/infra/models"

	"cloud.google.com/go/datastore"
	"github.com/juntaki/techbook-qrcode/src/domain"
	"gopkg.in/gorp.v2"
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

type QRCodeRepositorySQLImpl struct {
	db *gorp.DbMap
}

func NewQRCodeRepositorySQLImpl(db *gorp.DbMap) domain.QRCodeRepository {
	return &QRCodeRepositorySQLImpl{
		db: db,
	}
}

func (i *QRCodeRepositorySQLImpl) GetQRCodes(ctx context.Context) ([]*domain.QRCode, error) {
	var m []*models.Qrcode
	_, err := i.db.Select(&m, "select * from qrcode order by index asc")
	if err != nil {
		return nil, err
	}

	var ret []*domain.QRCode
	for _, mm := range m {
		ret = append(ret, &domain.QRCode{
			ID:    mm.Id,
			Index: mm.Index,
		})
	}
	return ret, nil
}

func (i *QRCodeRepositorySQLImpl) IsExistQRCode(ctx context.Context, code *domain.QRCode) bool {
	c, err := i.db.SelectInt("select count(*) from qrcode where id = ?", code.ID)
	if err != nil || c == 0 {
		return false
	}
	return true
}

func (i *QRCodeRepositorySQLImpl) SaveQRCodes(ctx context.Context, qrs []*domain.QRCode) error {
	for _, q := range qrs {
		m := &models.Qrcode{
			Id:    q.ID,
			Index: q.Index,
		}
		err := i.db.Insert(m)
		if err != nil {
			return err
		}
	}
	return nil
}
