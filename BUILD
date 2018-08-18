load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "w2buffer",
    srcs = [
        "w2buffer.go",
    ],
    importpath = "github.com/zhach/personalweb/w2buffer",
)

go_library(
    name = "datafiles",
    srcs = [
        "grab_datafiles.go",
    ],
    importpath = "github.com/zhach/personalweb/datafiles",
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
    data = [
        ":webfiles.zip",
    ],
    deps = [
        ":handlers",
        ":datafiles",
    ],
)

load("//defs:zipper.bzl", "zip_file")
zip_file(
    name = "webfiles",
    deps = [":assets"],
)

# Bring in assets into top directory
load("@io_bazel_rules_closure//closure:defs.bzl", "web_library")
web_library(
  name = "assets",
  srcs = [
    "//personalweb/components:index.html",
  ],
  path = "/",
)