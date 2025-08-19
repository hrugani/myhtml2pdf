# myhtml2pdf Project

This project provides a REST API for processing PDF files, offering two main endpoints: `/html2pdf` and `/merge`. It acts as a wrapper for two powerful command-line utilities: `wkhtmltopdf` for converting HTML to PDF, and `pdftk` for merging PDF documents.

The primary goal is to offer a minimalist and straightforward implementation for specific use cases, but the codebase can be easily adapted for more generic purposes.

## Features

-   `/html2pdf`: Converts an HTML file (along with its assets like images and CSS) into a PDF.
-   `/merge`: Merges multiple PDF files into a single document, sorted alphabetically by filename.

Both endpoints expect a `.zip` file containing all the necessary input files.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

-   **Go**: Version 1.18 or higher.
-   **wkhtmltopdf**: A command-line tool to render HTML into PDF.
-   **pdftk**: A command-line tool for manipulating PDF documents.

## Getting Started

Below are the instructions to get the project up and running on different platforms.

### Windows (10 & 11)

1.  **Install Dependencies**:
    -   Download and install `wkhtmltopdf` from the [official site](https://wkhtmltopdf.org/downloads.html).
    -   Download and install `pdftk` from the [official site](https://www.pdflabs.com/tools/pdftk-the-pdf-toolkit/).
    -   Ensure the installation directories for both tools are added to your system's `PATH` environment variable.

2.  **Build the Application**:
    Open a command prompt or PowerShell and run the following command from the project root to build the executable:
    ```sh
    go build -o mypdfservice.exe ./cmd/webapi/main.go
    ```

3.  **Run the Service**:
    Execute the compiled binary:
    ```sh
    ./mypdfservice.exe
    ```
    The server will start listening on the default port `8080`. To use a different port, pass it as an argument:
    ```sh
    ./mypdfservice.exe 9999
    ```

### Linux (Debian/Ubuntu & openSUSE)

1.  **Install Dependencies**:
    -   **For Debian/Ubuntu**:
        ```sh
        sudo apt-get update
        sudo apt-get install -y wkhtmltopdf pdftk
        ```
    -   **For openSUSE**:
        ```sh
        sudo zypper refresh
        sudo zypper install -y wkhtmltopdf pdftk
        ```

2.  **Build the Application**:
    Open a terminal and run the following command from the project root:
    ```sh
    go build -o mypdfservice ./cmd/webapi/main.go
    ```

3.  **Run the Service**:
    Execute the compiled binary:
    ```sh
    ./mypdfservice
    ```
    The server will start on port `8080`. To specify a different port:
    ```sh
    ./mypdfservice 9999
    ```

### Docker

A `docker-compose` setup is provided for easy local deployment.

1.  **Prerequisite**: Ensure you have Docker and Docker Compose installed.

2.  **Build and Run**:
    Navigate to the `cmd/docker` directory and run:
    ```sh
    cd cmd/docker
    docker-compose up --build
    ```
    This command will build the Docker image and start the service. The API will be accessible at `http://localhost:8080`.

### Google Cloud Run

You can deploy the application as a serverless container to Google Cloud Run.

1.  **Prerequisites**:
    -   Install the [Google Cloud CLI](https://cloud.google.com/sdk/docs/install).
    -   Authenticate and configure the CLI with your GCP project: `gcloud auth login` and `gcloud config set project YOUR_PROJECT_ID`.
    -   Enable the Cloud Build and Cloud Run APIs for your project.

2.  **Deploy**:
    From the project root directory, run the following command:
    ```sh
    gcloud builds submit --config cloudbuild.yaml .
    ```
    This command uses Cloud Build to build the container image, push it to the Google Container Registry, and deploy it as a new service on Cloud Run.

## API Usage

### `/html2pdf` Endpoint

-   **Method**: `POST`
-   **Description**: Converts an HTML file to a PDF.
-   **Input**: A `multipart/form-data` request containing a single `.zip` file. The zip file must include:
    -   One `.html` file.
    -   All image files (`.jpg`, `.jpeg`, `.png`, `.gif`) referenced by `<img>` tags in the HTML. Image paths in the `src` attribute must match the filenames in the zip archive.

### `/merge` Endpoint

-   **Method**: `POST`
-   **Description**: Merges multiple PDF files into one.
-   **Input**: A `multipart/form-data` request containing a single `.zip` file with all the `.pdf` files to be merged. The files will be merged in alphabetical order of their filenames.

## Testing with cURL

Here are examples of how to test the endpoints using cURL.

-   **Merge PDFs**:
    ```sh
    curl -X POST http://localhost:8080/merge \
      -F "files=@/path/to/your/merge-example.zip" \
      -H "Content-Type: multipart/form-data" \
      -o merged_output.pdf
    ```

-   **Convert HTML to PDF**:
    ```sh
    curl -X POST http://localhost:8080/html2pdf \
      -F "files=@/path/to/your/html-example.zip" \
      -H "Content-Type: multipart/form-data" \
      -o converted_output.pdf
    ```

## Logging

The application maintains a detailed debug log in `mypdfservice_debug.log`, created in the same directory as the executable. This log is rotated automatically to manage disk space:
-   The log file is limited to 10 MB.
-   When the limit is reached, the file is compressed (e.g., `mypdfservice_debug-TIMESTAMP.log.gz`).
-   A maximum of 20 old log files are kept.
-   Logs older than 90 days are automatically deleted.

You can also capture a summary log of console output by redirecting the standard output:
```sh
# Windows
./mypdfservice.exe > summary.log

# Linux
./mypdfservice > summary.log
```

## Acknowledgements

-   Many thanks to the developers of the **wkhtmltopdf** and **pdftk** projects. This application would not be possible without their excellent work.
-   This project was developed entirely using the **Gitpod** cloud development environment.

## License

// TODO