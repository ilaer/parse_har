go build  -o ghar.exe  -ldflags "-H windowsgui"  main.go
D:
cd D:\GoP\ghar
del ghar.zip
timeout /nobreak /t 3
go build  -o ghar.exe   -ldflags "-H windowsgui"
"C:\Program Files\WinRAR\WinRAR.exe" a -r -ep1 -idq -inul -y "ghar.zip" "ghar.exe"
timeout /nobreak /t 15