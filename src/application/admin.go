package application

import (
	"context"
	"fmt"

	"github.com/juntaki/techbook-qrcode/src/domain"
	"github.com/juntaki/techbook-qrcode/src/lib/qrcode"
)

func NewQRCodeServiceServer(
	qrcodeRepository domain.QRCodeRepository,
	techBookRepository domain.TechBookRepository,
) *QRCodeServiceServer {
	return &QRCodeServiceServer{
		qrcodeRepository:   qrcodeRepository,
		techBookRepository: techBookRepository,
	}
}

type QRCodeServiceServer struct {
	qrcodeRepository   domain.QRCodeRepository
	techBookRepository domain.TechBookRepository
}

func (i *QRCodeServiceServer) GetURL(ctx context.Context, input *qrcode.Empty) (*qrcode.URL, error) {
	url, err := i.techBookRepository.GetTechBookURL(ctx)
	if err != nil {
		return nil, err
	}
	return &qrcode.URL{
		Url: url,
	}, nil
}

func (i *QRCodeServiceServer) UpdateURL(ctx context.Context, input *qrcode.URL) (*qrcode.Empty, error) {
	err := i.techBookRepository.SetTechBookURL(ctx, input.Url)
	if err != nil {
		return nil, err
	}
	return &qrcode.Empty{}, nil
}

func (i *QRCodeServiceServer) GetQRCodes(
	ctx context.Context,
	input *qrcode.Empty,
) (*qrcode.QRCodeList, error) {
	list, err := i.qrcodeRepository.GetQRCodes(ctx)
	if err != nil {
		return nil, err
	}

	ret := &qrcode.QRCodeList{
		QRCodes: make([]*qrcode.QRCode, len(list)),
	}

	for i, l := range list {
		ret.QRCodes[i] = &qrcode.QRCode{
			Id:    fmt.Sprintf("%d:%s", l.Index, l.ID),
			Url:   l.GetURL(),
			Image: l.GetQRCode(),
		}
	}

	return ret, nil
}

func (i *QRCodeServiceServer) AddQRCodes(ctx context.Context, input *qrcode.Empty) (*qrcode.Empty, error) {
	qrs, err := i.qrcodeRepository.GetQRCodes(ctx)
	if err != nil {
		return nil, err
	}

	index := len(qrs)
	newQRs := []*domain.QRCode{}
	for i := 0; i < 50; i++ {
		newQRs = append(newQRs, domain.NewQRCode(index+i))
	}

	err = i.qrcodeRepository.SaveQRCodes(ctx, newQRs)
	if err != nil {
		return nil, err
	}
	return &qrcode.Empty{}, nil
}
