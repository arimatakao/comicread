#!/usr/bin/env bash

set -euo pipefail

REPO="arimatakao/comicread"
BIN_NAME="comicread"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"
VERSION_INPUT="latest"
AUTO_YES="false"
TMP_DIR=""

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

confirm_install() {
  local answer=""

  [ "$AUTO_YES" = "true" ] && return 0
  if [ -t 0 ]; then
    read -r -p "$1 [y/N]: " answer
  elif [ -r /dev/tty ]; then
    read -r -p "$1 [y/N]: " answer < /dev/tty
  else
    printf '%s\n' 'Error: confirmation requires a terminal; use --yes to continue.' >&2
    exit 1
  fi

  case "$answer" in
    y|Y|yes|YES) ;;
    *)
      printf '%s\n' 'Installation cancelled.'
      exit 0
      ;;
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

path_contains_install_dir() {
  case ":${PATH:-}:" in
    *":${INSTALL_DIR}:"*) return 0 ;;
    *) return 1 ;;
  esac
}

ensure_path_setup() {
  path_contains_install_dir && return 0

  local shell_name target_file line
  shell_name="$(basename "${SHELL:-}")"
  case "$shell_name" in
    bash) target_file="${HOME}/.bashrc" ;;
    zsh) target_file="${HOME}/.zshrc" ;;
    *) target_file="${HOME}/.profile" ;;
  esac

  mkdir -p "$(dirname "$target_file")"
  touch "$target_file"
  line="case \":\$PATH:\" in *\":${INSTALL_DIR}:\"*) ;; *) export PATH=\"${INSTALL_DIR}:\$PATH\" ;; esac"
  grep -Fqx "$line" "$target_file" || printf '\n%s\n' "$line" >> "$target_file"
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

  confirm_install "Install ${BIN_NAME} ${version} to ${INSTALL_DIR}?"
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

  printf 'comicread %s installed to %s/%s\n' "$version" "$INSTALL_DIR" "$BIN_NAME"
  if ! path_contains_install_dir; then
    printf 'Restart your terminal or run: export PATH="%s:$PATH"\n' "$INSTALL_DIR"
  fi
}

main "$@"
