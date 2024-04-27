# Go CLI Form

A quick form creator library

## Installation:

run `go get github.com/DimRev/go-cli-form`

To Create a quick cli form put init the form struct using `Form.Start()`

```go
func main(){
  Form := Form.Start()
    //... Rest of the form code
  Form.End()
}
```

After initializing the for you can use the form's input functions:

- `Form.TextInput(qst string) string` - Free form text input
- `Form.SelectInput(qst string, opts []string) string` - Use left/right arrows to navigate between the option values from the opts slice, use enter to submit the selection
- `Form.MultiSelectInput(qst string, opts []string) []string` - Using up/down to navigate the options from the options slice, use space to select the option, and enter to submit selection
