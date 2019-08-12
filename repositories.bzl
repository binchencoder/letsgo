load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "binchencoder_third_party_go",
        importpath = "github.com/binchencoder/third-party-go",
        urls = [
            "https://codeload.github.com/binchencoder/third-party-go/tar.gz/4e9c6ce6b9edd7289966dda9be983f12a063584c",
        ],
        strip_prefix = "third-party-go-4e9c6ce6b9edd7289966dda9be983f12a063584c",
        type = "tar.gz",
    )
    go_repository(
        name = "binchencoder_ease_gateway",
        importpath = "github.com/binchencoder/ease-gateway",
        urls = [
            "https://codeload.github.com/binchencoder/ease-gateway/tar.gz/544d50be5ccd1d8956eef3da33ed90ec7d6281e6",
        ],
        strip_prefix = "ease-gateway-544d50be5ccd1d8956eef3da33ed90ec7d6281e6",
        type = "tar.gz",
    )