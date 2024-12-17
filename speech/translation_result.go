package speech

import (
	"time"
	"unsafe"

	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
)

// #include <stdlib.h>
// #include <stdint.h>
// #include <speechapi_c_common.h>
// #include <speechapi_c_result.h>
// #include <speechapi_c_recognizer.h>
// #include <speechapi_c_translation_result.h>
//
import "C"

type TranslationRecognitionResult struct {
	handle      C.SPXHANDLE
	ResultID    string
	Reason      common.ResultReason
	Text        string
	Language    string
	Translation string
	Duration    time.Duration
	Offset      time.Duration
	Properties  *common.PropertyCollection
}

func (result TranslationRecognitionResult) Close() {
	result.Properties.Close()
	C.recognizer_result_handle_release(result.handle)
}

func NewTranslationRecognitionResultResultFromHandle(handle common.SPXHandle) (*TranslationRecognitionResult, error) {
	buffer := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(buffer))
	result := new(TranslationRecognitionResult)
	result.handle = uintptr2handle(handle)
	/* ResultID */
	ret := uintptr(C.result_get_result_id(result.handle, (*C.char)(buffer), 1024))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result.ResultID = C.GoString((*C.char)(buffer))
	/* Reason */
	var cReason C.Result_Reason
	ret = uintptr(C.result_get_reason(result.handle, &cReason))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result.Reason = (common.ResultReason)(cReason)
	/* Text */
	ret = uintptr(C.result_get_text(result.handle, (*C.char)(buffer), 1024))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result.Text = C.GoString((*C.char)(buffer))
	// Translation
	var languageSize C.size_t
	var textSize C.size_t
	ret = uintptr(C.translation_text_result_get_translation(result.handle, 0, nil, nil, &languageSize, &textSize))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	language := C.malloc(C.sizeof_char * languageSize)
	defer C.free(unsafe.Pointer(language))
	text := C.malloc(C.sizeof_char * textSize)
	defer C.free(unsafe.Pointer(text))
	ret = uintptr(C.translation_text_result_get_translation(result.handle, 0, (*C.char)(language), (*C.char)(text), &languageSize, &textSize))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result.Language = C.GoString((*C.char)(language))
	result.Translation = C.GoString((*C.char)(text))
	/* Duration */
	var cDuration C.uint64_t
	ret = uintptr(C.result_get_duration(result.handle, &cDuration))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result.Duration = time.Nanosecond * time.Duration(100*cDuration)
	/* Offset */
	var cOffset C.uint64_t
	ret = uintptr(C.result_get_offset(result.handle, &cOffset))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result.Offset = time.Nanosecond * time.Duration(100*cOffset)
	/* Properties */
	var propBagHandle C.SPXHANDLE
	ret = uintptr(C.result_get_property_bag(uintptr2handle(handle), &propBagHandle))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	result.Properties = common.NewPropertyCollectionFromHandle(handle2uintptr(propBagHandle))
	return result, nil
}
