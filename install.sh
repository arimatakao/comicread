#!/usr/bin/env bash

set -euo pipefail

REPO="arimatakao/comicread"
BIN_NAME="comicread"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
VERSION_INPUT="latest"
AUTO_YES="false"
TMP_DIR=""
REINSTALL_CONFIRMED="false"

cleanup() {
  [ -z "${TMP_DIR:-}" ] || rm -rf "$TMP_DIR"
}

usage() {
  cat <<'EOF'
Usage: bash install.sh [-y|--yes] [version]

Install the latest comicread release, or a specified version.

Examples:
  bash install.sh
  bash install.sh v1.2.3
  bash install.sh --yes
EOF
}

require_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    printf "Error: required command '%s' is not installed.\n" "$1" >&2
    exit 1
  fi
}

parse_args() {
  local positional=()

  while [ "$#" -gt 0 ]; do
    case "$1" in
      -y|--yes)
        AUTO_YES="true"
        ;;
      -h|--help)
        usage
        exit 0
        ;;
      -*)
        printf "Error: unknown option '%s'.\n" "$1" >&2
        usage >&2
        exit 1
        ;;
      *) positional+=("$1") ;;
    esac
    shift
  done

  if [ "${#positional[@]}" -gt 1 ]; then
    printf '%s\n' 'Error: too many version arguments.' >&2
    usage >&2
    exit 1
  fi

  if [ "${#positional[@]}" -eq 1 ]; then
    VERSION_INPUT="${positional[0]}"
  fi
}

read_answer() {
  ANSWER=""
  if [ -t 0 ]; then
    read -r -p "$1" ANSWER
  elif [ -r /dev/tty ]; then
    read -r -p "$1" ANSWER < /dev/tty
  else
    printf '%s\n' 'Error: interactive input requires a terminal; use --yes to continue.' >&2
    exit 1
  fi
}

confirm_install() {
  [ "$AUTO_YES" = "true" ] && return 0
  read_answer "$1 [y/N]: "

  case "$ANSWER" in
    y|Y|yes|YES) ;;
    *)
      printf '%s\n' 'Installation cancelled.'
      exit 0
      ;;
  esac
}

ask_yes_no() {
  read_answer "$1 [y/N]: "
  case "$ANSWER" in
    y|Y|yes|YES) return 0 ;;
    *) return 1 ;;
  esac
}

normalize_os() {
  case "$(uname -s)" in
    Linux) printf '%s\n' 'linux' ;;
    Darwin) printf '%s\n' 'darwin' ;;
    *)
      printf "Error: unsupported OS '%s' (only Linux and macOS are supported).\n" "$(uname -s)" >&2
      exit 1
      ;;
  esac
}

normalize_arch() {
  case "$(uname -m)" in
    x86_64|amd64) printf '%s\n' 'amd64' ;;
    aarch64|arm64) printf '%s\n' 'arm64' ;;
    i386|i686) printf '%s\n' '386' ;;
    *)
      printf "Error: unsupported architecture '%s'.\n" "$(uname -m)" >&2
      exit 1
      ;;
  esac
}

resolve_version() {
  if [ "$VERSION_INPUT" = "latest" ]; then
    local latest_url
    latest_url="$(curl -fsSLI -o /dev/null -w '%{url_effective}' "https://github.com/${REPO}/releases/latest")"
    basename "$latest_url"
  elif [[ "$VERSION_INPUT" == v* ]]; then
    printf '%s\n' "$VERSION_INPUT"
  else
    printf 'v%s\n' "$VERSION_INPUT"
  fi
}

normalize_version() {
  local value="${1#v}"
  value="${value%%[-+]*}"
  printf '%s\n' "$value"
}

version_to_sort_key() {
  local version major minor patch extra
  version="$(normalize_version "$1")"
  IFS='.' read -r major minor patch extra <<EOF
$version
EOF
  major="${major:-0}"
  minor="${minor:-0}"
  patch="${patch:-0}"
  case "${major}.${minor}.${patch}" in
    *[!0-9.]*|*..*) return 1 ;;
  esac
  printf '%010d%010d%010d\n' "$major" "$minor" "$patch"
}

is_version_newer() {
  local target_key current_key
  target_key="$(version_to_sort_key "$1")" || return 1
  current_key="$(version_to_sort_key "$2")" || return 1
  [ "$target_key" \> "$current_key" ]
}

get_installed_version() {
  local candidate version_output match

  for candidate in "${INSTALL_DIR}/${BIN_NAME}" "$(command -v "$BIN_NAME" 2>/dev/null || true)"; do
    [ -n "$candidate" ] && [ -x "$candidate" ] || continue
    version_output="$("$candidate" --version 2>/dev/null || true)"
    match="$(printf '%s\n' "$version_output" | grep -Eo 'v?[0-9]+(\.[0-9]+){1,3}([-.+][0-9A-Za-z.-]+)?' | head -n1)"
    if [ -n "$match" ]; then
      printf '%s\n' "$match"
      return 0
    fi
  done

  return 1
}

confirm_upgrade_if_needed() {
  local target_version="$1" installed_version=""
  installed_version="$(get_installed_version || true)"
  [ -n "$installed_version" ] || return 0

  if is_version_newer "$target_version" "$installed_version"; then
    confirm_install "${BIN_NAME} is already installed (version ${installed_version}). Do you want to update to ${target_version}?"
    REINSTALL_CONFIRMED="true"
    return 0
  fi

  if [ "$(normalize_version "$target_version")" = "$(normalize_version "$installed_version")" ]; then
    confirm_install "${BIN_NAME} is already installed (version ${installed_version}). Do you want to reinstall ${target_version}?"
    REINSTALL_CONFIRMED="true"
  fi
}

