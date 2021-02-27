window.onload = function() {
  document.body.addEventListener('keydown',
    event => {
      if (event.key === 'v' && event.ctrlKey) {
        response("keydown");
      }
    });
};
