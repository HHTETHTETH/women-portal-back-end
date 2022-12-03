package routers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	m "back/models"
	u "back/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/xid"
	"gopkg.in/mgo.v2/bson"
)

func heCheck(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	//*** Forward All Headers to http://khintkabar.com/api/server-he
	origin := context.Request.Header
	//http://khintkabar.com/api/server-he
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://khintkabar.com/api/server-he", nil)
	req.Header = origin
	fmt.Println(" Request header ", req.Header)
	response, err := client.Do(req)
	//response, err := http.Get("http://khintkabar.com/api/server-he")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusTemporaryRedirect,
			"message": "Forward to Login Page - The HTTP request failed with error",
			"page":    "login",
		})
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		msg := string(data)
		fmt.Println(" Message ", msg)

		if strings.Contains(msg, "This is not MPT Network") {
			//fmt.Println(" MSISDN Value is mssing, Forward to Login Page")
			//Forward Login Page, if MSISDN is missing
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusTemporaryRedirect,
				"message": "Forward to Login Page - This is not MPT Network",
				"page":    "login",
				//"msg":     msg,
				//"origin":  origin,
				//"rheader": req.Header,
			})

		} else {
			msisdn := ""
			i := strings.Index(msg, "Decrypted MSISDN")
			if i > -1 {
				//chars := msg[:i]
				msisdn = msg[i+18:]
				msisdn = msisdn[:len(msisdn)-1]
				fmt.Println(msisdn)
			} else {
				fmt.Println("Index not found")
			}
			phone := u.PhoneNumber(msisdn)
			//phone := "1028484"
			user, err := u.FindUserByPhone(phone)
			if err == nil {
				//fmt.Println(" user is already subscribed - ", user.Status)
				if user.OperationID != "NO" && user.OperationID != "ACI" && user.OperationID != "RD" && user.OperationID != "PD" {
					//Generate TOken
					token, _ := u.GenerateToken(phone)
					context.JSON(http.StatusOK, gin.H{
						"status":  http.StatusOK,
						"message": "Forward to Home Page",
						"page":    "/",
						"token":   token,
					})
					return

				} else {
					context.JSON(http.StatusOK, gin.H{
						"status":  http.StatusTemporaryRedirect,
						"message": "Forward to Landing Page",
						"page":    "/landing",
					})
					return
				}

			} else {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusTemporaryRedirect,
					"message": "Forward to Landing Page",
					"page":    "/landing",
				})
				return
			}

		}
	}
}

func maSubscribe(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	//Get HE
	response, err := http.Get("http://khintkabar.com/api/server-he")
	msg := ""
	phone := ""
	if err == nil {
		data, _ := ioutil.ReadAll(response.Body)
		msg = string(data)
		if strings.Contains(msg, "This is not MPT Network") {
		} else {
			msisdn := ""
			i := strings.Index(msg, "Decrypted MSISDN")
			if i > -1 {
				msisdn = msg[i+18:]
				msisdn = msisdn[:len(msisdn)-1]
			} else {
				fmt.Println("Index not found")
			}
			phone = u.PhoneNumber(msisdn)
		}
	}
	if phone == "" {
		phone = "898225533"
	}
	//API
	time.Sleep(3 * time.Second)
	guid := xid.New()
	transID := u.Int32ToString(guid.Counter())
	url := "http://macnt.mpt.com.mm/API/CGRequest?transID=" + transID + "&MSISDN=959" + phone + "&CpId=CSS&productID=9310&pName=Khint%20Kabar&pPrice=99&pVal=1&CpPwd=css@123&CpName=CSS&reqMode=WAP&reqType=Subscription&ismID=17&sRenewalPrice=99&sRenewalValidity=1&serviceType=T_KK_WEB_SUB_D&planId=T_KK_WEB_SUB_D_99&request_locale=mm&Wap_mdata=http://13.251.20.26/resources/loveCharacterImages/khint_logo.png"
	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusTemporaryRedirect,
		"message": "",
		"page":    url,
	})
}

