# JSON Parsing

[Why struct tags](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)

* Struct tags are small pieces of metadata attached to fields of a struct that provide instructions to other Go code that works with the struct

* Other Go code is then capable of examining these structs and extracting the values assigned to specific keys it requests

* Struct tags have no effect on the operation of your code without some other code that examines them

* To use struct tags to accomplish something, other Go code must be written to examine structs at runtime. The standard library has packages that use struct tags as part of their operation. The most popular of these is the encoding/json package.

* The JSON encoder in the standard library makes use of struct tags as annotations indicating to the encoder how you would like to name your fields in the JSON output. These JSON encoding and decoding mechanisms can be found in the encoding/json package.

* By placing the camel-cased version of the field names as the value to the json key, the encoder will use that name instead.

Followed the awesome tutorial at [Tutorial Edge](https://tutorialedge.net/golang/parsing-json-with-golang/), made some modifications for go 1.16 and some parsing url as well
