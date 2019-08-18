package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    importpath = "binchencoder.com/letsgo/strings",
)

go_test(
    name = "trace_test",
    size = "small",
    srcs = glob(["*_test.go"]),
    embed = [
        ":go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "convert_test.go",
        "rang_test.go",
    ],
    embed = [":go_default_library"],
)