path_contains_install_dir() {
  case ":${PATH:-}:" in
    *":${INSTALL_DIR}:"*) return 0 ;;
    *) return 1 ;;
  esac
}

shell_config_file() {
  local shell_name
  shell_name="$(basename "${SHELL:-}")"
  case "$shell_name" in
    bash) printf '%s\n' "${HOME}/.bashrc" ;;
    zsh) printf '%s\n' "${HOME}/.zshrc" ;;
    *) printf '%s\n' "${HOME}/.profile" ;;
  esac
}

append_shell_line() {
  local target_file="$1" line="$2"
  mkdir -p "$(dirname "$target_file")"
  touch "$target_file"
  grep -Fqx "$line" "$target_file" || printf '\n%s\n' "$line" >> "$target_file"
}

ensure_path_setup() {
  path_contains_install_dir && return 0

  local target_file line
  target_file="$(shell_config_file)"
  line="case \":\$PATH:\" in *\":${INSTALL_DIR}:\"*) ;; *) export PATH=\"${INSTALL_DIR}:\$PATH\" ;; esac"
  append_shell_line "$target_file" "$line"
}

configure_environment_value() {
  local target_file="$1" name="$2" values="$3"

  while :; do
    read_answer "  ${name} (${values}; leave blank to skip): "
    if [ -z "$ANSWER" ]; then
      return
    fi

    case "${name}:${ANSWER}" in
      COMICREAD_GRAPHICS:auto|COMICREAD_GRAPHICS:ascii|COMICREAD_GRAPHICS:dots|COMICREAD_GRAPHICS:kitty|COMICREAD_GRAPHICS:sixel|COMICREAD_GRAPHICS:iterm2|COMICREAD_VIEW:book-view|COMICREAD_VIEW:right-view|COMICREAD_VIEW:circle-view|COMICREAD_VIEW:right-circle-view|COMICREAD_LANG:en|COMICREAD_LANG:uk|COMICREAD_LANG:pl|COMICREAD_LANG:de|COMICREAD_LANG:fr|COMICREAD_LANG:es|COMICREAD_LANG:cs|COMICREAD_LANG:ro|COMICREAD_LANG:it|COMICREAD_LANG:ko|COMICREAD_LANG:ja|COMICREAD_LANG:id|COMICREAD_LANG:hi|COMICREAD_LANG:el|COMICREAD_LANG:tr|COMICREAD_LANG:kk|COMICREAD_LANG:ka)
        append_shell_line "$target_file" "export ${name}=${ANSWER}"
        return
        ;;
      *) printf 'Invalid value. Choose one of: %s\n' "$values" >&2 ;;
    esac
  done
}

configure_environment() {
  [ "$AUTO_YES" = "true" ] && return
  if ! ask_yes_no 'Configure comicread environment variables in your shell profile?'; then
    return
  fi

  local target_file
  target_file="$(shell_config_file)"
  printf 'Values will be saved in %s.\n' "$target_file"
  printf '%s\n' 'COMICREAD_GRAPHICS chooses how pages are rendered; auto detects terminal support.'
  printf '%s\n' 'COMICREAD_VIEW chooses the default page layout; leave it blank for single-page view.'
  printf '%s\n' 'COMICREAD_LANG chooses the language of the interface; the default is en.'
  configure_environment_value "$target_file" 'COMICREAD_GRAPHICS' 'auto ascii dots kitty sixel iterm2'
  configure_environment_value "$target_file" 'COMICREAD_VIEW' 'book-view right-view circle-view right-circle-view'
  configure_environment_value "$target_file" 'COMICREAD_LANG' 'en uk pl de fr es cs ro it ko ja id hi el tr kk ka'
}

install_binary() {
  local source="$1"

  mkdir -p "$INSTALL_DIR"
  if [ ! -w "$INSTALL_DIR" ]; then
    printf "Error: '%s' is not writable.\n" "$INSTALL_DIR" >&2
    printf 'Set a writable directory, for example: INSTALL_DIR=$HOME/.local/bin bash install.sh\n' >&2
    exit 1
  fi

  install -m 0755 "$source" "${INSTALL_DIR}/${BIN_NAME}"
}

main() {
  parse_args "$@"
  require_cmd curl
  require_cmd tar
  require_cmd install

  local os arch version archive url
  os="$(normalize_os)"
  arch="$(normalize_arch)"
  version="$(resolve_version)"
  archive="${BIN_NAME}_${version}_${os}_${arch}.tar.gz"
  url="https://github.com/${REPO}/releases/download/${version}/${archive}"

  confirm_upgrade_if_needed "$version"
  if [ "$REINSTALL_CONFIRMED" != "true" ]; then
    confirm_install "Install ${BIN_NAME} ${version} to ${INSTALL_DIR}?"
  fi
  TMP_DIR="$(mktemp -d)"
  trap cleanup EXIT HUP INT TERM

  printf 'Downloading %s...\n' "$archive"
  curl -fsSL "$url" -o "${TMP_DIR}/${archive}"
  printf 'Extracting %s...\n' "$archive"
  tar -xzf "${TMP_DIR}/${archive}" -C "$TMP_DIR"

  if [ ! -f "${TMP_DIR}/${BIN_NAME}" ]; then
    printf "Error: '%s' was not found in the archive.\n" "$BIN_NAME" >&2
    exit 1
  fi

  install_binary "${TMP_DIR}/${BIN_NAME}"
  ensure_path_setup
  configure_environment

  printf 'comicread %s installed to %s/%s\n' "$version" "$INSTALL_DIR" "$BIN_NAME"
  if ! path_contains_install_dir; then
    printf '%s\n' 'Restart your terminal to use comicread.'
  fi
}

main "$@"
