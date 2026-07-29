package i18n

var hiMessages = map[string]string{
	CLIFlagConfigUsage: "उपयोग की जाने वाली कॉन्फ़िगरेशन फ़ाइल", CLIFlagResetConfigUsage: "config.toml को डिफ़ॉल्ट पर रीसेट करें और बाहर निकलें", CLIFlagSetConfigUsage: "config.toml अपडेट करें: कुंजी=मान",
	ReaderStatusWaitingTerminalSize: "टर्मिनल आकार की प्रतीक्षा हो रही है",
	ReaderStatusTerminalTooSmall:    "टर्मिनल विंडो बहुत छोटी है",
	ReaderStatusLastPage:            "अंतिम पृष्ठ",
	ReaderStatusFirstPage:           "पहला पृष्ठ",
	ReaderStatusRenderError:         "रेंडरिंग त्रुटि: %v",
	ReaderStatusMaximumZoom:         "अधिकतम ज़ूम",
	ReaderStatusMinimumZoom:         "न्यूनतम ज़ूम",
	ReaderStatusInvalidPage:         "अमान्य पृष्ठ संख्या",

	ReaderViewTerminalTooSmall: "comicread: टर्मिनल विंडो बहुत छोटी है",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "पृष्ठ %d/%d",
	ReaderViewPageRange:        "पृष्ठ %d-%d/%d",
	ReaderViewRendering:        "रेंडर हो रहा है",
	ReaderViewBookmarks:        "बुकमार्क",
	ReaderViewNoBookmarks:      "(कोई बुकमार्क नहीं)",
	ReaderViewBookmarksHelp:    "ऊपर/नीचे चलें | enter खोलें | esc बंद करें",
	ReaderViewGoToPage:         "पृष्ठ पर जाएँ: %s",
	ReaderViewHelp: `कुंजियाँ

← →  पिछला / अगला पृष्ठ
↑ ↓  ज़ूम किए गए पृष्ठ को स्क्रॉल करें
+ -  ज़ूम बढ़ाएँ / घटाएँ
b    बुकमार्क जोड़ें / हटाएँ
v ← → पिछला / अगला बुकमार्क
c v  बुकमार्क
g123 enter  पृष्ठ पर जाएँ
q    बाहर निकलें

?    सहायता बंद करें`,

	FilepickerHeader:         "comicread — अध्याय चुनें\n%s\n\n",
	FilepickerNoEntries:      "  (कोई समर्थित प्रविष्टि नहीं)\n",
	FilepickerHelp:           "\n↑/↓ चलें\n← मूल निर्देशिका\n→ निर्देशिका में जाएँ\nenter फ़ाइल खोलें\ns चुनी हुई निर्देशिका चुनें\nf वर्तमान निर्देशिका को पसंदीदा में जोड़ें / हटाएँ\nF पसंदीदा निर्देशिका जोड़ें\nb पसंदीदा निर्देशिकाएँ\no निर्देशिका पर जाएँ\nq बाहर निकलें\n",
	FilepickerWindowTitle:    "comicread — फ़ाइल चुनें",
	FilepickerGoToPrompt:     "\nनिर्देशिका पर जाएँ: %s\n",
	FilepickerFavoritePrompt: "\nपसंदीदा निर्देशिका: %s\n",
	FilepickerFavorites:      "पसंदीदा निर्देशिकाएँ\n\n",
	FilepickerNoFavorites:    "  (कोई पसंदीदा निर्देशिका कॉन्फ़िगर नहीं है)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ चलें\nenter निर्देशिका पर जाएँ\nd पसंदीदा हटाएँ\nesc वापस जाएँ\n",
	FilepickerFavoriteErr:    "  पसंदीदा सहेजने में त्रुटि: %s\n",
	FilepickerGoToErr:        "  त्रुटि: %s\n",

	FilepickerErrResolveDir: "निर्देशिका %q निर्धारित नहीं की जा सकती: %w",
	FilepickerErrReadDir:    "निर्देशिका %q पढ़ी नहीं जा सकती: %w",
	FilepickerErrRunPicker:  "फ़ाइल चयनक चलाने में त्रुटि: %w",
	FilepickerErrEmptyPath:  "पथ खाली है",
	FilepickerErrNotDir:     "%q एक निर्देशिका नहीं है",

	LoadingViewOpening:     "%s खोला जा रहा है…",
	LoadingViewWindowTitle: "comicread — खोल रहा है",

	CLIErrGetWorkingDir:             "कार्य निर्देशिका प्राप्त नहीं की जा सकती: %w",
	CLIErrPickFile:                  "फ़ाइल नहीं चुनी जा सकती: %w",
	CLIErrRunTUI:                    "TUI चलाने में त्रुटि: %w",
	CLIErrParseArgs:                 "तर्कों को पार्स करने में त्रुटि: %w",
	CLIErrOpenChapter:               "अध्याय नहीं खोला जा सकता: %w",
	CLIErrOpenJournal:               "जर्नल नहीं खोला जा सकता: %w",
	CLIErrClearJournal:              "जर्नल साफ़ नहीं किया जा सकता: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal के लिए फ़ाइल या निर्देशिका आवश्यक है",
	CLIErrNoPages:                   "अध्याय में कोई पढ़ने योग्य छवि पृष्ठ नहीं है",
	CLIErrInspectInput:              "इनपुट %q की जाँच नहीं की जा सकती: %w",
	CLIErrUnsupportedFile:           "असमर्थित फ़ाइल %q: समर्थित प्रारूप CBZ, PDF, EPUB या छवियों की निर्देशिका हैं",
	CLIFlagGraphicsUsage:            "रेंडरर: auto, ascii, dots, kitty, sixel या iterm2",
	CLIFlagVersionUsage:             "संस्करण दिखाएँ और बाहर निकलें",
	CLIFlagUpdateUsage:              "अपडेट जाँचें और बाहर निकलें",
	CLIFlagEnvUsage:                 "comicread परिवेश दिखाएँ और बाहर निकलें",
	CLIFlagClearJournalUsage:        "फ़ाइल या निर्देशिका का स्थानीय जर्नल हटाएँ और बाहर निकलें",
	CLIFlagBookViewUsage:            "पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ",
	CLIFlagRightBookViewUsage:       "पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ",
	CLIFlagCircleBookViewUsage:      "ओवरलैप होते पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ",
	CLIFlagRightCircleBookViewUsage: "ओवरलैप होते पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ",
	CLIErrMultipleBookViews:         "केवल एक पुस्तक दृश्य विकल्प इस्तेमाल किया जा सकता है",
	CLIErrInvalidView:               "असमर्थित COMICREAD_VIEW मान %q (मान्य: book-view, right-view, circle-view या right-circle-view)",
	CLIFlagOpenUsage:                "फ़ाइल चयनक में खोलने की निर्देशिका (डिफ़ॉल्ट: COMICREAD_DIR या वर्तमान निर्देशिका)",
	CLIErrOpenNotDir:                "निर्देशिका खोलें %q: निर्देशिका नहीं है",
	CLIFlagWebUsage:                 "टर्मिनल UI के बजाय स्थानीय वेब रीडर शुरू करें",
	CLIErrWebArgs:                   "--web फ़ाइल या निर्देशिका आर्ग्युमेंट स्वीकार नहीं करता",
	WebServerStarted:                "comicread वेब रीडर %s पर चल रहा है (रोकने के लिए Ctrl+C दबाएँ)",
	WebErrListen:                    "वेब सर्वर शुरू करें: %w",
	WebErrServe:                     "वेब सर्वर चलाएँ: %w",
	CLIHelpHint:                     "सहायता के लिए 'comicread --help' चलाएँ",
	CLIUsage:                        "उपयोग: comicread [विकल्प] [फ़ाइल]",
	CLIUsageFull: `comicread — टर्मिनल के लिए एक न्यूनतम मंगा रीडर

उपयोग: comicread [विकल्प] [फ़ाइल]

विकल्प:
  --config string     उपयोग की जाने वाली कॉन्फ़िगरेशन फ़ाइल
  --graphics string   रेंडरर: auto, ascii, dots, kitty, sixel या iterm2 (डिफ़ॉल्ट "auto")
  --book-view         पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ
  --right-view        पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ
  --circle-view       ओवरलैप होते पृष्ठों के जोड़े बाएँ से दाएँ दिखाएँ
  --right-circle-view
                      ओवरलैप होते पृष्ठों के जोड़े दाएँ से बाएँ दिखाएँ
  --clear-journal    फ़ाइल या निर्देशिका का स्थानीय जर्नल हटाएँ और बाहर निकलें
  --reset-config     config.toml को डिफ़ॉल्ट पर रीसेट करें और बाहर निकलें
  --set-config value config.toml अपडेट करें: कुंजी=मान
  -o, --open string   फ़ाइल चयनक में खोलने की निर्देशिका (डिफ़ॉल्ट: COMICREAD_DIR या वर्तमान निर्देशिका)
  --update            अपडेट जाँचें और बाहर निकलें
  --web               टर्मिनल UI के बजाय स्थानीय वेब रीडर शुरू करें
  -v, --version       संस्करण दिखाएँ और बाहर निकलें
  -h, --help          यह सहायता दिखाएँ

यदि कोई फ़ाइल या निर्देशिका नहीं दी गई है, तो COMICREAD_DIR में इंटरैक्टिव फ़ाइल चयनक खुलेगा
(यदि यह मान्य निर्देशिका पर सेट है), अन्यथा वर्तमान निर्देशिका में खुलेगा।`,
	ReaderViewMetadata: "मेटाडेटा",
}
