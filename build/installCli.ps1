Param(
    [string] $version = "[RELEASE]"
)

if ($version -eq "[RELEASE]") {
    Write-Host "Downloading the latest version of JFrog CLI..."
} else {
    Write-Host "Downloading version $version of JFrog CLI..."
}
Start-Process powershell "-NoProfile -Command iwr https://releases.jfrog.io/artifactory/jfrog-cli/v2/[RELEASE]/jfrog-cli-windows-amd64/jfrog.exe -OutFile $env:SYSTEMROOT\system32\jf.exe" -Verb RunAs