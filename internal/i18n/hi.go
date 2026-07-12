package i18n

var hiMessages = map[string]string{
	ReaderStatusWaitingTerminalSize: "टर्मिनल आकार की प्रतीक्षा हो रही है",
	ReaderStatusTerminalTooSmall:    "टर्मिनल विंडो बहुत छोटी है",
	ReaderStatusLastPage:            "अंतिम पृष्ठ",
	ReaderStatusFirstPage:           "पहला पृष्ठ",
	ReaderStatusRenderError:         "रेंडरिंग त्रुटि: %v",
	ReaderStatusMaximumZoom:         "अधिकतम ज़ूम",
	ReaderStatusMinimumZoom:         "न्यूनतम ज़ूम",

	ReaderViewTerminalTooSmall: "comicread: टर्मिनल विंडो बहुत छोटी है",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "पृष्ठ %d/%d",
	ReaderViewPageRange:        "पृष्ठ %d-%d/%d",
	ReaderViewRendering:        "रेंडर हो रहा है",
	ReaderViewHelp: `कुंजियाँ

← →  पिछला / अगला पृष्ठ
↑ ↓  ज़ूम किए गए पृष्ठ को स्क्रॉल करें
+ -  ज़ूम बढ़ाएँ / घटाएँ
q    बाहर निकलें

?    सहायता बंद करें`,

	FilepickerHeader:      "comicread — अध्याय चुनें\n%s\n\n",
	FilepickerNoEntries:   "  (कोई समर्थित प्रविष्टि नहीं)\n",
	FilepickerHelp:        "\n↑/↓ चलें  |  ← मूल निर्देशिका  |  → निर्देशिका में जाएँ  |  enter फ़ाइल खोलें  |  s चुनी हुई निर्देशिका चुनें  |  q बाहर निकलें\n",
	FilepickerWindowTitle: "comicread — फ़ाइल चुनें",

	FilepickerErrResolveDir: "निर्देशिका %q निर्धारित नहीं की जा सकती: %w",
	FilepickerErrReadDir:    "निर्देशिका %q पढ़ी नहीं जा सकती: %w",
	FilepickerErrRunPicker:  "फ़ाइल चयनक चलाने में त्रुटि: %w",

	LoadingViewOpening:     "%s खोला जा रहा है…",
	LoadingViewWindowTitle: "comicread — खोल रहा है",

	CLIErrGetWorkingDir:             "कार्य निर्देशिका प्राप्त नहीं की जा सकती: %w",
	CLIErrPickFile:                  "फ़ाइल नहीं चुनी जा सकती: %w",
	CLIErrRunTUI:                    "TUI चलाने में त्रुटि: %w",
	CLIErrParseArgs:                 "तर्कों को पार्स करने में त्रुटि: %w",
	CLIErrOpenChapter:               "अध्याय नहीं खोला जा सकता: %w",
	CLIErrNoPages:                   "अध्याय में कोई पढ़ने योग्य छवि पृष्ठ नहीं है",
	CLIErrInspectInput:              "इनपुट %q की जाँच नहीं की जा सकती: %w",
	CLIErrUnsupportedFile:           "असमर्थित फ़ाइल %q: समर्थित प्रारूप CBZ, PDF, EPUB या छवियों की निर्देशिका हैं",
	CLIFlagGraphicsUsage:            "रेंडरर: auto, ascii, dots, kitty, sixel या iterm2",
	CLIFlagVersionUsage:             "संस्करण दिखाएँ और बाहर निकलें",
	CLIFlagEnvUsage:                 "comicread परिवेश दिखाएँ और बाहर निकलें",
	CLIFlagBookViewUsage:            "पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ",
	CLIFlagRightBookViewUsage:       "पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ",
	CLIFlagCircleBookViewUsage:      "ओवरलैप होते पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ",
	CLIFlagRightCircleBookViewUsage: "ओवरलैप होते पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ",
	CLIErrMultipleBookViews:         "केवल एक पुस्तक दृश्य विकल्प इस्तेमाल किया जा सकता है",
	CLIErrInvalidView:               "असमर्थित COMICREAD_VIEW मान %q (मान्य: book-view, right-view, circle-view या right-circle-view)",
	CLIHelpHint:                     "सहायता के लिए 'comicread --help' चलाएँ",
	CLIUsage:                        "उपयोग: comicread [विकल्प] [फ़ाइल]",
	CLIUsageFull: `comicread — टर्मिनल के लिए एक न्यूनतम मंगा रीडर

उपयोग: comicread [विकल्प] [फ़ाइल]

विकल्प:
  --graphics string   रेंडरर: auto, ascii, dots, kitty, sixel या iterm2 (डिफ़ॉल्ट "auto")
  --book-view         पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ
  --right-view        पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ
  --circle-view       ओवरलैप होते पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ
  --right-circle-view
                      ओवरलैप होते पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ
  --env               comicread परिवेश दिखाएँ और बाहर निकलें
  -v, --version       संस्करण दिखाएँ और बाहर निकलें
  -h, --help          यह सहायता दिखाएँ

यदि कोई फ़ाइल या निर्देशिका नहीं दी गई है, तो वर्तमान निर्देशिका में इंटरैक्टिव फ़ाइल चयनक खुलेगा।

परिवेश:
  COMICREAD_GRAPHICS  डिफ़ॉल्ट रेंडरर: auto, ascii, dots, kitty, sixel या iterm2
  COMICREAD_VIEW      डिफ़ॉल्ट दृश्य: book-view, right-view, circle-view या right-circle-view
  COMICREAD_LANG      संदेश की भाषा: en, uk, pl, de, fr, es, cs, ro, it, ko, ja, id, hi, el या tr (डिफ़ॉल्ट "en")`,
}
