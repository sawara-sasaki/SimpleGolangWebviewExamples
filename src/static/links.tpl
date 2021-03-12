{{define "base"}}
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
    <title>WebView Example</title>
    <style type="text/css">
{{template "background.css" .}}
body {
  display: flex;
  justify-content: space-between;
  overflow: auto;
}
body > div  {
  margin: 50px 30px;
}
ul {
  list-style: none;
}
a {
  text-decoration:none;
}
a:link,
a:visited {
  color: #1A73F0;
}
    </style>
  </head>
  <body>
    <div>
      <ul>
{{ range .Links }}
        <li><a href="{{ .Url }}">{{ .Title }}</a></li>
{{end}}
      </ul>
    </div>
    <div>
      <ul>
        <li><a href="#" onclick="local('index.tpl');">Top</a></li>
      </ul>
    </div>
  </body>
</html>
{{end}}
