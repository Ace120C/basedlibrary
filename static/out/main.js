"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
const SerachButton = document.querySelector(".button-select");
if (SerachButton) {
    SerachButton.addEventListener('click', fetchBook);
}
function fetchBook() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const input = document.querySelector(".searchbar").value.toLowerCase();
            const response = yield fetch('http://localhost:8080/dbOUT.json');
            if (!response.ok) {
                throw new Error("could not fetch the database");
            }
            const data = yield response.json();
            const matchedBook = data.find((book) => book.name.toLowerCase() === input);
            if (matchedBook) {
                window.location.href = matchedBook.LK;
            }
            else {
                alert('book not found.');
            }
        }
        catch (error) {
            console.error(error);
        }
    });
}
