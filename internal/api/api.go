package api

import (
    "encoding/json"
	 "encoding/csv"
    "net/http"
    "os"
	 "fmt"
	 "strings"
)

type OpenWeatherResponse struct {
    Lat            float64 `json:"lat"`
    Lon            float64 `json:"lon"`
    Timezone       string  `json:"timezone"`
    TimezoneOffset int     `json:"timezone_offset"`
    Current        struct {
        Dt         int64   `json:"dt"`
        Sunrise    int64   `json:"sunrise"`
        Sunset     int64   `json:"sunset"`
        Temp       float64 `json:"temp"`
        FeelsLike  float64 `json:"feels_like"`
        Pressure   int     `json:"pressure"`
        Humidity   int     `json:"humidity"`
        DewPoint   float64 `json:"dew_point"`
        Uvi        float64 `json:"uvi"`
        Clouds     int     `json:"clouds"`
        Visibility int     `json:"visibility"`
        WindSpeed  float64 `json:"wind_speed"`
        WindDeg    int     `json:"wind_deg"`
        Weather    []struct {
            ID          int    `json:"id"`
            Main        string `json:"main"`
            Description string `json:"description"`
            Icon        string `json:"icon"`
        } `json:"weather"`
    } `json:"current"`
    Minutely []struct {
        Dt            int64   `json:"dt"`
        Precipitation float64 `json:"precipitation"`
    } `json:"minutely"`
    Hourly []struct {
        Dt         int64   `json:"dt"`
        Temp       float64 `json:"temp"`
        FeelsLike  float64 `json:"feels_like"`
        Pressure   int     `json:"pressure"`
        Humidity   int     `json:"humidity"`
        DewPoint   float64 `json:"dew_point"`
        Uvi        float64 `json:"uvi"`
        Clouds     int     `json:"clouds"`
        Visibility int     `json:"visibility"`
        WindSpeed  float64 `json:"wind_speed"`
        WindDeg    int     `json:"wind_deg"`
        Weather    []struct {
            ID          int    `json:"id"`
            Main        string `json:"main"`
            Description string `json:"description"`
            Icon        string `json:"icon"`
        } `json:"weather"`
        Pop float64 `json:"pop"`
    } `json:"hourly"`
    Daily []struct {
        Dt         int64 `json:"dt"`
        Sunrise    int64 `json:"sunrise"`
        Sunset     int64 `json:"sunset"`
        Moonrise   int64 `json:"moonrise"`
        Moonset    int64 `json:"moonset"`
        MoonPhase  float64 `json:"moon_phase"`
        Temp       struct {
            Day   float64 `json:"day"`
            Min   float64 `json:"min"`
            Max   float64 `json:"max"`
            Night float64 `json:"night"`
            Eve   float64 `json:"eve"`
            Morn  float64 `json:"morn"`
        } `json:"temp"`
        FeelsLike struct {
            Day   float64 `json:"day"`
            Night float64 `json:"night"`
            Eve   float64 `json:"eve"`
            Morn  float64 `json:"morn"`
        } `json:"feels_like"`
        Pressure   int     `json:"pressure"`
        Humidity   int     `json:"humidity"`
        DewPoint   float64 `json:"dew_point"`
        WindSpeed  float64 `json:"wind_speed"`
        WindDeg    int     `json:"wind_deg"`
        Weather    []struct {
            ID          int    `json:"id"`
            Main        string `json:"main"`
            Description string `json:"description"`
            Icon        string `json:"icon"`
        } `json:"weather"`
        Clouds     int     `json:"clouds"`
        Pop        float64 `json:"pop"`
        Rain       float64 `json:"rain"`
        Uvi        float64 `json:"uvi"`
    } `json:"daily"`
    Alerts []struct {
        SenderName  string   `json:"sender_name"`
        Event       string   `json:"event"`
        Start       int64    `json:"start"`
        End         int64    `json:"end"`
        Description string   `json:"description"`
        Tags        []string `json:"tags"`
    } `json:"alerts"`
}

