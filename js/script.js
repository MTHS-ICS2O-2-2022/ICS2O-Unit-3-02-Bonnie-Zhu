// Created by: Bonnie Zhu
// Created on: April. 2023
// This file contains the JS functions for index.html

'use strict'
/**
* This function calculates your pay depending on the amount of hours you work (including tax).
*/
function calculate() {
  // input
  const Length = parseInt(document.getElementById('Length').value)
  const Width = parseInt(document.getElementById('Width').value)
  const Height = parseInt(document.getElementById('Height').value)

  // process
  const Volume = (Length * Width * Height) /3

  // output
  document.getElementById('Volume').innerHTML = 'Volume is: ' + Volume + ' mm3'
}

