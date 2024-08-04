package examples

import (
	"io/ioutil"
	// "io"
	// "os"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	// "github.com/alpacahq/alpaca-trade-api-go/v2/alpaca"
	// "time" 
	"fmt"
	"encoding/json"
	"C"
	// "log"

)

// data structure is date, open, close, low, high
// dates are read as string, so data must be ordered sequentially by date

func lineStyle(arrayTime []string, arrayRSI []float32) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "LineGraph", Subtitle: "This is the subtitle.",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: opts.Bool(true),
		}),
		// charts.WithDataZoomOpts(opts.DataZoom{
		// 	Start:      50,
		// 	End:        100,
		// 	XAxisIndex: []int{0},
		// }),
	)

	lengthArray := len(arrayRSI)
	items := make([]opts.LineData, 0)
	for i := 0; i < lengthArray; i++ {
		items = append(items, opts.LineData{Value: arrayRSI[i]})
	}

	line.SetXAxis(arrayTime).
		AddSeries("Category A", items)
	return line
}


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
// type ReadGraphingData2 struct {
//     TimeStamp string `json:"time"`
//     Open  float32    `json:"open"`
//     High float32 `json:"high"`
//     Low float32 `json:"low"`
//     Close  float32    `json:"close"`
// }

// func fromJSON(documentPtr *C.char) *C.char {
//     documentString := C.GoString(documentPtr)
//     var jsonDocument map[string]interface{}
//     err := json.Unmarshal([]byte(documentString), &jsonDocument)
//     if err != nil {
//         log.Println("Error parsing JSON:", err)
//         return C.CString("Error parsing JSON")
//     }
//     filePath := "test2GraphingData.json"
//     if _, err := os.Stat(filePath); os.IsNotExist(err) {
//         initialContent := []byte("[]")
//         if err := ioutil.WriteFile(filePath, initialContent, 0644); err != nil {
//             log.Fatalf("Error creating JSON file: %v", err)
//         }
//         // fmt.Println("File created:", filePath)
//     }
//     file, err := ioutil.ReadFile(filePath)
//     if err != nil {
//         log.Fatalf("Error reading JSON file: %v", err)
//     }
//     // Unmarshal the JSON data into a slice of Person structs
//     var graphingValues []ReadGraphingData2
//     if err := json.Unmarshal(file, &graphingValues); err != nil {
//         log.Fatalf("Error parsing JSON: %v", err)
//     }
//     newDataBytes, err := json.Marshal(jsonDocument)
//     if err != nil {
//         log.Fatalf("Error marshaling new data: %v", err)
//     }

//     var newGraphingData ReadGraphingData2
//     if err := json.Unmarshal(newDataBytes, &newGraphingData); err != nil {
//         log.Fatalf("Error unmarshaling new data: %v", err)
//     }
//     // Append the new data to the slice
//     graphingValues = append(graphingValues, newGraphingData)
//     // Marshal the updated slice back to JSON
//     updatedJSON, err := json.MarshalIndent(graphingValues, "", "    ")
//     if err != nil {
//         log.Fatalf("Error marshaling JSON: %v", err)
//     }
//     // Write the updated JSON back to the file
//     if err := ioutil.WriteFile(filePath, updatedJSON, 0644); err != nil {
//         log.Fatalf("Error writing JSON to file: %v", err)
//     }

//     fmt.Println("Successfully added new data to JSON file.")


