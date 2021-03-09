{{define "base"}}
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
    <title>WebView Example</title>
    <style type="text/css">
html {
  scroll-behavior: smooth;
  text-align: center;
}
body {
  background-image:url('{{template "sample.jpg" .}}');
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  max-width: 100%;
  max-height: 100%;
}
#main {
  width: 100vw;
  height: 100vh;
  display: flex;
  flex-direction: column;
  justify-content: center;
}
#memo-form {
  width: 60vw;
  margin: 20px auto;
  display: flex;
  flex-direction: row;
  justify-content: center;
}
#memo-textarea {
  width: 100%;
  height: 50vh;
}
#buttons {
  width: 60vw;
  height: 32px;
  margin: 0 auto;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
#file-buttons {
  height: 32px;
  display: flex;
}
#file-buttons > .memo-button {
  margin-right: 10px;
}
#link-buttons {
  height: 32px;
  display: flex;
  justify-content: flex-end;
}
#link-buttons > .memo-button {
  margin-left: 10px;
}
.memo-button {
  width: 5em;
  height: 30px;
  margin: 0;
  cursor: pointer;
  border: solid;
  border-radius: 3px;
  border-color: #777;
  border-width: thin;
  background-color: #F0F0F0;
  text-align: center;
  line-height: 30px;
}
    </style>
  </head>
  <body>
    <div id="main">
      <form id="memo-form">
        <textarea id="memo-textarea"></textarea>
      </form>
      <div id="buttons">
        <div id="file-buttons">
          <div class="memo-button" onclick="save();">Save</div>
          <div class="memo-button" onclick="load();">Load</div>
        </div>
        <div id="link-buttons">
          <div class="memo-button" onclick="local('index.tpl');">Top</div>
        </div>
      </div>
    </div>
    <script>
    var save = function() {
      write(document.getElementById("memo-textarea").value);
    };
    var load = function() {
      read().then(function(res) {
        document.getElementById("memo-textarea").value = res;
      });
    };
    </script>
  </body>
</html>
{{end}}