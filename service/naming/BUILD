package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    importpath = "github.com/binchencoder/letsgo/service/naming",
    deps = [
        "@com_github_binchencoder_gateway_proto//data:go_default_library",
    ],
)

go_test(
    name = "naming_test",
    size = "small",
    srcs = glob(["*_test.go"]),
    embed = [
        ":go_default_library",
    ],
    deps = [
        "@com_github_binchencoder_gateway_proto//data:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["naming_test.go"],
    embed = [":go_default_library"],
    deps = ["@com_github_binchencoder_gateway_proto//data:go_default_library"],
)
