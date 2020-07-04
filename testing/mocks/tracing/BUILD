package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["tracing.go"],
    importpath = "github.com/binchencoder/letsgo/testing/mocks/tracing",
    deps = [
        "@com_github_opentracing_opentracing_go//:go_default_library",
        "@com_github_stretchr_testify//mock:go_default_library",
    ],
)
