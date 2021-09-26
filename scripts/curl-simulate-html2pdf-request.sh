# simulates a post http request to /merge endpoint and uploading a mergepdfs-example-1.zip file as the input data
cd ../zip-input-files-examples 
curl -X POST http://localhost:8080/html2pdf?preffix=preffix1234 -F files=@html2pdf-example-1.zip -H "Content-Type: multipart/form-data" -o comverted.pdf
cd ../scripts