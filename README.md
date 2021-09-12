# myhtml2pdf

implements a REST api que process PDF files.  
  
Offers 2 endponts:  
/convert  
/concat  

Both endepoints receive a zipped file that should contains all the necessary inputs
to execute the desired pdf process.  
  
This web app is only a wrapper for 2 command lines applications that execute the wanted action.
For HTRML to PDF convertion the wkhtmltopdf command line application is used.
For PDF merging, the Application pdftk is used.

Both command line appllications offer a lot of options that allow more complex tranformations.
The mean goal of this project is a minimalist implementation  
to reach the necessities of a specific company.
So, here we spend efforts to make the things simple for some target use cases.
But, this base code can be used and adapted for other use cases or even
adapted to reach more generic goals. 


1. /convert  
   converts HTML file to PDF.  
   this endpoint receives a standard Multipart HTTP request which one should contains  
   
   1 zipped file with all files necessary to performe the HTML to PDF conversion.  
   So, inside of the zipped file must have:  

   1 file with the extension .html (HTML content to be converted).  
   
   N files of type images with all respective images referenced in the src attributes of IMG HTML tags.
   Here, it is important to assert all images pointed by each IMG tag are present into zipped files.
   All IMG tags must point to an image that exists into the zipped file and in the src attribute must
   be valued with the exact name of the image file in the zipped file.
   Even when a IMG tag doensn't have to show no-content, an transparent image must be present into
   the zipped file with the right name in the src atribute its respective IMG tag.
   Also it is import to notice, if a specific image should have to be shown saveral times, this image needs to be present
   only one time into the zipped file.  
   It is enough for the IMG tags to reference the same filename that corresponds to the repeated image.  

   Only images of the types .jpg, .jpeg, .png and .gif are acceptable.

   Returns 1 PDF file that contains the HTML content converted to PDF format.

2. /concat  
   concatenates N PDFs files into 1 pdf file.
   Also receives a Multpart HTTP request with only one zipped file attached:
   This zipped files must contains all the pdf files that should be concatenated.
   The order of concatenation obeys the alphabetical order of the PDF file names inside the zipped file.

   Returns 1 PDF file named "output.pdf" where into it all pdf files uploaded in the zipped file
   concatenated in alphabetical order of its respective pdf file names.  

   ## Deploying in Windows server

   ## Deploying in linux server

   ## Deploying using docker containers