func maUnsubscribe(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	//http://macnt.mpt.com.mm/API/CGUnsubscribe?CpId=CSS&MSISDN=959898225533&productID=9310&pName=Khint%20Kabar&CpPwd=css@123&CpName=CSS&reqMode=PIN&reqType=SUBSCRIPTION&transID=1000&request_locale=en&serviceType=T_KK_WEB_SUB_D&planId=T_KK_WEB_SUB_D_99&opId=101
	var requestData m.Token
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
		return
	}
	//fmt.Println(" ---- ", requestData.Token)
	token, _ := jwt.Parse(requestData.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tmp m.User
		mapstructure.Decode(claims, &tmp)
		fmt.Println(" ---- ", tmp)
		if tmp.Phone != "" {
			//DB User Status
			user, err := u.FindUserByPhone(tmp.Phone)
			if err == nil {
				guid := xid.New()
				transID := u.Int32ToString(guid.Counter())
				url := "http://macnt.mpt.com.mm/API/CGUnsubscribe?CpId=CSS&MSISDN=959" + user.Phone + "&productID=9310&pName=Khint%20Kabar&CpPwd=css@123&CpName=CSS&reqMode=PIN&reqType=SUBSCRIPTION&transID=" + transID + "&request_locale=my&serviceType=T_KK_WEB_SUB_D&planId=T_KK_WEB_SUB_D_99&opId=101"
				//fmt.Println(" ----- ", url)
				resp, err := http.Get(url)
				if err != nil {
					context.JSON(http.StatusOK, gin.H{
						"status":  http.StatusBadRequest,
						"message": "Unsuscribe ျပဳုလုပ္ျပင္း မေအာင္ျမင္ပါ",
					})
					return
				} else {
					cg := m.CgResponse{}
					data, _ := ioutil.ReadAll(resp.Body)
					xml.Unmarshal(data, &cg)
					if cg.ErrorCode == "0" && cg.ErrorDesc == "UNSUBSCRIPTION_REQ_SUBMITTED" {
						err = u.DeleteUser(user)
						if err == nil {
							context.JSON(http.StatusOK, gin.H{
								"status": http.StatusOK,
								"page":   "/landing",
							})
							return
						}
					} else {
						context.JSON(http.StatusOK, gin.H{
							"status":  http.StatusBadRequest,
							"message": "Unsuscribe ျပဳုလုပ္ျပင္း မေအာင္ျမင္ပါ",
						})
						return
					}
				}

			} else {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusBadRequest,
					"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
					"page":    "/landing",
				})
				//return
			}

		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
			"page":    "/",
		})
		return
	}
}

func maLogin(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.User
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
		return
	}
	//fmt.Println(" Request Data ", requestData.Phone)
	number := u.PhoneNumber(requestData.Phone)

	//Check in Database
	user, err := u.FindUserByPhone(number)
	if err != nil {
		if err.Error() == "not found" {
			//New User
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusTemporaryRedirect,
				"message": "ဤဖုန္းနံပါတ္မွာ ၀န္ေဆာင္မူးအားရယူထားျခင္းမရွိပါ",
				"page":    "landing",
			})
			return
		}
	} else {
		if user.OperationID != "NO" && user.OperationID != "ACI" && user.OperationID != "PCI" && user.OperationID != "RD" && user.OperationID != "PD" {
			//Generate OTP API
			guid := xid.New()
			transID := u.Int32ToString(guid.Counter())
			fmt.Println(" Transaction ID ", transID)
			url := "http://macnt.mpt.com.mm/API/GetOtp?mobile=959" + number + "&regUser=REGIS_MPT&regPassword=UkVHSVNfT1RQQDU0MzI=&otpMsgLang=2&serviceName=Khint%20Kabar&serviceDesc=Khint%20Kabar&CLI=8440&transId=" + transID + "&cpId=CSS&cpPassWord=css@123&email=&requestChannel=PIN"
			_, err := http.Get(url)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "OTP code အားေပးပို႔ျခင္း မေအာင္ျမင္ပါ",
				})
				return
			}
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "သင့္ဖုန္းနံပါတ္သို OTP code ေပးပို႔ၿပီးပါၿပီ။",
				"transID": transID,
			})

		} else {
			//Expired User - You have not subscribed this service
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusTemporaryRedirect,
				"message": "ဤဖုန္းနံပါတ္မွာ ၀န္ေဆာင္မူးအားရယူထားျခင္းမရွိပါ",
				"page":    "landing",
			})
			return
		}
	}
}

