package interfaces

import "github.com/Vefo1/Kvant_practice/internal/models" // Replace with your actual module name

// PredictService defines the interface for the prediction business logic
type PredictService interface {
	PredictHBA1C(data models.HBA1CPredictRequest) ([]byte, int, error)
	PredictLdll(data models.LdllPredictRequest) ([]byte, int, error)
	PredictFerr(data models.FerrPredictRequest) ([]byte, int, error)
	PredictLdl(data models.LdlPredictRequest) ([]byte, int, error)
	PredictTg(data models.TgPredictRequest) ([]byte, int, error)
	PredictHdl(data models.HdlPredictRequest) ([]byte, int, error)
}
