package examples

import (
	"io/ioutil"
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"time" 
	"fmt"
	"log"
	"encoding/json"

)

// data structure is date, open, close, low, high
// dates are read as string, so data must be ordered sequentially by date

func klineBase(arrayTime []string, arrayData [][]float32) *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for i := 0; i < len(arrayData); i++ {
		x = append(x, arrayTime[i])
		y = append(y, opts.KlineData{Value: arrayData[i]})
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Kline-example",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: opts.Bool(true),
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	kline.SetXAxis(x).AddSeries("kline", y)
	return kline
}

// func klineDataZoomInside() *charts.Kline {
// 	kline := charts.NewKLine()

// 	x := make([]string, 0)
// 	y := make([]opts.KlineData, 0)
// 	for i := 0; i < len(kd); i++ {
// 		x = append(x, kd[i].date)
// 		y = append(y, opts.KlineData{Value: kd[i].data})
// 	}

// 	kline.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "DataZoom(inside)",
// 		}),
// 		charts.WithXAxisOpts(opts.XAxis{
// 			SplitNumber: 20,
// 		}),
// 		charts.WithYAxisOpts(opts.YAxis{
// 			Scale: opts.Bool(true),
// 		}),
// 		charts.WithDataZoomOpts(opts.DataZoom{
// 			Type:       "inside",
// 			Start:      50,
// 			End:        100,
// 			XAxisIndex: []int{0},
// 		}),
// 	)

// 	kline.SetXAxis(x).AddSeries("kline", y)
// 	return kline
// }
// data zoom inside and slider
// func klineDataZoomBoth() *charts.Kline {
// 	kline := charts.NewKLine()

// 	x := make([]string, 0)
// 	y := make([]opts.KlineData, 0)
// 	for i := 0; i < len(kd); i++ {
// 		x = append(x, kd[i].date)
// 		y = append(y, opts.KlineData{Value: kd[i].data})
// 	}

// 	kline.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "DataZoom(inside&slider)",
// 		}),
// 		charts.WithXAxisOpts(opts.XAxis{
// 			SplitNumber: 20,
// 		}),
// 		charts.WithYAxisOpts(opts.YAxis{
// 			Scale: opts.Bool(true),
// 		}),
// 		charts.WithDataZoomOpts(opts.DataZoom{
// 			Type:       "inside",
// 			Start:      50,
// 			End:        100,
// 			XAxisIndex: []int{0},
// 		}),
// 		charts.WithDataZoomOpts(opts.DataZoom{
// 			Type:       "slider",
// 			Start:      50,
// 			End:        100,
// 			XAxisIndex: []int{0},
// 		}),
// 	)

// 	kline.SetXAxis(x).AddSeries("kline", y)
// 	return kline
// }

// func klineDataZoomYAxis() *charts.Kline {
// 	kline := charts.NewKLine()

// 	x := make([]string, 0)
// 	y := make([]opts.KlineData, 0)
// 	for i := 0; i < len(kd); i++ {
// 		x = append(x, kd[i].date)
// 		y = append(y, opts.KlineData{Value: kd[i].data})
// 	}

// 	kline.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "DataZoom(yAxis)",
// 		}),
// 		charts.WithXAxisOpts(opts.XAxis{
// 			SplitNumber: 20,
// 		}),
// 		charts.WithYAxisOpts(opts.YAxis{
// 			Scale: opts.Bool(true),
// 		}),
// 		charts.WithDataZoomOpts(opts.DataZoom{
// 			Type:       "slider",
// 			Start:      50,
// 			End:        100,
// 			YAxisIndex: []int{0},
// 		}),
// 	)

// 	kline.SetXAxis(x).AddSeries("kline", y)
// 	return kline
// }


// different style kline
// focus on this one --> markers should note when to buy and sell
	// doesnt ned to be distinguished by color
	// buy markers should display : "B"
	// sell markers should display : "S +gain" OR "S-loss" 
