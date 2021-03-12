window.onload = function() {
  document.body.addEventListener('keydown',
    event => {
      if (event.key === 'c' && event.ctrlKey) {
        log("Ctrl+C");
      } else if (event.key === 'v' && event.ctrlKey) {
        debug();
      } else if (event.key === 's' && event.ctrlKey) {
        if (!window.location.href.startsWith('data')) {
          src(window.location.href, document.documentElement.innerHTML);
        }
      }
    });
};