//     return C.CString("Going back  to python")
// }

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
			Title: "KlineGraph",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: opts.Bool(true),
		}),
		// charts.WithDataZoomOpts(opts.DataZoom{
		// 	Start:      50,
		// 	End:        100,
		// 	XAxisIndex: []int{0},
		// }),
	)

	type ReadMarkPointData struct {
        MarkLabel string `json:"label"`
        XCoordinate  string    `json:"time"`
        YCoordinate int `json:"price"`
        Profit string `json:"gain"`
    }
    // Open and read the file

	jsonFile, err := ioutil.ReadFile("./../Trading-main/TestingTradingData.json")
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
    
	// profit colors: green
	// loss colors: red
	// buy: triangle
	// sell: circle
 
	markPointColor := "blue"
	markPointSymbol := "triangle"

	kline.SetXAxis(x).AddSeries("kline", y)
	for i := 0; i < len(markPointValues); i++ {
		if markPointValues[i].MarkLabel == "buy" {
			markPointSymbol = "triangle"
			markPointColor = "black"
			// fmt.Println("CHECKING BUY MARKER")
			// fmt.Println("Reading from JSON FILE, getting profit: ", markPointValues[i].Profit)
		}
		if markPointValues[i].MarkLabel == "sell" {
			markPointSymbol = "circle"
			// fmt.Println("CHECKING SELL MARKER")
			if len(markPointValues[i].Profit) > 0 && markPointValues[i].Profit[0] == byte('-') {
	        markPointColor = "red"
		    }
		    if len(markPointValues[i].Profit) > 0 && markPointValues[i].Profit[0] != byte('-') {
		        markPointColor = "green"
		    }
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
					Show:     opts.Bool(false),
					Color:    "orange",
					Position: "inside",
				},
				ItemStyle: &opts.ItemStyle{
		            Color:  markPointColor, // Customize color
		            Opacity: 1,   // Adjust opacity if necessary
		        },
				Symbol: markPointSymbol,
				SymbolSize: 10,
	 		}),
			
			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
				Label: &opts.Label{
					Show: opts.Bool(true),
				},
			}),
			charts.WithItemStyleOpts(opts.ItemStyle{
				Color:        "#D3D3D3",
				Color0:       "#FAEBD7",
				BorderColor:  "#D3D3D3",
				BorderColor0: "#FAEBD7",
			}),
		)
	}
	return kline
	
}


type KlineExamples struct{}

func (KlineExamples) Examples(w http.ResponseWriter) {

	// assigns the data from createArray to separate arrays
	var dataPoints [][]float32
	var dataTime []string
	dataTime, dataPoints = createArray()

	var dataRSI []float32
	dataRSI = getRSI()
	// fmt.Println("Array from data package:", dataTime)
	// fmt.Println("Array from data package:", dataPoints)

	kline1 := klineStyle(dataTime, dataPoints)
    line2 := lineStyle(dataTime, dataRSI)

    // Synchronize data zoom options
    syncDataZoom := opts.DataZoom{
        Type:       "slider",
        Start:      50,
        End:        100,
        XAxisIndex: []int{0, 1}, // Synchronize both charts
    }

    kline1.SetGlobalOptions(
    	charts.WithTitleOpts(opts.Title{Title: "KlineGraph"}),
        charts.WithDataZoomOpts(syncDataZoom),
    )

    line2.SetGlobalOptions(
    	charts.WithTitleOpts(opts.Title{Title: "LineGraph"}),
        charts.WithDataZoomOpts(syncDataZoom),
    )
    fmt.Println("Rendering Charts Now")

    // Render the charts
    page := components.NewPage()
    page.AddCharts(kline1, line2)
	// page := components.NewPage()
	// page.AddCharts(
	// 	// klineBase(dataTime, dataPoints),
	// 	// klineDataZoomInside(),
	// 	// klineDataZoomBoth(),
	// 	// klineDataZoomYAxis(),
		
	// 	klineStyle(dataTime, dataPoints),
	// 	lineStyle(dataTime, dataRSI),
	// )


	// f, err := os.Create("examples/html/kline.html")
	// if err != nil {
	// 	panic(err)

	// }
	// page.Render(io.MultiWriter(f))
	page.Render(w)

    // Include JavaScript for synchronization
    js := `
        <script type="text/javascript">
            setTimeout(function() {
                var chart1 = echarts.getInstanceByDom(document.getElementById('chart_0'));
                var chart2 = echarts.getInstanceByDom(document.getElementById('chart_1'));

                function syncCharts(sourceChart, targetChart) {
                    sourceChart.on('dataZoom', function(params) {
                        targetChart.dispatchAction({
                            type: 'dataZoom',
                            start: params.start,
                            end: params.end,
                        });
                    });
                }

                syncCharts(chart1, chart2);
                syncCharts(chart2, chart1);
            }, 1000); // Delay to ensure charts are fully initialized
        </script>
    `
    w.Write([]byte(js))
}

