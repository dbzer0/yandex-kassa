#!/usr/bin/env bash
set -e

workdir=.cover
profile="$workdir/cover.out"
mode=count

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    for pkg in "$@"; do
        f="$workdir/$(echo $pkg | tr / -).cover"
        go test -covermode="$mode" -coverprofile="$f" "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
}

show_cover_report_func() {
    go tool cover -func="$profile"
}

show_cover_report_html() {
    go tool cover -html="$profile" -o coverage.html
}

echo "generate cover data to \"$workdir\"..."
generate_cover_data $(go list ./...)
echo "done!"

echo "runing unit tests..."
show_cover_report_func
echo "done!"

echo "generate cover report to \"coverage.html\"..."
show_cover_report_html 
echo "done!"

if [ ! -z ${UID} ]; then
    echo "change uid to ${UID}..."
    chown -R ${UID} coverage.html $workdir
fi

if [ ! -z ${GID} ]; then
    echo "change gid to ${GID}..."
    chown -R :${GID} coverage.html $workdir
fi
