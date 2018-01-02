package controllers

import (
	"encoding/json"
	"fmt"
	"bytes"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type ScannerController struct {
	beego.Controller
}

type QueryCondition struct {
	DeviceId  string `json:"device_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type OneAlgae struct {
	AlgaeName    int32   `json:"algae_name"`
	AlgaeCount   int32   `json:"algae_count"`
	AlgaeDensity float64 `json:"algae_density"`
}

type Result struct {
	DeviceId              string     `json:"device_id"`
	UserId                int32      `json:"user_id"`
	UserName              string     `json:"user_name"`
	ExperimentName        string     `json:"experiment_name"`
	Algaes                []OneAlgae `json:"algaes"`
	TotalAlgaeDensity     float64    `json:"total_algae_density"`
	TotalAlgaeCount       int32      `json:"total_algae_count"`
	AdvantageAlgaeName    int32      `json:"advantage_algae_name"`
	AdvantageAlgaeDensity float64    `json:"advantage_algae_density"`
	AdvantageAlgaePercent float32    `json:"advantage_algae_percent"`
	ScannerSampleVolume   float32    `json:"scanner_sample_volume"` //取样容量
	SampleVolume          float32    `json:"sample_volume"`         //玻片容量
	TotalVolume           float32    `json:"total_volume"`          //总容量
	DilutionMultiple      int32      `json:"dilution_multiple"`
	ViewCount             int32      `json:"view_count"`
	SamplePlace           string     `json:"sample_place"`
	SampleDate            string     `json:"sample_date"`
	ExperimentDateTime    string     `json:"experiment_datetime"`
	OneViewArea           float32    `json:"one_view_area"`
	SlideArea             float32    `json:"slide_area"`
}

type ResultJson struct {
	Result []Result `json:"scanner_results"`
}

const (
	//pageNums    = 20
	perPageNums = 20
)

var scannerCollection *mgo.Collection = nil
var queryCollectionCondition = QueryCondition{}

func init() {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		return
	}
	session.SetMode(mgo.Strong, true)
	scannerCollection = session.DB("ecology").C("scanner_result")
}

func (c *ScannerController) Get() {
	if nil == scannerCollection {
		fmt.Printf("error,scanner controller is nil\n")
		c.Abort("500")
	}
	c.TplName = "scanner/result.tpl"
	//c.Data["perPageNums"] = perPageNums
	//c.TplName = "scanner/result.tpl"
	//page := c.GetString("page")
	//if "" == page {
	//	json.Unmarshal(c.Ctx.Input.RequestBody, &queryCollectionCondition)
	//	c.Data["currentPageNum"] = 1
	//} else {
	//	currentPageNum, err := strconv.Atoi(page)
	//	if err != nil {
	//		fmt.Printf("error,currntPageNum is wrong value\n")
	//		c.Abort("412")
	//	}
	//	c.Data["currentPageNum"] = currentPageNum
	//}
	// query := mgo.Query{}
	//if queryCollectionCondition.deviceId == "" {
	//	query = *scannerCollection.Find(bson.M{/**"sampledate":bson.M{"$gte":queryCollectionCondition.startDate,
	//	"$lte":queryCollectionCondition.endDate}**/})
	//} else {
	//	query = *scannerCollection.Find(bson.M{"deviceid":queryCollectionCondition.deviceId,
	//		"sampledate":bson.M{"$gte":queryCollectionCondition.startDate,
	//		"$lte":queryCollectionCondition.endDate}})
	//}
	//count,err := query.Count()
	//if err != nil {
	//	fmt.Printf("error,query count failed,err is %v\n",err)
	//	c.Abort("412")
	//}
	//fmt.Printf("result count is %v\n",count)
	//c.Data["pageNums"] = count
	//var result []Result
	//if err := query.All(&result); err != nil {
	//	fmt.Printf("error,query all result failed,err is %v\n",err)
	//	c.Abort("412")
	//}
	//if "" == page {
	//	if count <= perPageNums {
	//		c.Data["currentResult"] = result
	//	} else {
	//		c.Data["currentResult"] = result[:perPageNums]
	//	}
	//}else {
	//	currentPageNum, err := strconv.Atoi(page)
	//	if err != nil {
	//		fmt.Printf("error,currntPageNum is wrong value\n")
	//		c.Abort("412")
	//	}
	//	if count <= (currentPageNum - 1) * perPageNums {
	//		fmt.Printf("error:count less than need \n")
	//		c.Abort("412")
	//	}
	//	if count > currentPageNum * perPageNums {
	//		c.Data["currentResult"] = result[(currentPageNum -1 ) * perPageNums:currentPageNum * perPageNums]
	//	} else {
	//		c.Data["currentResult"] = result[(currentPageNum -1 ) * perPageNums:]
	//	}
	//}
}

func (c *ScannerController) Post() {
	if nil == scannerCollection {
		fmt.Printf("error,scanner controller is nil\n")
		c.Abort("500")
	}
	//c.Data["pageNums"] = pageNums
	c.Data["perPageNums"] = perPageNums
	c.TplName = "scanner/result.tpl"
	page := c.GetString("page")
	if "" == page {
		//fmt.Printf("initialize query collection condition\n")
		//fmt.Printf("request body is :%v\n", c.Ctx.Input.RequestBody)
		json.Unmarshal(c.Ctx.Input.RequestBody, &queryCollectionCondition)
		c.Data["currentPageNum"] = 1
	} else {
		currentPageNum, err := strconv.Atoi(page)
		if err != nil {
			fmt.Printf("error,currntPageNum is wrong value\n")
			c.Abort("412")
		}
		c.Data["currentPageNum"] = currentPageNum
	}
	//fmt.Printf("device id is %v\n", queryCollectionCondition.DeviceId)
	//fmt.Printf("start time is %v\n", queryCollectionCondition.StartDate)
	//fmt.Printf("end time is %v\n", queryCollectionCondition.EndDate)
	query := mgo.Query{}
	if queryCollectionCondition.DeviceId == "" {
		query = *scannerCollection.Find(bson.M{ /**"sampledate":bson.M{"$gte":queryCollectionCondition.startDate,
		"$lte":queryCollectionCondition.endDate}**/})
	} else {
		query = *scannerCollection.Find(bson.M{"deviceid": queryCollectionCondition.DeviceId,
			"sampledate": bson.M{"$gte": queryCollectionCondition.StartDate,
				"$lte": queryCollectionCondition.EndDate}})
	}
	count, err := query.Count()
	if err != nil {
		fmt.Printf("error,query count failed,err is %v\n", err)
		c.Abort("412")
	}
	fmt.Printf("result count is %v\n", count)
	c.Data["pageNums"] = count
	var result []Result
	if err := query.All(&result); err != nil {
		fmt.Printf("error,query all result failed,err is %v\n", err)
		c.Abort("412")
	}
	resultJson := ResultJson{result}
	fmt.Printf("result json is %v\n", resultJson)
	resultByte, err := json.Marshal(resultJson)
	if err != nil {
		fmt.Printf("error,result json convert to byte failed,err is %v\n", err)
		c.Abort("412")
	}
	c.Data["json"] = bytes.NewBuffer(resultByte).String()
	c.ServeJSON()
}