func getRSI() ([]float32) {
	type ReadRSIData struct {
        RSI float32 `json:"rsi"`
    }
    // Open and read the file

	// jsonGraphingFile, err := ioutil.ReadFile("test1GraphingData.json")
	jsonRSIFile, err := ioutil.ReadFile("./../Trading-main/RecordingRSI.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        // return nil, err
    }

    // fmt.Println("Getting data values from 'jsonGraphingFile':", jsonGraphingFile)
    var rsiValues []ReadRSIData
    err = json.Unmarshal(jsonRSIFile, &rsiValues)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        // return nil, err
    }
    

	fmt.Println("TESTING RSI DATA, getting time json: ", rsiValues[0].RSI)

	lenRSI := len(rsiValues)
	fmt.Println("LENGTH OF RSI VALUES: ", lenRSI)

	rsiData := make([]float32, lenRSI)
    for i := 0; i < lenRSI; i++ {
    	rsiData[i] = rsiValues[i].RSI
    }

    fmt.Println("TESTING RSI DATA, getting time array: ", rsiData[0])

    return rsiData

}

// generates data and organizes it into separate arrays for graphing. 
// no more export createArray
func createArray() ([]string, [][]float32) {
	
	type ReadGraphingData struct {
        TimeStamp string `json:"time"`
        Open  float32    `json:"open"`
        High float32 `json:"high"`
        Low float32 `json:"low"`
        Close  float32    `json:"close"`
    }
    // Open and read the file

	// jsonGraphingFile, err := ioutil.ReadFile("test1GraphingData.json")
	jsonGraphingFile, err := ioutil.ReadFile("./../Trading-main/TestingGraphingData.json")
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        // return nil, err
    }

    // fmt.Println("Getting data values from 'jsonGraphingFile':", jsonGraphingFile)
    var graphingValues []ReadGraphingData
    err = json.Unmarshal(jsonGraphingFile, &graphingValues)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        // return nil, err
    }
    

	fmt.Println("TESTING GRAPHING DATA, getting time: ", graphingValues[0].TimeStamp)


    // creates an array slice of undefined length to hold the data
    numRows := len(graphingValues)
    const numCols = 4
    dataPoints := make([][]float32, numRows)

    // assigns data values to the array (formats the data)
    for i := 0; i < numRows; i++ {
        dataPoints[i] = make([]float32, numCols)
        for j := 0; j < numCols; j++ {
        	if j == 0 {
            	dataPoints[i][j] = graphingValues[i].Open
       		}
       		if j == 1 {
            	dataPoints[i][j] = graphingValues[i].Close
       		}
       		if j == 2 {
            	dataPoints[i][j] = graphingValues[i].Low
       		}
       		if j == 3 {
            	dataPoints[i][j] = graphingValues[i].High
       		}
        }
    }

    // creates an array to hold the date (x-axis / time)
    dataTime := make([]string, numRows)
    for i := 0; i < numRows; i++ {
    	dataTime[i] = graphingValues[i].TimeStamp
    }
    

    return dataTime, dataPoints  
}

// func renderPage(w http.ResponseWriter, _ *http.Request) {
//     page := createArray()
//     page.Render(w)
// }

// func generateLocal() {
//     http.HandleFunc("/", renderPage)
//     http.ListenAndServe(":8080", nil)
//     fmt.Println("RUNning at server from go")
// }

