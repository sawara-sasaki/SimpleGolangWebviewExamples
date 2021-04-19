var simpleGolangWebview = {};
simpleGolangWebview.isLocal = window.location.href.startsWith('data');
simpleGolangWebview.eventSetting = function() {
  const urlForm = document.getElementById("url-form");
  if (!!urlForm) {
    urlForm.addEventListener("submit", function() {
      navigate(document.getElementById("url-input").value);
    })
  }
  const changeColorButton = document.getElementById("change-color-button");
  if (!!changeColorButton) {
    changeColorButton.addEventListener("click", function() {
      var r = document.getElementById("red").value;
      var g = document.getElementById("green").value;
      var b = document.getElementById("blue").value;
      document.body.style.backgroundColor = "rgba(".concat(r, ",", g, ",", b, ",1)");
      log("rgba(".concat(r, ",", g, ",", b, ",1)"));
    });
  }
  const settingOpenButton = document.getElementById("setting-open-button");
  if (!!settingOpenButton) {
    settingOpenButton.addEventListener("click", function() {
      var settingOpenButtonContainerElem = document.getElementById("setting-open-button-container");
      settingOpenButtonContainerElem.style.display = "none";
      var settingElem = document.getElementById("setting");
      settingElem.style.display = "block";
    });
  }
  const memoLink = document.getElementById("memo");
  if (!!memoLink) {
    memoLink.addEventListener("click", function() {
      local('memo.tpl');
    });
  }
  const linksLink = document.getElementById("links");
  if (!!linksLink) {
    linksLink.addEventListener("click", function() {
      local('links.tpl');
    });
  }
  const topLink = document.getElementById("top");
  if (!!topLink) {
    topLink.addEventListener("click", function() {
      local('index.tpl');
    });
  }
  const saveButton = document.getElementById("save");
  if (!!saveButton) {
    saveButton.addEventListener("click", function() {
      write(document.getElementById("memo-textarea").value);
    });
  }
  const loadButton = document.getElementById("load");
  if (!!loadButton) {
    loadButton.addEventListener("click", function() {
      read().then(function(res) {
        document.getElementById("memo-textarea").value = res;
      });
    });
  }
};
window.onload = function() {
  if (!simpleGolangWebview.isLocal) {
    Array.prototype.forEach.call(document.getElementsByTagName('a'), (element) => element.setAttribute('target', ''));
  } else {
    simpleGolangWebview.eventSetting();
  }
  document.body.addEventListener('keydown',
    event => {
      if (event.key === 'c' && event.ctrlKey) {
        if (!simpleGolangWebview.isLocal) {
          saveCookie(window.location.href, document.cookie);
        }
      } else if (event.key === 'v' && event.ctrlKey) {
        debug();
      } else if (event.key === 's' && event.ctrlKey) {
        if (!simpleGolangWebview.isLocal) {
          saveSource(window.location.href, document.documentElement.innerHTML);
        }
      }
    });
  document.body.addEventListener('click',
    event => {
      if (event.target.tagName === 'A') {
        if ((!!event.target.download && event.target.download.length > 0) ||
            (!!event.target.dataset.gaClick && event.target.dataset.gaClick.startsWith('Repository, download zip'))){
          download(event.target.href);
          return false;
        }
      }
    }, {passive: false} );
};
