package routes

import (
	"net/http"
)

func GetLoaInsuranceCardif(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("cPulseInsuranceCardifCallbackExport"))
}
