package examples

import (
	"io"
	"math/rand"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var (
	itemCntLine = 6
	fruits      = []string{"Apple", "Banana", "Peach ", "Lemon", "Pear", "Cherry"}
)

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < itemCntLine; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func generateLineData(data []float32) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}

func lineOverlap() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "overlap rect-charts"}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems())
	line.Overlap(esEffectStyle())
	line.Overlap(scatterBase())
	return line
}

func lineBase() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic line example", Subtitle: "This is the subtitle."}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems())
	return line
}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "title and label options",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		SetSeriesOptions(
			charts.WithLineChartOpts(opts.LineChart{
				ShowSymbol: opts.Bool(true),
			}),
			charts.WithLabelOpts(opts.Label{
				Show: opts.Bool(true),
			}),
		)
	return line
}

// func lineSplitLine() *charts.Line {
// 	line := charts.NewLine()
// 	line.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "splitline options",
// 		}),
// 		charts.WithYAxisOpts(opts.YAxis{
// 			SplitLine: &opts.SplitLine{
// 				Show: opts.Bool(true),
// 			},
// 		}),
// 	)

// 	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems(),
// 		charts.WithLabelOpts(
// 			opts.Label{Show: opts.Bool(true)},
// 		))
// 	return line
// }

// func lineStep() *charts.Line {
// 	line := charts.NewLine()
// 	line.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "step style",
// 		}),
// 	)

// 	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
// 		SetSeriesOptions(charts.WithLineChartOpts(
// 			opts.LineChart{
// 				Step: true,
// 			}),
// 		)
// 	return line
// }

// func lineSmooth() *charts.Line {
// 	line := charts.NewLine()
// 	line.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "smooth style",
// 		}),
// 	)

// 	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
// 		SetSeriesOptions(charts.WithLineChartOpts(
// 			opts.LineChart{
// 				Smooth: opts.Bool(true),
// 			}),
// 		)
// 	return line
// }

// func lineArea() *charts.Line {
// 	line := charts.NewLine()
// 	line.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "area options",
// 		}),
// 	)

// 	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
// 		SetSeriesOptions(
// 			charts.WithLabelOpts(
// 				opts.Label{
// 					Show: opts.Bool(true),
// 				}),
// 			charts.WithAreaStyleOpts(
// 				opts.AreaStyle{
// 					Opacity: 0.2,
// 				}),
// 		)
// 	return line
// }

// func lineSmoothArea() *charts.Line {
// 	line := charts.NewLine()
// 	line.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{Title: "smooth area"}),
// 	)

// 	line.SetXAxis(fruits).AddSeries("Category A", generateLineItems()).
// 		SetSeriesOptions(
// 			charts.WithLabelOpts(opts.Label{
// 				Show: opts.Bool(true),
// 			}),
// 			charts.WithAreaStyleOpts(opts.AreaStyle{
// 				Opacity: 0.2,
// 			}),
// 			charts.WithLineChartOpts(opts.LineChart{
// 				Smooth: opts.Bool(true),
// 			}),
// 		)
// 	return line
// }

// func lineDemo() *charts.Line {
// 	line := charts.NewLine()

// 	line.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{
// 			Title: "Search Time: Hash table vs Binary search",
// 		}),
// 		charts.WithYAxisOpts(opts.YAxis{
// 			Name: "Cost time(ns)",
// 			SplitLine: &opts.SplitLine{
// 				Show: opts.Bool(true),
// 			},
// 		}),
// 		charts.WithXAxisOpts(opts.XAxis{
// 			Name: "Elements",
// 		}),
// 	)

// 	line.SetXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
// 		AddSeries("map", generateLineItems(),
// 			charts.WithLabelOpts(opts.Label{Show: opts.Bool(true), Position: "bottom"})).
// 		AddSeries("slice", generateLineData([]float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131}),
// 			charts.WithLabelOpts(opts.Label{Show: opts.Bool(true), Position: "top"})).
// 		SetSeriesOptions(
// 			charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{
// 				Name: "Average",
// 				Type: "average",
// 			}),
// 			charts.WithLineChartOpts(opts.LineChart{
// 				Smooth: opts.Bool(true),
// 			}),
// 			charts.WithMarkPointStyleOpts(opts.MarkPointStyle{
// 				Label: &opts.Label{
// 					Show:      opts.Bool(true),
// 					Formatter: "{a}: {b}",
// 				},
// 			}),
// 		)

// 	return line
// }

func lineSymbols() *charts.Line {

	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "symbol options",
			Subtitle: "tooltip with 'axis' trigger",
		}),
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}),
	)

	// Put data into instance
	line.SetXAxis(fruits).
		AddSeries("Category A", generateLineItems()).
		AddSeries("Category B", generateLineItems()).
		SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: opts.Bool(false), ShowSymbol: opts.Bool(true), SymbolSize: 15, Symbol: "diamond"},
		))

	return line
}

