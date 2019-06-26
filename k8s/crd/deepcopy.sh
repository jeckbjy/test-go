#!/usr/bin/env bash

#deepcopy-gen  --go-header-file ./hack/boilerplate.go.txt --input-dirs ./v1alpha1 -O zz_generated.deepcopy --bounding-dirs ./v1alpha1 -v 2
deepcopy-gen --input-dirs ./v1alpha1/ -o .  -O types.deepcopy --go-header-file ./hack/boilerplate.go.txt -v 2