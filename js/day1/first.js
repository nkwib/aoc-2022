import * as utils from './../utils.js';

const getReducedList = () => {
  let input = utils.readFile(1, "one");
  let elves = input.split('\n\n');
  elves = elves.map((elf) => {
    let list = elf.split('\n');
    let total = 0;

    list.forEach((item) => {
      total += parseInt(item);
    });

    return total;
  });
  return elves;
}

export const first = () => {
  let elves = getReducedList();
  let mostCaloriesElf = elves[0];
  elves.forEach((elf) => {
    if (elf > mostCaloriesElf) {
      mostCaloriesElf = elf;
    }
  })
  return mostCaloriesElf;
}

export const second = () => {
  let elves = getReducedList();
  elves.sort((a, b) => b - a);
  let firstThreeElves = elves.slice(0, 3);
  return firstThreeElves.reduce((a, b) => a + b, 0);;
}