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
        "@com_github_golang_glog//:go_default_library",
    ],
)
