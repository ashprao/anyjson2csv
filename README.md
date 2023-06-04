# anyjson2csv
Converts any json file to a csv

## Building the project
- Open a terminal and navigate to the directory containing the Makefile and your Go code.

- To compile for the current platform, run the command `make`.

- To compile for Linux, run the command `make build-linux`.

- To compile for Windows, run the command `make build-windows`.

- To compile for macOS, run the command `make build-macos`.

After running the appropriate make command, you will find the compiled binaries in the build directory with the following structure:

```
build/
  - linux/
      - anyjson2csv
  - windows/
      - anyjson2csv.exe
  - macos/
      - anyjson2csv
```
## Usage

```
Usage: /path/to/anyjson2csv [options] <source file>
Options:
  -all
        Includes json object and list columns
  -debug
        Enable debug mode
  -output string
        Destination file name (default "output.csv")
Only one argument referencing the source json file needs to be provided.
```

## Example output

For the following json input:

```json
[
    {
        "image_id": "1",
        "category": "category1",
        "attributes": [
            {
                "attribute": "attribute1",
                "tag": "tag1",
                "prev_tag": "tag2"
            },
            {
                "attribute": "attribute2",
                "tag": "tag3",
                "prev_tag": "tag4"
            },
            {
                "attribute": "attribute3",
                "tag": "tag5",
                "prev_tag": "tag6"
            }
        ]
    },
    {
        "image_id": "2",
        "category": "category1",
        "attributes": [
            {
                "attribute": "attribute1",
                "tag": "tag2",
                "prev_tag": "tag1"
            },
            {
                "attribute": "attribute2",
                "tag": "tag7",
                "prev_tag": ""
            },
            {
                "attribute": "attribute3",
                "tag": "tag8",
                "prev_tag": "tag5"
            }            
        ]
    },
    {
        "image_id": "3",
        "category": "category1",
        "attributes": [
            {
                "attribute": "attribute1",
                "tag": "tag9",
                "prev_tag": ""
            },
            {
                "attribute": "attribute2",
                "tag": "tag4",
                "prev_tag": ""
            },
            {
                "attribute": "attribute3",
                "tag": "tag5",
                "prev_tag": ""
            }            
        ]
    }
]
```
This is the output CSV after running `anyjson2csv`:

```
image_id,category,attribute,tag,prev_tag
1,category1,attribute1,tag1,tag2
,,attribute2,tag3,tag4
,,attribute3,tag5,tag6
2,category1,attribute1,tag2,tag1
,,attribute2,tag7,
,,attribute3,tag8,tag5
3,category1,attribute1,tag9,
,,attribute2,tag4,
,,attribute3,tag5,
```
When viewed in Excel or Sheets:

| image_id | category  | attribute  | tag  | prev_tag |
|----------|-----------|------------|------|----------|
| 1        | category1 | attribute1 | tag1 | tag2     |
|          |           | attribute2 | tag3 | tag4     |
|          |           | attribute3 | tag5 | tag6     |
|2|category1|attribute1|tag2|tag1|
|||attribute2|tag7||
|||attribute3|tag8|tag5|
|3|category1|attribute1|tag9||
|||attribute2|tag4||
|||attribute3|tag5||

**Note:** The order of the columns is not guaranteed in the generated CSV file. This will be addressed in an upcoming update.