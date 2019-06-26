#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")
SRC_ROOT=$SCRIPT_ROOT/..
CODEGEN_PKG=${GOPATH}/src/github.com/kubernetes/code-generator

"${CODEGEN_PKG}"/generate-groups.sh "deepcopy"      \
  generated ${SRC_ROOT}/v1alpha1 v1alpha1           \
  --output-base ${SRC_ROOT}/v1alpha1                \
  --go-header-file ${SCRIPT_ROOT}/boilerplate.go.txt

#  --output-base "$(dirname "${BASH_SOURCE[0]}")/../../.."
#  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt

# To use your own boilerplate text append:
#   --go-header-file "${SCRIPT_ROOT}"/hack/custom-boilerplate.go.txt
