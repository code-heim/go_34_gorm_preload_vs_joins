{{ define "index.tpl"}}
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <meta charset="utf-8">
            <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
            <meta name="description" content="">
            <meta name="author" content="">

            <title>Blogs App</title>

            <!-- Custom fonts for this template -->
            <link href="https://fonts.googleapis.com/css?family=Lato:300,400,700,300italic,400italic,700italic" rel="stylesheet" type="text/css">

            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
            <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-YvpcrYf0tY3lHB60NNkmXc5s9fDVZLESaAA55NDzOxhy9GkcIdslK1eN7N6jIeHz" crossorigin="anonymous"></script>

        </head>
        <body>
            <div class="container mt-4">
                <br/>
                <br/>
                <br/>
                <div>
                    
                    {{ with .notes }}
                        {{ range . }}
                            <div class="card" style="width: 18rem;">
                                <div class="card-body">
                                    <h5 class="card-title">{{ .Name }}</h5>
                                    <p class="card-text">{{ .Content }}</p>
                                    <p class="card-text"> - {{ .User.Name }} ({{ .User.Username }}) </p>
                                </div>
                            </div>
                            <br/>
                        {{ end }}
                    {{ end }}
                </div>
                <br/>

                <div>
                    {{ with .joinData }}
                        {{ range . }}
                            <div class="card" style="width: 18rem;">
                                <div class="card-body">
                                    <h5 class="card-title">{{ .Name }}</h5>
                                    <p class="card-text"> - {{ .Username }} </p>
                                </div>
                            </div>
                            <br/>
                        {{ end }}
                    {{ end }}
                </div>

                <div>
                    {{ with .noteCount }}
                        {{ range . }}
                            <div class="card" style="width: 18rem;">
                                <div class="card-body">
                                    <h5 class="card-title">{{ .Username }}</h5>
                                    <p class="card-text"> - Count --- {{ .NoteCount }} </p>
                                </div>
                            </div>
                            <br/>
                        {{ end }}
                    {{ end }}
                </div>
            </div>
        </body>
    </html>
{{ end }}