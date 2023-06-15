# Kubernetes Custom Resources Define
##  Define CRD

add comment code in 'types.go' for k8s code generator and controller-gen (kubebuilder)
```
extend.oam.dev
├── register.go
└── v1
    ├── doc.go
    ├── register.go
    └── types.go
```

## Generate  Client
./vendor/k8s.io/code-generator/generate-groups.sh all git.eniot.io/apaas/kubematches/pkg/client git.eniot.io/apaas/kubematches/pkg/apis  "extend.oam.dev:v1" --go-header-file "hack/boilerplate.go.txt"

## Gernerat CRD & Role
controller-gen rbac:roleName=kubematches crd paths="./pkg/apis/extend.oam.dev/v1/..." output:crd:artifacts:config=deploy/crd  output:rbac:artifacts:config=deploy/rbac

[https://book.kubebuilder.io/reference/markers.html](https://book.kubebuilder.io/reference/markers.html)
