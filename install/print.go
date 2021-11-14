package install

import (
	"encoding/json"
	"strings"

	"github.com/wonderivan/logger"
)

//Print is
func (s *DmosInstaller) Print(process ...string) {
	if len(process) == 0 {
		configJson, _ := json.Marshal(s)
		logger.Info("\n[globals]Dmos config is: ", string(configJson))
	} else {
		var sb strings.Builder
		for _, v := range process {
			sb.Write([]byte("==>"))
			sb.Write([]byte(v))
		}
		logger.Debug(sb.String())
	}

}
func (s *DmosInstaller) PrintFinish() {
	logger.Info("Dmos install success.")
}
