# Golang Covid Summary API

This is a Golang application that provides a basic API endpoint to retrieve a COVID-19 summary.

## Structure
Project structure is based on [pallat/gotemplate](https://github.com/pallat/gotemplate) which focuses on straghtforward un-layered structure.
```
project-root-directory/
    app/                        # containing things that will be shared among packages
        gin.go                  # gin wrapper (for hexagonal purpose)
        covid/                  # package covid for related covid summary stuff
            covid.go            # business logic
            api.go              # API Adapter (for hexagonal purpose)
            test_files          # unit test and mocks   
    main.go
    test/                       # contains end-to-end test file
```

## Remarks
- Writing end-to-end test with go `"testing"` is not convinice as its execute along unit tests while the application is not running, resulting in always "FAIL" result.
To actually run just the unit test, the `go test` command must be modify which might not be conform with the CI pipeline

- The project structure, while it is simple and straighforward, leaves a lot room for devs to be creative (and too flexible) which can sometimes leads to messy project.