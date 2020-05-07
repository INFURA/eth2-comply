const { spawn } = require('child_process')

const buildCommand = "bazel"
const standardBuildArgs = ["build"]

let buildProcs = []
let buildProc

const linuxAmd64BuildArgs = standardBuildArgs.concat([":linux_amd64"])
buildProc = spawn(buildCommand, linuxAmd64BuildArgs)
buildProcs = buildProcs.concat([buildProc])

const darwinAmd64BuildArgs = standardBuildArgs.concat([":darwin_amd64"])
buildProc = spawn(buildCommand, darwinAmd64BuildArgs)
buildProcs = buildProcs.concat([buildProc])

const windowsAmd64BuildArgs = standardBuildArgs.concat([":windows_amd64"])
buildProc = spawn(buildCommand, windowsAmd64BuildArgs)
buildProcs = buildProcs.concat([buildProc])

const testsBuildArgs = standardBuildArgs.concat([":tests"])
buildProc = spawn(buildCommand, testsBuildArgs)
buildProcs = buildProcs.concat([buildProc])

buildProcs.map(buildProc => {
  buildProc.stdout.on("data", data => {
    process.stdout.write(data.toString("utf-8"))
  })
  buildProc.stderr.on("data", data => {
    process.stderr.write(data.toString("utf-8"))
  })
})