func maOTPCheck(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.OTPMessage
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
		return
	}
	fmt.Println("Requested OTP Check", requestData)
	number := u.PhoneNumber(requestData.Phone)
	otp := requestData.Code
	tranID := requestData.TransID
	url := "http://macnt.mpt.com.mm/API/VerifyOtp?mobile=959" + number + "&regUser=REGIS_MPT&regPassword=UkVHSVNfT1RQQDU0MzI=&otpMsgLang=2&otp=" + otp + "&otpEmail=&transId=" + tranID + "&cpId=CSS&cpPassWord=css@123&requestChannel=PIN"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(" Error ", err)
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
	} else {
		cg := m.CgResponse{}
		data, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(data, &cg)
		if cg.ErrorCode == "0" && cg.ErrorDesc == "Verified" {
			// OTP Success
			//	Generate Token
			token, _ := u.GenerateToken(number)
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "",
				"token":   token,
				"page":    "/",
			})
			return
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusBadRequest,
				"message": "OTP Code မွားယြင္းေနပါသည္",
			})
			return
		}
	}

}

func maStatus(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Token
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
		return
	}
	//fmt.Println("", requestData.Token)
	token, _ := jwt.Parse(requestData.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tmp m.User
		mapstructure.Decode(claims, &tmp)
		if tmp.Phone != "" {
			//DB User Status
			user, err := u.FindUserByPhone(tmp.Phone)

			if err == nil {
				if user.OperationID == "SP" {
					//Every Request, when HTTP.StatusNonAUthrita , show dialog
					context.JSON(http.StatusOK, gin.H{
						"status":  http.StatusNonAuthoritativeInfo,
						"message": "Forward to Home Page But dont allow Content",
						"user":    user,
					})
				} else if user.OperationID == "YS" {
					//Every Request, when HTTP.StatusNonAUthrita , show dialog
					context.JSON(http.StatusOK, gin.H{
						"status":  http.StatusAlreadyReported,
						"message": "Forward to Home Page But dont allow Content  and click Check",
						"user":    user,
					})
				} else if user.OperationID != "NO" && user.OperationID != "ACI" && user.OperationID != "SAC" && user.OperationID != "RD" && user.OperationID != "PD" {
					context.JSON(http.StatusOK, gin.H{
						"status": http.StatusOK,
						"user":   user,
					})
				} else {
					context.JSON(http.StatusOK, gin.H{
						"status": http.StatusBadRequest,
					})
				}

			} else {
				context.JSON(http.StatusOK, gin.H{
					"status": http.StatusInternalServerError,
					"user":   tmp,
				})
				return
			}
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
	}
}

func maStatusByNumber(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.User
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
		return
	}
	//fmt.Println("- Subscribe number -", requestData.Msisdn)
	if requestData.Msisdn != "" {
		//DB User Status
		phone := u.PhoneNumber(requestData.Msisdn)
		user, err := u.FindUserByPhone(phone)
		if err == nil {
			token, _ := u.GenerateToken(phone)
			if user.OperationID == "SP" {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "Forward to Home Page But dont allow Content",
					"page":    "/",
					"token":   token,
				})
			} else if user.OperationID == "YS" {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "Forward to Home Page But dont allow Content  and click Check",
					"page":    "/",
					"token":   token,
				})
			} else if user.OperationID != "NO" && user.OperationID != "ACI" && user.OperationID != "SAC" && user.OperationID != "RD" && user.OperationID != "PD" {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "Forward to Home Page",
					"page":    "/",
					"token":   token,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusUnauthorized,
					"message": "Forward to Landing ( User Found ) " + user.OperationID,
					"page":    "landing",
				})
			}
		} else {
			context.JSON(http.StatusOK, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Forward to Landing (User Not Found)",
				"page":    "landing",
			})

		}
	} else {
		fmt.Println("Don't Found User")
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
			"user":   "",
		})
	}
}

func maProfile(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Token
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
		})
		return
	}
	token, _ := jwt.Parse(requestData.Token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tmp m.User
		mapstructure.Decode(claims, &tmp)
		if tmp.Phone != "" {
			user, err := u.FindUserByPhone(tmp.Phone)
			if err == nil {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusOK,
					"message": "လုပ္ေဆာင္ခ်က္ေအာင္ျမင္ပါသည္",
					"user":    user,
				})
			} else {
				context.JSON(http.StatusOK, gin.H{
					"status":  http.StatusInternalServerError,
					"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
					"user":    "",
				})
			}
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ္ေဆာင္ခ်က္ မေအာင္ျမင္ပါ",
			"user":    "",
		})
	}
}

