# Family Tree Command Line Tool

This command line tool allows users to define and manage their family tree. It provides options to add people, specify relationships, and query the family tree.

## Getting Started

To use the family tree command line tool, follow the instructions below.

### Prerequisites

- [Language Runtime] - The tool is implemented in [Language of your choice].

### Installation

1. Download the family tree tool from the [GitHub repository](https://github.com/yourusername/family-tree) or [Google Drive Link](https://yourgoogledrive/family-tree.zip).

2. Extract the contents of the ZIP file.

3. Navigate to the directory where you extracted the files.

### Usage

# Family Tree Command Line Tool

Design a command line system to help a user define their family tree. This program should be implemented in any programming language of your choice. The objective is to evaluate your design skills, consider future complexity and evolution, and follow functional programming or good object-oriented principles.

## Requirements

1. Design a command line tool named `family-tree` with the following options:
   - `family-tree add person <name>`: Add a person to the family tree.
     - Example: `family-tree add person Amit Dhakad`
   - `family-tree add relationship <name>`: Add a relationship to a person.
     - Example: `family-tree add relationship father` or `family-tree add relationship son`
   - `family-tree connect <name1> as <relationship> of <name2>`: Connect two people in the family tree with a specified relationship.
     - Example: `family-tree connect Amit Dhakad as son of KK Dhakad`

2. Queries: Based on the relationships created, the program should support the following queries:
   - `family-tree count sons of <name>`: Count the number of sons of a person.
   - `family-tree count daughters of <name>`: Count the number of daughters of a person.
   - `family-tree count wives of <name>`: Count the number of wives of a person.
   - `family-tree father of <name>`: Retrieve the name of the father of a person.

3. The family tree should ignore the date of birth (DoB).

4. Visual Representation: Pink should represent females.

## Implementation

You can use the provided code sample as a starting point for your implementation. This sample is written in the Go programming language.

## How to Run

To run the program:
1. Ensure you have Go installed on your computer. If not, download and install it from [here](https://golang.org/dl/).

2. Save the provided code to a file with a `.go` extension (e.g., `family_tree.go`).

3. Open your terminal or command prompt.

4. Navigate to the directory where you saved the `family_tree.go` file using the `cd` command.

5. Build the Go program with the following command: `go build family_tree.go`


6. Run the program with various commands and flags:
- To add a person:
  ```
  ./family_tree add -person "Amit Dhakad"
  ```
- To add a relationship:
  ```
  ./family_tree add -relationship "father"
  ```
- To connect two people:
  ```
  ./family_tree connect -name1 "Amit Dhakad" -relationship "son" -name2 "KK Dhakad"
  ```
- To count sons:
  ```
  ./family_tree count -type "son" -name "KK Dhakad"
  ```
- To get the father of a person:
  ```
  ./family_tree father -name "Amit Dhakad"
  ```

7. The program will display the results based on the commands you input. Replace `/path/to/your/code/directory` with the actual path to your code directory.

Feel free to modify and expand upon this implementation as needed for your coding interview.
