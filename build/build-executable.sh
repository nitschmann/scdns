#/bin/bash

supported_os=("darwin" "linux" "windows")
supported_arch=("386" "amd64")

os=$1

if [[ -z "$os" ]]; then
  echo "No OS given"
  exit 1
fi

os_valid=0

for i in "${supported_os[@]}"
do
  if [[ "$i" ==  "$os" ]]; then
    os_valid=1
  fi
done

if [[ "$os_valid" != 1 ]]; then
  echo "Given OS is not supported in this build script; Try it manually."
  exit 1
fi

version=$(cat VERSION)
build_version_os_path="./build/package/$version/$os"

for arch in "${supported_arch[@]}"
do
  arch_build_path="$build_version_os_path/$arch"
  executable_name="scdns"
  if [[ $os == "windows" ]]; then
    executable_name="$executable_name.exe"
  fi
  output_path="$arch_build_path/$executable_name"

  mkdir -p $arch_build_path

  env CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -o $output_path -v ./cmd/scdns/main.go
  if [ $? -ne 0 ]; then
    echo 'An error has occurred! Aborting the script execution...'
    exit 1
  fi
done
