package models

import "gopkg.in/mgo.v2/bson"

// User -
type User struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Msisdn      string        `json:"msisdn"`
	Phone       string        `json:"phone"`
	Status      string        `json:"status"` // unsubscribed, active, pending
	OperationID string        `json:"operationId"`
}

// OTPMessage -
type OTPMessage struct {
	Phone   string `json:"phone"`
	Code    string `json:"code"`
	TransID string `json:"transID"`
}

// CgResponse -
type CgResponse struct {
	ErrorCode string `xml:"error_code"`
	ErrorDesc string `xml:"errorDesc"`
	CgID      string `xml:"cgId"`
	TransID   string `xml:"transId"`
}

// Token -
type Token struct {
	Token string `json:"token"`
	//Phone string `json:"phone"`
}
