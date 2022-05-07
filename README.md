Just working through the basics as a refresher: 

1) (Introduction)[https://go.dev/doc/tutorial/create-module]
2) (Create a Module)[https://go.dev/doc/tutorial/create-module]

Worth remembering... locating local dependencies between local modules
`$ go mod edit -replace example.com/greetings=../greetings`
`$ go mod tidy`