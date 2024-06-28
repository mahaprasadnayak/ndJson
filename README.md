# ndJson: Newline-Delimited Json

**ndJson** (Newline-Delimited JSON) is a convenient format for storing or streaming structured data that may be processed one record at a time. This format is particularly useful for log files and for passing messages between cooperating processes. NDJSON works well with Unix-style text processing tools and shell pipelines.

## Why ndJson?

- **Stream Processing**: Process one record at a time, making it efficient for large datasets.
- **Compatibility**: Works seamlessly with Unix-style text processing tools.
- **Flexibility**: Ideal for log files and inter-process communication.

## Example ndJson

```ndJson
{"some":"thing"}
{"foo":17,"bar":false,"quux":true}
{"may":{"include":"nested","objects":["and","arrays"]}}

```
## example in Go that demonstrates how to read and decode ndJson data
```
	reqBody := `{"some":"thing"}
{"foo":17,"bar":false,"is_data":true}`

	ndjsonReader := ndjson.NewReader(strings.NewReader(reqBody))
	for ndjsonReader.Next() {
		var data interface{}
		if err := ndjsonReader.Decode(&data); err != nil {
			//handle error
			
		}
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			//handle error
		}
		fmt.Printf("Decoded data: %s\n", string(jsonData))
	}
```

## Key Features
- **Line Separator is \n**: This means \r\n is also supported because trailing whitespace is ignored when parsing JSON values.
- **Each Line is a Valid JSON Value**: The most common values are objects or arrays, but any JSON value is permitted. See json.org for more information about JSON values.
  
## Use Cases
- **Log Files**: Easily store and process logs.
- **Data Streaming**: Efficiently handle large streams of JSON data.
- **Inter-process Communication**: Pass messages between processes in a flexible format.

## Learn More
- **For more information, please visit** [ndJson.org](http://ndjson.org).
