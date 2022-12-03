package models

import "gopkg.in/mgo.v2/bson"

// CallBack -
/*type CallBack struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Msisdn    string        `json:"msisdn"`    //MSISDN
	Result    string        `json:"result"`    // Result  - SUCCESS
	Reason    string        `json:"reason"`    // Reason - Success_and_accepted_by_user
	ProductID string        `json:"productId"` // productld - 9310
	TransID   string        `json:"transID"`   // transID -
	TPCGID    string        `json:"TPCGID"`    // TPCGID
	Songname  string        `json:"songname"`  // Songname - null
}*/
type CallBack struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	CallingParty  string        `json:"callingParty"`  //callingParty - 959898225533
	ServiceID     string        `json:"serviceId"`     //serviceId - 9310
	ServiceType   string        `json:"serviceType:"`  //serviceType - T_KK_WEB_SUB_D
	RequestedPaln string        `json:"requestedPlan"` //requestedPlan - T_KK_WEB_SUB_D_99
	AppliededPlan string        `json:"appliededPlan"` //appliededPlan - T_KK_WEB_SUB_D_99
	RenewalPlan   string        `json:"renewalPlan"`   //renewalPlan - T_KK_WEB_SUB_D_99

	//ContentId - -1
	//category - -1
	//serviceNode- CSS
	//chargeAmount - 99.0
	//validityDays - 1
	OperationID string `json:"operationId"` //operationId - SN
	//bearerId - 106
	//resultCode - 0
	Result         string `json:"result"`    // result - Success
	StartTime      string `json:"startTime"` //startTime - Tue Apr 23 15:12:44 GMT+06:30 2019
	SequenceNo     string `json:"sequenceNo"`
	ProcessingTime string `json:"processingTime"`
	//sequenceNo - 967226
	//keyword - OFF
	//processingTime - 190423151242
}
