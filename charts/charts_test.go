package charts_test

import (
	"image/color"
	"os"
	"testing"

	"github.com/nikolaydubina/calendarheatmap/charts"
)

var counts map[string]int = map[string]int{
	"2020-05-17": 13,
	"2020-05-18": 5,
	"2020-05-19": 8,
	"2020-05-20": 5,
	"2020-05-21": 5,
	"2020-05-22": 3,
	"2020-05-23": 5,
	"2020-05-24": 6,
	"2020-05-25": 3,
	"2020-05-26": 5,
	"2020-05-27": 8,
	"2020-05-28": 2,
	"2020-05-29": 2,
	"2020-05-30": 8,
	"2020-05-31": 5,
	"2020-06-01": 1,
	"2020-06-02": 3,
	"2020-06-03": 1,
	"2020-06-04": 3,
	"2020-06-05": 1,
	"2020-06-06": 3,
	"2020-06-07": 5,
	"2020-06-09": 1,
	"2020-06-10": 2,
	"2020-06-12": 9,
	"2020-06-13": 7,
	"2020-06-14": 4,
	"2020-06-15": 1,
	"2020-06-17": 1,
	"2020-06-20": 2,
	"2020-06-21": 1,
	"2020-06-23": 2,
	"2020-06-24": 2,
	"2020-06-25": 3,
	"2020-06-26": 3,
	"2020-06-27": 2,
	"2020-06-28": 1,
	"2020-06-29": 1,
	"2020-06-30": 2,
}

func save(t *testing.T, conf charts.HeatmapConfig, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		t.Errorf(err.Error())
	}
	if err := charts.WriteHeatmap(conf, f); err != nil {
		t.Errorf(err.Error())
	}
	if err := f.Close(); err != nil {
		t.Errorf(err.Error())
	}
}

func TestBasicData(t *testing.T) {
	os.Setenv("CALENDAR_HEATMAP_ASSETS_PATH", "assets")

	var colorscale = charts.BasicColorScale{
		color.RGBA{247, 252, 240, 255},
		color.RGBA{224, 243, 219, 255},
		color.RGBA{204, 235, 197, 255},
		color.RGBA{168, 221, 181, 255},
		color.RGBA{123, 204, 196, 255},
		color.RGBA{78, 179, 211, 255},
		color.RGBA{43, 140, 190, 255},
		color.RGBA{8, 104, 172, 255},
		color.RGBA{8, 64, 129, 255},
	}

	t.Run("basic", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "png",
		}
		save(t, conf, "testdata/basic.png")
	})

	t.Run("korean", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
			DrawMonthSeparator: true,
			DrawLabels:         true,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Locale:             "ko_KR",
			Format:             "png",
		}
		save(t, conf, "testdata/korean.png")
	})

	t.Run("no separator", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
			DrawMonthSeparator: false,
			DrawLabels:         true,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "png",
		}
		save(t, conf, "testdata/noseparator.png")
	})

	t.Run("no labels", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "png",
		}
		save(t, conf, "testdata/nolabels.png")
	})

	t.Run("no separator, no labels", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
			DrawMonthSeparator: false,
			DrawLabels:         false,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "png",
		}
		save(t, conf, "testdata/noseparator_nolabels.png")
	})

	t.Run("empty data", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "png",
		}
		save(t, conf, "testdata/empty_data.png")
	})

	t.Run("nil data", func(t *testing.T) {
		conf := charts.HeatmapConfig{
			Counts:             counts,
			ColorScale:         colorscale,
			DrawMonthSeparator: true,
			DrawLabels:         false,
			Margin:             30,
			BoxSize:            150,
			TextWidthLeft:      350,
			TextHeightTop:      200,
			TextColor:          color.RGBA{100, 100, 100, 255},
			BorderColor:        color.RGBA{200, 200, 200, 255},
			Format:             "png",
		}
		save(t, conf, "testdata/nil_data.png")
	})
}
