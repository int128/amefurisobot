package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/int128/amefuriso/domain"
	"github.com/int128/amefuriso/usecases/interfaces/mock_usecases"
)

func TestGetImage_ServeHTTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_usecases.NewMockGetImage(ctrl)
	usecase.EXPECT().
		Do(gomock.Not(nil), domain.ImageID("FOO")).
		Return(&domain.Image{ContentType: "image/png"}, nil)

	req := httptest.NewRequest("GET", "/images/FOO.png", nil)
	w := httptest.NewRecorder()
	h := Handlers{GetImage: GetImage{usecase}}
	h.NewRouter().ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Code wants 200 but %v", w.Code)
	}
	contentType := w.Header().Get("content-type")
	if contentType != "image/png" {
		t.Errorf("content-type wants image/png but %v", contentType)
	}
}

func TestGetImage_ServeHTTP_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usecase := mock_usecases.NewMockGetImage(ctrl)
	usecase.EXPECT().
		Do(gomock.Not(nil), domain.ImageID("FOO")).
		Return(nil, domain.ErrNoSuchImage{ID: domain.ImageID("FOO")})

	req := httptest.NewRequest("GET", "/images/FOO.png", nil)
	w := httptest.NewRecorder()
	h := Handlers{GetImage: GetImage{usecase}}
	h.NewRouter().ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Code wants 200 but %v", w.Code)
	}
}
