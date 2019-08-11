package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = glob(
        ["*.go"],
        exclude = ["*_test.go"],
    ),
    importpath = "github.com/binchencoder/letsgo/token",
    deps = ["@binchencoder_third_party_go//vendor/github.com/dgrijalva/jwt-go:go_default_library"],
)

go_test(
    name = "token_tests",
    srcs = ["token_test.go"],
    embed = [
        ":go_default_library",
    ],
    deps = ["@binchencoder_third_party_go//vendor/github.com/smartystreets/goconvey/convey:go_default_library"],
)
