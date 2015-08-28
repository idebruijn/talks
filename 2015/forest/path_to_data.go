<?xml version="1.0"?>
<Book title="Learning Go the hard way">
	<Chapter title="Foreword"></Child>
	<Chapter title="Introduction"></Child>					
</Root>


chapters := XMLPath(t, r, "Book//Chapter")


{
	"Volumes": [
		{"Name":"C"},
		{"Name":"D"},
		{"Name":"E"}
	]
}


e := JSONPath(t, r, ".Volumes.2.Name")


