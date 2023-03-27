package main

type Stuff struct {
	inn int
}

type Person struct {
	Name string
}

type SecretAgent struct {
	Person
	Stuff
}