/*
func handleCallback(context *gin.Context) {
	cb := m.CallBack{}
	msisdn, ok := context.Request.URL.Query()["MSISDN"]
	if ok {
		fmt.Println(" MSISDN ", msisdn)
		cb.Msisdn = msisdn[0][1:]
	}
	result, ok := context.Request.URL.Query()["Result"]
	if ok {
		fmt.Println(" Result ", result)
		cb.Result = result[0]
	}
	reason, ok := context.Request.URL.Query()["Reason"]
	if ok {
		//	fmt.Println(" Reason ", reason)
		cb.Reason = reason[0]
	}
	productld, ok := context.Request.URL.Query()["productId"]
	if ok {
		//	fmt.Println(" productld ", productld)
		cb.ProductID = productld[0]
	}
	transID, ok := context.Request.URL.Query()["transID"]
	if ok {
		//	fmt.Println(" transID ", transID)
		cb.TransID = transID[0]
	}
	TPCGID, ok := context.Request.URL.Query()["TPCGID"]
	if ok {
		//	fmt.Println(" TPCGID ", TPCGID)
		cb.TPCGID = TPCGID[0]
	}
	songname, ok := context.Request.URL.Query()["songname"]
	if ok {
		//fmt.Println(" songname ", songname)
		cb.Songname = songname[0]
	}
	user, err := u.FindUserByPhone(cb.Msisdn)
	if err == nil {
		fmt.Println(" Found Existing User ", user)
		if cb.Result == "SUCCESS" {
			user.Status = "active"
			user.Reason = cb.Reason
		} else {
			user.Status = "pending"
		}
		err = u.UpdateUser(user)
		fmt.Println(" After Update ", user)
	} else {
		fmt.Println(" Create New User")
		nuser := m.User{}
		nuser.ID = bson.NewObjectId()
		nuser.Phone = cb.Msisdn
		if cb.Result == "SUCCESS" {
			nuser.Status = "active"
			nuser.Reason = cb.Reason
		} else {
			nuser.Status = "pending"
		}
		err = u.InsertUser(nuser)
	}
	err = u.InsertCBTran(cb)
}*/

//Private Call ( VPN )
func handleCallbackAPI2(context *gin.Context) {

	cb := m.CallBack{}
	msisdn, ok := context.Request.URL.Query()["callingParty"]
	if ok {
		cb.CallingParty = u.PhoneNumber(msisdn[0])
	}

	serviceID, ok := context.Request.URL.Query()["serviceId"]
	if ok {
		cb.ServiceID = serviceID[0]
	}

	serviceType, ok := context.Request.URL.Query()["serviceType"]
	if ok {
		cb.ServiceType = serviceType[0]
	}

	requestedPlan, ok := context.Request.URL.Query()["requestedPlan"]
	if ok {
		cb.RequestedPaln = requestedPlan[0]
	}

	appliededPlan, ok := context.Request.URL.Query()["appliededPlan"]
	if ok {
		cb.AppliededPlan = appliededPlan[0]
	}

	renewalPlan, ok := context.Request.URL.Query()["renewalPlan"]
	if ok {
		cb.RenewalPlan = renewalPlan[0]
	}

	operationID, ok := context.Request.URL.Query()["operationId"]
	if ok {
		cb.OperationID = operationID[0]
	}

	result, ok := context.Request.URL.Query()["result"]
	if ok {
		cb.Result = result[0]
	}

	startTime, ok := context.Request.URL.Query()["startTime"]
	if ok {
		cb.StartTime = startTime[0]
	}

	sequenceNo, ok := context.Request.URL.Query()["sequenceNo"]
	if ok {
		cb.SequenceNo = sequenceNo[0]
	}

	user, err := u.FindUserByPhone(cb.CallingParty)
	if err == nil {
		fmt.Println(" Found Existing User ", user)
		if cb.Result == "Success" {
			user.Status = "active"
		} else {
			user.Status = "pending"
		}
		user.OperationID = cb.OperationID
		err = u.UpdateUser(user)
	} else {
		fmt.Println(" Create New User")
		nuser := m.User{}
		nuser.ID = bson.NewObjectId()
		nuser.Phone = cb.CallingParty
		if cb.Result == "Success" {
			nuser.Status = "active"
		} else {
			nuser.Status = "pending"
		}
		nuser.OperationID = cb.OperationID
		err = u.InsertUser(nuser)
	}
	err = u.InsertCBTran(cb)
}

func callbacks(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	cbs, err := u.FindAllCBTrans()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"message": "လုပ်ဆောင်ချက်မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   cbs,
	})
}

