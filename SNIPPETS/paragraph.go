package main

import (
    "fmt"
    "os"
    "html/template"
)

func main() {

    const (
        paragraph_hypothesis = 1 << iota
        paragraph_attachment = 1 << iota
        paragraph_menu       = 1 << iota
    )

    const text = "{{.Paratype | printpara}}\n" // A simple test template

    type Paragraph struct {
        Paratype int
    }

    var paralist = []*Paragraph{
        &Paragraph{paragraph_hypothesis},
        &Paragraph{paragraph_attachment},
        &Paragraph{paragraph_menu},
    }

    t := template.New("testparagraphs")

    printPara := func(paratype int) string {
        text := ""
        switch paratype {
        case paragraph_hypothesis:
            text = "This is a hypothesis\n"
        case paragraph_attachment:
            text = "This is an attachment\n"
        case paragraph_menu:
            text = "Menu\n1:\n2:\n3:\n\nPick any option:\n"
        }
        return text
    }

    template.Must(t.Funcs(template.FuncMap{"printpara": printPara}).Parse(text))

    for _, p := range paralist {
        err := t.Execute(os.Stdout, p)
        if err != nil {
            fmt.Println("executing template:", err)
        }
    }
}
