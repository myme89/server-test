package ultils

import (
	"encoding/json"
	"net/http"
	"server-test/server-gateway/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, resp map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

func ErrorHandlerGRPC(code codes.Code, message string, codeHttp int) error {
	st := status.New(code, message)

	detail := &pb.ErrorInfo{
		Status:   "error",
		CodeHttp: int64(codeHttp),
	}
	st, _ = st.WithDetails(detail)

	return st.Err()
}
