package handler

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Vefo1/Kvant_practice/internal/interfaces"
	"github.com/Vefo1/Kvant_practice/internal/models"
	"github.com/Vefo1/Kvant_practice/pkg/logger"
)

// Handler handles HTTP requests
type Handler struct {
	predictService interfaces.PredictService
	logger         *logger.Logger
}

// NewHandler creates a new Handler instance
func NewHandler(service interfaces.PredictService, log *logger.Logger) *Handler {
	return &Handler{
		predictService: service,
		logger:         log,
	}
}

// Helper function to safely parse float64 from query parameters
func (h *Handler) parseFloat(paramName string, query url.Values) (float64, error) {
	valStr := query.Get(paramName)
	if valStr == "" {
		return 0.0, fmt.Errorf("parameter '%s' is missing", paramName)
	}
	val, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		return 0.0, fmt.Errorf("error parsing parameter '%s': %w", paramName, err)
	}
	return val, nil
}

// Helper function to safely parse int from query parameters
func (h *Handler) parseInt(paramName string, query url.Values) (int, error) {
	valStr := query.Get(paramName)
	if valStr == "" {
		return 0, fmt.Errorf("parameter '%s' is missing", paramName)
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, fmt.Errorf("error parsing parameter '%s': %w", paramName, err)
	}
	return val, nil
}

