package components

// Go has great support for templates, so we will use that to build the people
// table that gets displayed in the UI.
//
// The template itself in the /templates/peopleTableTemplate.html file, but
// for the demo, we embed it in the binary via the embedded_resources.go file.
import (
	"bytes"
	"fmt"
	"hello_there/repo"
	"html/template"
)

// ----------------------------------------------------------------------------
// used to create the list of people the template parses to create the html
// syntax for the people table
type peopleTableEntry struct {
	Num     int
	Name    string
	Surname string
	Age     int
}

// ----------------------------------------------------------------------------
// uses the people data to build the HTML elements required to display them
func BuildPeopleTable(data *[]repo.Person, htmlTemplate string) string {
	peopleData := make([]peopleTableEntry, 0)

	for pos, person := range *data {
		peopleData = append(peopleData, peopleTableEntry{
			Num:     pos + 1,
			Name:    person.FirstName,
			Surname: person.LastName,
			Age:     person.Age})
	}

	template := template.Must(template.New("table").Parse(htmlTemplate))

	var buffer bytes.Buffer
	err := template.Execute(&buffer, peopleData)
	if err != nil {
		fmt.Println(err)
		return "<p>...Error processing template...</p>"
	}

	return buffer.String()
}
