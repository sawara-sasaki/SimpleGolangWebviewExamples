{{define "base"}}
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
    <title>WebView Example</title>
    <style type="text/css">
html {
  scroll-behavior: smooth;
}
body {
  background-image:url('{{template "sample.jpg" .}}');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  max-width: 100%;
  max-height: 100%;
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
        <li><a href="https://github.com/">GitHub</a></li>
        <li><a href="https://github.com/webview/webview">webview</a></li>
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
