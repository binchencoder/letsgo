// Package version prints the version information set by Bazel.
//
// To make it work, one has to:
// a) add "letsgo.Init()" as the first line in main() function:
// b) add the following line in the go_binary build target:
//    linkstamp = "github.com/binchencoder/letsgo/version",
//
// To set the BUILD_EMBED_LABEL, add flag --embed_label in bazel build command:
//    bazel build --embed_label=1.2.3 <taget>
//
// For Jingoal's custom build info, add flag --workspace_status_command in
// bazel build command:
//    bazel build --workspace_status_command=./status.sh <taget>
// here status.sh outputs the build info as "key value" lines.
package version
