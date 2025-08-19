# Gemini Code Assistant Context

This document provides context for the Gemini Code Assistant to help it understand the project and provide more relevant assistance.

## Project Overview

This project is an HTML to PDF converter written in Golang. It provides both a command-line interface (CLI) and a web API for converting HTML files to PDF. The project also includes functionality to merge multiple PDF files into a single document.

The core of the project relies on two external command-line tools:
*   **wkhtmltopdf**: For converting HTML to PDF.
*   **pdftk**: For merging PDF files.

## Key Technologies

*   **Golang**: The primary programming language.
*   **wkhtmltopdf**: External tool for HTML to PDF conversion.
*   **pdftk**: External tool for PDF merging.

## Project Structure

The project is organized into the following main directories:

*   `cmd/`: Contains the entry points for the two applications.
    *   `cli/`: The command-line interface application.
    *   `webapi/`: The web API application.
*   `services/`: Contains the business logic for the application's services.
    *   `Html2pdfService.go`: Service for converting HTML to PDF.
    *   `mergePdfsService.go`: Service for merging PDF files.
    *   `imagesEmbedder.go`: Service for embedding images into HTML.
    *   `pdfMerger.go`: Wrapper for the `pdftk` tool.
    *   `wkhtmltopdfConverter.go`: Wrapper for the `wkhtmltopdf` tool.
*   `webapi/`: Contains the web-specific components for the web API.
    *   `app/`: Application setup and URL mappings.
    *   `controller/`: HTTP controllers for handling web requests.
*   `scripts/`: Contains utility scripts for building and testing the application.
*   `zip-input-files-examples/`: Contains example ZIP files for testing the services.

## How to Run

### Web API

To run the web API, execute the following command:

```bash
go run cmd/webapi/main.go
```

The API will be available at `http://localhost:8080`.

### CLI

To run the CLI, you can use the following command:

```bash
go run cmd/cli/main.go --help
```

This will display the available commands and options for the CLI.

## How to Deploy to Google Cloud Run

This project can be deployed to Google Cloud Run using Google Cloud Build.

1.  **Enable the Cloud Build and Cloud Run APIs** in your Google Cloud project.
2.  **Set your project ID**:
    ```bash
    gcloud config set project YOUR_PROJECT_ID
    ```
3.  **Run the build**:
    ```bash
    gcloud builds submit --config cloudbuild.yaml .
    ```

This will build the container image, push it to Google Container Registry, and deploy it to Google Cloud Run.

## Available Services

### HTML to PDF Conversion

*   **Web API Endpoint**: `POST /html2pdf`
*   **Description**: Converts an HTML file to a PDF. The request should be a multipart/form-data request with a ZIP file containing the HTML file and any associated assets (images, CSS, etc.).

### PDF Merging

*   **Web API Endpoint**: `POST /mergepdfs`
*   **Description**: Merges multiple PDF files into a single PDF. The request should be a multipart/form-data request with a ZIP file containing the PDF files to be merged.
