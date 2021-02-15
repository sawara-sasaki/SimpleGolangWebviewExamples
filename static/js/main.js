var changeColor = function() {
  var r = document.getElementById("red").value;
  var g = document.getElementById("green").value;
  var b = document.getElementById("blue").value;
  document.body.style.backgroundColor = "rgba(".concat(r, ",", g, ",", b, ",1)");
}
