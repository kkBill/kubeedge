/*
Copyright 2020 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dtclient

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/golang/mock/gomock"

	"github.com/kubeedge/kubeedge/edge/mocks/beego"
	"github.com/kubeedge/kubeedge/edge/pkg/common/dbm"
)

// errFailedDBOperation is common DB operation fail error
//var errFailedDBOperation = errors.New("Failed DB Operation")

// deviceMeta is global variable for passing as test parameter
var deviceMeta = DeviceMeta{
	Key:   "TestKey",
	Value: "TestValue",
}

func TestSaveDeviceMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the test case
		name string
		// returnInt is first return of mock interface ormerMock
		returnInt int64
		// returnErr is second return of mock interface ormerMock which is also expected error
		returnErr error
	}{{
		// Success Case
		name:      "SuccessCase",
		returnInt: int64(1),
		returnErr: nil,
	}, {
		// Failure Case
		name:      "FailureCase",
		returnInt: int64(1),
		returnErr: errFailedDBOperation,
	},
	}
	
	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			ormerMock.EXPECT().Insert(gomock.Any()).Return(test.returnInt, test.returnErr).Times(1)
			err := SaveDeviceMeta(&deviceMeta)
			if test.returnErr != err {
				t.Errorf("Save DeviceMeta Case failed : wanted error %v and got error %v", test.returnErr, err)
			}
		})
	}
}

// TestDeleteDeviceMetaByKey is function to test DeleteDeviceMetaByKey
func TestDeleteDeviceMetaByKey(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// deleteReturnInt is the first return of mock interface querySeterMock's delete function
		deleteReturnInt int64
		// deleteReturnErr is the second return of mock interface querySeterMocks's delete function also expected error
		deleteReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		deleteReturnInt:  int64(1),
		deleteReturnErr:  nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		deleteReturnInt:  int64(0),
		deleteReturnErr:  errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			querySeterMock.EXPECT().Delete().Return(test.deleteReturnInt, test.deleteReturnErr).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			err := DeleteDeviceMetaByKey("test")
			if test.deleteReturnErr != err {
				t.Errorf("Delete Meta By Key Case failed : wanted %v and got %v", test.deleteReturnErr, err)
			}
		})
	}
}

// TestUpdateDeviceMeta is function to test UpdateDeviceMeta
func TestUpdateDeviceMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// returnInt is first return of mock interface ormerMock
		returnInt int64
		// returnErr is second return of mock interface ormerMock which is also expected error
		returnErr error
	}{{
		// Success Case
		name:      "SuccessCase",
		returnInt: int64(1),
		returnErr: nil,
	}, {
		// Failure Case
		name:      "FailureCase",
		returnInt: int64(0),
		returnErr: errFailedDBOperation,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			ormerMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(test.returnInt, test.returnErr).Times(1)
			err := UpdateDeviceMeta(&deviceMeta)
			if test.returnErr != err {
				t.Errorf("Update Meta Case failed : wanted %v and got %v", test.returnErr, err)
			}
		})
	}
}

// TestInsertOrUpdate is function to test InsertOrUpdate
func TestInsertOrUpdate(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	rawSeterMock := beego.NewMockRawSeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// returnSQL is first return of mock interface rawSeterMock's Exec function
		returnSQL sql.Result
		// returnErr is second return of mock interface rawSeterMock's Exec function which is also expected error
		returnErr error
		// returnRaw is the return of mock interface ormerMock's Raw function
		returnRaw orm.RawSeter
	}{{
		// Success Case
		name:      "SuccessCase",
		returnSQL: nil,
		returnErr: nil,
		returnRaw: rawSeterMock,
	}, {
		// Failure Case
		name:      "FailureCase",
		returnSQL: nil,
		returnErr: errFailedDBOperation,
		returnRaw: rawSeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			rawSeterMock.EXPECT().Exec().Return(test.returnSQL, test.returnErr).Times(1)
			ormerMock.EXPECT().Raw(gomock.Any(), gomock.Any()).Return(test.returnRaw).Times(1)
			err := InsertOrUpdate(&deviceMeta)
			if test.returnErr != err {
				t.Errorf("Insert or Update Meta Case failed : wanted %v and got %v", test.returnErr, err)
			}
		})
	}
}

// TestUpdateDeviceMetaField is function to test UpdateDeviceMetaField
func TestUpdateDeviceMetaField(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// updateReturnInt is the first return of mock interface querySeterMock's update function
		updateReturnInt int64
		// updateReturnErr is the second return of mock interface querySeterMocks's update function also expected error
		updateReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(1),
		updateReturnErr:  nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(0),
		updateReturnErr:  errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			querySeterMock.EXPECT().Update(gomock.Any()).Return(test.updateReturnInt, test.updateReturnErr).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			err := UpdateDeviceMetaField("test", "test", "test")
			if test.updateReturnErr != err {
				t.Errorf("Update Meta Field Case failed : wanted %v and got %v", test.updateReturnErr, err)
			}
		})
	}
}

// TestUpdateDeviceMetaFields is function to test UpdateDeviceMetaFields
func TestUpdateDeviceMetaFields(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// updateReturnInt is the first return of mock interface querySeterMock's update function
		updateReturnInt int64
		// updateReturnErr is the second return of mock interface querySeterMocks's update function also expected error
		updateReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(1),
		updateReturnErr:  nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		updateReturnInt:  int64(0),
		updateReturnErr:  errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			querySeterMock.EXPECT().Update(gomock.Any()).Return(test.updateReturnInt, test.updateReturnErr).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			err := UpdateDeviceMetaFields("test", nil)

			if test.updateReturnErr != err {
				t.Errorf("Update Meta Fields Case failed : wanted %v and got %v", test.updateReturnErr, err)
			}
		})
	}
}

// TestQueryDeviceMeta is function to test QueryDeviceMeta
func TestQueryDeviceMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// allReturnInt is the first return of mock interface querySeterMock's all function
		allReturnInt int64
		// allReturnErr is the second return of mock interface querySeterMocks's all function also expected error
		allReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{{
		// Success Case
		name:             "SuccessCase",
		filterReturn:     querySeterMock,
		allReturnInt:     int64(1),
		allReturnErr:     nil,
		queryTableReturn: querySeterMock,
	}, {
		// Failure Case
		name:             "FailureCase",
		filterReturn:     querySeterMock,
		allReturnInt:     int64(0),
		allReturnErr:     errFailedDBOperation,
		queryTableReturn: querySeterMock,
	},
	}

	// fakeDao is used to set the argument of All function
	fakeDao := new([]DeviceMeta)
	fakeDaoArray := make([]DeviceMeta, 1)
	fakeDaoArray[0] = DeviceMeta{Key: "Test"}
	fakeDao = &fakeDaoArray

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().All(gomock.Any()).SetArg(0, *fakeDao).Return(test.allReturnInt, test.allReturnErr).Times(1)
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			meta, err := QueryDeviceMeta("test", "test")
			if test.allReturnErr != err {
				t.Errorf("Query Meta Case Failed : wanted error %v and got error %v", test.allReturnErr, err)
				return
			}

			if err == nil {
				if len(*meta) != 1 {
					t.Errorf("Query Meta Case failed: wanted length 1 and got length %v", len(*meta))
				}
			}
		})
	}
}

// TestQueryAllDeviceMeta is function to test QueryAllDeviceMeta
func TestQueryAllDeviceMeta(t *testing.T) {
	//Initialize Global Variables (Mocks)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	ormerMock := beego.NewMockOrmer(mockCtrl)
	querySeterMock := beego.NewMockQuerySeter(mockCtrl)
	dbm.DBAccess = ormerMock

	cases := []struct {
		// name is name of the testcase
		name string
		// filterReturn is the return of mock interface querySeterMock's filter function
		filterReturn orm.QuerySeter
		// allReturnInt is the first return of mock interface querySeterMock's all function
		allReturnInt int64
		// allReturnErr is the second return of mock interface querySeterMocks's all function also expected error
		allReturnErr error
		// queryTableReturn is the return of mock interface ormerMock's QueryTable function
		queryTableReturn orm.QuerySeter
	}{
		{
			// Success Case
			name:             "SuccessCase",
			filterReturn:     querySeterMock,
			allReturnInt:     int64(1),
			allReturnErr:     nil,
			queryTableReturn: querySeterMock,
		},
		{
			// Failure Case
			name:             "FailureCase",
			filterReturn:     querySeterMock,
			allReturnInt:     int64(0),
			allReturnErr:     errFailedDBOperation,
			queryTableReturn: querySeterMock,
		},
	}

	// fakeDao is used to set the argument of All function
	fakeDao := new([]DeviceMeta)
	fakeDaoArray := make([]DeviceMeta, 1)
	fakeDaoArray[0] = DeviceMeta{Key: "Test", Value: "Test"}
	fakeDao = &fakeDaoArray

	// run the test cases
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			querySeterMock.EXPECT().All(gomock.Any()).SetArg(0, *fakeDao).Return(test.allReturnInt, test.allReturnErr).Times(1)
			querySeterMock.EXPECT().Filter(gomock.Any(), gomock.Any()).Return(test.filterReturn).Times(1)
			ormerMock.EXPECT().QueryTable(gomock.Any()).Return(test.queryTableReturn).Times(1)
			meta, err := QueryAllDeviceMeta("test", "test")
			if test.allReturnErr != err {
				t.Errorf("Query All Meta Case Failed : wanted error %v and got error %v", test.allReturnErr, err)
				return
			}

			if err == nil {
				if len(*meta) != 1 {
					t.Errorf("Query All Meta Case failed: wanted length 1 and got length %v", len(*meta))
				}
			}
		})
	}
}

// TestIsNonUniqueNameError is function to test IsNonUniqueNameError().
func TestIsNonUniqueNameError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantBool bool
	}{
		{
			name:     "Suffix-are not unique",
			err:      errors.New("The fields are not unique"),
			wantBool: true,
		},
		{
			name:     "Contains-UNIQUE constraint failed",
			err:      errors.New("Failed-UNIQUE constraint failed"),
			wantBool: true,
		},
		{
			name:     "Contains-constraint failed",
			err:      errors.New("The input constraint failed"),
			wantBool: true,
		},
		{
			name:     "OtherError",
			err:      errors.New("Failed"),
			wantBool: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotBool := IsNonUniqueNameError(test.err)
			if gotBool != test.wantBool {
				t.Errorf("IsNonUniqueError() failed, Got = %v, Want = %v", gotBool, test.wantBool)
			}
		})
	}
}
