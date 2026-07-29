package i18n

var esMessages = map[string]string{
	CLIFlagConfigUsage: "archivo de configuración que se utilizará", CLIFlagResetConfigUsage: "restablecer config.toml y salir", CLIFlagSetConfigUsage: "actualizar config.toml: clave=valor",
	ReaderStatusWaitingTerminalSize: "esperando el tamaño del terminal",
	ReaderStatusTerminalTooSmall:    "la ventana del terminal es demasiado pequeña",
	ReaderStatusLastPage:            "última página",
	ReaderStatusFirstPage:           "primera página",
	ReaderStatusRenderError:         "error de renderizado: %v",
	ReaderStatusMaximumZoom:         "zoom máximo",
	ReaderStatusMinimumZoom:         "zoom mínimo",
	ReaderStatusInvalidPage:         "número de página no válido",

	ReaderViewTerminalTooSmall: "comicread: la ventana del terminal es demasiado pequeña",
	ReaderViewWindowTitle:      "comicread — %s",
	ReaderViewPages:            "páginas %d/%d",
	ReaderViewPageRange:        "páginas %d-%d/%d",
	ReaderViewRendering:        "renderizando",
	ReaderViewBookmarks:        "Marcadores",
	ReaderViewNoBookmarks:      "(sin marcadores)",
	ReaderViewBookmarksHelp:    "arriba/abajo mover | enter abrir | esc cerrar",
	ReaderViewGoToPage:         "Ir a la página: %s",
	ReaderViewHelp: `Teclas

← →  página anterior / siguiente
↑ ↓  desplazar una página ampliada
+ -  ampliar / reducir
b    añadir / quitar marcador
v ← → marcador anterior / siguiente
c v  marcadores
g123 enter  ir a la página
q    salir

?    cerrar ayuda`,

	FilepickerHeader:         "comicread — seleccionar un capítulo\n%s\n\n",
	FilepickerNoEntries:      "  (no hay entradas compatibles)\n",
	FilepickerHelp:           "\n↑/↓ mover\n← directorio superior\n→ entrar al directorio\nenter abrir archivo\ns seleccionar el directorio resaltado\nf añadir / quitar el directorio actual de favoritos\nF añadir directorio favorito\nb directorios favoritos\no ir a un directorio\nq salir\n",
	FilepickerWindowTitle:    "comicread — elegir un archivo",
	FilepickerGoToPrompt:     "\nIr al directorio: %s\n",
	FilepickerFavoritePrompt: "\nDirectorio favorito: %s\n",
	FilepickerFavorites:      "Directorios favoritos\n\n",
	FilepickerNoFavorites:    "  (no hay directorios favoritos configurados)\n",
	FilepickerFavoritesHelp:  "\n↑/↓ mover\nenter ir al directorio\nd quitar favorito\nesc volver\n",
	FilepickerFavoriteErr:    "  error al guardar favoritos: %s\n",
	FilepickerGoToErr:        "  error: %s\n",

	FilepickerErrResolveDir: "no se puede resolver el directorio %q: %w",
	FilepickerErrReadDir:    "no se puede leer el directorio %q: %w",
	FilepickerErrRunPicker:  "error al ejecutar el selector de archivos: %w",
	FilepickerErrEmptyPath:  "la ruta está vacía",
	FilepickerErrNotDir:     "%q no es un directorio",

	LoadingViewOpening:     "abriendo %s…",
	LoadingViewWindowTitle: "comicread — abriendo",

	CLIErrGetWorkingDir:             "no se puede obtener el directorio de trabajo: %w",
	CLIErrPickFile:                  "no se puede elegir el archivo: %w",
	CLIErrRunTUI:                    "error al ejecutar la TUI: %w",
	CLIErrParseArgs:                 "error al analizar los argumentos: %w",
	CLIErrOpenChapter:               "no se puede abrir el capítulo: %w",
	CLIErrOpenJournal:               "no se puede abrir el registro: %w",
	CLIErrClearJournal:              "no se puede borrar el registro: %w",
	CLIErrClearJournalRequiresInput: "--clear-journal requiere un archivo o directorio",
	CLIErrNoPages:                   "el capítulo no contiene páginas de imagen legibles",
	CLIErrInspectInput:              "no se puede inspeccionar la entrada %q: %w",
	CLIErrUnsupportedFile:           "archivo no compatible %q: los formatos compatibles son CBZ, PDF, EPUB o un directorio de imágenes",
	CLIFlagGraphicsUsage:            "renderizador: auto, ascii, dots, kitty, sixel o iterm2",
	CLIFlagVersionUsage:             "mostrar la versión y salir",
	CLIFlagUpdateUsage:              "buscar actualizaciones y salir",
	CLIFlagEnvUsage:                 "mostrar el entorno de comicread y salir",
	CLIFlagClearJournalUsage:        "eliminar el registro local de un archivo o directorio y salir",
	CLIFlagBookViewUsage:            "mostrar pares de páginas de izquierda a derecha",
	CLIFlagRightBookViewUsage:       "mostrar pares de páginas de derecha a izquierda",
	CLIFlagCircleBookViewUsage:      "mostrar pares de páginas superpuestas de izquierda a derecha",
	CLIFlagRightCircleBookViewUsage: "mostrar pares de páginas superpuestas de derecha a izquierda",
	CLIErrMultipleBookViews:         "solo se puede usar una opción de vista de libro",
	CLIErrInvalidView:               "valor de COMICREAD_VIEW no compatible %q (se espera: book-view, right-view, circle-view o right-circle-view)",
	CLIFlagOpenUsage:                "directorio para abrir en el selector de archivos (predeterminado: COMICREAD_DIR o el directorio actual)",
	CLIErrOpenNotDir:                "abrir directorio %q: no es un directorio",
	CLIFlagWebUsage:                 "iniciar un lector web local en lugar de la interfaz de terminal",
	CLIErrWebArgs:                   "--web no acepta un argumento de archivo o directorio",
	WebServerStarted:                "el lector web de comicread se ejecuta en %s (pulse Ctrl+C para detenerlo)",
	WebErrListen:                    "iniciar servidor web: %w",
	WebErrServe:                     "ejecutar servidor web: %w",
	CLIHelpHint:                     "ejecute 'comicread --help' para ver la ayuda",
	CLIUsage:                        "uso: comicread [opciones] [archivo]",
	CLIUsageFull: `comicread — un lector de manga minimalista para la terminal

uso: comicread [opciones] [archivo]

opciones:
  --config string     archivo de configuración que se utilizará
  --graphics string   renderizador: auto, ascii, dots, kitty, sixel o iterm2 (predeterminado "auto")
  --book-view         mostrar pares de páginas de izquierda a derecha
  --right-view        mostrar pares de páginas de derecha a izquierda
  --circle-view       mostrar pares de páginas superpuestas de izquierda a derecha
  --right-circle-view
                      mostrar pares de páginas superpuestas de derecha a izquierda
  --clear-journal    eliminar el registro local de un archivo o directorio y salir
  --reset-config     restablecer config.toml y salir
  --set-config value actualizar config.toml: clave=valor
  -o, --open string   directorio para abrir en el selector de archivos (predeterminado: COMICREAD_DIR o el directorio actual)
  --update            buscar actualizaciones y salir
  --web               iniciar un lector web local en lugar de la interfaz de terminal
  -v, --version       mostrar la versión y salir
  -h, --help          mostrar esta ayuda

Si no se proporciona ningún archivo o directorio, se abre un selector de archivos interactivo en COMICREAD_DIR
(si está configurado como un directorio válido) o, de lo contrario, en el directorio actual.`,
	ReaderViewMetadata: "Metadatos",
}
