package tests

import ()

const (
	testDataDir                 = "../test_pb_data"
	OnRecordsListRequest        = "OnRecordsListRequest"
	OnRecordViewRequest         = "OnRecordViewRequest"
	OnModelAfterUpdate          = "OnModelAfterUpdate"
	OnModelBeforeUpdate         = "OnModelBeforeUpdate"
	OnRecordAfterUpdateRequest  = "OnRecordAfterUpdateRequest"
	OnRecordBeforeUpdateRequest = "OnRecordBeforeUpdateRequest"
	OnModelAfterCreate          = "OnModelAfterCreate"
	OnModelBeforeCreate         = "OnModelBeforeCreate"
	OnRecordAfterCreateRequest  = "OnRecordAfterCreateRequest"
	OnRecordBeforeCreateRequest = "OnRecordBeforeCreateRequest"
	OnModelAfterDelete          = "OnModelAfterDelete"
	OnModelBeforeDelete         = "OnModelBeforeDelete"
	OnRecordAfterDeleteRequest  = "OnRecordAfterDeleteRequest"
	OnRecordBeforeDeleteRequest = "OnRecordBeforeDeleteRequest"
)

var (
	authHeaders AuthHeaders
	testData    TestData
)