type NASAPowerResponse struct {
    CDD10_SUM                float64            `json:"CDD10_SUM"`
    CDD18_SUM                float64            `json:"CDD18_SUM"`
    CDH23_SUM                float64            `json:"CDH23_SUM"`
    CDH26_SUM                float64            `json:"CDH26_SUM"`
    CM                       float64            `json:"CM"`
    DB_004                   float64            `json:"DB_004"`
    DB_010                   float64            `json:"DB_010"`
    DB_980                   float64            `json:"DB_980"`
    DB_990                   float64            `json:"DB_990"`
    DB_996                   float64            `json:"DB_996"`
    DB_AVG                   float64            `json:"DB_AVG"`
    DB_MAX_AVG               float64            `json:"DB_MAX_AVG"`
    DB_MAX_STD               float64            `json:"DB_MAX_STD"`
    DB_MIN_AVG               float64            `json:"DB_MIN_AVG"`
    DB_MIN_STD               float64            `json:"DB_MIN_STD"`
    DB_RANGE                 float64            `json:"DB_RANGE"`
    DB_STD                   float64            `json:"DB_STD"`
    DP_004                   float64            `json:"DP_004"`
    DP_010                   float64            `json:"DP_010"`
    DP_980                   float64            `json:"DP_980"`
    DP_990                   float64            `json:"DP_990"`
    DP_996                   float64            `json:"DP_996"`
    EN_980                   float64            `json:"EN_980"`
    EN_990                   float64            `json:"EN_990"`
    EN_996                   float64            `json:"EN_996"`
    EXT_WS_010               float64            `json:"EXT_WS_010"`
    EXT_WS_025               float64            `json:"EXT_WS_025"`
    EXT_WS_050               float64            `json:"EXT_WS_050"`
    HDD10_SUM                float64            `json:"HDD10_SUM"`
    HDD18_SUM                float64            `json:"HDD18_SUM"`
    MCDB_004                 float64            `json:"MCDB_004"`
    MCDB_010                 float64            `json:"MCDB_010"`
    MCDB_CM_990              float64            `json:"MCDB_CM_990"`
    MCDB_CM_996              float64            `json:"MCDB_CM_996"`
    MCDB_DP_980              float64            `json:"MCDB_DP_980"`
    MCDB_DP_990              float64            `json:"MCDB_DP_990"`
    MCDB_DP_996              float64            `json:"MCDB_DP_996"`
    MCDB_EN_980              float64            `json:"MCDB_EN_980"`
    MCDB_EN_990              float64            `json:"MCDB_EN_990"`
    MCDB_EN_996              float64            `json:"MCDB_EN_996"`
    MCDB_WB_980              float64            `json:"MCDB_WB_980"`
    MCDB_WB_990              float64            `json:"MCDB_WB_990"`
    MCDB_WB_996              float64            `json:"MCDB_WB_996"`
    MCHR_004                 float64            `json:"MCHR_004"`
    MCHR_010                 float64            `json:"MCHR_010"`
    MCHR_DP_980              float64            `json:"MCHR_DP_980"`
    MCHR_DP_990              float64            `json:"MCHR_DP_990"`
    MCHR_DP_996              float64            `json:"MCHR_DP_996"`
    MCWB_980                 float64            `json:"MCWB_980"`
    MCWB_990                 float64            `json:"MCWB_990"`
    MCWB_996                 float64            `json:"MCWB_996"`
    MCWD_004                 float64            `json:"MCWD_004"`
    MCWD_DB_996              float64            `json:"MCWD_DB_996"`
    MCWS_004                 float64            `json:"MCWS_004"`
    MCWS_DB_996              float64            `json:"MCWS_DB_996"`
    PRECTOTCORR_ANNUAL_MAX   float64            `json:"PRECTOTCORR_ANNUAL_MAX"`
    PRECTOTCORR_ANNUAL_MIN   float64            `json:"PRECTOTCORR_ANNUAL_MIN"`
    PRECTOTCORR_ANNUAL_STD   float64            `json:"PRECTOTCORR_ANNUAL_STD"`
    PRECTOTCORR_SUM          float64            `json:"PRECTOTCORR_SUM"`
    T2MWET_MAX_VALUE         float64            `json:"T2MWET_MAX_VALUE"`
    WB_980                   float64            `json:"WB_980"`
    WB_990                   float64            `json:"WB_990"`
    WB_996                   float64            `json:"WB_996"`
    WB_MAX_AVG               float64            `json:"WB_MAX_AVG"`
    WB_MAX_STD               float64            `json:"WB_MAX_STD"`
    WB_MIN_AVG               float64            `json:"WB_MIN_AVG"`
    WB_MIN_STD               float64            `json:"WB_MIN_STD"`
    WM                       float64            `json:"WM"`
    WS10M_ANNUAL             float64            `json:"WS10M_ANNUAL"`
    WS_CM_990                float64            `json:"WS_CM_990"`
    WS_CM_996                float64            `json:"WS_CM_996"`
    CDD10                    map[string]float64 `json:"CDD10"`
    CDD18                    map[string]float64 `json:"CDD18"`
    CDH23                    map[string]float64 `json:"CDH23"`
    CDH26                    map[string]float64 `json:"CDH26"`
    CLRSKY_SFC_SW_DIFF_NOON_HOUR map[string]float64 `json:"CLRSKY_SFC_SW_DIFF_NOON_HOUR"`
    CLRSKY_SFC_SW_DNI_NOON_HOUR  map[string]float64 `json:"CLRSKY_SFC_SW_DNI_NOON_HOUR"`
    DB_RANGE_MONTH           map[string]float64 `json:"DB_RANGE_MONTH"`
    DB_TEMP_900              map[string]float64 `json:"DB_TEMP_900"`
    DB_TEMP_950              map[string]float64 `json:"DB_TEMP_950"`
    DB_TEMP_980              map[string]float64 `json:"DB_TEMP_980"`
    DB_TEMP_996              map[string]float64 `json:"DB_TEMP_996"`
    HDD10                    map[string]float64 `json:"HDD10"`
    HDD18                    map[string]float64 `json:"HDD18"`
    MCDBR_DB                 map[string]float64 `json:"MCDBR_DB"`
    MCDBR_WB                 map[string]float64 `json:"MCDBR_WB"`
    MCDB_TEMP_900            map[string]float64 `json:"MCDB_TEMP_900"`
    MCDB_TEMP_950            map[string]float64 `json:"MCDB_TEMP_950"`
    MCDB_TEMP_980            map[string]float64 `json:"MCDB_TEMP_980"`
    MCDB_TEMP_996            map[string]float64 `json:"MCDB_TEMP_996"`
    MCWBR_DB                 map[string]float64 `json:"MCWBR_DB"`
    MCWBR_WB                 map[string]float64 `json:"MCWBR_WB"`
    MCWB_TEMP_900            map[string]float64 `json:"MCWB_TEMP_900"`
    MCWB_TEMP_950            map[string]float64 `json:"MCWB_TEMP_950"`
    MCWB_TEMP_980            map[string]float64 `json:"MCWB_TEMP_980"`
    MCWB_TEMP_996            map[string]float64 `json:"MCWB_TEMP_996"`
    PRECTOTCORR              map[string]float64 `json:"PRECTOTCORR"`
    PRECTOTCORR_MONTH_MAX    map[string]float64 `json:"PRECTOTCORR_MONTH_MAX"`
    PRECTOTCORR_MONTH_MIN    map[string]float64 `json:"PRECTOTCORR_MONTH_MIN"`
    PRECTOTCORR_MONTH_STD    map[string]float64 `json:"PRECTOTCORR_MONTH_STD"`
    RADIATION_AVG            map[string]float64 `json:"RADIATION_AVG"`
    RADIATION_STD            map[string]float64 `json:"RADIATION_STD"`
    T2M_AVG                  map[string]float64 `json:"T2M_AVG"`
    T2M_STD                  map[string]float64 `json:"T2M_STD"`
    WB_TEMP_900              map[string]float64 `json:"WB_TEMP_900"`
    WB_TEMP_950              map[string]float64 `json:"WB_TEMP_950"`
    WB_TEMP_980              map[string]float64 `json:"WB_TEMP_980"`
    WB_TEMP_996              map[string]float64 `json:"WB_TEMP_996"`
    WS10M_MONTH              map[string]float64 `json:"WS10M_MONTH"`
    DB_EXTREME_MAX           map[string]float64 `json:"DB_EXTREME_MAX"`
    DB_EXTREME_MIN           map[string]float64 `json:"DB_EXTREME_MIN"`
    WB_EXTREME_MAX           map[string]float64 `json:"WB_EXTREME_MAX"`
    WB_EXTREME_MIN           map[string]float64 `json:"WB_EXTREME_MIN"`
    Metadata                 struct {
        Lon       float64 `json:"lon"`
        Lat       float64 `json:"lat"`
        Elevation float64 `json:"elevation"`
        Timezone  int     `json:"timezone"`
        Pressure  float64 `json:"pressure"`
        Start     int     `json:"start"`
        End       int     `json:"end"`
    } `json:"metadata"`
}

