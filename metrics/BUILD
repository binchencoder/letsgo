package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["metrics.go"],
    importpath = "binchencoder.com/letsgo/metrics",
    deps = [
        "@com_github_golang_glog//:go_default_library",
        "@com_github_prometheus_client_golang//prometheus:go_default_library",
    ],
)
