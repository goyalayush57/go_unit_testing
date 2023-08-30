package external

import (
	"fmt"
	"gophercon_2023/unit_testing/externaldependency/mocks"
	"gophercon_2023/unit_testing/externaldependency/model"
	"testing"

	"github.com/golang/mock/gomock"
)

//Linux : https://github.com/golang/mock
//Mac : https://ports.macports.org/port/go-mockgen/

//mockgen -destination=externaldependency/mocks/mock_cache.go -package=external "gophercon_2023/unit_testing/externaldependency" CacheService

func Test_register_Register_Impl(t *testing.T) {
	ctrl := gomock.NewController(t) //helps manage the lifecycle of a mock object
	dbMock := mocks.NewMockDatastoreService(ctrl)
	cacheMock := mocks.NewMockCacheService(ctrl)
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
		setup   func() //Test Fixture //Not Auto generated
	}{
		{
			name: "Username already exist",
			fields: fields{
				db:    dbMock,
				cache: cacheMock,
			},
			args: args{
				name:  "alreadyTaken",
				email: "xyz@gmail.com",
				pass:  "NahiBatauga",
			},
			setup: func() {
				cacheMock.EXPECT().Get("alreadyTaken").Return("123")
			},
			wantErr: true,
		},
		{
			name: "Error in inserting data to db",
			fields: fields{
				db:    dbMock,
				cache: cacheMock,
			},
			args: args{
				name:  "validName",
				email: "xyz@gmail.com",
				pass:  "BolaNaNahiBatauga",
			},
			setup: func() {
				cacheMock.EXPECT().Get("validName").Return("")
				dbMock.EXPECT().Insert(gomock.Any()).Return(model.Row{}, fmt.Errorf("error in insertion"))
			},
			wantErr: true,
		},
		{
			name: "Register New User",
			fields: fields{
				db:    dbMock,
				cache: cacheMock,
			},
			args: args{
				name:  "first last",
				email: "xyz@gmail.com",
				pass:  "gazabKeZiddiHoYr",
			},
			setup: func() {
				cacheMock.EXPECT().Get("first last").Return("")
				dbMock.EXPECT().Insert(gomock.Any()).Return(model.Row{Name: "first last", Email: "xyz@gmail.com", ID: 3223}, nil)
				cacheMock.EXPECT().Put("first last", 3223)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.setup()
		t.Run(tt.name, func(t *testing.T) {
			a := &register{
				db:    tt.fields.db,
				cache: tt.fields.cache,
			}
			if err := a.Register(tt.args.name, tt.args.email, tt.args.pass); (err != nil) != tt.wantErr {
				t.Errorf("register.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_register_Register(t *testing.T) {
	type fields struct {
		db    DatastoreService
		cache CacheService
	}
	type args struct {
		name        string
		email       string
		encryptPass string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		setup   func()
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt.setup()
		t.Run(tt.name, func(t *testing.T) {
			a := &register{
				db:    tt.fields.db,
				cache: tt.fields.cache,
			}
			if err := a.Register(tt.args.name, tt.args.email, tt.args.encryptPass); (err != nil) != tt.wantErr {
				t.Errorf("register.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
