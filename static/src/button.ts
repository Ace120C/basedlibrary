const button = document.querySelector('.button-select') as HTMLButtonElement;

function buttonSelected(){
  button.style.backgroundColor = 'rgb(65,74,76)'
}

function buttonUnselected(){
  button.style.backgroundColor = 'rgb(0,0,0)'
}

button.addEventListener('mouseenter', buttonSelected);
button.addEventListener('mouseleave', buttonUnselected);


