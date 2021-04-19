{{define "base"}}
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
    <meta http-equiv="Content-Security-Policy" content="script-src 'unsafe-inline' 'unsafe-eval';">
    <title>WebView Example</title>
    <style type="text/css">
{{template "main.css" .}}
    </style>
  </head>
  <body>
    <div id="main">
      <h1>WebView Example</h1>
      <form id="url-form">
        <input id="url-input" type="text">
      </form>
      <div id="buttons">
        <a href="https://github.com/"><div class="link-button">G</div></a>
        <span><div id="memo" class="link-button">M</div></span>
        <span><div id="links" class="link-button">L</div></span>
      </div>
    </div>
    <div id="setting-open-button-container">
      <div id="setting-open-button">Setting</div>
    </div>
    <div id="setting">
      <form id="setting-form">
        <label>rgb:</label>
        <input id="red" type="number" max="255" min="0" value="0">
        <input id="green" type="number" max="255" min="0" value="0">
        <input id="blue" type="number" max="255" min="0" value="0">
        <div id="change-color-button">Change</div>
      </form>
    </div>
  </body>
</html>
{{end}}