func welcome(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	q := context.Request.URL.Query()
	subPhone := q["MSISDN"][0]

	//welcome?MSISDN=9898225533&Result=SUCCESS&Reason=Success_and_accepted_by_user&productId=9310&transID=111000100&TPCGID=190603161549023823&Songname=null
	//He Check here
	//MPT Number - eg 09898225533
	//Subscribed Number - eg 423001231
	//Check Subscribe Number and Mobile Data
	//If equal , ok
	//if not equal, adding paramter in URL

	origin := context.Request.Header
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://khintkabar.com/api/server-he", nil)
	req.Header = origin
	//fmt.Println(" Request header ", req.Header)
	response, err := client.Do(req)
	if err != nil {
		//fmt.Printf("The HTTP request failed with error %s\n", err)
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusTemporaryRedirect,
			"message": "Forward to Login Page - The HTTP request failed with error",
			"page":    "login",
		})
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		msg := string(data)
		//	fmt.Println(" Message ", msg)
		if strings.Contains(msg, "This is not MPT Network") {
			time.Sleep(3 * time.Second)
			context.Redirect(http.StatusTemporaryRedirect, "http://khintkabar.com?MSISDN=95"+subPhone)
		} else {
			msisdn := ""
			i := strings.Index(msg, "Decrypted MSISDN")
			if i > -1 {
				//chars := msg[:i]
				msisdn = msg[i+18:]
				msisdn = msisdn[:len(msisdn)-1]
				fmt.Println(msisdn)
			} else {
				fmt.Println("Index not found")
			}
			phone := u.PhoneNumber(msisdn)
			phone = "9" + phone
			//subPhone, ok := context.URL.Query()["key"]
			if phone == subPhone {
				//MSISDN
				time.Sleep(3 * time.Second)
				context.Redirect(http.StatusTemporaryRedirect, "http://khintkabar.com")
			} else {
				//Append Parameter
				time.Sleep(3 * time.Second)
				context.Redirect(http.StatusTemporaryRedirect, "http://khintkabar.com?MSISDN=95"+subPhone)
			}
		}
	}

}

func resourceBlocker(context *gin.Context) {
	context.Redirect(http.StatusTemporaryRedirect, "http://khintkabar.com")
}

func maForceMARedirect(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	url := ""
	serviceKK := "Khint Kabar"
	serviceAK := "Achit Kabar"
	var requestData m.ClickMap

	// exrtract vendor and domain
	params := context.Request.URL.Query()
	vendor := params.Get("v")
	domain := params.Get("d")
	click := params.Get("c")
	// generate transactionID
	guid := xid.New()
	transID := vendor + "_" + u.Int32ToString(guid.Counter())
	// generate correct ma link
	urlAK := "http://macnt.mpt.com.mm/API/CGRequest?transID=" + transID + "&MSISDN=959898225533&CpId=CSS&productID=9300&pName=Achit%20Kabar&pPrice=99&pVal=1&CpPwd=css@123&CpName=CSS&reqMode=WAP&reqType=Subscription&ismID=17&sRenewalPrice=99&sRenewalValidity=1&Wap_mdata=http://13.251.20.26/resources/loveCharacterImages/achitkabar_logo.png&serviceType=T_LAR_WAP_SUB_D&planId=T_LAR_WAP_SUB_D_99&request_locale=my"
	urlKK := "http://macnt.mpt.com.mm/API/CGRequest?transID=" + transID + "&MSISDN=959898225533&CpId=CSS&productID=9310&pName=Khint%20Kabar&pPrice=99&pVal=1&CpPwd=css@123&CpName=CSS&reqMode=WAP&reqType=Subscription&ismID=17&sRenewalPrice=99&sRenewalValidity=1&serviceType=T_KK_WEB_SUB_D&planId=T_KK_WEB_SUB_D_99&request_locale=mm&Wap_mdata=http://13.251.20.26/resources/loveCharacterImages/khint_logo.png"

	if domain == "AK" {
		url = urlAK
		requestData.ServiceID = serviceAK
		context.Set("Reffer", "http://achitkabar.com/")
	} else {
		url = urlKK
		requestData.ServiceID = serviceKK
		context.Set("Reffer", "http://Khintkabar.com/")
	}
	// save click on clickMap
	requestData.Click = click
	requestData.SequenceNo = transID
	requestData.Vendor = vendor
	err := u.InsertClickMap(requestData)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}
	// send it back
	/*
		context.JSON(http.StatusOK, gin.H{
			"status":  http.StatusTemporaryRedirect,
			"message": "",
			"page":    url,
		})
	*/
	context.Redirect(http.StatusTemporaryRedirect, url)
}
