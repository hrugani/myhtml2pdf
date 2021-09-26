# myhtml2pdf

implements a REST api that process PDF files.  
  
Offers 2 endpoints:  
/html2pdf  
/merge  

Both endpoints receive a zipped file that should contains all the necessary inputs
to execute the desired pdf process.  
  
This web app is only a wrapper for 2 execelent command lines applications that execute PDF actions.
For HTML to PDF convertion the wkhtmltopdf command line application is used.
For PDF merging, the Application pdftk is used.
We are very grateful to all developers that contributed to these 2 projects. Good job!!!!

Both command line applications offer a lot of options that allow more complex tranformations.
The mean goal of this project is a minimalist implementation  
to reach the necessities of a specific company.
So, here we spend efforts to make the things simple for some target use cases.
But, this base code can be used and adapted for other use cases or even
adapted to reach more generic goals. 


### How to use the 2 Endpoints:   

1. /convert  
   converts HTML file to PDF.  
   this endpoint receives a standard Multipart HTTP request which one should contains  
   
   1 zipped file with all files necessary to performe the HTML to PDF conversion.  
   So, inside of the zipped file must have:  

   1 file with the extension .html (HTML content to be converted).  
   
   N files of type images with all respective images referenced in the src attributes of IMG HTML tags.
   Here, it is important to assert all images pointed by each IMG tag are present into zip file.
   All IMG tags must point to an image that exists into the zipped file and in the src attribute must
   be valued with the exact name of the image file in the zipped file.
   Even when a IMG tag doensn't have to show no-content, an transparent image must be present into
   the zipped file with the right name in the src atribute its respective IMG tag.
   Also it is import to notice, if a specific image should have to be shown saveral times, this image needs to be present
   only one time into the zipped file.  
   It is enough for the IMG tags to reference the same filename that corresponds to the repeated image.  

   Only images of the types .jpg, .jpeg, .png and .gif are acceptable.

   Returns 1 PDF file that contains the HTML content converted to PDF format.

2. /merge  
   concatenates N PDFs files into 1 pdf file.
   Also receives a Multpart HTTP request with only one zipped file attached:
   This zipped files must contains all the pdf files that should be concatenated.
   The order of concatenation obeys the alphabetical order of the PDF file names inside the zipped file.

   Returns 1 PDF file named "output.pdf" where into it all pdf files uploaded in the zipped file
   concatenated in alphabetical order of its respective pdf file names. 


### CURL examples for testing


Merging PDFs files:
This curl command must be executed cmd/web/api directory as the current dir
curl -X POST http://localhost:8080/merge?preffix=evol1234 \
  -F files=@examples/pdf-example-1.zip \
  -H "Content-Type: multipart/form-data" \
  -o merged.pdf

Converting HTML to PDF:
This curl command also must be executed cmd/web/api directory as the current dir
curl -X POST http://localhost:8080/html2pdf?preffix=evol1234 \
  -F files=@examples/example-4.zip \
  -H "Content-Type: multipart/form-data" \
  -o converted.pdf


Obs: you can create your own .zip files for testing
Always You must take in account:
1. the .zip file content for html2pdf must contain:  
   Only 1 file .html whitch one should contains the HTML content to be converted to PDF.  
   All the images files that are pointed by each IMG src attribute present into the html Body.  
   Only images of type .png, .jpeg, .jpg will be accepted.  
   All IMG src must point to an actual image file present into .zip file.  
   When you want to show nothing in a determined HTML-IMG-tag,  
   it is necessary you put into the .zip file a tranparent image without any graphical information  
   and values the IMG src attribute with the name of that transparent image.  
   When a IMG src atribute points to a not existent image in the .zip file an error  
   will be launched and the wanted convertion will not be done.  
   The gold rule is to put into the .zip all elements that is necessary to
   provide a consistent HTML-to-PDF convertion.  
   The presence of unecessary images are allowed, but it isn't a good practice.  
   Please, avoid to use special chars in the .html and images file names.
   Avoid spaces in file names too.
   Prefer pure ASCII characteres    

2. the .zip file for pdf merging must contain:
   All .pdf files which ones should be merged.
   They will be merged in alphabetical order.
   Please, avoid to use special chars in the names of .pdf files
   Avoid spaces in the .pdf file names too.
   Prefer pure ASCII characteres    


   ## Deploying in Windows server

   On windows machines, the better approach is to have all binary files
   and its respective DLLs present into the same directory.  
   *Simply puts all of them into the same directory and you will be ready to go.  
     

   There are 3 programs:
   1) **mypdfservices.exe**:
      This is the binary genereted by this project in golang.
      it can be generated executing /scripts/update-win-dist.sh
       The main command line into this script is:  
      **GOOS=windows go build -o mypdfservices.exe main.go**  
      This is the only binary file that should be executed to make all services up and running.
      This executable file receives an optional parameter used to change the default IP Port.
      When this program is executed without any parameter, by default, the server will respond on 8080 port.  


   2) **pdftk.exe** (and its DLL: libiconv2.dll)  
      this binaries files can be found in the cmd/webapi folder  
      of this project.
      This program is executed inside mypdfservices.
      Only the presence of it into de instalation directory is necessary  

   3) **wkhtmltopdf.exe** (and its DDL: wkhtmltox.dll)  
      this binaries files also can be found in the cmd/webapi folder  
      of this project.  
      This program is called inside mypdfservices.
      Only the presence of it into de instalation directory is necessary  

   To make the service active, simply execute the mypdfservices file.
   By default, this program will be listening at IP port 8080 (HTTP Post requests)
   Whether you want to use another IP port, you can pass the desired port number as the first parameter  
   ex: mypdfservices 9134
   This command line will change the default port 8080 to 9134  

   To make tests easy 2 files .html are provided.
   The files are **Test_HTML2PDF.html** and **Test_MergePDFs.html**
   Both contain a minimalistic HTML-FORM.
   When they are opened by any browser running on the same machine where our mypdfservices are running
   the browser executes a correct HTTP POST request in port 8080 (default port).
   Whether the default port 8080 was changed using a parameter of mypdfserves program
   these html files must be modified accordingly.


   ## Deploying in linux server

   the pdftk must be installed in the linux systems, using the package manager
   For Debien and Ubuntu linux flavors you can run the followinf commands:  
   **sudo apt-get update**  
   **sudo apt-get -y install pdftk**    
  
   For html2pdf it is better get the binary file that can be found in cmd/webapi folder 

   ## Deploying using docker containers
