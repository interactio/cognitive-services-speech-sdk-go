// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See LICENSE.md file in the project root for full license information.

package speech

import (
	"strings"

	"github.com/Microsoft/cognitive-services-speech-sdk-go/common"
)

// #include <stdlib.h>
// #include <stdint.h>
// #include <speechapi_c_common.h>
// #include <speechapi_c_result.h>
// #include <speechapi_c_property_bag.h>
//
import "C"

var voices = map[string]string{
	"af-ZA":  "af-ZA-AdriNeural",
	"sq-AL":  "sq-AL-AnilaNeural",
	"am-ET":  "am-ET-MekdesNeural",
	"ar-EG":  "ar-EG-SalmaNeural",
	"hy-AM":  "hy-AM-AnahitNeural",
	"as-IN":  "as-IN-YashicaNeural",
	"az-AZ":  "az-AZ-BanuNeural",
	"eu-ES":  "eu-ES-AinhoaNeural",
	"bn-IN":  "bn-IN-TanishaaNeural",
	"bs-BA":  "bs-BA-VesnaNeural",
	"bg-BG":  "bg-BG-KalinaNeural",
	"my-MM":  "my-MM-NilarNeural",
	"ca-ES":  "ca-ES-JoanaNeural",
	"zh-CN":  "zh-CN-guangxi-YunqiNeural",
	"hr-HR":  "hr-HR-GabrijelaNeural",
	"cs-CZ":  "cs-CZ-VlastaNeural",
	"da-DK":  "da-DK-ChristelNeural",
	"nl-NL":  "nl-NL-FennaNeural",
	"en-US":  "en-US-AvaMultilingualNeural",
	"et-EE":  "et-EE-AnuNeural",
	"fil-PH": "fil-PH-BlessicaNeural",
	"fi-FI":  "fi-FI-SelmaNeural",
	"fr-FR":  "fr-FR-VivienneMultilingualNeural",
	"gl-ES":  "gl-ES-SabelaNeural",
	"ka-GE":  "ka-GE-EkaNeural",
	"de-DE":  "de-DE-SeraphinaMultilingualNeural",
	"el-GR":  "el-GR-AthinaNeural",
	"gu-IN":  "gu-IN-DhwaniNeural",
	"he-IL":  "he-IL-HilaNeural",
	"hi-IN":  "hi-IN-AaravNeural",
	"hu-HU":  "hu-HU-NoemiNeural",
	"is-IS":  "is-IS-GudrunNeural",
	"id-ID":  "id-ID-GadisNeural",
	"ga-IE":  "ga-IE-OrlaNeural",
	"zu-ZA":  "zu-ZA-ThandoNeural",
	"it-IT":  "it-IT-AlessioMultilingualNeural",
	"ja-JP":  "ja-JP-NanamiNeural",
	"jv-ID":  "jv-ID-SitiNeural",
	"kn-IN":  "kn-IN-SapnaNeural",
	"kk-KZ":  "kk-KZ-AigulNeural",
	"km-KH":  "km-KH-SreymomNeural",
	"sw-KE":  "sw-KE-ZuriNeural",
	"ko-KR":  "ko-KR-HyunsuMultilingualNeural",
	"lo-LA":  "lo-LA-KeomanyNeural",
	"lv-LV":  "lv-LV-EveritaNeural",
	"lt-LT":  "lt-LT-OnaNeural",
	"mk-MK":  "mk-MK-MarijaNeural",
	"ms-MY":  "ms-MY-YasminNeural",
	"ml-IN":  "ml-IN-SobhanaNeural",
	"mt-MT":  "mt-MT-GraceNeural",
	"mr-IN":  "mr-IN-AarohiNeural",
	"mn-MN":  "mn-MN-YesuiNeural",
	"ne-NP":  "ne-NP-HemkalaNeural",
	"nb-NO":  "nb-NO-PernilleNeural",
	"or-IN":  "or-IN-SubhasiniNeural",
	"ps-AF":  "ps-AF-LatifaNeural",
	"fa-IR":  "fa-IR-DilaraNeural",
	"pl-PL":  "pl-PL-AgnieszkaNeural",
	"pt-PT":  "pt-PT-RaquelNeural",
	"pa-IN":  "pa-IN-OjasNeural",
	"ro-RO":  "ro-RO-AlinaNeural",
	"ru-RU":  "ru-RU-SvetlanaNeural",
	"sr-RS":  "sr-RS-SophieNeural",
	"si-LK":  "si-LK-ThiliniNeural",
	"sk-SK":  "sk-SK-ViktoriaNeural",
	"sl-SI":  "sl-SI-PetraNeural",
	"so-SO":  "so-SO-UbaxNeural",
	"es-ES":  "es-ES-ArabellaMultilingualNeural",
	"sv-SE":  "sv-SE-SofieNeural",
	"ta-IN":  "ta-IN-PallaviNeural",
	"te-IN":  "te-IN-ShrutiNeural",
	"th-TH":  "th-TH-PremwadeeNeural",
	"tr-TR":  "tr-TR-EmelNeural",
	"uk-UA":  "uk-UA-PolinaNeural",
	"ur-IN":  "ur-IN-GulNeural",
	"uz-UZ":  "uz-UZ-MadinaNeural",
	"vi-VN":  "vi-VN-HoaiMyNeural",
	"cy-GB":  "cy-GB-NiaNeural",
}

