package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "init.go",
        "pprof.go",
    ],
    importpath = "github.com/binchencoder/letsgo/runtime/pprof",
    deps = [
        "@binchencoder_third_party_go//vendor/github.com/golang/glog:go_default_library",
    ],
)
