package functions

import (
	"gpt-worker/internal/interfaces"
	"gpt-worker/pkg/functions"
)

var AllFunctions = []interfaces.Function{
	functions.NewListDirFunction(),
	functions.NewReadFileFunction(),
	functions.NewCreateFileFunction(),
	functions.NewZipFunction(),
	functions.NewUnzipFunction(),
}
