package i18n

var idMessages = map[string]string{
	CLIFlagConfigUsage: "berkas konfigurasi yang digunakan", CLIFlagResetConfigUsage: "atur ulang config.toml ke bawaan lalu keluar", CLIFlagSetConfigUsage: "perbarui config.toml: kunci=nilai",
	ReaderStatusWaitingTerminalSize: "menunggu ukuran terminal",
	ReaderStatusTerminalTooSmall:    "jendela terminal terlalu kecil",
	ReaderStatusLastPage:            "halaman terakhir",
	ReaderStatusFirstPage:           "halaman pertama",
	ReaderStatusRenderError:         "kesalahan perenderan: %v",
	ReaderStatusMaximumZoom:         "zoom maksimum",
	ReaderStatusMinimumZoom:         "zoom minimum",
	ReaderStatusInvalidPage:         "nomor halaman tidak valid",

	ReaderViewTerminalTooSmall: "comicread: jendela terminal terlalu kecil",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "halaman %d/%d",
	ReaderViewPageRange:        "halaman %d-%d/%d",
	ReaderViewRendering:        "merender",
	ReaderViewBookmarks:        "Bookmark",
	ReaderViewNoBookmarks:      "(tidak ada bookmark)",
	ReaderViewBookmarksHelp:    "atas/bawah pindah | enter buka | esc tutup",
	ReaderViewGoToPage:         "Pergi ke halaman: %s",
	ReaderViewHelp: `Tombol

← →  halaman sebelumnya / berikutnya
↑ ↓  gulir halaman yang diperbesar
+ -  perbesar / perkecil
b    tambah / hapus bookmark
v ← → bookmark sebelumnya / berikutnya
c v  bookmark
g123 enter  pergi ke halaman
q    keluar

?    tutup bantuan`,

	FilepickerHeader:         "comicread — pilih bab\n%s\n\n",
	FilepickerNoEntries:      "  (tidak ada entri yang didukung)\n",
	FilepickerHelp:           "\n↑/↓ pindah\n← direktori induk\n→ masuk direktori\nenter buka berkas\ns pilih direktori yang disorot\nf tambah / hapus direktori saat ini dari favorit\nF tambah direktori favorit\nb direktori favorit\no pergi ke direktori\nq keluar\n",
	FilepickerWindowTitle:    "comicread — pilih berkas",
	FilepickerGoToPrompt:     "\nPergi ke direktori: %s\n",
	FilepickerFavoritePrompt: "\nDirektori favorit: %s\n",
	FilepickerFavorites:      "Direktori favorit\n\n",
	FilepickerNoFavorites:    "  (tidak ada direktori favorit yang dikonfigurasi)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ pindah\nenter pergi ke direktori\nd hapus favorit\nesc kembali\n",
	FilepickerFavoriteErr:    "  kesalahan menyimpan favorit: %s\n",
	FilepickerGoToErr:        "  kesalahan: %s\n",

	FilepickerErrResolveDir: "tidak dapat menentukan direktori %q: %w",
	FilepickerErrReadDir:    "tidak dapat membaca direktori %q: %w",
	FilepickerErrRunPicker:  "kesalahan saat menjalankan pemilih berkas: %w",
	FilepickerErrEmptyPath:  "jalur kosong",
	FilepickerErrNotDir:     "%q bukan direktori",

	LoadingViewOpening:     "membuka %s…",
	LoadingViewWindowTitle: "comicread — membuka",

	CLIErrGetWorkingDir:             "tidak dapat memperoleh direktori kerja: %w",
	CLIErrPickFile:                  "tidak dapat memilih berkas: %w",
	CLIErrRunTUI:                    "kesalahan saat menjalankan TUI: %w",
	CLIErrParseArgs:                 "kesalahan saat mengurai argumen: %w",
	CLIErrOpenChapter:               "tidak dapat membuka bab: %w",
	CLIErrOpenJournal:               "tidak dapat membuka jurnal: %w",
	CLIErrClearJournal:              "tidak dapat menghapus jurnal: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal memerlukan berkas atau direktori",
	CLIErrNoPages:                   "bab tidak berisi halaman gambar yang dapat dibaca",
	CLIErrInspectInput:              "tidak dapat memeriksa masukan %q: %w",
	CLIErrUnsupportedFile:           "berkas tidak didukung %q: format yang didukung adalah CBZ, PDF, EPUB, atau direktori gambar",
	CLIFlagGraphicsUsage:            "perender: auto, ascii, dots, kitty, sixel, atau iterm2",
	CLIFlagVersionUsage:             "tampilkan versi lalu keluar",
	CLIFlagUpdateUsage:              "periksa pembaruan lalu keluar",
	CLIFlagEnvUsage:                 "tampilkan lingkungan comicread lalu keluar",
	CLIFlagClearJournalUsage:        "hapus jurnal lokal untuk berkas atau direktori lalu keluar",
	CLIFlagBookViewUsage:            "tampilkan pasangan halaman dari kiri ke kanan",
	CLIFlagRightBookViewUsage:       "tampilkan pasangan halaman dari kanan ke kiri",
	CLIFlagCircleBookViewUsage:      "tampilkan pasangan halaman bertumpuk dari kiri ke kanan",
	CLIFlagRightCircleBookViewUsage: "tampilkan pasangan halaman bertumpuk dari kanan ke kiri",
	CLIErrMultipleBookViews:         "hanya satu opsi tampilan buku yang dapat digunakan",
	CLIErrInvalidView:               "nilai COMICREAD_VIEW tidak didukung %q (harus: book-view, right-view, circle-view, atau right-circle-view)",
	CLIFlagOpenUsage:                "direktori untuk dibuka di pemilih berkas (bawaan: COMICREAD_DIR atau direktori saat ini)",
	CLIErrOpenNotDir:                "buka direktori %q: bukan direktori",
	CLIFlagWebUsage:                 "mulai pembaca web lokal alih-alih antarmuka terminal",
	CLIErrWebArgs:                   "--web tidak menerima argumen berkas atau direktori",
	WebServerStarted:                "pembaca web comicread berjalan di %s (tekan Ctrl+C untuk berhenti)",
	WebErrListen:                    "mulai server web: %w",
	WebErrServe:                     "jalankan server web: %w",
	CLIHelpHint:                     "jalankan 'comicread --help' untuk bantuan",
	CLIUsage:                        "penggunaan: comicread [opsi] [berkas]",
	CLIUsageFull: `comicread — pembaca manga terminal minimal

penggunaan: comicread [opsi] [berkas]

opsi:
  --config string     berkas konfigurasi yang digunakan
  --graphics string   perender: auto, ascii, dots, kitty, sixel, atau iterm2 (bawaan "auto")
  --book-view         tampilkan pasangan halaman dari kiri ke kanan
  --right-view        tampilkan pasangan halaman dari kanan ke kiri
  --circle-view       tampilkan pasangan halaman bertumpuk dari kiri ke kanan
  --right-circle-view
                      tampilkan pasangan halaman bertumpuk dari kanan ke kiri
  --clear-journal    hapus jurnal lokal untuk berkas atau direktori lalu keluar
  --reset-config     atur ulang config.toml ke bawaan lalu keluar
  --set-config value perbarui config.toml: kunci=nilai
  -o, --open string   direktori untuk dibuka di pemilih berkas (bawaan: COMICREAD_DIR atau direktori saat ini)
  --update            periksa pembaruan lalu keluar
  --web               mulai pembaca web lokal alih-alih antarmuka terminal
  -v, --version       tampilkan versi lalu keluar
  -h, --help          tampilkan bantuan ini

Jika tidak ada berkas atau direktori yang diberikan, pemilih berkas interaktif akan terbuka di COMICREAD_DIR
(jika diatur ke direktori yang valid) atau di direktori saat ini.`,
	ReaderViewMetadata: "Metadata",
}