// adds markpoints to graph
func lineChartScaled(arrayTime []string, arrayClose []float32, arrayRSI []float32) *charts.Line {
	line := charts.NewLine()
	type ReadMarkPointData struct {
        MarkLabel string `json:"label"`
        XCoordinate  string    `json:"time"`
        YCoordinate int `json:"price"`
        Profit string `json:"gain"`
    }

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
 
	markPointColor := "blue"
	markPointSymbol := "triangle"

	lengthArray := len(arrayClose)
	items := make([]opts.LineData, 0)
	for i := 0; i < lengthArray; i++ {
		items = append(items, opts.LineData{Value: arrayClose[i], YAxisIndex:1})
	}
	lengthRSI := len(arrayRSI)
	itemsRSI := make([]opts.LineData, 0)
	for i := 0; i < lengthRSI; i++ {
		itemsRSI = append(itemsRSI, opts.LineData{Value: arrayRSI[i], YAxisIndex:0})
	}

	// add series to graph
	line.SetXAxis(arrayTime).
		// AddSeries("Close Data", items)
		AddSeries("RSI Data", itemsRSI,
			charts.WithLineChartOpts(opts.LineChart{YAxisIndex: 0}),
		).
		AddSeries("Close Data", items,
			charts.WithLineChartOpts(opts.LineChart{YAxisIndex: 1}),
		)
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "LineGraph",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			SplitNumber: 20,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:  "RSI",
			Scale: opts.Bool(true),
			Min: 5,
			Max: 150,
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
				LineStyle: &opts.LineStyle{
					Color: "blue",
					Type:  "dashed",
				},
			},
		}),
		// charts.WithYAxisOpts(opts.YAxis{
		// 	Scale: opts.Bool(true),
		// 	Min: 210,
		// 	Max:230,
		// }),
		// charts.WithYAxisOpts(opts.YAxis{
  //           Name: "Series 1 Values",
  //           Min:  0,
  //           Max:  100,
  //       }),
  //       charts.WithYAxisOpts(opts.YAxis{
  //           Name:      "Series 2 Values",
  //           Min:       200,
  //           Max:       300,
  //           NameLocation:  "end",
  //           // AxisLine:  &opts.AxisLine{Show: true},
  //           // AxisLabel: &opts.AxisLabel{Show: true},
  //       }),
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
	)

	line.ExtendYAxis(opts.YAxis{
		Name:  "Close",
		// Type:  "value",
		Show:  opts.Bool(true),
		Scale: opts.Bool(true), // only available when min and max are set to specific values 
		Min: 160,
		Max: "dataMax",
		SplitLine: &opts.SplitLine{
			Show: opts.Bool(true),
			LineStyle: &opts.LineStyle{
				Color: "green",
				Type:  "dashed",
			},
		},
		//GridIndex: 1, // y index 1 // not required
	})
	for i := 0; i < len(markPointValues); i++ {
		// fmt.Println("Getting the Profit: ", markPointValues[i].Profit)
		if markPointValues[i].MarkLabel == "buy" {
			markPointSymbol = "triangle"
			markPointColor = "black"
			// fmt.Println("Reading from JSON FILE, getting price: ", markPointValues[i].YCoordinate)
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
		line.SetSeriesOptions(
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
		)
	}
	return line
}

// set graph with multiple series
func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "multi lines",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Theme: "shine",
		}),
	)

	line.SetXAxis(fruits).
		AddSeries("Category  A", generateLineItems()).
		AddSeries("Category  B", generateLineItems()).
		AddSeries("Category  C", generateLineItems()).
		AddSeries("Category  D", generateLineItems())
	return line
}

func lineChartUnscaled(arrayTime []string, arrayClose []float32, arrayRSI []float32) *charts.Line {
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
		charts.WithDataZoomOpts(opts.DataZoom{
			Start:      50,
			End:        100,
			XAxisIndex: []int{0},
		}),
		// charts.WithDataZoomOpts(opts.DataZoom{
		// 	Start:      50,
		// 	End:        100,
		// 	YAxisIndex: []int{0},
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
	// fmt.Println("Printing the markPointValues: ", markPointValues)

	lengthArray := len(arrayClose)
	items := make([]opts.LineData, 0)
	for i := 0; i < lengthArray; i++ {
		items = append(items, opts.LineData{Value: arrayClose[i]})
	}
	lengthRSI := len(arrayRSI)
	itemsRSI := make([]opts.LineData, 0)
	for i := 0; i < lengthRSI; i++ {
		itemsRSI = append(itemsRSI, opts.LineData{Value: arrayRSI[i]})
	}

	// add series to graph
	line.SetXAxis(arrayTime).
		// AddSeries("Close Data", items)
		AddSeries("Close Data", items).
		AddSeries("RSI Data", itemsRSI)
	for i := 0; i < len(markPointValues); i++ {
		// fmt.Println("Getting the Profit: ", markPointValues[i].Profit)
		if markPointValues[i].MarkLabel == "buy" {
			markPointSymbol = "triangle"
			markPointColor = "black"
			// fmt.Println("Reading from JSON FILE, getting price: ", markPointValues[i].YCoordinate)
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
		line.SetSeriesOptions(
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
		)
	}
	return line
}

type LineExamples struct{}

func (LineExamples) Examples() {
	var dataRSI []float32
	dataRSI = getRSIdata()

	var dataPoints []float32
	var dataTime []string
	dataTime, dataPoints = getTradingdata()

	page := components.NewPage()
	page.AddCharts(
		// lineBase(),
		// lineShowLabel(),
		// lineSymbols(),
		lineChartScaled(dataTime, dataPoints, dataRSI),
		// lineChartUnscaled(dataTime, dataPoints, dataRSI),
		// lineSplitLine(),
		// lineStep(),
		// lineSmooth(),
		// lineArea(),
		// lineSmoothArea(),
		// lineOverlap(),
		// lineMulti(),
		// lineDemo(),
	)
	f, err := os.Create("examples/html/line.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}

func getRSIdata() ([]float32) {
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

func getTradingdata() ([]string, []float32) {
	
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
    dataPoints := make([]float32, numRows)

    // assigns data values to the array (formats the data)
    for i := 0; i < numRows; i++ {
    	dataPoints[i] = graphingValues[i].Close
    }

    // creates an array to hold the date (x-axis / time)
    dataTime := make([]string, numRows)
    for i := 0; i < numRows; i++ {
    	dataTime[i] = graphingValues[i].TimeStamp
    }
    

    return dataTime, dataPoints  
}
