# ATTS Breed Statistics

ATTS Breed Statistics is a Go project designed to facilitate the extraction and analysis of temperament test results from the American Temperament Test Society's website [https://atts.org/](https://atts.org/). As the original site lacks comprehensive sorting and filtering functionalities, this tool aims to enable users to effectively compare breeds and their corresponding test scores.

## Purpose

The primary objective of this project is to automate the retrieval of temperament test data and provide users with the ability to manipulate, sort, and filter this data to gain insights into various dog breeds' behavioral tendencies. The project's functionalities enable the extraction of data from the website, storage in a structured format, and subsequent parsing and analysis.

## Functionality

The project's main functionalities include:

- Web scraping of temperament test results from the ATTS website.
- Storage of extracted data in a structured format within the `scores.txt` file.
- Parsing, sorting, and filtering of test results to provide users with customizable data views.
- Exporting the parsed and sorted data to the `outputs.txt` file for convenient analysis and comparison.

## Usage

To utilize the ATTS Breed Statistics tool, follow the instructions outlined in the "Running Locally" section. This will enable you to extract, manipulate, and analyze the temperament test data as desired, providing valuable insights into the behavioral traits of various dog breeds.

Feel free to modify and enhance the tool's functionalities to cater to your specific analysis requirements and preferences.

This enhanced overview provides a more detailed explanation of the project's purpose, functionality, and usage. Adjust the content as needed to accurately reflect your project's scope and objectives.

## Running Locally

- **Fork and Clone the Repository**: First, fork the repository on GitHub and then clone it onto your local computer using the following command:

```bash
git clone https://github.com/your-username/your-repo.git
```

- **Install Go**: Make sure that Go is installed on your computer. You can download it from the [official Go website](https://go.dev/doc/install).

- **Initialize Go Modules**: Navigate to the project directory and use the following command to initialize Go modules:

```bash
go mod init github.com/your-username/your-repo
```

- **Run without Building an Executable File**: To run the project without building an executable file, use the following command:

```bash
go run .
```

- **Build an Executable**: To build an executable, navigate to the project directory and use the following commands:

```bash
go build
```

This will create an executable file. To execute the file, use the following command:

```bash
./atts_breed_scores
```

Make sure to replace `atts_breed_scores` with the actual name of your executable file. If there are any additional setup steps or configuration requirements, make sure to include them in your README as well.

## Future Updates

ATTS Breed Statistics is currently a side project aimed at facilitating the learning process for Go development. The intention is to expand the project's capabilities and user-friendliness in the future. Some of the planned updates include the following:

- **Enhanced User Interaction**: Implementing command-line arguments for the executable to facilitate both terminal and file output options. This update will allow users to specify the desired file name and location for output, enhancing the flexibility and usability of the tool.

- **Versatile Data Formats**: Providing support for various data formats, including structured tables and JSON, to accommodate diverse user preferences. This enhancement will allow users to seamlessly integrate the extracted data into their existing workflows or systems.

- **Advanced Sorting and Filtering**: Introducing additional parameters for sorting and filtering to enable users to retrieve specific data tailored to their unique use cases. This functionality will empower users to perform targeted analyses and comparisons, making it easier to derive meaningful insights from the temperament test data.

These planned updates are aimed at making ATTS Breed Statistics more versatile, user-friendly, and adaptable to different usage scenarios. The project will continue to evolve to meet the needs of both learners and professionals interested in exploring dog breed temperament data.
