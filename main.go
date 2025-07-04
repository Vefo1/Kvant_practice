package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os" // Import os for environment variables
	"strconv"
	"time" // Import time for timeouts
)

const (
	targetAPIURL = "https://apiml.labhub.online/api/v1/predict/hba1c"
	// bearerToken will be obtained from an environment variable
)

// PredictRequest represents the data structure for the external API request
type PredictRequest struct {
	UID    string  `json:"uid"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	RDW    float64 `json:"rdw"`
	WBC    float64 `json:"wbc"`
	RBC    float64 `json:"rbc"`
	HGB    float64 `json:"hgb"`
	HCT    float64 `json:"hct"`
	MCV    float64 `json:"mcv"`
	MCH    float64 `json:"mch"`
	MCHC   float64 `json:"mchc"`
	PLT    float64 `json:"plt"`
	NEU    float64 `json:"neu"`
	EOS    float64 `json:"eos"`
	BAS    float64 `json:"bas"`
	LYM    float64 `json:"lym"`
	MON    float64 `json:"mon"`
	SOE    float64 `json:"soe"`
	CHOL   float64 `json:"chol"`
	GLU    float64 `json:"glu"`
}

func predictHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the bearer token from environment variables
	bearerToken := os.Getenv("PREDICT_API_BEARER_TOKEN")
	if bearerToken == "" {
		http.Error(w, "Configuration Error: PREDICT_API_BEARER_TOKEN not set", http.StatusInternalServerError)
		log.Println("Configuration Error: PREDICT_API_BEARER_TOKEN not set")
		return
	}

	query := r.URL.Query()

	// Helper function to safely parse float64 from query parameters
	parseFloat := func(paramName string) (float64, error) {
		valStr := query.Get(paramName)
		if valStr == "" {
			return 0.0, fmt.Errorf("Parameter '%s' is missing", paramName)
		}
		val, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return 0.0, fmt.Errorf("Error parsing parameter '%s': %v", paramName, err)
		}
		return val, nil
	}

	// Helper function to safely parse int from query parameters
	parseInt := func(paramName string) (int, error) {
		valStr := query.Get(paramName)
		if valStr == "" {
			return 0, fmt.Errorf("Parameter '%s' is missing", paramName)
		}
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return 0, fmt.Errorf("Error parsing parameter '%s': %v", paramName, err)
		}
		return val, nil
	}

	// Parse parameters with error handling
	age, err := parseInt("age")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gender, err := parseInt("gender")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rdw, err := parseFloat("rdw")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	wbc, err := parseFloat("wbc")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rbc, err := parseFloat("rbc")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hgb, err := parseFloat("hgb")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	hct, err := parseFloat("hct")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mcv, err := parseFloat("mcv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mch, err := parseFloat("mch")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mchc, err := parseFloat("mchc")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	plt, err := parseFloat("plt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	neu, err := parseFloat("neu")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	eos, err := parseFloat("eos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bas, err := parseFloat("bas")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lym, err := parseFloat("lym")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mon, err := parseFloat("mon")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	soe, err := parseFloat("soe")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	chol, err := parseFloat("chol")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	glu, err := parseFloat("glu")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the request object for the external API
	predictReq := PredictRequest{
		UID:    "web-client", // Assuming UID is fixed for this proxy
		Age:    age,
		Gender: gender,
		RDW:    rdw,
		WBC:    wbc,
		RBC:    rbc,
		HGB:    hgb,
		HCT:    hct,
		MCV:    mcv,
		MCH:    mch,
		MCHC:   mchc,
		PLT:    plt,
		NEU:    neu,
		EOS:    eos,
		BAS:    bas,
		LYM:    lym,
		MON:    mon,
		SOE:    soe,
		CHOL:   chol,
		GLU:    glu,
	}

	// Marshal the struct into JSON
	jsonBody, err := json.Marshal(predictReq)
	if err != nil {
		http.Error(w, "JSON serialization error: "+err.Error(), http.StatusInternalServerError)
		log.Printf("JSON serialization error: %v", err)
		return
	}

	// Create a new POST request to the external API
	req, err := http.NewRequest(http.MethodPost, targetAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, "Error creating request to external API: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating request to external API: %v", err)
		return
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+bearerToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	// Configure HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // 10-second timeout for the external API request
	}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error sending request to external API: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error sending request to external API: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response from the external API
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response from external API: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error reading response from external API: %v", err)
		return
	}

	// Set the status code and headers from the external API's response
	w.WriteHeader(resp.StatusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // Assuming the external API returns JSON
	w.Write(body)                                                     // Pass the external API's response directly to the client
}

func main() {
	// Register the handler for the /predict route
	http.HandleFunc("/predict", predictHandler)

	// Start the server on port 8080
	port := ":8080"
	log.Printf("Server started on http://localhost%s/predict", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
