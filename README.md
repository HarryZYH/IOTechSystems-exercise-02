# IOTechSystems-exercise-02
This is a solution to exercise02 in [IOTech Coding Exercises](https://github.com/IOTechSystems/exercises) written in Go. A dependency [gojsonschema](https://github.com/xeipuuv/gojsonschema) is used for validating the output format.

## Running the program
1. Verify that you have Go installed on your computer by checking running the following command. If you don't, you can download it from the official website: https://golang.org/dl/
    ```sh
    go version
    ```
2. Pull the repository to local. Open a terminal and naviaget to the directory where you saved the file.
    ```sh
    cd IOTechSystems-exercise-02
    ```

  - To run
    ```sh
    go run main.go
    ```
  - Build and run
    ```sh
    go build
    ./solution
    ```
    
3. You should see a [`output.json`](./output.json) file generated and two following messages displyed in the terminal.
  - "Ouput written to output.json"
    Verify the answers are written into the file.
  - "The output file is valid"
    Verify that the output is following the schema in [`schema.json`](./output-schema/schema.json)

## Test
1. Run the test cases by running the following command.
    ```sh
    go test
    ```