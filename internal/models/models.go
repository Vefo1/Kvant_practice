package models

// HBA1CPredictRequest represents the data structure for the hba1c prediction external API request
type HBA1CPredictRequest struct {
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

// LdllPredictRequest represents the data structure for the ldll prediction external API request
type LdllPredictRequest struct {
	UID    string  `json:"uid"`
	Age    int     `json:"age"`
	Gender int     `json:"gender"`
	CHOL   float64 `json:"chol"`
	HDL    float64 `json:"hdl"`
	TG     float64 `json:"tg"`
}

// FerrPredictRequest represents the data structure for the ferr prediction external API request
type FerrPredictRequest struct {
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
	CRP    float64 `json:"crp"` // Specific to Ferr
}

// LdlPredictRequest represents the data structure for the ldl prediction external API request
type LdlPredictRequest struct {
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

// TgPredictRequest represents the data structure for the tg prediction external API request
type TgPredictRequest struct {
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

// HdlPredictRequest represents the data structure for the hdl prediction external API request
type HdlPredictRequest struct {
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
