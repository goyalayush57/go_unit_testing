package external

import (
	"fmt"
	"gophercon_2023/unittesting/externaldependency/mocks"
	"gophercon_2023/unittesting/externaldependency/model"
	"testing"

	"github.com/golang/mock/gomock"
)

//Linux : https://github.com/golang/mock
//Mac : https://ports.macports.org/port/go-mockgen/

//mockgen -destination=externaldependency/mocks/mock_cache.go -package=external "gophercon_2023/unittesting/externaldependency" CacheService

func Test_register_Register_Impl(t *testing.T) {
	type fields struct {
		db    DatastoreService
		cache CacheService
	}
	type args struct {
		name  string
		email string
		pass  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		//Test Fixture //Not Auto generated
		setup func() (*gomock.Controller, *mocks.MockCacheService, *mocks.MockDatastoreService)
	}{
		{
			name: "Username already exist in cache",
			args: args{
				name:  "alreadyTaken",
				email: "xyz@gmail.com",
				pass:  "NahiBatauga",
			},
			setup: func() (*gomock.Controller, *mocks.MockCacheService, *mocks.MockDatastoreService) {
				ctrl := gomock.NewController(t) //helps manage the lifecycle of a mock object
				dbMock := mocks.NewMockDatastoreService(ctrl)
				cacheMock := mocks.NewMockCacheService(ctrl)
				cacheMock.EXPECT().Get("alreadyTaken").Return("123")
				return ctrl, cacheMock, dbMock
			},
			wantErr: true,
		},
		{
			name: "Error in inserting data to db",
			args: args{
				name:  "validName",
				email: "xyz@gmail.com",
				pass:  "BolaNaNahiBatauga",
			},
			setup: func() (*gomock.Controller, *mocks.MockCacheService, *mocks.MockDatastoreService) {
				ctrl := gomock.NewController(t) //helps manage the lifecycle of a mock object
				dbMock := mocks.NewMockDatastoreService(ctrl)
				cacheMock := mocks.NewMockCacheService(ctrl)
				cacheMock.EXPECT().Get("validName").Return("")
				dbMock.EXPECT().Insert(gomock.Any()).Return(model.Row{}, fmt.Errorf("error in insertion"))
				return ctrl, cacheMock, dbMock
			},
			wantErr: true,
		},
		{
			name: "Register New User",
			args: args{
				name:  "first last",
				email: "xyz@gmail.com",
				pass:  "gazabKeZiddiHoYr",
			},
			setup: func() (*gomock.Controller, *mocks.MockCacheService, *mocks.MockDatastoreService) {
				ctrl := gomock.NewController(t) //helps manage the lifecycle of a mock object
				dbMock := mocks.NewMockDatastoreService(ctrl)
				cacheMock := mocks.NewMockCacheService(ctrl)
				cacheMock.EXPECT().Get("first last").Return("")
				dbMock.EXPECT().Insert(gomock.Any()).Return(model.Row{Name: "first last", Email: "xyz@gmail.com", ID: 3223}, nil)
				cacheMock.EXPECT().Put("first last", 3223)
				return ctrl, cacheMock, dbMock
			},
			wantErr: false, //exact error
		},
	}
	for _, tt := range tests {
		_, cacheMock, dbMock := tt.setup()
		//defer ctrl.Finish()
		t.Run(tt.name, func(t *testing.T) {
			a := &register{
				db:    dbMock,
				cache: cacheMock,
			}
			if err := a.Register(tt.args.name, tt.args.email, tt.args.pass); (err != nil) != tt.wantErr {
				t.Errorf("register.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
