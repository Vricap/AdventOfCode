const str = `3x1x4
3x5x1
2x4x6
20x23x3
1x17x21
21x21x27
13x1x24
7x30x20
21x9x18
23x26x6
22x9x29
17x6x21
28x28x29
19x25x26
9x27x21
5x26x8
11x19x1
10x1x18
29x4x8
21x2x22
14x12x8`;

const array = str.split("\n");
const x = [];

for (let i = 0; i < array.length; i++) {
  x.push(array[i].split("x"));
}

for (let i = 0; i < x.length; i++) {
  console.log(i);
}

let total = 0;

for (let i = 0; i < x.length; i++) {
  let count = 0;
  let smallest = Number(x[i][0]);
  for (let j = 0; j < x[i].length; j++) {
    if (Number(x[i][j]) < smallest) {
      // console.log(Number(x[i][j]));
      smallest = Number(x[i][j]);
    }
    if (j === x[i].length - 1) {
      count += 2 * Number(x[i][j]) * Number(x[i][0]) + smallest;
      break;
    }
    count += 2 * Number(x[i][j]) * Number(x[i][j + 1]);
  }
  total += count;
  console.log(total);
}
console.log(total);
