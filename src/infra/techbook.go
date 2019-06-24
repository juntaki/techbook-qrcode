package infra

import (
	"context"
	"github.com/juntaki/techbook-qrcode/src/infra/models"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/juntaki/techbook-qrcode/src/domain"
	"gopkg.in/gorp.v2"
)

var cached domain.TechBook
var validUntil time.Time

type TechBook struct {
	URL string
}

type TechBookRepositoryDatastoreImpl struct {
	db *datastore.Client
}

func NewTechBookRepositoryDatastoreImpl(db *datastore.Client) domain.TechBookRepository {
	return &TechBookRepositoryDatastoreImpl{
		db: db,
	}
}

func (i *TechBookRepositoryDatastoreImpl) SetTechBookURL(ctx context.Context, TechBookURL string) error {
	key := datastore.NameKey("TechBook", "master", nil)
	_, err := i.db.Put(ctx, key, &TechBook{
		URL: TechBookURL,
	})
	log.Printf("Set URL %s", TechBookURL)

	validUntil = time.Now()
	log.Printf("Cache invalidate. Valid until now = %s", validUntil)

	if err != nil {
		return err
	}
	return nil
}

func (i *TechBookRepositoryDatastoreImpl) GetTechBookURL(ctx context.Context) (string, error) {
	key := datastore.NameKey("TechBook", "master", nil)
	val := TechBook{}

	err := i.db.Get(ctx, key, &val)
	if err != nil {
		return "", err
	}
	return val.URL, nil
}

func (i *TechBookRepositoryDatastoreImpl) GetTechBook(ctx context.Context) (domain.TechBook, error) {
	if validUntil.After(time.Now()) && len(cached) > 0 {
		log.Printf("Return from cache. valid until: %s", validUntil)
		return cached, nil
	}

	log.Printf("Get Techbook from URL")
	key := datastore.NameKey("TechBook", "master", nil)
	dst := TechBook{}
	err := i.db.Get(ctx, key, &dst)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(dst.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cached = ret
	validUntil = time.Now().Add(6 * time.Hour)
	log.Printf("Cache is valid until %s", validUntil)

	return ret, nil
}



type TechBookRepositorySQLImpl struct {
	db *gorp.DbMap
}

func NewTechBookRepositorySQLImpl(db *gorp.DbMap) domain.TechBookRepository {
	return &TechBookRepositorySQLImpl{
		db: db,
	}
}

func (i *TechBookRepositorySQLImpl) SetTechBookURL(ctx context.Context, TechBookURL string) error {
	t := &models.Techbook{
		Id:         0,
		CurrentUrl: TechBookURL,
	}
	_,_ = i.db.Delete(t)
	err := i.db.Insert(t)
	log.Printf("Set URL %s", TechBookURL)

	validUntil = time.Now()
	log.Printf("Cache invalidate. Valid until now = %s", validUntil)

	if err != nil {
		return err
	}
	return nil
}

func (i *TechBookRepositorySQLImpl) GetTechBookURL(ctx context.Context) (string, error) {
	t := models.Techbook{}
	_,err := i.db.Select(t, "select * from techbook")
	if err != nil {
		return "", err
	}
	return t.CurrentUrl, nil
}

func (i *TechBookRepositorySQLImpl) GetTechBook(ctx context.Context) (domain.TechBook, error) {
	if validUntil.After(time.Now()) && len(cached) > 0 {
		log.Printf("Return from cache. valid until: %s", validUntil)
		return cached, nil
	}

	log.Printf("Get Techbook from URL")

	dst, err := i.GetTechBookURL(ctx)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(dst)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	ret, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cached = ret
	validUntil = time.Now().Add(6 * time.Hour)
	log.Printf("Cache is valid until %s", validUntil)

	return ret, nil
}
