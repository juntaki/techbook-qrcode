//+build wireinject

package main

import (
	"cloud.google.com/go/datastore"
	"github.com/google/wire"
	"github.com/juntaki/techbook-qrcode/src/application"
	"github.com/juntaki/techbook-qrcode/src/infra"
	gorp "gopkg.in/gorp.v2"
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

func InitializeQRCodeServiceServerSQL(
	db *gorp.DbMap,
) *application.QRCodeServiceServer {
	wire.Build(
		application.NewQRCodeServiceServer,
		infra.NewQRCodeRepositorySQLImpl,
		infra.NewTechBookRepositorySQLImpl,
	)
	return &application.QRCodeServiceServer{}
}

func InitializeTechBookServerSQL(
	db *gorp.DbMap,
) *application.TechBookServer {
	wire.Build(
		application.NewTechBookServer,
		infra.NewQRCodeRepositorySQLImpl,
		infra.NewTechBookRepositorySQLImpl,
	)
	return &application.TechBookServer{}
}