// handlePredictionRequest is a generic helper to process prediction requests
// Changed dataCreator func signature to accept url.Values
func (h *Handler) handlePredictionRequest(w http.ResponseWriter, r *http.Request, predictFunc func(data interface{}) ([]byte, int, error), dataCreator func(query url.Values) (interface{}, error)) {
	if r.Method != http.MethodGet {
		h.logger.Warn("Method Not Allowed: %s", r.Method)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query() // query is of type url.Values

	reqData, err := dataCreator(query) // Pass url.Values to dataCreator
	if err != nil {
		h.logger.Error("Bad Request: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respBody, statusCode, err := predictFunc(reqData)
	if err != nil {
		h.logger.Error("Service prediction failed: %v", err)
		if statusCode == 0 {
			statusCode = http.StatusInternalServerError
		}
		http.Error(w, "Internal Server Error", statusCode)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(respBody)
	h.logger.Info("Request successfully proxied and response sent. Status: %d", statusCode)
}

// HBA1CPredictHandler handles the GET /predict/hba1c request
func (h *Handler) HBA1CPredictHandler(w http.ResponseWriter, r *http.Request) {
	h.handlePredictionRequest(w, r,
		func(data interface{}) ([]byte, int, error) {
			return h.predictService.PredictHBA1C(data.(models.HBA1CPredictRequest))
		},
		func(query url.Values) (interface{}, error) {
			age, err := h.parseInt("age", query)
			if err != nil {
				return nil, err
			}
			gender, err := h.parseInt("gender", query)
			if err != nil {
				return nil, err
			}
			rdw, err := h.parseFloat("rdw", query)
			if err != nil {
				return nil, err
			}
			wbc, err := h.parseFloat("wbc", query)
			if err != nil {
				return nil, err
			}
			rbc, err := h.parseFloat("rbc", query)
			if err != nil {
				return nil, err
			}
			hgb, err := h.parseFloat("hgb", query)
			if err != nil {
				return nil, err
			}
			hct, err := h.parseFloat("hct", query)
			if err != nil {
				return nil, err
			}
			mcv, err := h.parseFloat("mcv", query)
			if err != nil {
				return nil, err
			}
			mch, err := h.parseFloat("mch", query)
			if err != nil {
				return nil, err
			}
			mchc, err := h.parseFloat("mchc", query)
			if err != nil {
				return nil, err
			}
			plt, err := h.parseFloat("plt", query)
			if err != nil {
				return nil, err
			}
			neu, err := h.parseFloat("neu", query)
			if err != nil {
				return nil, err
			}
			eos, err := h.parseFloat("eos", query)
			if err != nil {
				return nil, err
			}
			bas, err := h.parseFloat("bas", query)
			if err != nil {
				return nil, err
			}
			lym, err := h.parseFloat("lym", query)
			if err != nil {
				return nil, err
			}
			mon, err := h.parseFloat("mon", query)
			if err != nil {
				return nil, err
			}
			soe, err := h.parseFloat("soe", query)
			if err != nil {
				return nil, err
			}
			chol, err := h.parseFloat("chol", query)
			if err != nil {
				return nil, err
			}
			glu, err := h.parseFloat("glu", query)
			if err != nil {
				return nil, err
			}

			return models.HBA1CPredictRequest{
				UID: "web-client", Age: age, Gender: gender, RDW: rdw, WBC: wbc, RBC: rbc, HGB: hgb, HCT: hct, MCV: mcv, MCH: mch, MCHC: mchc, PLT: plt, NEU: neu, EOS: eos, BAS: bas, LYM: lym, MON: mon, SOE: soe, CHOL: chol, GLU: glu,
			}, nil
		})
}

// LdllPredictHandler handles the GET /predict/ldll request
func (h *Handler) LdllPredictHandler(w http.ResponseWriter, r *http.Request) {
	h.handlePredictionRequest(w, r,
		func(data interface{}) ([]byte, int, error) {
			return h.predictService.PredictLdll(data.(models.LdllPredictRequest))
		},
		func(query url.Values) (interface{}, error) {
			age, err := h.parseInt("age", query)
			if err != nil {
				return nil, err
			}
			gender, err := h.parseInt("gender", query)
			if err != nil {
				return nil, err
			}
			chol, err := h.parseFloat("chol", query)
			if err != nil {
				return nil, err
			}
			hdl, err := h.parseFloat("hdl", query)
			if err != nil {
				return nil, err
			}
			tg, err := h.parseFloat("tg", query)
			if err != nil {
				return nil, err
			}

			return models.LdllPredictRequest{
				UID: "web-client", Age: age, Gender: gender, CHOL: chol, HDL: hdl, TG: tg,
			}, nil
		})
}

// FerrPredictHandler handles the GET /predict/ferr request
func (h *Handler) FerrPredictHandler(w http.ResponseWriter, r *http.Request) {
	h.handlePredictionRequest(w, r,
		func(data interface{}) ([]byte, int, error) {
			return h.predictService.PredictFerr(data.(models.FerrPredictRequest))
		},
		func(query url.Values) (interface{}, error) {
			age, err := h.parseInt("age", query)
			if err != nil {
				return nil, err
			}
			gender, err := h.parseInt("gender", query)
			if err != nil {
				return nil, err
			}
			rdw, err := h.parseFloat("rdw", query)
			if err != nil {
				return nil, err
			}
			wbc, err := h.parseFloat("wbc", query)
			if err != nil {
				return nil, err
			}
			rbc, err := h.parseFloat("rbc", query)
			if err != nil {
				return nil, err
			}
			hgb, err := h.parseFloat("hgb", query)
			if err != nil {
				return nil, err
			}
			hct, err := h.parseFloat("hct", query)
			if err != nil {
				return nil, err
			}
			mcv, err := h.parseFloat("mcv", query)
			if err != nil {
				return nil, err
			}
			mch, err := h.parseFloat("mch", query)
			if err != nil {
				return nil, err
			}
			mchc, err := h.parseFloat("mchc", query)
			if err != nil {
				return nil, err
			}
			plt, err := h.parseFloat("plt", query)
			if err != nil {
				return nil, err
			}
			neu, err := h.parseFloat("neu", query)
			if err != nil {
				return nil, err
			}
			eos, err := h.parseFloat("eos", query)
			if err != nil {
				return nil, err
			}
			bas, err := h.parseFloat("bas", query)
			if err != nil {
				return nil, err
			}
			lym, err := h.parseFloat("lym", query)
			if err != nil {
				return nil, err
			}
			mon, err := h.parseFloat("mon", query)
			if err != nil {
				return nil, err
			}
			soe, err := h.parseFloat("soe", query)
			if err != nil {
				return nil, err
			}
			crp, err := h.parseFloat("crp", query) // Specific to Ferr
			if err != nil {
				return nil, err
			}

			return models.FerrPredictRequest{
				UID: "web-client", Age: age, Gender: gender, RDW: rdw, WBC: wbc, RBC: rbc, HGB: hgb, HCT: hct, MCV: mcv, MCH: mch, MCHC: mchc, PLT: plt, NEU: neu, EOS: eos, BAS: bas, LYM: lym, MON: mon, SOE: soe, CRP: crp,
			}, nil
		})
}

// LdlPredictHandler handles the GET /predict/ldl request
func (h *Handler) LdlPredictHandler(w http.ResponseWriter, r *http.Request) {
	h.handlePredictionRequest(w, r,
		func(data interface{}) ([]byte, int, error) {
			return h.predictService.PredictLdl(data.(models.LdlPredictRequest))
		},
		func(query url.Values) (interface{}, error) {
			age, err := h.parseInt("age", query)
			if err != nil {
				return nil, err
			}
			gender, err := h.parseInt("gender", query)
			if err != nil {
				return nil, err
			}
			rdw, err := h.parseFloat("rdw", query)
			if err != nil {
				return nil, err
			}
			wbc, err := h.parseFloat("wbc", query)
			if err != nil {
				return nil, err
			}
			rbc, err := h.parseFloat("rbc", query)
			if err != nil {
				return nil, err
			}
			hgb, err := h.parseFloat("hgb", query)
			if err != nil {
				return nil, err
			}
			hct, err := h.parseFloat("hct", query)
			if err != nil {
				return nil, err
			}
			mcv, err := h.parseFloat("mcv", query)
			if err != nil {
				return nil, err
			}
			mch, err := h.parseFloat("mch", query)
			if err != nil {
				return nil, err
			}
			mchc, err := h.parseFloat("mchc", query)
			if err != nil {
				return nil, err
			}
			plt, err := h.parseFloat("plt", query)
			if err != nil {
				return nil, err
			}
			neu, err := h.parseFloat("neu", query)
			if err != nil {
				return nil, err
			}
			eos, err := h.parseFloat("eos", query)
			if err != nil {
				return nil, err
			}
			bas, err := h.parseFloat("bas", query)
			if err != nil {
				return nil, err
			}
			lym, err := h.parseFloat("lym", query)
			if err != nil {
				return nil, err
			}
			mon, err := h.parseFloat("mon", query)
			if err != nil {
				return nil, err
			}
			soe, err := h.parseFloat("soe", query)
			if err != nil {
				return nil, err
			}
			chol, err := h.parseFloat("chol", query)
			if err != nil {
				return nil, err
			}
			glu, err := h.parseFloat("glu", query)
			if err != nil {
				return nil, err
			}

			return models.LdlPredictRequest{
				UID: "web-client", Age: age, Gender: gender, RDW: rdw, WBC: wbc, RBC: rbc, HGB: hgb, HCT: hct, MCV: mcv, MCH: mch, MCHC: mchc, PLT: plt, NEU: neu, EOS: eos, BAS: bas, LYM: lym, MON: mon, SOE: soe, CHOL: chol, GLU: glu,
			}, nil
		})
}

// TgPredictHandler handles the GET /predict/tg request
func (h *Handler) TgPredictHandler(w http.ResponseWriter, r *http.Request) {
	h.handlePredictionRequest(w, r,
		func(data interface{}) ([]byte, int, error) {
			return h.predictService.PredictTg(data.(models.TgPredictRequest))
		},
		func(query url.Values) (interface{}, error) {
			age, err := h.parseInt("age", query)
			if err != nil {
				return nil, err
			}
			gender, err := h.parseInt("gender", query)
			if err != nil {
				return nil, err
			}
			rdw, err := h.parseFloat("rdw", query)
			if err != nil {
				return nil, err
			}
			wbc, err := h.parseFloat("wbc", query)
			if err != nil {
				return nil, err
			}
			rbc, err := h.parseFloat("rbc", query)
			if err != nil {
				return nil, err
			}
			hgb, err := h.parseFloat("hgb", query)
			if err != nil {
				return nil, err
			}
			hct, err := h.parseFloat("hct", query)
			if err != nil {
				return nil, err
			}
			mcv, err := h.parseFloat("mcv", query)
			if err != nil {
				return nil, err
			}
			mch, err := h.parseFloat("mch", query)
			if err != nil {
				return nil, err
			}
			mchc, err := h.parseFloat("mchc", query)
			if err != nil {
				return nil, err
			}
			plt, err := h.parseFloat("plt", query)
			if err != nil {
				return nil, err
			}
			neu, err := h.parseFloat("neu", query)
			if err != nil {
				return nil, err
			}
			eos, err := h.parseFloat("eos", query)
			if err != nil {
				return nil, err
			}
			bas, err := h.parseFloat("bas", query)
			if err != nil {
				return nil, err
			}
			lym, err := h.parseFloat("lym", query)
			if err != nil {
				return nil, err
			}
			mon, err := h.parseFloat("mon", query)
			if err != nil {
				return nil, err
			}
			soe, err := h.parseFloat("soe", query)
			if err != nil {
				return nil, err
			}
			chol, err := h.parseFloat("chol", query)
			if err != nil {
				return nil, err
			}
			glu, err := h.parseFloat("glu", query)
			if err != nil {
				return nil, err
			}

			return models.TgPredictRequest{
				UID: "web-client", Age: age, Gender: gender, RDW: rdw, WBC: wbc, RBC: rbc, HGB: hgb, HCT: hct, MCV: mcv, MCH: mch, MCHC: mchc, PLT: plt, NEU: neu, EOS: eos, BAS: bas, LYM: lym, MON: mon, SOE: soe, CHOL: chol, GLU: glu,
			}, nil
		})
}

// HdlPredictHandler handles the GET /predict/hdl request
func (h *Handler) HdlPredictHandler(w http.ResponseWriter, r *http.Request) {
	h.handlePredictionRequest(w, r,
		func(data interface{}) ([]byte, int, error) {
			return h.predictService.PredictHdl(data.(models.HdlPredictRequest))
		},
		func(query url.Values) (interface{}, error) {
			age, err := h.parseInt("age", query)
			if err != nil {
				return nil, err
			}
			gender, err := h.parseInt("gender", query)
			if err != nil {
				return nil, err
			}
			rdw, err := h.parseFloat("rdw", query)
			if err != nil {
				return nil, err
			}
			wbc, err := h.parseFloat("wbc", query)
			if err != nil {
				return nil, err
			}
			rbc, err := h.parseFloat("rbc", query)
			if err != nil {
				return nil, err
			}
			hgb, err := h.parseFloat("hgb", query)
			if err != nil {
				return nil, err
			}
			hct, err := h.parseFloat("hct", query)
			if err != nil {
				return nil, err
			}
			mcv, err := h.parseFloat("mcv", query)
			if err != nil {
				return nil, err
			}
			mch, err := h.parseFloat("mch", query)
			if err != nil {
				return nil, err
			}
			mchc, err := h.parseFloat("mchc", query)
			if err != nil {
				return nil, err
			}
			plt, err := h.parseFloat("plt", query)
			if err != nil {
				return nil, err
			}
			neu, err := h.parseFloat("neu", query)
			if err != nil {
				return nil, err
			}
			eos, err := h.parseFloat("eos", query)
			if err != nil {
				return nil, err
			}
			bas, err := h.parseFloat("bas", query)
			if err != nil {
				return nil, err
			}
			lym, err := h.parseFloat("lym", query)
			if err != nil {
				return nil, err
			}
			mon, err := h.parseFloat("mon", query)
			if err != nil {
				return nil, err
			}
			soe, err := h.parseFloat("soe", query)
			if err != nil {
				return nil, err
			}
			chol, err := h.parseFloat("chol", query)
			if err != nil {
				return nil, err
			}
			glu, err := h.parseFloat("glu", query)
			if err != nil {
				return nil, err
			}

			return models.HdlPredictRequest{
				UID: "web-client", Age: age, Gender: gender, RDW: rdw, WBC: wbc, RBC: rbc, HGB: hgb, HCT: hct, MCV: mcv, MCH: mch, MCHC: mchc, PLT: plt, NEU: neu, EOS: eos, BAS: bas, LYM: lym, MON: mon, SOE: soe, CHOL: chol, GLU: glu,
			}, nil
		})
}
