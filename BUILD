load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "w2buffer",
    srcs = [
        "w2buffer.go",
    ],
    importpath = "github.com/zhach/personalweb/w2buffer",
)

go_library(
    name = "handlers",
    srcs = [
        "handlers.go",
    ],
    deps = [
        ":w2buffer",
    ],
    importpath = "github.com/zhach/personalweb/handlers",
)

go_binary(
    name = "webserver",
    srcs = [
        "main.go"
    ],
    deps = [
        ":handlers",
    ],
)