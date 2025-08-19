#!/usr/bin/env bash
set -euo pipefail

# Build all example modules under categories/** by locating go.mod files.
# Safer than hard-coding specific category/example names.

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
REPO_ROOT=$(cd "${SCRIPT_DIR}/.." && pwd)
CATEGORIES_DIR="${REPO_ROOT}/categories"

if [[ ! -d "${CATEGORIES_DIR}" ]]; then
    echo "Error: categories directory not found at ${CATEGORIES_DIR}" >&2
    exit 1
fi

echo "Building all examples under: ${CATEGORIES_DIR}"

successes=()
failures=()

# Optional filter: pass one or more path globs relative to categories/ to limit builds
filters=("$@")

# We'll iterate modules using find + while loop for better portability (works on macOS bash 3.x)
found_any=false

should_build() {
    local moddir="$1"
    # If no filters provided, build everything
    if [[ ${#filters[@]} -eq 0 ]]; then
        return 0
    fi
    # If filters provided, match any of them
    local rel
    rel=${moddir#"${CATEGORIES_DIR}/"}
    for f in "${filters[@]}"; do
        if [[ "$rel" == $f* ]]; then
            return 0
        fi
    done
    return 1
}

while IFS= read -r -d '' modfile; do
    found_any=true
    moddir=$(dirname "${modfile}")
    if ! should_build "${moddir}"; then
        continue
    fi

    relpath=${moddir#"${REPO_ROOT}/"}
    printf "\n==> Building %s\n" "${relpath}"

    pushd "${moddir}" >/dev/null
    # Ensure module deps are tidy
    rm go.sum
    if ! go mod tidy; then
        echo "[FAIL] ${relpath}: go mod tidy failed" >&2
        failures+=("${relpath}")
        popd >/dev/null
        continue
    fi

    # Build all packages within the module
    if ! go build ./...; then
        echo "[FAIL] ${relpath}: build failed" >&2
        failures+=("${relpath}")
        popd >/dev/null
        continue
    fi

    # Clean common output binaries if any were produced
    bin_candidates=(
        "main" 
        "main.exe" 
        "$(basename "${moddir}")" 
        "$(basename "${moddir}").exe"
    )
    for b in "${bin_candidates[@]}"; do
        if [[ -f "${b}" ]]; then
            rm -f "${b}"
        fi
    done

    echo "[OK] ${relpath}"
    successes+=("${relpath}")
    popd >/dev/null
done < <(find "${CATEGORIES_DIR}" -type f -name go.mod -print0)

if [[ "${found_any}" == false ]]; then
    echo "No modules found under ${CATEGORIES_DIR}" >&2
    exit 1
fi

printf "\nBuild summary:\n"
printf "  OK:   %d\n" "${#successes[@]}"
printf "  FAIL: %d\n" "${#failures[@]}"

if [[ ${#failures[@]} -gt 0 ]]; then
    printf "\nFailures:\n" >&2
    for f in "${failures[@]}"; do
        echo "  - ${f}" >&2
    done
    exit 1
fi

exit 0
