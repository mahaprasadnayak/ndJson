# ndJson



NDJSON is a convenient format for storing or ```streaming structured data that may be processed one record at a time```. It works well with unix-style text processing tools and shell pipelines. It's a great format for log files. It's also a flexible format for passing messages between cooperating processes.

##. Example NDJSON
~~~~~
 {"some":"thing"}
 {"foo":17,"bar":false,"quux":true}
 {"may":{"include":"nested","objects":["and","arrays"]}}
~~~~~
(with `\n` line separators)

1.Line Separator is '\n'
This means '\r\n' is also supported because trailing white space is ignored when parsing JSON values.


2. Each Line is a Valid JSON Value
The most common values will be objects or arrays, but any JSON value is permitted. See json.org for more information about JSON values.
## please visit for more info "ndjson.org"
