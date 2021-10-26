# myhtml2pdf Project

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

## Logging police section:

The system automatically will save in the mypdfservice_debug.log file detailed information abount all requestes in json format. 
Then we can read programmatically the log file and parse it
usdin any program language that has a paerser for json ( corrently
all languages have good josn libraries for that).

The file **mypdfservice_debug.log** is created automacally in the same folder where the executable **mypdfservice** binary file (in windows **mypdfservices.exe**) is located.

To avoid high disk consuming space, the system limits the log file to 10 MegaBytes.
Always when the log file reaches the size of 10 MBytes, 
the current log file is renamed and zipped.
In addition, a new log file named **mypdfservices_debug.log** is created
immediatly where the new logging records will be continually saved. 

The name of Gzipped old log files has its name like that:  
**mypdfservice_debug-2021-10-12T22-39-39.489.log.gz**
the characteres between ***mypdfservice-*** and ***.log.gz***  
represents the date/time when the new log file was created.

The system also allows up to 20 old log files (the zipped ones).
Finally, old log files the have more than 90 days age also will be deleted.

## Additional Summary Logging Section:

We can get an additional summary login file that capture all info that
the system sends to the consule. A Summary of all http request the 
receives. This infomation we have in more detaled way in the main
log file **mypdfservices_debug.log**.
The name of this second log file the user that starts the service 
can decide in the command line used to start the service using
redirecting from the commandline shell. (>).

Ex: command line exemple for windows
mypdfservices.exe > mysummary.log

In this example, the name of the sumamary logging file will be mysummary.log and it will be located in the sam folder the command line have been executed.

Obs:
Using the command line also is possible to change the TCP port
where the application will lestinig for http requests.

Example: 
mypdfservice 9999 > my_summary.log

In this example the TCP port will be 9999 and the summary logging file name will be my_summary.log.

This summary loggin file doesn't have any logging police assigned to it.
Only the main logging file described preveouly has polices for mitigate 
high disk space consumming.


## How to use the 2 Endpoints Section:   

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


## CURL examples for testing Section:


Merging PDFs files:
This curl command must be executed in the *cmd/web/api* directory as the current dir and inside it also is necessary to create a folder named *examples* and save into it a *pdf-merge-example-1.zip* file with all pdfs files that have to be merged.
curl -X POST http://localhost:8080/merge?preffix=evol1234 \
  -F files=@examples/pdf-merge-example-1.zip \
  -H "Content-Type: multipart/form-data" \
  -o merged.pdf

Converting HTML to PDF:
This curl command also must be executed in *cmd/web/api* directory as the current dir and, inside it also it is necessary to create a folder named *examples* and save into it a *pdf-conversion-example-1.zip* file with all typical files that is necessary to execute a good html->pdf conversion.
curl -X POST http://localhost:8080/html2pdf?preffix=evol1234 \
  -F files=@examples/pdf-conversion-example-1.zip \
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


## Deploying in Windows server Section:

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




## Deploying in linux server Section:

   the pdftk must be installed in the linux systems, using the package manager
   For Debien and Ubuntu linux flavors you can run the followinf commands:  
   **sudo apt-get update**  
   **sudo apt-get -y install pdftk**    
  
   For html2pdf it is better get the binary file that can be found in cmd/webapi folder 


## Deploying using docker containers Section:

   //todo yet ... 



## Theankful Section:

1. I am thankful for all people that make part of the wkhtnltopdf project. Without your good work, this project would be dificult to do.

2. Also, Thankful for all people that make part of the pdftk project.
Over here, the same we said about the wkhtmltopdf project is valid.
Without your good work, this project would be difficult to do.  

3. we are also thankful for the Gitpod Team. This project was written entirely using GItpod free tier resources. I did all my work without the necessity to install anything on my machine. Also, I used only a simple Internet Browser. I am an advocate of development in the Cloud in this strict way to do it. Gitpod is on the right path and this first project test had a positive result in that direction.  


## Licences Section:  

   // todo yet ...