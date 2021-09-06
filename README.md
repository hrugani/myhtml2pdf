# myhtml2pdf

implements a REST api que process PDF files.  
  
Offers 2 endponts:  
/convert  
/concat  

Both endepoints receive a zipped file that should contains all the necessary inputs
to execute the desired pdf process.  

1. /convert  
   converts HTML file to PDF.  
   this enpoint receives a standard Multipart HTTP request which one should contains  
   
   1 zipped file with all files necessary to performe the HTML to PDF conversion.  
   So, inside of the zipped file must have:  

   1 file with the extension .html qith the HTML content to be coneverted.  
   
   N files of type images qith all respective images referenced in the src attribute of IMG HTML tags.
   here, it is important to assert all images pointed by each IMG tag are present into zipped files.
   All IMG tags must point to an image that exists into the zipped file and in the src attribute must
   be valued with the exact name o the file in the zipped file.
   Even when a IMG tag doensn't have to show no-content, an transparent image must be present into
   the zipped file with the right name in the src atribute of the repsective IMG tag.
   If a specific image should have to be shown saveral times. This image needs to be present
   only one time into the zipped file. It is enough for the IMG tags to reference the same filename that corresponds to the repeated image.  
     
2. /concat  
   concatenates N PDFs files into 1 pdf file.
   Also receives a Multpart HTTP request win only one zipped file attached:
   This zipped files must contains all the pdf files that should be concatenated.
   The order of concatenation obeys the alphabetical order of the PDF file names inside the zipped file.

   Returns 1 PDF file named "output.pdf" where into it all pdf files uploaded in the zipped file
   concatenated in alphabetical order of the respective pdf file names.  