{{define "background.css"}}
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
}
{{end}}
{{define "main.css"}}
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
{{end}}
