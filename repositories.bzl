load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "binchencoder_third_party_go",
        commit = "884a585d57639840ae3a617bf51443951bde4724",
        importpath = "gitee.com/binchencoder/third-party-go",
    )
    go_repository(
        name = "binchencoder_ease_gateway",
        commit = "779db23c32e8e52a93721d479752bbf2426e053b",
        importpath = "gitee.com/binchencoder/ease-gateway",
    )