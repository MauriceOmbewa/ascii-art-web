# ASCII ART WEB
## Description
Ascii-art-web is a web-based implementation of the ascii-art project. It allows users to input text and select a banner to generate ascii art.

## Implementation Details: Algorithm

The program reads user input, processes it, and generates ASCII art based on the selected banner. Here is the overview of the algorithm:

    User Input:
        The user enters text in the input field.
        They select a banner (Thinkertoy, Shadow, or Standard).

    Processing:
        The program splits the input text into lines.
        For each character in the input, it looks up the corresponding ASCII art representation from the chosen banner file.
        The program combines the lines of ASCII art to create the final output.

    Rendering:
        The server renders the HTML template (ascii-art.html) with the input text and generated ASCII art.
        The client sees the result in their web browser.

## Usage: How to run

1.  Clone this repository to your local machine.
```go
    git clone https://learn.zone01kisumu.ke/git/skisenge/ascii-art-web
```
2.  Ensure you have Go installed.
3.  Navigate to the project directory.
4.  Run the following command to start the server:
```bash
    cd asciiwebproject
    go run main.go
```
5.  Open a web browser and visit `http://localhost:8080`.
    - This usage is also applicable:
    ```bash
    go run main.go [port]
    ```
      
    - For example,
    ```bash
    go run main.go 8000
    ```
6.  Enter your text and choose a banner to see the generated ASCII art by clicking on the `Generate Ascii Art` button
7.  Click `Back to HomePage` to navigate back to test other entries.

## Author
- [@mombewa](https://learn.zone01kisumu.ke/git/mombewa)
