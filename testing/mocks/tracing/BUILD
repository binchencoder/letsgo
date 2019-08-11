package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "tracing.go",
    ],
    importpath = "github.com/binchencoder/letsgo/testing/mocks/tracing",
    deps = [
        "@binchencoder_third_party_go//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "@binchencoder_third_party_go//vendor/github.com/stretchr/testify/mock:go_default_library",
    ],
)
