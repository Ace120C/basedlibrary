
const SerachButton = document.querySelector(".button-select") as HTMLButtonElement;

if (SerachButton) {
  SerachButton.addEventListener('click', fetchBook)
}

async function fetchBook() {

try {
 const input = (document.querySelector(".searchbar")as HTMLInputElement).value.toLowerCase();
 const response = await fetch('http://localhost:8080/dbOUT.json');

 if (!response.ok) {
  throw new Error("could not fetch the database");
 }

const data = await response.json();

const matchedBook = data.find((book: { name: string }) => book.name.toLowerCase() === input);

if (matchedBook) {
  window.location.href = matchedBook.LK;
} else {
  alert('book not found.')
}

} catch (error) {
 console.error(error);
  
 }
}
