/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-12-11
 * @fileoverview: This file searches for a word in a given sentence 
 */


// INPUT
let sentence: string = prompt("Please enter a sentence?") || "";
let word: string = prompt("Please enter a word to search for in your sentence?") || "";

// asssign variables using inputs 
let found: boolean = false;
let sentenceLower = sentence.toLowerCase();
let wordLower = word.toLowerCase();

// PROCESS
for (let i = 0; i <= sentenceLower.length - wordLower.length; i++) {
  let match = true;

  for (let j = 0; j < wordLower.length; j++) {
    if (sentenceLower.charAt(i + j) !== wordLower.charAt(j)) {
      match = false;
    }
  }

  if (match === true) {
    found = true;
  }
}

// OUTPUT
if (found === true) {
  console.log("Hooray, the word " + word + " was found in the sentence: " + sentence);
} else {
  console.log("Sorry, the word " + word + " was not found in the sentence: " + sentence);
}

console.log("\nDone.");