// ask sean for api (for the imported data) 
// change datasets to imported values rather than set numbers
// 
func klineStyle(arrayTime []string, arrayData [][]float32) *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([]opts.KlineData, 0)
	for i := 0; i < len(arrayData); i++ {
		x = append(x, arrayTime[i])
		y = append(y, opts.KlineData{Value: arrayData[i]})
	}

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "different style",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: opts.Bool(true),
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	type ReadMarkPointData struct {
        MarkLabel string `json:"label"`
        XCoordinate  string    `json:"x"`
        YCoordinate int `json:"y"`
        Profit string `json:"prof"`
    }
    // Open and read the file

	jsonFile, err := ioutil.ReadFile("test1MarkPoint.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        // return nil, err
    }

    var markPointValues []ReadMarkPointData
    err = json.Unmarshal(jsonFile, &markPointValues)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        // return nil, err
    }

    fmt.Println("Reading from JSON FILE: ", markPointValues)


	// customPointX2 := "2022-10-02T05:00:00Z"
	// customPointY2 := 19500
	markPointColor := "blue"

	kline.SetXAxis(x).AddSeries("kline", y)
	for i := 0; i < len(markPointValues); i++ {
		if markPointValues[i].MarkLabel == "buy" {
			markPointColor = "blue"
		}
		if markPointValues[i].MarkLabel == "sell" {
			markPointColor = "purple"
		}
		kline.SetSeriesOptions(
			// charts.WithMarkPointNameTypeItemOpts(opts.MarkPointNameTypeItem{
			// 	Name:     "low",
			// 	Type:     "min",
			// 	ValueDim: "lowest",
			// }),

			charts.WithMarkPointNameCoordItemOpts(opts.MarkPointNameCoordItem{
				Name: markPointValues[i].MarkLabel, // lable when hovering over markpoint
				Coordinate: []interface{}{markPointValues[i].XCoordinate, markPointValues[i].YCoordinate}, // coordinates of mark; string, int
				Value: markPointValues[i].Profit, // value displayed on top of markpoint
				Label: &opts.Label{
					Show:     opts.Bool(true),
					Color:    "orange",
					Position: "inside",
				},
				ItemStyle: &opts.ItemStyle{
		            Color:  markPointColor, // Customize color
		            Opacity: 1,   // Adjust opacity if necessary
		        },
				Symbol: "pin",
	 		}),
			
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show: opts.Bool(true),
				},
			}),
			charts.WithItemStyleOpts(opts.ItemStyle{
				Color:        "#ec0000",
				Color0:       "#00da3c",
				BorderColor:  "#8A0000",
				BorderColor0: "#008F28",
			}),
		)
	}
	return kline
}

type KlineExamples struct{}

func (KlineExamples) Examples() {

	// assigns the data from createArray to separate arrays
	var dataPoints [][]float32
	var dataTime []string
	dataTime, dataPoints = createArray()
	// fmt.Println("Array from data package:", dataTime)
	// fmt.Println("Array from data package:", dataPoints)

	page := components.NewPage()
	page.AddCharts(
		klineBase(dataTime, dataPoints),
		// klineDataZoomInside(),
		// klineDataZoomBoth(),
		// klineDataZoomYAxis(),
		klineStyle(dataTime, dataPoints),
	)


	f, err := os.Create("examples/html/kline.html")
	if err != nil {
		panic(err)

	}
	page.Render(io.MultiWriter(f))
}


// generates data from alpaca api and organizes it into separate arrays for graphing. 
func createArray() ([]string, [][]float32) {
	// importing data from alapaca api
	client := marketdata.NewClient(marketdata.ClientOpts{})
	request := marketdata.GetCryptoBarsRequest{
	  TimeFrame: marketdata.OneMin, // change data frequency by changing this to OneDay, OneMinute, etc
	  Start:     time.Date(2022, 9, 1, 0, 0, 0, 0, time.UTC),
	  End:       time.Date(2022, 9, 35, 0, 0, 0, 0, time.UTC),
	}

	bars, err := client.GetCryptoBars("BTC/USD", request)
	if err != nil {
		panic(err)
	}
	
	// extracts data from 'bars' and assigns it to the json 'rawData'
    rawData, err := json.MarshalIndent(bars, "", "   ")
    if err != nil {
    	log.Fatalf("marshaling error: %s", err)
    }

    // keys for the 'rawData' json file
    type DataStruct struct {
        TimeStamp string `json:"t"`
        Open  float32    `json:"o"`
        High float32 `json:"h"`
        Low float32 `json:"l"`
        Close  float32    `json:"c"`
        Volume float32 `json:"v"`
        TradeCount  float32    `json:"n"`
        VWAP float32 `json:"vw"`
    }

    // unmarshals the json into the data structure to be organized
    var dataSet []DataStruct
    err = json.Unmarshal([]byte(rawData), &dataSet)
    if err != nil {
        fmt.Println("error:", err)
        // return
    }

    // creates an array slice of undefined length to hold the data
    numRows := len(dataSet)
    const numCols = 4
    dataPoints := make([][]float32, numRows)

    // assigns data values to the array (formats the data)
    for i := 0; i < numRows; i++ {
        dataPoints[i] = make([]float32, numCols)
        for j := 0; j < numCols; j++ {
        	if j == 0 {
            	dataPoints[i][j] = dataSet[i].Open
       		}
       		if j == 1 {
            	dataPoints[i][j] = dataSet[i].Close
       		}
       		if j == 2 {
            	dataPoints[i][j] = dataSet[i].Low
       		}
       		if j == 3 {
            	dataPoints[i][j] = dataSet[i].High
       		}
        }
    }

    // creates an array to hold the date (x-axis / time)
    dataTime := make([]string, numRows)
    for i := 0; i < numRows; i++ {
    	dataTime[i] = dataSet[i].TimeStamp
    }
    return dataTime, dataPoints  
}

