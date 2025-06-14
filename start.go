package loadbby

import (
	"net/http"

	"github.com/gpr3211/nthropy/memory"
)

type EntrhopyService struct {
	Reg    *memory.Registry
	Server *http.Server
}

func NewEnthService() *EntrhopyService { return &EntrhopyService{} }
