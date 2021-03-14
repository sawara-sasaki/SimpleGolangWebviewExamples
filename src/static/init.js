var simpleGolangWebview = {};
simpleGolangWebview.isLocal = window.location.href.startsWith('data');
window.onload = function() {
  if (!simpleGolangWebview.isLocal) {
    Array.prototype.forEach.call(document.getElementsByTagName('a'), (element) => element.setAttribute('target', ''));
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
