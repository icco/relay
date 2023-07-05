package lib

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/icco/gutil/logging"
	"go.uber.org/zap"
)

var (
	funcs = []DataParseFunc{
		jsonToSonarr,
		jsonToLidarr,
		jsonToGoogleCloud,
		jsonToGoogleCloudBuild,
		jsonToPlex,
		jsonToInflux,
		jsonToGitLabPush,
	}

	log = logging.Must(logging.NewLogger("relay"))
)

// DataParseFunc is the generic func type for parsing input.
type DataParseFunc func([]byte) DataType

// DataType is the interface all parsed data must match.
type DataType interface {
	Valid() bool
	Message() string
}

// BufferToMessage takes in a message buffer and returns a message string.
func BufferToMessage(buf []byte) string {

	var msg string
	log.Debugw("attempting to parse", "body", string(buf))

	for _, f := range funcs {
		if data := f(buf); data != nil {
			if data.Valid() {
				msg += data.Message()
			}
		}
	}

	if msg == "" {
		var f map[string]string
		if err := json.Unmarshal(buf, &f); err != nil {
			log.Infow("decoding json to map", zap.Error(err))
			return ""
		}

		var keys []string
		for k := range f {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			msg += fmt.Sprintf("%s: %s\n", k, f[k])
		}
	}

	return msg
}
