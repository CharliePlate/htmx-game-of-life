let isMouseDown = false;

function checkMouseState() {
  return isMouseDown;
}

document.addEventListener("mousedown", function () {
  isMouseDown = true;
});

document.addEventListener("mouseup", function () {
  isMouseDown = false;
});