func fetchOpenWeatherData(lat, lon string) (*OpenWeatherResponse, error) {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		 return nil, fmt.Errorf("OPENWEATHER_API_KEY not set in environment")
	}

	url := "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lon + "&appid=" + apiKey + "&units=metric"
	resp, err := http.Get(url)
	if err != nil {
		 return nil, err
	}
	defer resp.Body.Close()

	var data OpenWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		 return nil, err
	}
	return &data, nil
}

func fetchNASAPowerData(lat, lon, start, end string) (*NASAPowerResponse, error) {
    url := "https://power.larc.nasa.gov/api/application/indicators/point?latitude=" + lat + "&longitude=" + lon + "&start=" + start + "&end=" + end + "&format=JSON"
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var data NASAPowerResponse
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }
    return &data, nil
}

func getLatLonFromCountry(country string) (string, string, error) {
	file, err := os.Open("geocoding_results.csv")
	if err != nil {
		 return "", "", err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		 return "", "", err
	}

	for _, record := range records {
		 if strings.EqualFold(record[0], country) {
			  return record[1], record[2], nil
		 }
	}

	return "", "", fmt.Errorf("country not found")
}

func GetCombinedData(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	if country == "" {
		 http.Error(w, "country parameter is required", http.StatusBadRequest)
		 return
	}

	lat, lon, err := getLatLonFromCountry(country)
	if err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	}

	startDate := "2001"
	endDate := "2016"

	openWeatherData, err := fetchOpenWeatherData(lat, lon)
	if err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	}

	nasaPowerData, err := fetchNASAPowerData(lat, lon, startDate, endDate)
	if err != nil {
		 http.Error(w, err.Error(), http.StatusInternalServerError)
		 return
	}

	combinedData := struct {
		 OpenWeather *OpenWeatherResponse `json:"open_weather"`
		 NASAPower   *NASAPowerResponse   `json:"nasa_power"`
	}{
		 OpenWeather: openWeatherData,
		 NASAPower:   nasaPowerData,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(combinedData)
}