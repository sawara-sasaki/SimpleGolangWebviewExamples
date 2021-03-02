window.onload = function() {
  document.body.addEventListener('keydown',
    event => {
      if (event.key === 'c' && event.ctrlKey) {
        log("Ctrl+C");
      } else if (event.key === 'v' && event.ctrlKey) {
        log("Ctrl+V");
      }
    });
};
