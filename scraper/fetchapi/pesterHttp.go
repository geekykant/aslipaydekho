package fetchapi

import (
	"sync"

	"github.com/geekykant/aslipaydekho/scraper/config"
	"github.com/sethgrid/pester"
)

var lock = &sync.Mutex{}
var pesterInstance *pester.Client

func GetPesterCientInstance() *pester.Client {
	if pesterInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if pesterInstance == nil {
			pesterInstance = pester.New()
			config.SetHttpReqConfig(pesterInstance)
		}
	}
	return pesterInstance
}
