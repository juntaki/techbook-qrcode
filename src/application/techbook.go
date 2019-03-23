package application

import (
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/go-chi/chi"
	"github.com/juntaki/techbook-qrcode/src/domain"
)

type TechBookServer struct {
	router             http.Handler
	techBookRepository domain.TechBookRepository
	qrcodeRepository   domain.QRCodeRepository
}

func NewTechBookServer(
	techBookRepository domain.TechBookRepository,
	qrcodeRepository domain.QRCodeRepository,
) *TechBookServer {
	router := chi.NewRouter()

	s := &TechBookServer{
		techBookRepository: techBookRepository,
		qrcodeRepository:   qrcodeRepository,
	}
	router.Get("/{code}", s.GetTechbook)
	s.router = router
	return s
}

func (s *TechBookServer) GetTechbook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := chi.URLParam(r, "code")

	_, err := uuid.Parse(code)
	if err != nil {
		log.Printf("Not valid code %s %s", code, err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	isValid := s.qrcodeRepository.IsExistQRCode(ctx, &domain.QRCode{ID: code})
	if !isValid {
		log.Printf("Not found code %s", code)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	pdf, err := s.techBookRepository.GetTechBook(ctx)
	if err != nil {
		log.Printf("Cannot get techbook %s %s", err, code)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=600")
	w.Header().Set("Content-Disposition", "attachment; filename=m3techbook.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Length", strconv.Itoa(len(pdf)))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(pdf)
	if err != nil {
		log.Printf("Cannot write techbook %s %s", err, code)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Download success %s", code)
}

func (s *TechBookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
