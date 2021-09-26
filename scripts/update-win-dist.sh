# updates the windows distribution folder 

# IMPORTANT: all scripts in the scripts folder must be executed with the /scripts folder as the current directory.
# This procedure will guarantee the relative reference to files and folder works properly

# changes currenty dir
cd ../cmd/webapi
# build the .exe binary
GOOS=windows go build -o ../../dist/windows-dist/mypdfservices.exe main.go
# updates latest changes on Test_*.html files in wondows-dist project folder 
cp Test_*.html ../../dist/windows-dist
# backs /scripts folder as current dir
cd ../../scripts

# IMPORTANAT Obs:
# The binaries of the other programs and its respective DLLs exist only in the windows-dist
# They were built for windows platform using versions of them.
# Probably, they will be out of date quickly.
# If you want to update this binaries you should provide that manually in the prodution stage.
# The goal of this project is only treat the golang app that wraps this binaries into a web application. 