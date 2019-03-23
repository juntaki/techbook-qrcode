package application

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/juntaki/fix"
	"github.com/juntaki/techbook-qrcode/src/infra"
	"github.com/juntaki/techbook-qrcode/src/lib/qrcode"
	"google.golang.org/appengine/aetest"
)

var qrCodeServiceServer *QRCodeServiceServer
var ctx context.Context

func TestMain(m *testing.M) {
	var err error
	ctx, _, err = aetest.NewContext()
	if err != nil {
		panic(err)
	}
	db, err := datastore.NewClient(ctx, "test")
	if err != nil {
		panic(err)
	}

	qrCodeRepository := infra.NewQRCodeRepositoryDatastoreImpl(db)
	techBookRepository := infra.NewTechBookRepositoryDatastoreImpl(db)
	qrCodeServiceServer = NewQRCodeServiceServer(qrCodeRepository, techBookRepository)

	code := m.Run()
	os.Exit(code)
}

func TestQRCodeServiceServer_GetURL(t *testing.T) {
	ret2, err := qrCodeServiceServer.GetQRCodes(ctx, &qrcode.Empty{})
	if err != nil {
		t.Fatal(err)
	}

	err = fix.Fix(ret2)
	if err != nil {
		t.Fatal(err)
	}
}
