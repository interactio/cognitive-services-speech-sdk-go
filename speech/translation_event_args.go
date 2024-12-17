package speech

import (
	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
)

// #include <stdlib.h>
// #include <speechapi_c_recognizer.h>
import "C"

type TranslationRecognitionEventArgs struct {
	RecognitionEventArgs
	handle C.SPXHANDLE
	Result TranslationRecognitionResult
}

func (event TranslationRecognitionEventArgs) Close() {
	event.RecognitionEventArgs.Close()
	event.Result.Close()
}

func NewTranslationRecognitionEventArgsFromHandle(handle common.SPXHandle) (*TranslationRecognitionEventArgs, error) {
	base, err := NewRecognitionEventArgsFromHandle(handle)
	if err != nil {
		return nil, err
	}
	event := new(TranslationRecognitionEventArgs)
	event.RecognitionEventArgs = *base
	event.handle = uintptr2handle(handle)
	var resultHandle C.SPXHANDLE
	ret := uintptr(C.recognizer_recognition_event_get_result(event.handle, &resultHandle))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result, err := NewTranslationRecognitionResultResultFromHandle(handle2uintptr(resultHandle))
	if err != nil {
		return nil, err
	}
	event.Result = *result
	return event, nil
}

type TranslationRecognitionEventHandler func(event TranslationRecognitionEventArgs)
