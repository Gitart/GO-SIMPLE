func ParseTemplates() (*template.Template, error) {
    templateBuilder := template.New("")
    if t, _ := templateBuilder.ParseGlob("/*/*/*/*/*.tmpl"); t != nil {
        templateBuilder = t
    }
    if t, _ := templateBuilder.ParseGlob("/*/*/*/*.tmpl"); t != nil {
        templateBuilder = t
    }
    if t, _ := templateBuilder.ParseGlob("/*/*/*.tmpl"); t != nil {
        templateBuilder = t
    }
    if t, _ := templateBuilder.ParseGlob("/*/*.tmpl"); t != nil {
        templateBuilder = t
    }
    return templateBuilder.ParseGlob("/*.tmpl")
}
