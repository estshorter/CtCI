Param($no)
$ErrorActionPreference = "Stop"

$name = "p" + $no.ToString("00")
$cwd = Get-Location
$dir = Join-Path $cwd.ToString() $name
$file = Join-Path $dir ($name + ".go")
New-Item $dir -ItemType Directory
"package " + $name > $file