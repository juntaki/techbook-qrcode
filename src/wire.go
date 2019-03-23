//+build wireinject

package main

import (
	"cloud.google.com/go/datastore"
	"github.com/google/wire"
	"github.com/juntaki/techbook-qrcode/src/application"
	"github.com/juntaki/techbook-qrcode/src/infra"
)

func InitializeQRCodeServiceServer(
	db *datastore.Client,
) *application.QRCodeServiceServer {
	wire.Build(
		application.NewQRCodeServiceServer,
		infra.NewQRCodeRepositoryDatastoreImpl,
		infra.NewTechBookRepositoryDatastoreImpl,
	)
	return &application.QRCodeServiceServer{}
}

func InitializeTechBookServer(
	db *datastore.Client,
) *application.TechBookServer {
	wire.Build(
		application.NewTechBookServer,
		infra.NewQRCodeRepositoryDatastoreImpl,
		infra.NewTechBookRepositoryDatastoreImpl,
	)
	return &application.TechBookServer{}
}
