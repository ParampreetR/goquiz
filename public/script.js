let buttons = document.getElementsByClassName("option-btn");

for (let button of buttons) {
  button.addEventListener("click", (e) => {
    e.preventDefault();
    console.log(e.target.dataset.index);
    e.target.parentElement.querySelector("#answer").value =
      e.target.dataset.index;

    e.target.parentElement.querySelectorAll("button").forEach((elmt) => {
      console.log(elmt.classList);
      elmt.classList.remove("option-selected");
    });

    e.target.classList.add("option-selected");
  });
}
