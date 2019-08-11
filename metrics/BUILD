package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "metrics.go",
    ],
    importpath = "github.com/binchencoder/letsgo/metrics",
    deps = [
        "@binchencoder_third_party_go//vendor/github.com/golang/glog:go_default_library",
        "@binchencoder_third_party_go//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
    ],
)