// VoiceInfo contains information about result from voices list of speech synthesizers.
type VoiceInfo struct {
	handle C.SPXHANDLE

	// Name specifies the voice name.
	Name string

	// Locale specifies the locale of the voice
	Locale string

	// ShortName specifies the voice name in short format
	ShortName string

	// LocalName specifies the local name of the voice
	LocalName string

	// Gender specifies the gender of the voice.
	Gender common.SynthesisVoiceGender

	// VoiceType specifies the voice type.
	VoiceType common.SynthesisVoiceType

	// StyleList specifies the styles the voice supports.
	StyleList []string

	// VoicePath specifies the voice path
	VoicePath string

	// Collection of additional properties.
	Properties *common.PropertyCollection
}

// Close releases the underlying resources
func (result VoiceInfo) Close() {
	result.Properties.Close()
	C.voice_info_handle_release(result.handle)
}

// NewVoiceInfoFromHandle creates a VoiceInfo from a handle (for internal use)
func NewVoiceInfoFromHandle(handle common.SPXHandle) (*VoiceInfo, error) {
	voiceInfo := new(VoiceInfo)
	voiceInfo.handle = uintptr2handle(handle)
	/* Name */
	value := C.voice_info_get_name(voiceInfo.handle)
	voiceInfo.Name = C.GoString(value)
	C.property_bag_free_string(value)
	/* Locale */
	value = C.voice_info_get_locale(voiceInfo.handle)
	voiceInfo.Locale = C.GoString(value)
	C.property_bag_free_string(value)
	/* ShortName */
	value = C.voice_info_get_short_name(voiceInfo.handle)
	voiceInfo.ShortName = C.GoString(value)
	C.property_bag_free_string(value)
	/* LocalName */
	value = C.voice_info_get_local_name(voiceInfo.handle)
	voiceInfo.LocalName = C.GoString(value)
	C.property_bag_free_string(value)
	/* StyleList */
	value = C.voice_info_get_style_list(voiceInfo.handle)
	voiceInfo.StyleList = strings.Split(C.GoString(value), "|")
	C.property_bag_free_string(value)
	/* VoiceType */
	var cVoiceType C.Synthesis_VoiceType
	ret := uintptr(C.voice_info_get_voice_type(voiceInfo.handle, &cVoiceType))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	voiceInfo.VoiceType = (common.SynthesisVoiceType)(cVoiceType)
	/* VoicePath */
	value = C.voice_info_get_voice_path(voiceInfo.handle)
	voiceInfo.VoicePath = C.GoString(value)
	C.property_bag_free_string(value)
	/* Properties */
	var propBagHandle C.SPXHANDLE
	ret = uintptr(C.voice_info_get_property_bag(uintptr2handle(handle), &propBagHandle))
	if ret != C.SPX_NOERROR {
		return nil, common.NewCarbonError(ret)
	}
	voiceInfo.Properties = common.NewPropertyCollectionFromHandle(handle2uintptr(propBagHandle))
	gender := voiceInfo.Properties.GetPropertyByString("Gender", "")
	if gender == "Female" {
		voiceInfo.Gender = common.Female
	} else if gender == "Male" {
		voiceInfo.Gender = common.Male
	} else {
		voiceInfo.Gender = common.GenderUnknown
	}
	return voiceInfo, nil
}
