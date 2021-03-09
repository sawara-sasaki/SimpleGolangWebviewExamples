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
a {
  text-decoration:none;
}
a:link,
a:visited {
  color: #333;
}
label {
  background-color: #EEE;
}
#main {
  width: 100vw;
  height: calc(100vh - 50px);
  display: flex;
  flex-direction: column;
  justify-content: center;
}
#url-form {
  width: 50vw;
  height: 30px;
  margin: 20px auto;
  display: flex;
  flex-direction: row;
  justify-content: center;
}
#url-input {
  width: 100%;
  margin: 0;
}
#buttons {
  width: 200px;
  height: 50px;
  margin: 20px auto;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.link-button {
  width: 50px;
  height: 50px;
  cursor: pointer;
  border: none;
  border-radius: 25px;
  background-color: #EEE;
  font-size: 20px;
  line-height: 50px;
}
#setting {
  width: 100vw;
  height: 50px;
  display: none;
}
#setting-open-button-container {
  width: 100vw;
  height: 50px;
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}
#setting-open-button {
  width: 70px;
  height: 30px;
  margin-right: 30px;
  cursor: pointer;
  border: none;
  border-radius: 15px;
  background-color: #EEE;
  line-height: 30px;
}
#change-color-button {
  display: inline-block;
  width: 70px;
  height: 20px;
  cursor: pointer;
  border: solid;
  border-width: thin;
  border-radius: 5px;
  border-color: #777;
  background-color: #EEE;
  font-size: 0.9em;
  line-height: 20px;
}
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
        <a href="#" onclick="local('memo.tpl');"><div class="link-button">M</div></a>
        <a href="#" onclick="local('links.tpl');"><div class="link-button">L</div></a>
      </div>
    </div>
    <div id="setting-open-button-container">
      <div id="setting-open-button" onclick="settingOpen();">Setting</div>
    </div>
    <div id="setting">
      <form id="setting-form">
        <label>rgb:</label>
        <input id="red" type="number" max="255" min="0" value="0">
        <input id="green" type="number" max="255" min="0" value="0">
        <input id="blue" type="number" max="255" min="0" value="0">
        <div id="change-color-button" onclick="changeColor();">Change</div>
      </form>
    </div>
    <script>
    const urlForm = document.getElementById("url-form");
    urlForm.addEventListener("submit", function() {
      navigate(document.getElementById("url-input").value);
    })
    var changeColor = function() {
      var r = document.getElementById("red").value;
      var g = document.getElementById("green").value;
      var b = document.getElementById("blue").value;
      document.body.style.backgroundColor = "rgba(".concat(r, ",", g, ",", b, ",1)");
      log("rgba(".concat(r, ",", g, ",", b, ",1)"));
    };
    var settingOpen = function() {
      var settingOpenButtonContainerElem = document.getElementById("setting-open-button-container");
      settingOpenButtonContainerElem.style.display = "none";
      var settingElem = document.getElementById("setting");
      settingElem.style.display = "block";
    };
    </script>
  </body>
</html>
{{end}}